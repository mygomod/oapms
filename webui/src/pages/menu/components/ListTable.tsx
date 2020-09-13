// @BeeOverwrite NO
// @BeeGenerateTime 20200820_230345
import {Card, message, Tag, Button, Divider, Modal, Layout, Tree, Select, Form} from 'antd';
import {connect, history} from 'umi';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import request from "@/utils/request";
import ProTable, { ProColumns, ActionType } from '@ant-design/pro-table';
import React, {useState, useRef, Fragment, useEffect} from 'react';
import ListForm from "./ListForm"
import { PlusOutlined } from '@ant-design/icons';

const apiUrl = "/api/admin/menu";

const reqList = (params) => {
  return request(apiUrl+"/list", {
    params,
  });
}

const reqTree = (aid: string) => {
  return request("/api/admin/menu/tree?aid="+aid)
}

const handleCreate = async (values) => {
  const hide = message.loading('正在添加');
  try {
    const resp = await request("/api/admin/menu/create" , {
      method: 'POST',
      data: {
        ...values,
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

const handleUpdate = async (values) => {
  const hide = message.loading('正在更新');
  try {
    const resp = await request("/api/admin/menu/update" , {
      method: 'POST',
      data: {
        ...values,
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

const ListTable: React.FC<{}> = (props) => {
  const [createModalVisible, handleCreateModalVisible] = useState<boolean>(false);
  const [updateModalVisible, handleUpdateModalVisible] = useState<boolean>(false);
  const [updateValues, setUpdateValues] = useState({});

  const actionRef = useRef<ActionType>();
  const treeData = props.treeData;
  const setTreeData = props.setTreeData;
  const setPid = props.setPid;
  const setAppId = props.setAppId;
  const params = props.params

  const { appArr, dispatch } = props;

  /**
   * constructor
   */

  useEffect(() => {
    if (dispatch) {
      dispatch({
        type: 'global/fetchSelectAppArr'
      })
    }
  }, []);

  const columns = [
    {
      title: "id",
      dataIndex: "id",
      key: "id",
      hideInSearch:true,
      hideInForm:true,
    },

    {
      title: "应用名称",
      dataIndex: "appId",
      key: "appId",
      hideInSearch:true,
    },
    {
      title: "菜单名称",
      dataIndex: "name",
      key: "name",
    },
    {
      title: "路由",
      dataIndex: "path",
      key: "path",
    },
    {
      title: "标识",
      dataIndex: "pmsCode",
      key: "pmsCode",
    },
    {
      title: "上级菜单id",
      dataIndex: "pid",
      key: "pid",
      hideInSearch:true,
    },
    {
      title: "类型",
      dataIndex: "menuType",
      key: "menuType",
      hideInSearch:true,
      valueEnum: {
        1:{"text":"菜单"},
        2:{"text":"按钮"},
      }
    },
    {
      title: "状态",
      dataIndex: "state",
      key: "state",
      hideInSearch:true,
      valueEnum: {
        0: { text: '停用', status: 'error' },
        1: { text: '启用', status: 'processing' },
      },
    },
    {
      title: "图标",
      dataIndex: "icon",
      key: "icon",
      hideInSearch:true,
    },
    {
      title: "排序",
      dataIndex: "orderNum",
      key: "orderNum",
      hideInSearch:true,
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
      render: (value, record) => (
        <Fragment>
          <a
            onClick={() => {
              console.log("fuck",{
                aid: params.aid,
                ...record,
              })
              handleUpdateModalVisible(true);
              setUpdateValues({
                aid: params.appId,
                ...record,
              });
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
                      id: record.id,
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

  const convertMenuTreeToTreeData = (treeData) =>  {
    return treeData.map((item) => {
      return {
        title: item.name || '',
        key: item.id || '',
        children: convertMenuTreeToTreeData(item.children || [])
      }
    })
  }

  const handleChange = (value) => {
    setAppId(value)
    history.push({
      pathname: "/pms/menu",
      search: `?aid=${value}`,
    });
  }

  return (
    <PageHeaderWrapper>
      <Layout>
        <Layout.Sider
          width={280}
          style={{
            background: '#fff',
            borderRight: '1px solid lightGray',
            padding: 15,
            overflow: 'auto',
          }}
        >
          <Tree
            onSelect={(keys) => {
              setPid(keys[0])
            }}
            treeData={convertMenuTreeToTreeData(treeData)}
          >
          </Tree>
        </Layout.Sider>
        <Layout.Content>
          <Card bordered={false}>
            <ProTable
              params={params}
              onReset={()=>{
                setPid(null)
              }}
              actionRef={actionRef}
              request={(params, sorter, filter) => {
                reqTree(params.appId).then((response)=>{
                  if (response.code == 0) {
                    setTreeData(response.data)
                  }else {
                    setTreeData([])
                  }
                })
                return reqList({ ...params, sorter, filter })
              }}
              columns={columns}
              rowKey={(record) => record.id}
              toolBarRender={action => [
                <Select
                  showSearch
                  optionFilterProp="children"
                  placeholder="请选择"
                  style={{ width: 200 }}
                  disabled={false}
                  onChange={handleChange}
                  value={params.appId}
                >
                  {
                    (appArr || []).map((item,index)=>{
                      return <Option key={index} value={item.key}>{item.title}</Option>
                    })
                  }
                </Select>,
                <Button type="primary" onClick={() => handleCreateModalVisible(true)}>
                  <PlusOutlined /> 新建
                </Button>,
              ]}
            />
          </Card>
        </Layout.Content>
      </Layout>
      <ListForm
        formTitle={"创建权限"}
        onSubmit={async (value) => {
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
        onSubmit={async (value) => {
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

export default connect(({ global }) => ({
  appArr: global.appArr,
}))(ListTable);
