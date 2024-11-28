import { createLazyFileRoute } from "@tanstack/react-router"

export const Route = createLazyFileRoute("/admin/settings")({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/admin/settings"!</div>
}
