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

let isRefreshing = false
let failedQueue = []

const processQueue = (error, token = null) => {
  failedQueue.forEach(prom => {
    if (error) {
      prom.reject(error)
    } else {
      prom.resolve(token)
    }
  })

  failedQueue = []
}

api.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config

    if (error.response.status === 401 && !originalRequest._retry) {
      if (isRefreshing) {
        return new Promise((resolve, reject) => {
          failedQueue.push({ resolve, reject })
        }).then(token => {
          return api(originalRequest)
        }).catch(err => {
          return Promise.reject(err)
        })
      }

      originalRequest._retry = true
      isRefreshing = true

      try {
        await refreshToken()
        // No necesitamos manejar el token aquí, ya que se establece en las cookies
        processQueue(null)
        return api(originalRequest)
      } catch (refreshError) {
        processQueue(refreshError, null)
        // Redirigir al login y limpiar el estado de autenticación
        clearAuthState()
        Router.push('/login')
        return Promise.reject(refreshError)
      } finally {
        isRefreshing = false
      }
    }

    return Promise.reject(error)
  }
)

export const clearAuthState = () => {
  Cookies.remove('access-token')
  Cookies.remove('refresh-token')
  // Aquí puedes agregar cualquier otra limpieza necesaria
}

export default api