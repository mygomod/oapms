// https://umijs.org/config/
import { defineConfig } from 'umi';
import defaultSettings from './defaultSettings';
import proxy from './proxy';

const { REACT_APP_ENV } = process.env;

export default defineConfig({
  hash: true,
  antd: {},
  dva: {
    hmr: true,
  },
  locale: {
    // default zh-CN
    default: 'zh-CN',
    antd: true,
    // default true, when it is true, will use `navigator.language` overwrite default
    baseNavigator: true,
  },
  dynamicImport: {
    loading: '@/components/PageLoading/index',
  },
  targets: {
    ie: 11,
  },
  // umi routes: https://umijs.org/docs/routing
  routes: [
    {
      path: '/user',
      component: '../layouts/UserLayout',
      routes: [
        {
          name: 'login',
          path: '/user/login',
          component: './user/login'
        }
      ]
    },
    {
      path: '/',
      component: '../layouts/SecurityLayout',
      routes: [
        {
          path: '/',
          component: '../layouts/BasicLayout',
          authority: ['admin', 'user'],
          routes: [
            {
              path: '/',
              redirect: '/account'
            },
            {
              path: '/account',
              name: '个人',
              icon: 'smile',
              routes:[
                { path:"/account", redirect:"/account/welcome"},
                { name:"欢迎", icon:"RocketOutlined", path:"/account/welcome",component: './workspace/index',}
              ]

            },
            {
              path: '/admin',
              name: 'admin',
              icon: 'crown',
              component: './Admin',
              authority: ['admin'],
              routes: [
                {
                  path: '/admin/sub-page',
                  name: 'sub-page',
                  icon: 'smile',
                  component: './Welcome',
                  authority: ['admin']
                }
              ]
            },
            {
              path: '/oa',
              name: '授权',
              icon: 'crown',
              routes:[
                { path:"/oa", redirect:"/oa/app"},
                { name:"应用", icon:"RocketOutlined", path:"/oa/app", component: "./app/list"},
                { name:"用户", icon:"RocketOutlined", path:"/oa/user", component: "./user/list"},
              ]
            },
            {
              path: '/pms',
              name: '权限',
              icon: 'crown',
              routes:[
                { path:"/pms", redirect:"/pms/menu"},
                { name:"菜单", icon:"RocketOutlined", path:"/pms/menu", component: "./menu/list"},
                { name:"角色", icon:"RocketOutlined", path:"/pms/role", component: "./role/list"}
              ]
            },
            {
              component: './404'
            }
          ]
        },
        {
          component: './404'
        }
      ]
    },
    {
      component: './404'
    }
  ],
  // Theme for antd: https://ant.design/docs/react/customize-theme-cn
  theme: {
    // ...darkTheme,
    'primary-color': defaultSettings.primaryColor,
  },
  // @ts-ignore
  title: false,
  ignoreMomentLocale: true,
  proxy: proxy[REACT_APP_ENV || 'dev'],
  manifest: {
    basePath: '/',
  },
});
