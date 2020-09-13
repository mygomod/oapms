import request from "@/utils/request";


export async function MenuTree(): Promise<any> {
  return request('/api/admin/menu/tree');
}
