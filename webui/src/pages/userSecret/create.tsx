// @BeeOverwrite YES
// @BeeGenerateTime 20200831_101837
import {Form, Input, Button, Select, Card, message} from 'antd';
import {PageHeaderWrapper} from "@ant-design/pro-layout";
import React from "react";
import CommonForm from "./formconfig"
import request from "@/utils/request";
const apiUrl = "/api/admin/userSecret";

export default class Base extends React.Component {
  reqCreate(params: any) {
      return request(apiUrl+"/create", {
        method: 'POST',
        data: { ...params},
      });
  }
  render() {


    return (
      <PageHeaderWrapper>
        <Card>
          <CommonForm request={this.reqCreate} />
        </Card>
      </PageHeaderWrapper>
    );
  }
}
