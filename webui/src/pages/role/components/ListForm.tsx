import React, {  useEffect } from 'react';
import { Form, Input,  Modal, Select} from 'antd';
import RoleMenu from "./RoleMenu"
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
  const { modalVisible, onCancel, onSubmit, initialValues, formTitle, selectData } = props;
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
      width={1000}
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
        name="name"
        label="角色名称"
      >
        <Input />
      </Form.Item>
        <Form.Item
          name="appId"
          label="应用id"
        >
          <Select
            allowClear
            showSearch
            optionFilterProp="children"
            placeholder="请选择"
            style={{ width: 200 }}
            disabled={false}
          >
            {
              (selectData || []).map((item,index)=>{
                return <Option key={index} value={item.key}>{item.title}</Option>
              })
            }
          </Select>
        </Form.Item>
        <Form.Item
          name="menus"
          label="选择菜单权限"
        >
          <RoleMenu />
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default ListForm;
