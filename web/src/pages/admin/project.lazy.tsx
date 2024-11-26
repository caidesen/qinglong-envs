import { createLazyFileRoute } from '@tanstack/react-router'

export const Route = createLazyFileRoute('/admin/project')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/admin/panel"!</div>
}
