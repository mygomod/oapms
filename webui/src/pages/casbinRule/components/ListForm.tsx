// @BeeOverwrite YES
// @BeeGenerateTime 20200831_174645
import {Form, Input, Modal} from 'antd';
import React, { useEffect } from "react";
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
                  name="id"
                  label="ID"
                >
                  <Input />
                </Form.Item>
            
            
            
              <Form.Item
                  name="p"
                  label="策略、用户组"
                >
                  <Input />
                </Form.Item>
            
            
            
              <Form.Item
                  name="v0"
                  label="v0"
                >
                  <Input />
                </Form.Item>
            
            
            
              <Form.Item
                  name="v1"
                  label="v1"
                >
                  <Input />
                </Form.Item>
            
            
            
              <Form.Item
                  name="v2"
                  label="v2"
                >
                  <Input />
                </Form.Item>
            
            
            
              <Form.Item
                  name="v3"
                  label="v3"
                >
                  <Input />
                </Form.Item>
            
            
            
              <Form.Item
                  name="v4"
                  label="v4"
                >
                  <Input />
                </Form.Item>
            
            
            
              <Form.Item
                  name="v5"
                  label="v5"
                >
                  <Input />
                </Form.Item>
            
            
        </Form>
      </Modal>
  );
};
export default ListForm;
