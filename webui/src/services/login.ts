import request from '@/utils/request';

export interface LoginParamsType {
  nickname: string;
  password: string;
  mfa: string;
}

export async function oauthLogin(params) {
  return request('/api/admin/oauth/login', {
    method: 'POST',
    data: params,
  });
}

export async function oauthLogout() {
  return request('/api/admin/oauth/logout');
}


export async function getFakeCaptcha(mobile: string) {
  return request(`/api/login/captcha?mobile=${mobile}`);
}
