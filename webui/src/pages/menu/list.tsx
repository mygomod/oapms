// @BeeOverwrite NO
// @BeeGenerateTime 20200820_230345
import ProTable, { ProColumns, ActionType } from '@ant-design/pro-table';
import React, {useState, useRef, Fragment, useEffect} from 'react';
import ListTable from "./components/ListTable"


export default class List extends React.Component {
  state = {
    done: false,
    treeData: [],
    appId: 1,
    pid: null,
  }


  componentDidMount() {
    const { location: { query: { aid } } } = this.props;
    // NaN
    let aidInt = parseInt(aid)
    if (aidInt > 0) {
      this.setState({
        aid: aidInt,
      })
    }else {
      this.setState({
        aid: 1,
      })
    }
  }

  setTreeData = (treeData) => {
    this.setState({
      treeData,
    })
  };

  setAppId = (appId) => {
    this.setState({
      appId,
    })
  };

  setPid = (pid) => {
    this.setState({
      pid,
    })
  };

  render() {
    const { treeData,appId,pid } = this.state

    return(
      <div>
       <ListTable
          treeData={treeData}
          setTreeData={this.setTreeData}
          setAppId={this.setAppId}
          setPid={this.setPid}
          params={{
            appId,
            pid,
          }}
        />
      </div>
    )
  }
}
