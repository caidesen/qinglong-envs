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

function RouteComponent() {
  const getPanelsQuery = useQuery({
    queryKey: ["get /panels"],
    queryFn: async () => api.getPanels(),
  })
  return (
    <div>
      <Table>
        <TableHeader>
          <TableColumn>名称</TableColumn>
          <TableColumn>url</TableColumn>
          <TableColumn>client_id</TableColumn>
          <TableColumn>client_secret</TableColumn>
        </TableHeader>
        <TableBody emptyContent="No rows to display.">{[]}</TableBody>
      </Table>
      {getPanelsQuery.data?.toString()}
    </div>
  )
}
