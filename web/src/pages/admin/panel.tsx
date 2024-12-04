import { createFileRoute } from "@tanstack/react-router"
import { useQuery } from "@tanstack/react-query"
import * as api from "@/apis"
import { Table, TableBody, TableColumn, TableHeader } from "@nextui-org/table"
import { Spinner } from "@nextui-org/spinner"

export const Route = createFileRoute("/admin/panel")({
  component: RouteComponent,
  loader() {
    return new Promise((resolve) => {
      setTimeout(resolve, 2000)
    })
  },
  pendingComponent: () => (
    <div className="pt-[10vh] text-center">
      <Spinner></Spinner>
    </div>
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
    <div>
      <Table>
        <TableHeader columns={columus}>
          {(it) => <TableColumn>{it.label}</TableColumn>}
        </TableHeader>
        <TableBody emptyContent="No rows to display.">{[]}</TableBody>
      </Table>
      {getPanelsQuery.data?.toString()}
    </div>
  )
}
