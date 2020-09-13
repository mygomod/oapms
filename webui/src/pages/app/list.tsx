// @BeeOverwrite NO
// @BeeGenerateTime 20200820_230055
import { Card, message, Tag, Button, Divider, Modal } from 'antd';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import request from "@/utils/request";
import ProTable, { ProColumns, ActionType } from '@ant-design/pro-table';
import React, { useState, useRef,Fragment } from 'react';
import ListForm from "./components/ListForm"
import { PlusOutlined } from '@ant-design/icons';

const apiUrl = "/api/admin/app";
// @ts-ignore
const reqList = (params) => {
   return request(apiUrl+"/list", {
        params,
   });
}

const handleCreate = async (values: any) => {
  const hide = message.loading('正在添加');
  try {
    const resp = await request("/api/admin/app/create" , {
      method: 'POST',
      data: {
        ...values,
        "state":parseInt(values.state),
      },
    })
    if (resp.code !== 0) {
      hide();
      message.error('添加失败，错误信息：'+resp.msg);
      return true
    }
    hide();
    message.success('添加成功');
    return true;
  } catch (error) {
    hide();
    message.error('添加失败请重试！');
    return false;
  }
};

const handleUpdate = async (values: any) => {
  const hide = message.loading('正在更新');
  try {
    const resp = await request("/api/admin/app/update" , {
      method: 'POST',
      data: {
        ...values,
        "state":parseInt(values.state),
      },
    })
    if (resp.code !== 0) {
      hide();
      message.error('更新失败，错误信息：'+resp.msg);
      return true
    }
    hide();
    message.success('更新成功');
    return true;
  } catch (error) {
    hide();
    message.error('更新失败请重试！');
    return false;
  }
};


const TableList: React.FC<{}> = () => {
  const actionRef = useRef<ActionType>();
  const [createModalVisible, handleCreateModalVisible] = useState<boolean>(false);
  const [updateModalVisible, handleUpdateModalVisible] = useState<boolean>(false);
  const [updateValues, setUpdateValues] = useState({});

    const columns = [
        {
            title: "应用id",
            dataIndex: "aid",
            key: "aid",
        },
        {
          title: "名称",
          dataIndex: "name",
          key: "name",
        },
        {
            title: "客户端",
            dataIndex: "clientId",
            key: "clientId",
        },

        {
            title: "秘钥",
            dataIndex: "secret",
            key: "secret",
        },
        {
            title: "跳转地址",
            dataIndex: "redirectUri",
            key: "redirectUri",
        },
        {
          title: "访问地址",
          dataIndex: "url",
          key: "url",
        },
        {
            title: "额外信息",
            dataIndex: "extra",
            key: "extra",
        },
        {
            title: "调用次数",
            dataIndex: "callNo",
            key: "callNo",
        },
        {
            title: "状态",
            dataIndex: "state",
            key: "state",
            valueEnum: {
              0: { text: '停用', status: 'error' },
              1: { text: '启用', status: 'processing' },
            },
        },
        {
            title: "更新时间",
            dataIndex: "utime",
            key: "utime",
            hideInSearch:true,
        },

      {
        title: '操作',
        dataIndex: 'operating',
        key: 'operating',
        valueType:"option",
        render: (value: any, record: React.SetStateAction<{}>) => (
          <Fragment>
            <a
              onClick={() => {
                handleUpdateModalVisible(true);
                setUpdateValues(record);
              }}
            >
              编辑
            </a>
            <Divider type="vertical" />
            <a
              onClick={() => {
                Modal.confirm({
                  title: '确认删除？',
                  okText: '确认',
                  cancelText: '取消',
                  onOk: () => {
                    request(apiUrl+"/delete", {
                       method: 'POST',
                       data: {
                          aid: record.aid,
                       },
                     }).then((res) => {
                      if (res.code !== 0) {
                        message.error(res.msg);
                        return false;
                      }
                      actionRef.current?.reloadAndRest();
                      return true;
                    });
                  },
                });
              }}
            >
              删除
            </a>
          </Fragment>
        ),
      },
    ];
    return (
      <PageHeaderWrapper>
        <Card>
          <ProTable
            actionRef={actionRef}
            request={(params, sorter, filter) => reqList({ ...params, sorter, filter })}
            columns={columns}
            rowKey={(record: { aid: any; }) => record.aid}
            toolBarRender={action => [
              <Button type="primary" onClick={() => handleCreateModalVisible(true)}>
                <PlusOutlined /> 新建
              </Button>,
              <Button
                onClick={() => {
                  action.resetPageIndex();
                }}
                type="default"
                style={ {
                  marginRight: 8,
                } }
              >
                回到第一页
              </Button>,
            ]}
          />
        </Card>
        <ListForm
          formTitle={"创建权限"}
          onSubmit={async (value: any) => {
            const success = handleCreate(value);
            if (success) {
              handleCreateModalVisible(false);
              if (actionRef.current) {
                actionRef.current.reload();
              }
            }
          }}
          onCancel={() => handleCreateModalVisible(false)}
          modalVisible={createModalVisible}
        />
        <ListForm
          formTitle={"编辑权限"}
          onSubmit={async (value: any) => {
            const success = await handleUpdate(value);
            if (success) {
              handleUpdateModalVisible(false);
              setUpdateValues({});
              if (actionRef.current) {
                actionRef.current.reload();
              }
            }
          }}
          onCancel={() => {
            setUpdateValues({})
            handleUpdateModalVisible(false)
          }}
          modalVisible={updateModalVisible}
          initialValues={updateValues}
        />
      </PageHeaderWrapper>
    );
}
export default TableList;
