// @BeeOverwrite YES
// @BeeGenerateTime 20200831_101837
import {Form, Input, Button, Select, Card, message} from 'antd';
import {PageHeaderWrapper} from "@ant-design/pro-layout";
import React from "react";
import FormConfig from "./formconfig"
import request from "@/utils/request";
const apiUrl = "/api/admin/app";

export default class Base extends React.Component {
  state = {
    data: null,
  };

  async componentDidMount() {
    const { location: { query: { aid } } } = this.props;
    this.reqInfo(aid).then(res => {
      if (res.code !== 0) {
        message.error(res.msg);
        return false;
      }
      this.setState({
        data: res.data,
      });
      return true;
    });
  }

  reqInfo(id) {
    return request(apiUrl+"/info?aid="+id);
  }

  render() {
    const { data } = this.state;
    return (
      <PageHeaderWrapper>
        <Card>
          { data && <FormConfig  initialValues={data} mode="view" />}
        </Card>
      </PageHeaderWrapper>
    );
  }
}
