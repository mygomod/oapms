import request from "@/utils/request";

export async function AccountAppList(): Promise<any> {
  return request('/api/admin/account/app');
}


export async function AccountMenus(): Promise<any> {
  return request('/api/admin/account/menu');
}


export async function RoleAll(aid: string): Promise<any> {
  return request(`/api/admin/role/all?aid=${aid}`);
}



export async function UserGetRole(aid: any, uid: any) {
  return request(`/api/admin/user/getRole?aid=${aid}&uid=${uid}`);
}
