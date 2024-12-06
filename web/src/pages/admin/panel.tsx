import { createFileRoute } from "@tanstack/react-router"
import { useQuery } from "@tanstack/react-query"
import * as api from "@/apis"
import { Button } from "@/components/ui/button.tsx"
import { FiExternalLink, FiMoreHorizontal, FiPlus } from "react-icons/fi"
import {
  ColumnDef,
  createColumnHelper,
  flexRender,
  getCoreRowModel,
  useReactTable,
} from "@tanstack/react-table"
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table.tsx"
import { Badge } from "@/components/ui/badge.tsx"
import { cn } from "@/lib/utils.ts"

export const Route = createFileRoute("/admin/panel")({
  component: RouteComponent,
  pendingComponent: () => (
    <div className="pt-[10vh] text-center">loading...</div>
  ),
})

interface DataTableProps<TData, TValue> {
  columns: ColumnDef<TData, TValue>[]
  data: TData[]
}

export function DataTable<TData, TValue>({
  columns,
  data,
}: DataTableProps<TData, TValue>) {
  const table = useReactTable({
    data,
    columns,
    getCoreRowModel: getCoreRowModel(),
  })

  return (
    <div className="">
      <Table className="border-b">
        <TableHeader>
          {table.getHeaderGroups().map((headerGroup) => (
            <TableRow key={headerGroup.id}>
              {headerGroup.headers.map((header) => {
                return (
                  <TableHead
                    key={header.id}
                    className={cn(
                      "font-bold",
                      header.id === "actions" && "w-0"
                    )}
                  >
                    {header.isPlaceholder
                      ? null
                      : flexRender(
                          header.column.columnDef.header,
                          header.getContext()
                        )}
                  </TableHead>
                )
              })}
            </TableRow>
          ))}
        </TableHeader>
        <TableBody>
          {table.getRowModel().rows?.length ? (
            table.getRowModel().rows.map((row) => (
              <TableRow
                key={row.id}
                data-state={row.getIsSelected() && "selected"}
              >
                {row.getVisibleCells().map((cell) => (
                  <TableCell key={cell.id}>
                    {flexRender(cell.column.columnDef.cell, cell.getContext())}
                  </TableCell>
                ))}
              </TableRow>
            ))
          ) : (
            <TableRow>
              <TableCell colSpan={columns.length} className="h-24 text-center">
                No results.
              </TableCell>
            </TableRow>
          )}
        </TableBody>
      </Table>
    </div>
  )
}

function RouteComponent() {
  const getPanelsQuery = useQuery({
    queryKey: ["get /panels"],
    queryFn: async () => api.getPanels(),
  })
  const columnHelper = createColumnHelper<api.Panel>()
  const columns: ColumnDef<api.Panel, string>[] = [
    columnHelper.accessor("id", { header: "ID" }),
    columnHelper.accessor("name", { header: "名称" }),
    columnHelper.accessor("url", {
      header: "访问地址",
      cell: (val) => (
        <Button variant="link" size="sm" className="p-0" asChild>
          <a href={val.getValue()}>
            {val.renderValue()}
            <FiExternalLink />
          </a>
        </Button>
      ),
    }),
    columnHelper.accessor("clientId", {
      header: "Client ID",
      cell: () => "***",
    }),
    columnHelper.accessor("clientSecret", {
      header: "Client Secret",
      cell: () => "***",
    }),
    columnHelper.display({
      id: "actions",
      cell: () => (
        <div className="w-min">
          <Button variant="ghost" size="icon">
            <FiMoreHorizontal className="size-4" />
          </Button>
        </div>
      ),
    }),
  ]
  return (
    <div className="flex-col flex gap-4">
      <div className="flex gap-4 items-end">
        <div className="flex-1">
          <div className="font-bold text-2xl">青龙面板</div>
          <div className="text-gray-500 mt-2">在这里添加你的青龙面板</div>
        </div>
        <div>
          <Button>
            <FiPlus />
            <span>添加面板</span>
          </Button>
        </div>
      </div>
      <div>
        <Badge variant="secondary">{getPanelsQuery.data?.length}个面板</Badge>
      </div>
      <DataTable columns={columns} data={getPanelsQuery.data ?? []} />
    </div>
  )
}
