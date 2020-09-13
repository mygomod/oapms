/**
 * Ant Design Pro v4 use `@ant-design/pro-layout` to handle Layout.
 * You can view component api by:
 * https://github.com/ant-design/ant-design-pro-layout
 */
import ProLayout, {
  MenuDataItem,
  BasicLayoutProps as ProLayoutProps,
  Settings,
  DefaultFooter,
} from '@ant-design/pro-layout';
import React, { useEffect } from 'react';
import { Link, useIntl, connect, Dispatch, history } from 'umi';
import { GithubOutlined } from '@ant-design/icons';
import { Result, Button } from 'antd';
import Authorized from '@/utils/Authorized';
import RightContent from '@/components/GlobalHeader/RightContent';
import { ConnectState } from '@/models/connect';
import { getAuthorityFromRouter } from '@/utils/utils';
import logo from '../assets/logo.svg';

const noMatch = (
  <Result
    status={403}
    title="403"
    subTitle="Sorry, you are not authorized to access this page."
    extra={
      <Button type="primary">
        <Link to="/user/login">Go Login</Link>
      </Button>
    }
  />
);
export interface BasicLayoutProps extends ProLayoutProps {
  breadcrumbNameMap: {
    [path: string]: MenuDataItem;
  };
  route: ProLayoutProps['route'] & {
    authority: string[];
  };
  settings: Settings;
  dispatch: Dispatch;
}
export type BasicLayoutContext = { [K in 'location']: BasicLayoutProps[K] } & {
  breadcrumbNameMap: {
    [path: string]: MenuDataItem;
  };
};
/**
 * use Authorized check all menu item
 */
//
// const menuDataRender = (menuList: MenuDataItem[]): MenuDataItem[] =>
//   menuList.map((item) => {
//     console.log("menuList",menuList)
//     const localItem = {
//       ...item,
//       children: item.children ? menuDataRender(item.children) : undefined,
//     };
//     console.log("localItem",localItem)
//
//     return Authorized.check(item.authority, localItem, null) as MenuDataItem;
//   });

const menuDataRender = (menuList: MenuDataItem[]): MenuDataItem[] => {
    menuList = [{"id":1,"pid":0,"appId":1,"name":"权限","path":"/pms","key":"/pms","pmsCode":"","pmsExtend":"","menuType":1,"icon":"","orderNum":0,"state":1,"ctime":0,"utime":0,"actions":null,"children":[{"id":2,"pid":1,"appId":1,"name":"菜单","path":"/pms/menu","key":"/pms/menu","pmsCode":"","pmsExtend":"","menuType":1,"icon":"","orderNum":0,"state":1,"ctime":0,"utime":0,"actions":null,"children":null,"pro_layout_parentKeys":["/pms"]},{"id":5,"pid":1,"appId":1,"name":"角色","path":"/pms/role","key":"/pms/role","pmsCode":"","pmsExtend":"","menuType":1,"icon":"","orderNum":0,"state":1,"ctime":0,"utime":0,"actions":null,"children":null,"pro_layout_parentKeys":["/pms"]}],"pro_layout_parentKeys":null}]
    return menuList
}


const defaultFooterDom = (
  <DefaultFooter
    copyright={`${new Date().getFullYear()} 蚂蚁金服体验技术部出品`}
    links={[
      {
        key: 'Ant Design Pro',
        title: 'Ant Design Pro',
        href: 'https://pro.ant.design',
        blankTarget: true,
      },
      {
        key: 'github',
        title: <GithubOutlined />,
        href: 'https://github.com/ant-design/ant-design-pro',
        blankTarget: true,
      },
      {
        key: 'Ant Design',
        title: 'Ant Design',
        href: 'https://ant.design',
        blankTarget: true,
      },
    ]}
  />
);

const BasicLayout: React.FC<BasicLayoutProps> = (props) => {
  const {
    dispatch,
    children,
    settings,
    location = {
      pathname: '/',
    },
    menus,
  } = props;
  /**
   * constructor
   */

  useEffect(() => {
    if (dispatch) {
      dispatch({
        type: 'user/fetchCurrent',
      });

      dispatch({
        type: 'global/fetchMenus'
      })
    }
  }, []);
  /**
   * init variables
   */

  const handleMenuCollapse = (payload: boolean): void => {
    if (dispatch) {
      dispatch({
        type: 'global/changeLayoutCollapsed',
        payload,
      });
    }
  }; // get children authority

  const authorized = getAuthorityFromRouter(props.route.routes, location.pathname || '/') || {
    authority: undefined,
  };
  const { formatMessage } = useIntl();

  console.log("menus",menus)
  return (
    <div> { menus && <ProLayout
      logo={logo}
      formatMessage={formatMessage}
      onCollapse={handleMenuCollapse}
      onMenuHeaderClick={() => history.push('/')}
      menuItemRender={(menuItemProps, defaultDom) => {
        if (menuItemProps.isUrl || !menuItemProps.path) {
          return defaultDom;
        }
        return <Link to={menuItemProps.path}>{defaultDom}</Link>;
      }}
      breadcrumbRender={(routers = []) => [
        {
          path: '/',
          breadcrumbName: formatMessage({ id: 'menu.home' }),
        },
        ...routers,
      ]}
      itemRender={(route, params, routes, paths) => {
        const first = routes.indexOf(route) === 0;
        return first ? (
          <Link to={paths.join('/')}>{route.breadcrumbName}</Link>
        ) : (
          <span>{route.breadcrumbName}</span>
        );
      }}
      footerRender={() => defaultFooterDom}
      menuDataRender={() => menus}
      // menuDataRender={menuDataRender}
      rightContentRender={() => <RightContent />}
      {...props}
      {...settings}
    >
      <Authorized authority={authorized!.authority} noMatch={noMatch}>
        {children}
      </Authorized>
    </ProLayout>}
    </div>
  );
};

export default connect(({ global, settings }: ConnectState) => ({
  collapsed: global.collapsed,
  menus: global.menus,
  settings,
}))(BasicLayout);
