/* eslint-disable */
// @ts-ignore
import * as API from './types';
import request from '@/utils/http';

/** 获取面板列表 获取面板列表 GET /panels */
export async function getPanels(options?: { [key: string]: unknown }) {
  return request<API.Panel[]>('/panels', {
    method: 'GET',
    ...(options || {}),
  });
}

/** 更新面板 更新面板 PUT /panels */
export async function putPanels(
  body: API.UpdatePanelInput,
  options?: { [key: string]: unknown }
) {
  return request<API.Panel>('/panels', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 创建面板 创建面板 POST /panels */
export async function postPanels(
  body: API.CreatePanelInput,
  options?: { [key: string]: unknown }
) {
  return request<unknown>('/panels', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** 获取面板 获取面板 GET /panels/${param0} */
export async function getPanelsId(
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.getPanelsIdParams,
  options?: { [key: string]: unknown }
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Panel>(`/panels/${param0}`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** 删除面板 删除面板 DELETE /panels/${param0} */
export async function deletePanelsId(
  // 叠加生成的Param类型 (非body参数openapi默认没有生成对象)
  params: API.deletePanelsIdParams,
  options?: { [key: string]: unknown }
) {
  const { id: param0, ...queryParams } = params;
  return request<unknown>(`/panels/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {}),
  });
}
