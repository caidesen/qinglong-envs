import { createLazyFileRoute } from "@tanstack/react-router"
import { Button } from "@nextui-org/button"

export const Route = createLazyFileRoute("/")({
  component: Index,
})

function Index() {
  return (
    <div className="p-2">
      <Button>123</Button>
      <h3>Welcome Home!</h3>
    </div>
  )
}
