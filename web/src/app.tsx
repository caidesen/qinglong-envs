import { createRouter, RouterProvider } from "@tanstack/react-router"
import { routeTree } from "./routeTree.gen.ts"
import { QueryClient, QueryClientProvider } from "@tanstack/react-query"

const router = createRouter({ routeTree: routeTree })
declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router
  }
}
const queryClient = new QueryClient()
export default function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router} />
    </QueryClientProvider>
  )
}
