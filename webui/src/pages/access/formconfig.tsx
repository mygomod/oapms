// @BeeOverwrite YES
// @BeeGenerateTime 20200820_230345
import {Form, Input, Button, Select, Card, message} from 'antd';
import React from "react";
import {history} from "umi";

const formItemLayout = {
  labelCol: {
    xs: { span: 18 },
    sm: { span: 8 },
  },
  wrapperCol: {
    xs: { span: 18 },
    sm: { span: 16 },
  },
};

const tailFormItemLayout = {
  wrapperCol: {
    xs: {
      span: 24,
      offset: 0,
    },
    sm: {
      span: 16,
      offset: 8,
    },
  },
};

export default function CommonForm (props) {
  const [form] = Form.useForm();
  let initialValues = props.initialValues
  let request = props.request
  let mode = props.mode

  const onFinish = (values) => {
    request({
      ...values,
    }).then(res => {
      if (res.code !== 0) {
        message.error(res.msg);
        return false;
      }
      message.success(res.msg);
      history.goBack()
      return true;
    });
  };

  return (
    <Form
      {...formItemLayout}
      form={form}
      name="register"
      onFinish={onFinish}
      scrollToFirstError
      initialValues={initialValues}
    >
        
        
          <Form.Item
              name="id"
              label="ID"
            >
              <Input />
            </Form.Item>
        
        
        
          <Form.Item
              name="client"
              label="client"
            >
              <Input />
            </Form.Item>
        
        
        
          <Form.Item
              name="authorize"
              label="authorize"
            >
              <Input />
            </Form.Item>
        
        
        
          <Form.Item
              name="previous"
              label="previous"
            >
              <Input />
            </Form.Item>
        
        
        
          <Form.Item
              name="accessToken"
              label="access_token"
            >
              <Input />
            </Form.Item>
        
        
        
          <Form.Item
              name="refreshToken"
              label="refresh_token"
            >
              <Input />
            </Form.Item>
        
        
        
          <Form.Item
              name="expiresIn"
              label="expires_in"
            >
              <Input />
            </Form.Item>
        
        
        
          <Form.Item
              name="scope"
              label="scope"
            >
              <Input />
            </Form.Item>
        
        
        
          <Form.Item
              name="redirectUri"
              label="redirect_uri"
            >
              <Input />
            </Form.Item>
        
        
        
          <Form.Item
              name="extra"
              label="extra"
            >
              <Input />
            </Form.Item>
        
        
        
          <Form.Item
              name="ctime"
              label="创建时间"
            >
              <Input />
            </Form.Item>
        
        
      <Form.Item {...tailFormItemLayout}>
        {mode !== 'view' && <Button type="primary" htmlType="submit">
          提交
        </Button>}
        <Button
          onClick={()=>{
            history.goBack()
          }}
        >
          返回
        </Button>
      </Form.Item>
    </Form>
  );
};
