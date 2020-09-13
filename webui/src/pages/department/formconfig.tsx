// @BeeOverwrite YES
// @BeeGenerateTime 20200820_200335
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
              name="name"
              label="名称"
            >
              <Input />
            </Form.Item>
        
        
        
          <Form.Item
              name="pid"
              label="上级部门id"
            >
              <Input />
            </Form.Item>
        
        
        
          <Form.Item
              name="orderNum"
              label="排序"
            >
              <Input />
            </Form.Item>
        
        
        
          <Form.Item
              name="extendField"
              label="扩展字段"
            >
              <Input />
            </Form.Item>
        
        
        
          <Form.Item
              name="intro"
              label="介绍"
            >
              <Input />
            </Form.Item>
        
        
        
          <Form.Item
              name="createdAt"
              label="创建时间"
            >
              <Input />
            </Form.Item>
        
        
        
          <Form.Item
              name="updatedAt"
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
