import _ from 'lodash';
import "./event-timeline.css";

import React, { Component }  from 'react';
import PropTypes from 'prop-types';
import Timeline, {
    TimelineMarkers,
    CustomMarker,
} from 'react-calendar-timeline';
import moment from 'moment';
import 'moment-timezone';
import {CancelToken} from 'axios';
import api from '../core/api-service.jsx';
import VeloValueRenderer from '../utils/value.jsx';
import Dropdown from 'react-bootstrap/Dropdown';
import T from '../i8n/i8n.jsx';
import ToolTip from '../widgets/tooltip.jsx';
import Table from 'react-bootstrap/Table';

import DeleteTimelineRanges from './delete.jsx';
import Button from 'react-bootstrap/Button';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

import { ColumnToggle } from '../core/paged-table.jsx';

import {
    getFormatter,
    PrepareData,
} from '../core/table.jsx';


// We want to work in UTC everywhere
moment.tz.setDefault("Etc/UTC");

const normalizeTimestamp = (value) => {
    if (_.isNumber(value) && value > 0) {
        if (value > 20000000000) {
            value /= 1000;
        }

        if (value > 20000000000) {
            value /= 1000;
        }

        return value * 1000;
    }
    return value;
};


class EventTableRenderer  extends Component {
    static propTypes = {
        renderers: PropTypes.object,
        columns: PropTypes.array,
        rows: PropTypes.array,
        toggles: PropTypes.object,
        env: PropTypes.object,
    }

    state = {
        download: false,
    }

    componentDidUpdate(prevProps, prevState, snapshot) {
    }

    defaultFormatter = (cell, row, rowIndex) => {
        return <VeloValueRenderer value={cell}/>;
    }

    activeColumns = ()=>{
        let res = [];
        _.each(this.props.columns, c=>{
            if(!this.props.toggles[c]) {
                res.push(c);
            }
        });
        return res;
    }

    renderHeader = (column, idx)=>{
        let column_name = column;
        if (column_name == "_ts") {
            column_name = "Server Time";
        }

        return <th key={idx}>{ T(column_name) } </th>;
    }

    getColumnRenderer = column => {
        if(this.props.renderers && this.props.renderers[column]) {
            return this.props.renderers[column];
        }

        let column_types = this.state.column_types;
        if (!_.isArray(column_types)) {
            return this.defaultFormatter;
        }

        for (let i=0; i<column_types.length; i++) {
            if (column === column_types[i].name) {
                return getFormatter(column_types[i].type, column_types[i].name);
            }
        }
        return this.defaultFormatter;
    }

    renderCell = (column, row, rowIdx) => {
        let t = this.props.toggles[column];
        if(t) {return undefined;};

        let cell = row[column];
        let renderer = this.getColumnRenderer(column);

        return <td key={column}>
                 { renderer(cell, row, this.props.env)}
               </td>;
    };

    selectRow = (row, idx)=>{
        this.setState({selected_row: row, selected_row_idx: idx});
    }

    renderRow = (row, idx)=>{
        let selected_cls = "row-selected";

        if(this.state.selected_row_idx !== idx) {
            selected_cls = "";
        }

        return (
            <tr key={idx}
                onClick={x=>this.selectRow(row, idx)}
                className={selected_cls}>
              {_.map(this.activeColumns(), c=>this.renderCell(c, row, idx))}
            </tr>);
    }

    render() {
        if (!this.props.rows || !this.props.columns) {
            return <div></div>;
        }

        return (
            <>
              <Table className="paged-table">
                <thead>
                  <tr className="paged-table-header">
                    {_.map(this.activeColumns(), this.renderHeader)}
                  </tr>
                </thead>
                <tbody className="fixed-table-body">
                  {_.map(this.props.rows, this.renderRow)}
                </tbody>
              </Table>
            </>);
    }
}

export default class EventTimelineViewer extends React.Component {
    static propTypes = {
        // Render the toolbar buttons in our parent component.
        toolbar: PropTypes.func,
        client_id: PropTypes.string,
        artifact: PropTypes.string,
        mode: PropTypes.string,
        renderers: PropTypes.object,
        column_types: PropTypes.array,
        time_range_setter: PropTypes.func,
    };

    componentDidMount = () => {
        this.source = CancelToken.source();
        this.ts_source = CancelToken.source();
        this.fetchAvailableTimes();
    }

    componentDidUpdate(prevProps, prevState, snapshot) {
        if (!_.isEqual(prevState.start_time, this.state.start_time)) {
            this.fetchRows();
            return true;
        };

        if (!_.isEqual(prevState.table_start, this.state.table_start)) {
            return true;
        }

        if (!_.isEqual(prevProps.mode, this.props.mode)) {
            this.fetchRows();
            return true;
        };

        if (!_.isEqual(prevProps.artifact, this.props.artifact)) {
            this.fetchAvailableTimes();
            return true;
        };

        if (!_.isEqual(prevState.row_count, this.state.row_count)) {
            this.fetchRows();
            return true;
        };

        return false;
    }

    componentWillUnmount() {
        this.source.cancel();
        this.ts_source.cancel();
    }

    fetchAvailableTimes = () => {
        this.ts_source.cancel();
        this.ts_source = CancelToken.source();

        let client_id = this.props.client_id || "server";

        api.post("v1/ListAvailableEventResults", {
            client_id: client_id,
            artifact: this.props.artifact,
        }, this.source.token).then(resp => {
            if (resp.cancel || !resp.data.logs) return;

            let av_t = resp.data.logs[0].row_timestamps;
            if (this.props.mode === "Logs") {
                av_t = resp.data.logs[0].log_timestamps;
            }

            if (av_t && av_t.length > 0) {
                let ts = av_t[av_t.length-1]*1000;
                this.setState({
                    start_time: ts || 0,
                    table_start: ts,
                    end_time: ts + 60*60*24*1000,
                });
                this.centerPage();
            }
            this.setState({
                available_timestamps: resp.data.logs[0].row_timestamps || [],
                available_log_timestamps: resp.data.logs[0].log_timestamps || [],
                toggles: {},
            });

            this.fetchRows();
        });
    }

    fetchRows = () => {
        let url = "v1/GetTable";

        this.source.cancel();
        this.source = CancelToken.source();

        this.setState({loading: true});

        let mode = "CLIENT_EVENT";
        if (this.props.mode === "Logs") {
            mode = "CLIENT_EVENT_LOGS";
        }

        let params = {
            client_id: this.props.client_id,
            artifact: this.props.artifact,
            type: mode,
            start_time: parseInt(this.state.start_time / 1000),
            end_time: 2000000000,
            rows: this.state.row_count || 10,
        };

        api.get(url, params, this.source.token).then((response) => {
            if (response.cancel) {
                return;
            }

            let pageData = PrepareData(response.data);

            // The first time is the _ts column of the first row
            let table_start = null;
            let table_end = null;

            let rows = pageData.rows;
            if (rows.length > 0) {
                table_start = normalizeTimestamp(rows[0]["_ts"]);
                table_end = normalizeTimestamp(rows[rows.length-1]["_ts"]);
            }

            let columns = ["_ts"];
            _.each(pageData.columns, x=>{
                if(x!=="_ts") {
                    columns.push(x);
                }
            });

            this.setState({
                table_start: table_start,
                table_end:  table_end,
                columns: columns,
                rows: pageData.rows,
            });

            if(_.isEmpty(this.state.toggles)) {
                let toggles = {};
                _.each(columns, x=>{
                    if(_.isString(x) &&
                       x.length>0 && x[0] === "_" && x !== "_ts") {
                        toggles[x]=true;
                    } else {
                        toggles[x]=false;
                    }
                });

                this.setState({toggles: toggles});
            }

            // Re-render the toolbar each time we fetch a new row.
            this.props.toolbar(this.renderToolbar);
        });
    };

    pageSizeSelector = () => {
        let options = [10, 20, 50, 100];

        return <Dropdown
                 className="page-size-dropdown btn-group"

                 onSelect={(value)=>{
                     this.setState({row_count: value});
                 }}
                 >
                 <Dropdown.Toggle variant="default" id="row_count_selector">
                   {this.state.row_count}
                 </Dropdown.Toggle>

                 <Dropdown.Menu>
                   {_.map(options, (x, idx)=>{
                       return <Dropdown.Item
                                key={idx}
                                href="#" disabled={x===this.state.row_count}
                                eventKey={x}>{x}
                              </Dropdown.Item>;
                   })}
                 </Dropdown.Menu>
               </Dropdown>;
    }

    centerPage = () => {
        let page_size = this.state.visibleTimeEnd - this.state.visibleTimeStart;
        if (page_size === 0) {
            page_size = 60*60*24*1000;
        }
        let visibleTimeStart = this.state.table_start - page_size/2;
        let visibleTimeEnd = this.state.table_start + page_size/2;
        this.setState({
            start_time: this.state.table_start,
            visibleTimeStart: visibleTimeStart,
            visibleTimeEnd: visibleTimeEnd,
        });
        this.props.time_range_setter(visibleTimeStart, visibleTimeEnd);
    }

    // Jump to the previous page.
    prevPage = ()=>{
        this.setState({
            start_time: (this.state.start_time - 60*60*24*1000) || 0,
        });
        this.fetchRows();
    }

    nextPage = ()=>{
        if (this.state.table_end > 0) {
            let page_size = this.state.visibleTimeEnd - this.state.visibleTimeStart;
            this.setState({
                start_time: (this.state.table_end + 1000) || 0,
            });

            // Only scroll the timeline once we go past the view port.
            if (this.state.table_end > this.state.visibleTimeEnd) {
                let visibleTimeStart = this.state.table_end + 1;
                let visibleTimeEnd = this.state.table_end + page_size;
                this.setState({
                    visibleTimeStart: visibleTimeStart,
                    visibleTimeEnd: visibleTimeEnd,
                });
                this.props.time_range_setter(visibleTimeStart, visibleTimeEnd);
            }

            this.fetchRows();
        }
    }

    startDownload = (type) => {
        return true;
    }

    renderToolbar = () => {
        let mode = "CLIENT_EVENT";
        if (this.props.mode === "Logs") {
            mode = "CLIENT_EVENT_LOGS";
        }

        let start_time = moment(this.state.visibleTimeStart).format();
        let end_time = moment(this.state.visibleTimeEnd).format();
        let basename = `${this.props.artifact}-${start_time}-${end_time}-${this.props.client_id}`;

        // For JSON do not expand the columns - just pass them through
        // as they are.
        let downloads_json = {
            client_id: this.props.client_id,
            artifact: this.props.artifact,
            type: mode,
            start_time: parseInt(this.state.visibleTimeStart / 1000),
            end_time: parseInt(this.state.visibleTimeEnd / 1000),
            rows: 1,
            download_format: "json",
            download_filename: basename,
        };

        // For CSV we want all the columns expanded. This is not ideal
        // as it only includes the columns in this page.
        let downloads_csv = Object.assign({}, downloads_json);
        downloads_csv.download_format = "csv";
        downloads_csv.columns = this.state.columns;

        return <>
                 <ColumnToggle onToggle={(c)=>{
                     // Do not make a copy here because set state is
                     // not immediately visible and this will be called
                     // for each column.
                     let toggles = this.state.toggles;
                     toggles[c] = !toggles[c];
                     this.setState({toggles: toggles});
                 }}
                               columns={this.state.columns}
                               toggles={this.state.toggles} />
                 {this.pageSizeSelector()}
                 <Dropdown className="btn-group">
                   <Dropdown.Toggle variant="default">
                     <FontAwesomeIcon icon="download"/>
                   </Dropdown.Toggle>
                   <Dropdown.Menu>
                     <Dropdown.Item as="a"
                       href={api.href("/api/v1/DownloadTable", downloads_csv)}
                       variant="default" type="button">
                       <FontAwesomeIcon icon="file-csv"/>
                       <span className="button-label">
                         <div className="download-format">CSV</div>
                         <div className="download-time-range">
                           {start_time} - {end_time}
                         </div>
                       </span>
                     </Dropdown.Item>
                     <Dropdown.Item as="a"
                       href={api.href("/api/v1/DownloadTable", downloads_json)}
                       variant="default" type="button">
                       <FontAwesomeIcon icon="file-code"/>
                       <span className="button-label">
                         <div className="download-format">JSON</div>
                         <div className="download-time-range">
                           {start_time} - {end_time}
                         </div>
                       </span>
                     </Dropdown.Item>
                   </Dropdown.Menu>
                 </Dropdown>
                 <ToolTip tooltip={T("Delete")}>
                   <Button onClick={() => this.setState({showDeleteDialog: true})}
                           variant="default">
                     <FontAwesomeIcon icon="trash"/>
                   </Button>
                 </ToolTip>
                 <ToolTip tooltip={T("Previous")}>
                   <Button onClick={() => this.prevPage()}
                           variant="default">
                     <FontAwesomeIcon icon="backward"/>
                   </Button>
                 </ToolTip>
                 <ToolTip tooltip={T("Center")}>
                   <Button onClick={() => this.centerPage()}
                           variant="default">
                     <FontAwesomeIcon icon="crosshairs"/>
                   </Button>
                 </ToolTip>
                 <ToolTip tooltip={T("Next")}>
                   <Button onClick={() => this.nextPage()}
                           variant="default">
                     <FontAwesomeIcon icon="forward"/>
                   </Button>
                 </ToolTip>
               </>;
    }

    handleTimeChange = (visibleTimeStart, visibleTimeEnd) => {
        this.setState({
            visibleTimeStart,
            visibleTimeEnd,
            scrolling: true
        });
        this.props.time_range_setter(visibleTimeStart, visibleTimeEnd);
    };

    state = {
        start_time: 0,
        table_start: 0,
        table_end: 0,
        row_count: 10,
        visibleTimeStart: moment().startOf("day").valueOf(),
        visibleTimeEnd: moment().startOf("day").add(1, "day").valueOf(),

        available_timestamps: [],
        available_log_timestamps: [],

        showDeleteDialog: false,

        toggles: {},
    }

    render() {
        let groups =  [
            {
                id: -1,
                title: "Table View",
            },
            {
                id: 1,
                title: "Available",
            },
            {
                id: 2,
                title: "Logs",
            },
        ];

        let items = [{
            id:-1, group: -1,
            start_time: moment(this.state.table_start),
            end_time: moment(this.state.table_end),
            canMove: false,
            canResize: false,
            canChangeGroup: false,
            itemProps: {
                className: 'timeline-table-item',
                style: {
                    background: undefined,
                    color: undefined,
                },
            },
        }];


        let adder = (ts, group_id)=>{
            ts = ts || 0;
            items.push({
                id: items.length, group: group_id,
                ts: ts,
                start_time: ts * 1000,
                end_time: (ts + 60*60*24)*1000,
                canMove: false,
                canResize: false,
                canChangeGroup: false,
                itemProps: {
                    className: 'timeline-table-item',
                    style: {
                        background: undefined,
                        color: undefined,
                    },
                },
            });
        };

        _.each(this.state.available_timestamps, x=>adder(x, 1));
        _.each(this.state.available_log_timestamps, x=>adder(x, 2));

        let visible_start_time = this.state.visibleTimeStart || 0;
        let visible_end_time = this.state.visibleTimeEnd || 200000000;

        // Disable buffer to prevent horizontal scroll. This seems to
        // interact badly with MacOS trackpads.
        return <>
                 {this.state.showDeleteDialog &&
                  <DeleteTimelineRanges
                    client_id={this.props.client_id}
                    artifact={this.props.artifact}
                    start_time={visible_start_time}
                    end_time={visible_end_time}
                    onClose={()=>{
                        this.setState({showDeleteDialog: false});

                        // Trigger a refresh of the table
                        this.fetchAvailableTimes();
                    }}
                  />}

                 <Timeline
                   groups={groups}
                   items={items}
                   defaultTimeStart={moment().add(-1, "day")}
                   defaultTimeEnd={moment().add(1, "day")}
                   itemTouchSendsClick={true}
                   minZoom={5*60*1000}
                   buffer={1}
                   dragSnap={1000}
                   onCanvasClick={(groupId, time, e) => {
                       if(time) {
                           this.setState({start_time: time});
                       }
                   }}
                   onItemSelect={(itemId, e, time) => {
                       if(time) {
                           this.setState({start_time: time});
                       }
                       return false;
                   }}
                   onItemClick={(itemId, e, time) => {
                       if(time) {
                           this.setState({start_time: time});
                       }
                       return false;
                   }}
                   visibleTimeStart={this.state.visibleTimeStart}
                   visibleTimeEnd={this.state.visibleTimeEnd}
                   onTimeChange={this.handleTimeChange}
                 >
                   <TimelineMarkers>
                     <CustomMarker
                       date={this.state.start_time || Date.now()} >
                       { ({ styles, date }) => {
                           styles.backgroundColor = undefined;
                           styles.width = undefined;
                           return <div style={styles}
                                       className="timeline-marker"
                                  />;
                       }}
                     </CustomMarker>
                   </TimelineMarkers>
                 </Timeline>
                 <EventTableRenderer
                   renderers={this.props.renderers}
                   toggles={this.state.toggles}
                   rows={this.state.rows}
                   columns={this.state.columns}
                 />
               </>;
    }
};
