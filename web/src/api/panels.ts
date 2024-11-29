import { buildQuery, request } from "../utils/http.ts"
import { Pagination, PaginationParams } from "../type/model"

export interface Panel {
  id: number
  createdAt: Date
  updatedAt: Date
  name: string
  url: string
  clientId: string
  clientSecret: string
}

export function listPanels(input: PaginationParams) {
  return request
    .get("/api/panels" + buildQuery(input))
    .json<Pagination<Panel>>()
}
