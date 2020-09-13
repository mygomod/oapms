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
                  name="appId"
                  label="应用id"
                >
                  <Input />
                </Form.Item>
            
            
            
              <Form.Item
                  name="pid"
                  label="菜单id"
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
                  name="pmsCode"
                  label="权限标识"
                >
                  <Input />
                </Form.Item>
            
            
            
              <Form.Item
                  name="pmsRule"
                  label="数据规则"
                >
                  <Input />
                </Form.Item>
            
            
            
              <Form.Item
                  name="pmsType"
                  label="1=分类 2=数据权限"
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
                  name="intro"
                  label="说明"
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
            
            
        </Form>
      </Modal>
  );
};
export default ListForm;
