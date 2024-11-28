import {
  createFileRoute,
  Link,
  LinkProps,
  Outlet,
  useRouterState,
} from "@tanstack/react-router"
import { Tab, Tabs } from "@nextui-org/tabs"
import * as React from "react"
import { FiDatabase, FiLayout, FiSettings } from "react-icons/fi"
import { type IconType } from "react-icons"

export const Route = createFileRoute("/admin")({
  component: RouteComponent,
})

const Nav: React.FC = () => {
  const pathname = useRouterState().location.pathname
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
  return (
    <Tabs
      variant="underlined"
      aria-label="admin tabs"
      size="lg"
      color="primary"
      selectedKey={pathname}
      classNames={{
        tabList: "p-0",
        tab: "h-12",
      }}
    >
      {navList.map(({ to, title, icon: Icon }) => (
        <Tab
          key={to}
          href={to}
          title={
            <div className="flex items-center gap-2 font-bold">
              <Icon className="text-lg" />
              {title}
            </div>
          }
          as={Link}
        ></Tab>
      ))}
    </Tabs>
  )
}

function RouteComponent() {
  return (
    <>
      <div className="bg-gray-100 border-b pt-4 mb-6">
        <div className="container mx-auto">
          <div className="text-lg font-bold">XXX</div>
          <div className="-mx-2 mt-4">
            <Nav></Nav>
          </div>
        </div>
      </div>
      <div className="container mx-auto">
        <Outlet />
      </div>
    </>
  )
}
