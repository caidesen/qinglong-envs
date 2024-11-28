import { buildQuery, request } from "../utils/http.ts"

export interface PaginationParams {
  current: number
  pageSize: number
}

export interface Pagination<T> {
  current: number
  total: number
  data: T[]
}

export interface Panel {
  id: number
  createdAt: Date
  updatedAt: Date
  name: string
  url: string
  clientId: string
  clientSecret: string
}

export function ListPanels(input: PaginationParams) {
  return request
    .get("/api/panels" + buildQuery(input))
    .json<Pagination<Panel>>()
}
