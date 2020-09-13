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
            
            
        </Form>
      </Modal>
  );
};
export default ListForm;
