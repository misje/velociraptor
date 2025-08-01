/*
Velociraptor - Dig Deeper
Copyright (C) 2019-2025 Rapid7 Inc.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/
package authenticators

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Velocidex/ordereddict"
	"github.com/gorilla/csrf"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"www.velocidex.com/golang/velociraptor/acls"
	api_proto "www.velocidex.com/golang/velociraptor/api/proto"
	api_utils "www.velocidex.com/golang/velociraptor/api/utils"
	utils "www.velocidex.com/golang/velociraptor/api/utils"
	config_proto "www.velocidex.com/golang/velociraptor/config/proto"
	"www.velocidex.com/golang/velociraptor/constants"
	"www.velocidex.com/golang/velociraptor/gui/velociraptor"
	"www.velocidex.com/golang/velociraptor/json"
	"www.velocidex.com/golang/velociraptor/logging"
	"www.velocidex.com/golang/velociraptor/services"
)

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

type GoogleAuthenticator struct {
	config_obj    *config_proto.Config
	authenticator *config_proto.Authenticator
}

func (self *GoogleAuthenticator) LoginHandler() string {
	return "/auth/google/login"
}

// The URL that will be used to log in.
func (self *GoogleAuthenticator) LoginURL() string {
	return "/auth/google/login"
}

func (self *GoogleAuthenticator) CallbackHandler() string {
	return "/auth/google/callback"
}

func (self *GoogleAuthenticator) CallbackURL() string {
	return "/auth/google/callback"
}

func (self *GoogleAuthenticator) ProviderName() string {
	return "Google"
}

func (self *GoogleAuthenticator) AddHandlers(mux *api_utils.ServeMux) error {
	mux.Handle(api_utils.GetBasePath(self.config_obj, self.LoginHandler()),
		IpFilter(self.config_obj, self.oauthGoogleLogin()))
	mux.Handle(api_utils.GetBasePath(self.config_obj, self.CallbackHandler()),
		IpFilter(self.config_obj, self.oauthGoogleCallback()))

	return nil
}

func (self *GoogleAuthenticator) AddLogoff(mux *api_utils.ServeMux) error {
	installLogoff(self.config_obj, mux)
	return nil
}

func (self *GoogleAuthenticator) IsPasswordLess() bool {
	return true
}

func (self *GoogleAuthenticator) RequireClientCerts() bool {
	return false
}

func (self *GoogleAuthenticator) AuthRedirectTemplate() string {
	return self.authenticator.AuthRedirectTemplate
}

// Check that the user is proerly authenticated.
func (self *GoogleAuthenticator) AuthenticateUserHandler(
	parent http.Handler,
	permission acls.ACL_PERMISSION,
) http.Handler {

	return authenticateUserHandle(
		self.config_obj,
		permission,
		func(w http.ResponseWriter, r *http.Request, err error, username string) {
			reject_with_username(self.config_obj, w, r, err,
				username, self.LoginURL(), self.ProviderName())
		},
		parent)
}

func (self *GoogleAuthenticator) oauthGoogleLogin() http.Handler {

	return api_utils.HandlerFunc(nil,
		func(w http.ResponseWriter, r *http.Request) {
			googleOauthConfig, _ := self.GetGenOauthConfig()

			// Create oauthState cookie
			oauthState, err := r.Cookie("oauthstate")
			if err != nil {
				oauthState = generateStateOauthCookie(self.config_obj, w)
			}

			u := googleOauthConfig.AuthCodeURL(oauthState.Value, oauth2.ApprovalForce)
			http.Redirect(w, r, u, http.StatusTemporaryRedirect)
		})
}

func generateStateOauthCookie(
	config_obj *config_proto.Config,
	w http.ResponseWriter) *http.Cookie {
	// Do not expire from the browser - we will expire it anyway.
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	_, _ = rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{
		Name:     "oauthstate",
		Path:     utils.GetBasePath(config_obj),
		Value:    state,
		Secure:   true,
		HttpOnly: true,
		Expires:  expiration}
	http.SetCookie(w, &cookie)
	return &cookie
}

func (self *GoogleAuthenticator) oauthGoogleCallback() http.Handler {

	return api_utils.HandlerFunc(nil,
		func(w http.ResponseWriter, r *http.Request) {
			// Read oauthState from Cookie
			oauthState, _ := r.Cookie("oauthstate")
			if oauthState == nil || r.FormValue("state") != oauthState.Value {
				logging.GetLogger(self.config_obj, &logging.GUIComponent).
					Error("invalid oauth google state")
				http.Redirect(w, r, utils.Homepage(self.config_obj),
					http.StatusTemporaryRedirect)
				return
			}

			data, err := self.getUserDataFromGoogle(r.Context(), r.FormValue("code"))
			if err != nil {
				logging.GetLogger(self.config_obj, &logging.GUIComponent).
					WithFields(logrus.Fields{
						"err": err.Error(),
					}).Error("getUserDataFromGoogle")
				http.Redirect(w, r, utils.Homepage(self.config_obj),
					http.StatusTemporaryRedirect)
				return
			}

			user_info := &api_proto.VelociraptorUser{}
			err = json.Unmarshal(data, &user_info)
			if err != nil {
				logging.GetLogger(self.config_obj, &logging.GUIComponent).
					WithFields(logrus.Fields{
						"err": err.Error(),
					}).Error("getUserDataFromGoogle")
				http.Redirect(w, r, utils.Homepage(self.config_obj),
					http.StatusTemporaryRedirect)
				return
			}

			// Update the user picture in the datastore if we can - it
			// will be populated from there for each GetUserUITraits
			// call. This keeps our cookie smaller.
			setUserPicture(r.Context(), user_info.Email, user_info.Picture)

			// Sign and get the complete encoded token as a string using the secret
			cookie, err := getSignedJWTTokenCookie(
				self.config_obj, self.authenticator,
				&Claims{
					Username: user_info.Email,
				}, r)
			if err != nil {
				logging.GetLogger(self.config_obj, &logging.GUIComponent).
					WithFields(logrus.Fields{
						"err": err.Error(),
					}).Error("getUserDataFromGoogle")
				http.Redirect(w, r, utils.Homepage(self.config_obj),
					http.StatusTemporaryRedirect)
				return
			}

			http.SetCookie(w, cookie)
			http.Redirect(w, r, utils.Homepage(self.config_obj),
				http.StatusTemporaryRedirect)
		})
}

func (self *GoogleAuthenticator) GetGenOauthConfig() (*oauth2.Config, error) {
	res := &oauth2.Config{
		RedirectURL:  api_utils.GetPublicURL(self.config_obj, self.CallbackURL()),
		ClientID:     self.authenticator.OauthClientId,
		ClientSecret: self.authenticator.OauthClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	return res, nil
}

func (self *GoogleAuthenticator) getUserDataFromGoogle(
	ctx context.Context, code string) ([]byte, error) {

	// Use code to get token and get user info from Google.
	googleOauthConfig, _ := self.GetGenOauthConfig()

	token, err := googleOauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(
		io.LimitReader(response.Body, constants.MAX_MEMORY))
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}

	return contents, nil
}

func installLogoff(config_obj *config_proto.Config, mux *api_utils.ServeMux) {
	mux.Handle(utils.GetBasePath(config_obj, "/app/logoff.html"),
		IpFilter(config_obj,
			api_utils.HandlerFunc(nil,
				func(w http.ResponseWriter, r *http.Request) {
					params := r.URL.Query()
					old_username, ok := params["username"]
					username := ""
					if ok && len(old_username) == 1 {
						err := services.LogAudit(r.Context(),
							config_obj, old_username[0], "LogOff", ordereddict.NewDict())
						if err != nil {
							logger := logging.GetLogger(
								config_obj, &logging.FrontendComponent)
							logger.Error("LogAudit: LogOff %v", old_username[0])
						}
						username = old_username[0]
					}

					// Clear the cookie
					http.SetCookie(w, &http.Cookie{
						Name:     "VelociraptorAuth",
						Path:     utils.GetBaseDirectory(config_obj),
						Value:    "deleted",
						Secure:   true,
						HttpOnly: true,
						Expires:  time.Unix(0, 0),
					})

					renderLogoffMessage(config_obj, w, username)
				})))
}

func authenticateUserHandle(
	config_obj *config_proto.Config,
	permission acls.ACL_PERMISSION,
	reject_cb func(w http.ResponseWriter, r *http.Request,
		err error, username string),
	parent http.Handler) http.Handler {

	logger := GetLoggingHandler(config_obj)(parent)

	return api_utils.HandlerFunc(parent,
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-CSRF-Token", csrf.Token(r))

			claims, err := getDetailsFromCookie(config_obj, r)
			if err != nil {
				reject_cb(w, r, err, claims.Username)
				return
			}

			username := claims.Username

			// Now check if the user is allowed to log in.
			users := services.GetUserManager()
			user_record, err := users.GetUser(r.Context(), username, username)
			if err != nil {
				reject_cb(w, r, fmt.Errorf("Invalid user: %v", err), username)
				return
			}

			// Does the user have access to the specified org?
			err = CheckOrgAccess(config_obj, r, user_record, permission)
			if err != nil {
				reject_cb(w, r, fmt.Errorf("Insufficient permissions: %v", err), user_record.Name)
				return
			}

			// Checking is successful - user authorized. Here we
			// build a token to pass to the underlying GRPC
			// service with metadata about the user.
			user_info := &api_proto.VelociraptorUser{
				Name: user_record.Name,
			}

			// NOTE: This context is NOT the same context that is received
			// by the API handlers. This context sits on the incoming side
			// of the GRPC gateway. We stuff our data into the
			// GRPC_USER_CONTEXT of the context and the code will convert
			// this value into a GRPC metadata.

			// Must use json encoding because grpc can not handle
			// binary data in metadata.
			serialized, _ := json.Marshal(user_info)
			ctx := context.WithValue(
				r.Context(), constants.GRPC_USER_CONTEXT, string(serialized))

			// Need to call logging after auth so it can access
			// the contextKeyUser value in the context.
			logger.ServeHTTP(w, r.WithContext(ctx))
		}).AddChild("GetLoggingHandler")
}

func reject_with_username(
	config_obj *config_proto.Config,
	w http.ResponseWriter, r *http.Request,
	err error, username, login_url, provider string) {

	// Log failed login to the audit log only if there is an actual
	// user. First redirect will have username blank.
	if username != "" {
		err := services.LogAudit(r.Context(),
			config_obj, username, "User rejected by GUI",
			ordereddict.NewDict().
				Set("remote", r.RemoteAddr).
				Set("method", r.Method).
				Set("url", r.URL.String()).
				Set("err", err.Error()))
		if err != nil {
			logger := logging.GetLogger(
				config_obj, &logging.FrontendComponent)
			logger.Error("LogAudit: User rejected by GUI %v %v",
				username, r.RemoteAddr)
		}

	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusUnauthorized)

	renderRejectionMessage(config_obj,
		r, w, err, username, []velociraptor.AuthenticatorInfo{
			{
				LoginURL:     api_utils.PublicURL(config_obj, login_url),
				ProviderName: provider,
			},
		})
}
