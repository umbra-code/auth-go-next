import axios from 'axios'
import Router from 'next/router'
import Cookies from 'js-cookie'

const API_URL = 'http://localhost:8080' // Reemplaza con la URL de tu API

const api = axios.create({
  baseURL: API_URL,
  withCredentials: true,
})

export const login = async (username, password) => {
  return await api.post('/auth/login', { username, password })
}

export const register = async (username, password) => {
  return await api.post('/auth/register', { username, password })
}

export const refreshToken = async () => {
  return await api.post('/auth/refresh')
}

export const getUserProfile = async () => {
  return await api.get('/user/profile')
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
    if (axios.isCancel(error)) {
      return Promise.reject(error)
    }

    const originalRequest = error.config

    if (!error.response) {
      console.error('Network error:', error)
      return Promise.reject({
        isNetworkError: true,
        message: 'Network error. Please check your internet connection.'
      })
    }

    if (error.response.status === 401 && !originalRequest._retry) {
      if (isRefreshing) {
        return new Promise((resolve, reject) => {
          failedQueue.push({ resolve, reject })
        }).then(() => {
          return api(originalRequest)
        }).catch(err => {
          return Promise.reject(err)
        })
      }

      originalRequest._retry = true
      isRefreshing = true

      try {
        await refreshToken()
        processQueue(null)
        return api(originalRequest)
      } catch (refreshError) {
        processQueue(refreshError, null)
        clearAuthState()
        Router.push('/auth/login')
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
}

export default api