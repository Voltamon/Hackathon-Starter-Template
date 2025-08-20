import { useState } from 'preact/hooks'
import { fetchHello } from './api'

export function App() {
  const [message, setMessage] = useState<string | null>(null)
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  async function onClick() {
    setLoading(true)
    setError(null)
    try {
      const msg = await fetchHello()
      setMessage(msg)
    } catch (err: any) {
      setError(err?.message || String(err))
    } finally {
      setLoading(false)
    }
  }

  return (
    <section>
      <button onClick={onClick}>
        {loading ? 'Loading...' : 'Fetch from backend'}
      </button>

      {message && (
        <div>
          <h3>Backend says:</h3>
          <pre>{message}</pre>
        </div>
      )}

      {error && (
        <div className="error">
          <p>Error: {error}</p>
        </div>
      )}
    </section>
  )
}
