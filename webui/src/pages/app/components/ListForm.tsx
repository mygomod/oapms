import React, { useEffect } from 'react';
import {Form, Input, Modal, Radio } from 'antd';
interface ListFormProps {
  modalVisible: boolean;
  formTitle: string;
  initialValues: {};
  onSubmit: () => void;
  onCancel: () => void;
}
const formLayout = {
  labelCol: { span: 7 },
  wrapperCol: { span: 13 },
};

const ListForm: React.FC<ListFormProps> = (props) => {
  const { modalVisible, onCancel, onSubmit, initialValues, formTitle } = props;
  const [form] = Form.useForm();


  useEffect(() => {
    if (form && !modalVisible) {
      form.resetFields();
    }
  }, [modalVisible]);

  useEffect(() => {
    if (initialValues) {
      let state = "0"
      if (initialValues.state == 1) {
        state = "1"
      }
      form.setFieldsValue({
        ...initialValues,
        "state":state,
      });
    }
  }, [initialValues]);

  const handleSubmit = () => {
    if (!form) return;
    form.submit();
  };

  const modalFooter = { okText: '保存', onOk: handleSubmit, onCancel:onCancel }

  return (
    <Modal
      destroyOnClose
      title={formTitle}
      visible={modalVisible}
      onCancel={() => onCancel()}
      {...modalFooter}
    >
        <Form
          {...formLayout}
          form={form}
          onFinish={onSubmit}
          scrollToFirstError
        >
        <Form.Item
          name="aid"
          label="aid"
          hidden
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
          name="redirectUri"
          label="跳转地址"
        >
          <Input />
        </Form.Item>
        <Form.Item
          name="state"
          label="状态"
        >
          <Radio.Group>
            <Radio value="1">开启</Radio>
            <Radio value="0">停用</Radio>
          </Radio.Group>
        </Form.Item>
        <Form.Item
          name="extra"
          label="额外信息"
        >
          <Input />
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default ListForm;
