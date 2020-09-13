import { stringify } from 'querystring';
import { history, Reducer, Effect } from 'umi';

import { oauthLogin, oauthLogout } from '@/services/login';
import { setAuthority } from '@/utils/authority';
import { getPageQuery } from '@/utils/utils';
import {message} from "antd";

export interface StateType {
  status?: 'ok' | 'error';
  type?: string;
  currentAuthority?: 'user' | 'guest' | 'admin';
}

export interface LoginModelType {
  namespace: string;
  state: StateType;
  effects: {
    login: Effect;
    logout: Effect;
  };
  reducers: {
    changeLoginStatus: Reducer<StateType>;
  };
}

const Model: LoginModelType = {
  namespace: 'login',

  state: {
    status: undefined,
  },

  effects: {
    *login({ payload }, { call, put }) {
      const params = getPageQuery();
      const response = yield call(oauthLogin, {
        ...payload,
        params,
      });
      yield put({
        type: 'changeLoginStatus',
        payload: response,
      });
      // Login successfully
      if (response.code === 0) {
        const urlParams = new URL(window.location.href);
        const params = getPageQuery();
        let { redirect } = params as { redirect: string };
        if (redirect) {
          const redirectUrlParams = new URL(redirect);
          if (redirectUrlParams.origin === urlParams.origin) {
            redirect = redirect.substr(urlParams.origin.length);
            if (redirect.match(/^\/.*#/)) {
              redirect = redirect.substr(redirect.indexOf('#') + 1);
            }
          } else {
            window.location.href = '/';
            return;
          }
        }
        history.replace(redirect || '/');
      }else  if (response.code === 301) {
        window.location.href = response.data.redirect_uri
      }else {
        message.error(response.msg);
      }



    },

    *logout(_, { call, put }) {
      const response = yield call(oauthLogout);
      if (response.code === 0) {
        yield put({
          type: 'changeLoginStatus',
          payload: {
            status: false,
            data: {
              currentAuthority: 'guest',
            },
          },
        });
        const { redirect } = getPageQuery();
        // Note: There may be security issues, please note
        if (window.location.pathname !== '/user/login' && !redirect) {
          history.replace({
            pathname: '/user/login',
            search: stringify({
              redirect: window.location.href,
            }),
          });
        }
      }



    },
  },

  reducers: {
    changeLoginStatus(state, { payload }) {
      setAuthority(payload.currentAuthority);
      return {
        ...state,
        status: payload.status,
        type: payload.type,
      };
    },
  },
};

export default Model;
