import request from '@/utils/request';

export async function AppSelect(): Promise<any> {
  return request('/api/admin/app/select');
}

export async function AppSelectArr(): Promise<any> {
  return request('/api/admin/app/selectArr');
}
