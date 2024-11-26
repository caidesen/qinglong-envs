import { createRootRoute, Outlet, useRouter } from "@tanstack/react-router"
import React from "react"
import { NextUIProvider } from "@nextui-org/system"

const Provider: React.FC<React.PropsWithChildren> = (props) => {
  const router = useRouter()
  console.log("x")
  return (
    <NextUIProvider
      locale="zh-CN"
      navigate={(to) => {
        alert(1)
        return router.navigate({ to })
      }}
      useHref={(to) => {
        return router.buildLocation({ to }).href
      }}
    >
      {props.children}
    </NextUIProvider>
  )
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
