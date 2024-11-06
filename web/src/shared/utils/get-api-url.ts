export const getApiUrl = (protocol = "http") =>
  `${protocol}://${process.env.VITE_API_URL || "localhost:8080"}`