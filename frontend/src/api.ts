export async function fetchHello(): Promise<string> {
  const res = await fetch('http://localhost:8080/api/pocketbase', { method: 'GET' })
  if (!res.ok) {
    throw new Error(`Request failed: ${res.status}`)
  }
  const data = await res.json()
  return data.pocketbase
}
