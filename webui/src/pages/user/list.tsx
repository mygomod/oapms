// @BeeOverwrite NO
// @BeeGenerateTime 20200820_230345
import React from 'react';
import ListTable from "./components/ListTable"


export default class List extends React.Component {
  state = {
    aid: 1,
  };


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

  setAid = (aid: any) => {
    this.setState({
      aid,
    })
  };


  render() {
    const { aid } = this.state
    return(
      <div>
        <ListTable
          setAid={this.setAid}
          aid={aid}
        />
      </div>
    )
  }
}
