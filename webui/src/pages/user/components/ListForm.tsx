// @BeeOverwrite YES
// @BeeGenerateTime 20200831_174645
import {Form, Input, Modal, Radio} from 'antd';
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
                  name="uid"
                  label="uid"
                  hidden
                >
                  <Input />
              </Form.Item>
              <Form.Item
                  name="nickname"
                  label="昵称"
                >
                  <Input />
              </Form.Item>
              <Form.Item
                  name="email"
                  label="邮箱"
                >
                  <Input />
              </Form.Item>
              <Form.Item
                  name="password"
                  label="密码"
                >
                  <Input />
              </Form.Item>
              <Form.Item
                name="state"
                label="状态"
              >
                <Radio.Group>
                  <Radio value={1}>开启</Radio>
                  <Radio value={0}>停用</Radio>
                </Radio.Group>
              </Form.Item>
        </Form>
      </Modal>
  );
};
export default ListForm;
