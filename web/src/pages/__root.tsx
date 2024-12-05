import { createRootRoute, Outlet } from "@tanstack/react-router"
import React from "react"

import "@mantine/core/styles.css"
import { MantineProvider } from "@mantine/core"

const Provider: React.FC<React.PropsWithChildren> = (props) => {
  return <MantineProvider>{props.children}</MantineProvider>
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
