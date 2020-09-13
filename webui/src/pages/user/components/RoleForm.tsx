import {Form, Input, Modal, Checkbox,Row } from 'antd';
import React, { useEffect } from "react";

interface RoleFormProps {
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



const RoleForm: React.FC<RoleFormProps> = (props) => {
  const { modalVisible, onCancel, onSubmit, initialValues, formTitle } = props;
  const [form] = Form.useForm();


  useEffect(() => {
    if (form && !modalVisible) {
      form.resetFields();
    }
  }, [modalVisible]);

  useEffect(() => {
    if (initialValues) {
      form.setFieldsValue({
        ...initialValues,
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
      {...modalFooter}
    >
      <Form
        {...formLayout}
        form={form}
        onFinish={onSubmit}
        scrollToFirstError
      >
        <Form.Item
          name="uid"
          label="uid"
          hidden
        >
          <Input />
        </Form.Item>
        <Form.Item
          name="aid"
          label="aid"
          hidden
        >
          <Input />
        </Form.Item>
        <Form.Item
          name="roleIds"
          label="roleIds"
        >
        <Checkbox.Group>
          {(initialValues.roleAll||[]).map(item => {
            return <Row>
              <Checkbox value={item.id}>{item.name}</Checkbox>
            </Row>
          })}
        </Checkbox.Group>
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default RoleForm;
