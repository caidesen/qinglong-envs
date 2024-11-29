export interface PaginationParams {
  current: number
  pageSize: number
}

export interface Pagination<T> {
  current: number
  total: number
  data: T[]
}
