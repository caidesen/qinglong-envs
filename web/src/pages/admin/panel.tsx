import { createFileRoute } from "@tanstack/react-router"
import { useQuery } from "@tanstack/react-query"
import * as api from "@/apis"

export const Route = createFileRoute("/admin/panel")({
  component: RouteComponent,
  pendingComponent: () => (
    <div className="pt-[10vh] text-center">loading...</div>
  ),
})
type TableColumnType = {
  key: keyof api.Panel
  label: string
}

function RouteComponent() {
  const getPanelsQuery = useQuery({
    queryKey: ["get /panels"],
    queryFn: async () => api.getPanels(),
  })
  const columus: TableColumnType[] = [
    {
      key: "id",
      label: "ID",
    },
    {
      key: "name",
      label: "名称",
    },
    {
      key: "clientId",
      label: "clientId",
    },
    {
      key: "clientSecret",
      label: "clientSecret",
    },
  ]
  return (
    <div className="flex-col flex gap-4">
      <div className="flex gap-3">
        <div className="flex-1"></div>
        <div></div>
      </div>
    </div>
  )
}
