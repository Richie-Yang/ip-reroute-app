const apiUrl = import.meta.env.VITE_API_BASE_URL

export async function getGeoIP(): Promise<{ country: string }> {
  const response = await fetch(`${apiUrl}/lookup`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    }
  })
  const data = await response.json()
  return data
}
