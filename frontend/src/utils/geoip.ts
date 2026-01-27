export async function getGeoIP(): Promise<{ country: string }> {
  const response = await fetch('http://localhost:8080/lookup', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    }
  })
  const data = await response.json()
  return data
}
