import { createLazyFileRoute } from '@tanstack/react-router'

export const Route = createLazyFileRoute('/admin/about')({
  component: About,
})

function About() {
  return '123'
}
