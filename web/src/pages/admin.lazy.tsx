import {
  createLazyFileRoute,
  Link,
  LinkProps,
  Outlet,
} from "@tanstack/react-router"
import {
  FiHash,
  FiLayout,
  FiMail,
  FiServer,
  FiSettings,
  FiUsers,
} from "react-icons/fi"
import { type IconType } from "react-icons"
import { Button } from "@/components/ui/button"
import { FaRegCircleUser } from "react-icons/fa6"

export const Route = createLazyFileRoute("/admin")({
  component: RouteComponent,
})
const navList = [
  {
    to: "/admin/panel",
    title: "面板",
    icon: FiServer,
  },
  {
    to: "/admin/project",
    title: "项目",
    icon: FiLayout,
  },
  {
    to: "/admin/project",
    title: "用户",
    icon: FiUsers,
  },
  {
    to: "/admin/project",
    title: "变量",
    icon: FiHash,
  },
  {
    to: "/admin/project",
    title: "推送",
    icon: FiMail,
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
      <div className="pt-4 mb-6 bg-muted">
        <div className="container mx-auto">
          <div className="flex justify-between items-center">
            <div className="text-xl font-bold">QingLong Envs</div>
            <Button className="-mr-4" variant="ghost">
              管理员
              <FaRegCircleUser />
            </Button>
          </div>
          <div className="mt-10 pb-2 -mx-3 flex gap-4">
            {navList.map((item) => (
              <Button
                className="text-[1rem] group"
                variant="ghost"
                size="sm"
                asChild
              >
                <Link
                  to={item.to}
                  key={item.to}
                  activeProps={{ className: "group active" }}
                >
                  <div className="flex relative items-center gap-2 w-max group-[.active]:text-primary">
                    <item.icon className="!size-[18px]" />
                    {item.title}
                    <div className="w-full h-0.5 bg-primary absolute rounded-full -bottom-3.5 hidden group-[.active]:block"></div>
                  </div>
                </Link>
              </Button>
            ))}
          </div>
        </div>
      </div>
      <div className="container mx-auto">
        <Outlet />
      </div>
    </>
  )
}
