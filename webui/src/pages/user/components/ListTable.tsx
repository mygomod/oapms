// @BeeOverwrite NO
// @BeeGenerateTime 20200831_174645
import {Card, message, Tag, Button, Divider, Modal, Select} from 'antd';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import request from "@/utils/request";
import ProTable, { ProColumns, ActionType } from '@ant-design/pro-table';
import React, {useState, useRef, Fragment, useEffect} from 'react';
import ListForm from "./ListForm"
import RoleForm from "./RoleForm"
import { PlusOutlined } from '@ant-design/icons';
import moment from "moment";
import {RoleAll, UserGetRole} from "@/services/api";
import {connect, history} from "umi";

const apiUrl = "/api/admin/user";

const reqList = (params) => {
   return request(apiUrl+"/list", {
        params,
   });
}


const handleCreate = async (values) => {
  const hide = message.loading('正在添加');
  try {
    const resp = await request(apiUrl+"/create" , {
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
    const resp = await request(apiUrl+"/update" , {
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

const handleRole = async (values) => {
  const hide = message.loading('正在更新');
  try {
    const resp = await request(apiUrl+"/setRole" , {
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
  const actionRef = useRef<ActionType>();
  const [createModalVisible, handleCreateModalVisible] = useState<boolean>(false);
  const [updateModalVisible, handleUpdateModalVisible] = useState<boolean>(false);
  const [roleModalVisible, handleRoleModalVisible] = useState<boolean>(false);
  const [roleValues, setRoleValues] = useState({});
  const [updateValues, setUpdateValues] = useState({});

  const { appArr, dispatch, setAid, aid } = props;
  /**
   * constructor
   */

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


  const handleChange = (value) => {
    setAid(value)
    history.push({
      pathname: "/oa/user",
      search: `?aid=${value}`,
    });
  };

  const columns = [
        {
            title: "uid",
            dataIndex: "uid",
            key: "uid",
        },
        {
            title: "昵称",
            dataIndex: "nickname",
            key: "nickname",
        },
        {
            title: "邮箱",
            dataIndex: "email",
            key: "email",
        },
        {
            title: "头像",
            dataIndex: "avatar",
            key: "avatar",
            hideInSearch: true,
        },
        {
            title: "状态",
            dataIndex: "state",
            key: "state",
            valueEnum: {
              0: { text: '停用', status: 'Default' },
              1: { text: '启用', status: 'Processing' },
            },
        },
        {
            title: "性别",
            dataIndex: "gender",
            key: "gender",
            hideInSearch: true,
            valueEnum: {
              0: { text: '未知'},
              1: { text: '男'},
              2: { text: '女'},
            },
        },
        {
            title: "上次登录ip",
            dataIndex: "lastLoginIp",
            key: "lastLoginIp",
            hideInSearch: true,
        },
        {
            title: "上次登录时间",
            dataIndex: "lastLoginTime",
            key: "lastLoginTime",
            hideInSearch: true,
            render(val) {
              return moment(val,"X").format('YYYY-MM-DD HH:mm:ss')
            },
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
                handleUpdateModalVisible(true);
                setUpdateValues(record);
              }}
            >
              编辑
            </a>
            <Divider type="vertical" />
            <a
              onClick={() => {
                Promise.all([RoleAll(aid),UserGetRole(aid,record.uid)]).then((resultList) => {
                  handleRoleModalVisible(true);
                  setRoleValues({
                    aid: aid,
                    uid: record.uid,
                    roleAll: resultList[0].data,
                    roleIds: resultList[1].data
                  });
                })
              }}
            >
              设置角色
            </a>
            <Divider type="vertical" />
            <a
              onClick={() => {
                Modal.confirm({
                  title: '确认发送邮件google验证码？',
                  okText: '确认',
                  cancelText: '取消',
                  onOk: () => {
                    request(apiUrl+"/sendGoogleCode", {
                      method: 'POST',
                      data: {
                        uid: record.uid,
                      },
                    }).then((res) => {
                      if (res.code !== 0) {
                        message.error(res.msg);
                        return false;
                      }
                      return true;
                    });
                  },
                });
              }}
            >
              发送google验证码
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
                          uid: record.uid,
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
            rowKey={(record) => record.uid}
            toolBarRender={action => [
                <Select
                  showSearch
                  optionFilterProp="children"
                  placeholder="请选择"
                  style={{ width: 200 }}
                  disabled={false}
                  onChange={handleChange}
                  value={aid}
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
        <ListForm
          formTitle={"创建"}
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
          formTitle={"编辑"}
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
        <RoleForm
          formTitle={"设置角色"}
          onSubmit={async (value) => {
            const success = await handleRole(value);
            if (success) {
              handleRoleModalVisible(false);
              setRoleValues({})
            }
          }}
          initialValues={roleValues}
          onCancel={() => {
            handleRoleModalVisible(false)
            setRoleValues({})
          }}
          modalVisible={roleModalVisible}
        />
      </PageHeaderWrapper>
    );
}
export default connect(({ global }) => ({
  appArr: global.appArr,
}))(ListTable);
