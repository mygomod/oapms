import React, { useEffect } from 'react';
import { Form, Input, Modal, Select} from 'antd';
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

  console.log("initialValues",initialValues)

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
        name="id"
        label="id"
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
          name="pid"
          label="上级菜单id"
        >
          <Input />
        </Form.Item>
        <Form.Item
          name="name"
          label="菜单名称"
        >
          <Input />
        </Form.Item>
        <Form.Item
          name="path"
          label="路由"
        >
          <Input />
        </Form.Item>
        <Form.Item
          name="pmsCode"
          label="标识"
        >
          <Input />
        </Form.Item>
        <Form.Item
          name="menuType"
          label="类型"
        >
          <Input />
        </Form.Item>
        <Form.Item
          name="icon"
          label="图标"
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
          name="state"
          label="状态"
        >
          <Input />
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default ListForm;
