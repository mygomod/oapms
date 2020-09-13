// @BeeOverwrite YES
// @BeeGenerateTime 20200831_101837
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
              label="id"
            >
              <Input />
            </Form.Item>
        
        
        
          <Form.Item
              name="uid"
              label="uid"
            >
              <Input />
            </Form.Item>
        
        
        
          <Form.Item
              name="secret"
              label="秘钥"
            >
              <Input />
            </Form.Item>
        
        
        
          <Form.Item
              name="isBind"
              label="是否绑定"
            >
              <Input />
            </Form.Item>
        
        
        
          <Form.Item
              name="ctime"
              label="创建时间"
            >
              <Input />
            </Form.Item>
        
        
        
          <Form.Item
              name="utime"
              label="更新时间"
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
