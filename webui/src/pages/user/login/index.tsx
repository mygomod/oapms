import { Alert } from 'antd';
import React, { useState } from 'react';
import { Link, connect, Dispatch } from 'umi';
import { StateType } from '@/models/login';
import { LoginParamsType } from '@/services/login';
import { ConnectState } from '@/models/connect';
import LoginForm from './components/Login';

import styles from './style.less';

const { Tab, UserName, Password, Mfa, Submit } = LoginForm;
interface LoginProps {
  dispatch: Dispatch;
  userLogin: StateType;
  submitting?: boolean;
}

const LoginMessage: React.FC<{
  content: string;
}> = ({ content }) => (
  <Alert
    style={{
      marginBottom: 24,
    }}
    message={content}
    type="error"
    showIcon
  />
);



class LoginComponent extends React.Component {
  // componentWillMount() {   //初始化时在页面加载完成前执行
  //   this.appendMeta()
  // }
  //
  // componentWillReceiveProps(){  //刷新页面时执行
  //   this.appendMeta()
  // }
  //
  // appendMeta = () => {
  //   //在head标签插入meta标签，解决在生产环境链接失效问题
  //   // const metaTag = document.getElementsByTagName('meta');
  //   let isHasTag = true;
  //   // for (let i = 0; i < metaTag.length; i++) {   //避免重复插入meta标签
  //   //   const name = metaTag[i].name;
  //   //   console.log("name",name)
  //   //
  //   //   if (name == 'referrer') {
  //   //     isHasTag = false;
  //   //   }
  //   // }
  //   //
  //   // console.log("append meta",isHasTag)
  //   // isHasTag = true
  //   if (isHasTag) {
  //
  //     const headItem = document.head;
  //
  //     let oMeta = document.createElement('meta');
  //
  //     oMeta.setAttribute('name', 'referrer');
  //
  //     oMeta.content = 'origin';
  //
  //     headItem.appendChild(oMeta)
  //
  //   }
  // }


  render() {
    return (
      <Login submitting={this.props.submitting} userLogin={this.props.userLogin} dispatch={this.props.dispatch}/>
    )
  }
}




const Login: React.FC<LoginProps> = (props) => {
  const { userLogin = {}, submitting } = props;
  const { status, type: loginType } = userLogin;
  const [type, setType] = useState<string>('account');


  const handleSubmit = (values: LoginParamsType) => {
    console.log("values",values)
    const { dispatch } = props;
    dispatch({
      type: 'login/login',
      payload: { ...values, type },
    });
  };
  return (
    <div className={styles.main}>
      <LoginForm activeKey={type} onTabChange={setType} onSubmit={handleSubmit}>
        <Tab key="account" tab="登录">
          {status === 'error' && loginType === 'account' && !submitting && (
            <LoginMessage content="账户或密码错误（admin/ant.design）" />
          )}

          <UserName
            name="nickname"
            placeholder="用户名: "
            rules={[
              {
                required: true,
                message: '请输入用户名!',
              },
            ]}
          />
          <Password
            name="password"
            placeholder="密码: "
            rules={[
              {
                required: true,
                message: '请输入密码！',
              },
            ]}
          />
          <Mfa
            name="mfa"
            placeholder="密保: "
          />
        </Tab>
        <Submit loading={submitting}>登录</Submit>
      </LoginForm>
    </div>
  );
};

export default connect(({ login, loading }: ConnectState) => ({
  userLogin: login,
  submitting: loading.effects['login/login'],
}))(LoginComponent);
