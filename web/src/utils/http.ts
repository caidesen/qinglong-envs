// export interface Request {
//   <T>(url: string, options: RequestInit): Promise<T>
// }
//
// async function fetchResultHandle(res: Response) {
//   const isJSONBody = res.headers
//     .get("content-type")
//     ?.includes("application/json")
//   const body = isJSONBody ? await res.json() : undefined
//   if (!res.ok) {
//     const msg = body ? body.message : `request fall: ${res.statusText}`
//     throw new Error(msg)
//   }
//   if (!isJSONBody) return res
//   return body
// }
//
// function timeoutAbort(signal: AbortSignal | null | undefined) {
//   const abortController = new AbortController()
//   setTimeout(() => abortController.abort("timeout"), 1000 * 5)
//   if (signal)
//     signal.onabort = (ev) => {
//       abortController.abort((ev.target as AbortSignal)?.reason ?? "aborted")
//     }
//   return abortController.signal
// }
//
// export const request: Request = async (url, options = {}) => {
//   try {
//     const res = await fetch("/api" + url, {
//       ...options,
//       signal: timeoutAbort(options.signal),
//     })
//     return await fetchResultHandle(res)
//   } catch (error) {
//     if (error instanceof Error) {
//       throw error
//     } else {
//       throw new Error(String(error))
//     }
//   }
// }
//
export function buildQuery(params: Record<string, any>) {
  return Object.entries(params)
    .map(([key, value]) => `${key}=${value}`)
    .join("&")
}

import ky from "ky"

export const request = ky.extend({
  prefixUrl: "/api",
})
