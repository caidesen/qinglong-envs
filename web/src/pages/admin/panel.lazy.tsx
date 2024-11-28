import { createLazyFileRoute } from "@tanstack/react-router"
import { useQuery } from "@tanstack/react-query"
import { request } from "../../utils/http.ts"

export const Route = createLazyFileRoute("/admin/panel")({
  component: RouteComponent,
})
function RouteComponent() {
  useQuery({
    queryKey: ["/panels"],
    queryFn: async () =>
      request("/panels", {
        method: "GET",
      }),
  })
  return <div>Hello "/admin/panel"!</div>
}
