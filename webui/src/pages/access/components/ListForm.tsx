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
            
            
        </Form>
      </Modal>
  );
};
export default ListForm;
