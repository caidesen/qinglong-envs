import { createRootRoute, Outlet } from "@tanstack/react-router"
import React from "react"

const Provider: React.FC<React.PropsWithChildren> = (props) => {
  return <div>{props.children}</div>
}

const RootLayout: React.FC = () => {
  return (
    <>
      <Provider>
        <Outlet />
      </Provider>
    </>
  )
}

export const Route = createRootRoute({
  component: RootLayout,
})
