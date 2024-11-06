export const getApiUrl = (protocol = "http") => {
  const url = `${protocol}://${import.meta.env.VITE_API_URL || "localhost:8080"}`
  console.log("getApiUrl", url)
  return url
}