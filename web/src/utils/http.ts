import { HTTPError } from "@/apis"

export type RequestOptions = RequestInit & {
  data?: Record<string, unknown>
  params?: Record<string, unknown>
}

class APIError extends Error {
  response?: Response
  data?: HTTPError

  constructor(response: Response, body?: HTTPError) {
    super(body?.message)
    this.response = response
    this.data = body
  }
}

function safetyParseJSONBody(resp: Response) {
  if (!resp.headers.get("Content-Type")?.includes("application/json"))
    return null
  try {
    return resp.json()
  } catch {
    return null
  }
}

const baseURL = new URL("/api/v1", document.baseURI).toString()

async function request<T>(url: string, opt: RequestOptions = {}): Promise<T> {
  const { data, params, ...optWithoutData } = opt
  const fetchURL = new URL(baseURL + url)
  if (params)
    for (const [key, value] of Object.entries(params)) {
      fetchURL.searchParams.set(key, String(value))
    }
  const resp = await fetch(fetchURL, {
    ...optWithoutData,
    body: data ? JSON.stringify(data) : undefined,
  })
  const body = await safetyParseJSONBody(resp)
  if (!resp.ok) {
    // error handle
    throw new APIError(resp, body)
  }
  return body as T
}

export default request
