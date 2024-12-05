import { createLazyFileRoute, LinkProps, Outlet } from "@tanstack/react-router"
import * as React from "react"
import { FiDatabase, FiLayout, FiSettings } from "react-icons/fi"
import { type IconType } from "react-icons"
import { Button } from "../components/ui/button.tsx"

export const Route = createLazyFileRoute("/admin")({
  component: RouteComponent,
})
const navList = [
  {
    to: "/admin/panel",
    title: "面板",
    icon: FiDatabase,
  },
  {
    to: "/admin/project",
    title: "项目",
    icon: FiLayout,
  },
  {
    to: "/admin/settings",
    title: "设置",
    icon: FiSettings,
  },
] as const satisfies ReadonlyArray<{
  to: LinkProps["to"]
  title: string
  icon: IconType
}>

function RouteComponent() {
  return (
    <>
      <div className=" border-b pt-4 mb-6">
        <div className="container mx-auto flex justify-between items-center">
          <div className="text-lg font-bold">XXX</div>
          <div className="">admin</div>
        </div>
        <div className="container mx-auto">
          <Button
            className="hover:bg-accent hover:text-accent-foreground"
            variant="ghost"
          >
            Ghost
          </Button>
        </div>
      </div>
      <div className="container mx-auto">
        <Outlet />
      </div>
    </>
  )
}
