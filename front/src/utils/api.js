import axios from 'axios'

const API_URL = 'http://localhost:8080' // Reemplaza con la URL de tu API

const api = axios.create({
  baseURL: API_URL,
  withCredentials: true, // Esto es importante para manejar las cookies
})

export const login = async (username, password) => {
  const response = await api.post('/auth/login', { username, password })
  return response.data
}

export const register = async (username, password) => {
  const response = await api.post('/auth/register', { username, password })
  return response.data
}

export const refreshToken = async () => {
  const response = await api.post('/auth/refresh')
  return response.data
}

export const getUserProfile = async () => {
  const response = await api.get('/user/profile')
  return response.data
}

// Interceptor para manejar la renovación del token
api.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config
    if (error.response.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true
      try {
        await refreshToken()
        return api(originalRequest)
      } catch (refreshError) {
        // Si no se puede refrescar el token, redirigir al login
        window.location.href = '/login'
        return Promise.reject(refreshError)
      }
    }
    return Promise.reject(error)
  }
)

export default api