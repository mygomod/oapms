import { Settings as ProSettings } from '@ant-design/pro-layout';

type DefaultSettings = ProSettings & {
  pwa: boolean;
};

const proSettings: DefaultSettings = {
  "navTheme": "light",
  "primaryColor": "#1890ff",
  "layout": "mix",
  "contentWidth": "Fluid",
  "fixedHeader": false,
  "fixSiderbar": true,
  "menu": {
    "locale": true
  },
  "title": "用户权限系统",
  "pwa": false,
  "iconfontUrl": "",
  "splitMenus": true
};

export type { DefaultSettings };

export default proSettings;
