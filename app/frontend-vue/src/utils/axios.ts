import axios from 'axios'

const axiosInstance = axios.create({
  baseURL: 'http://127.0.0.1:8080', // 替换为实际API地址
  timeout: 5000,
  headers: {
    'Content-Type': 'application/json'
  },
  withCredentials: true
})

// 请求拦截器
axiosInstance.interceptors.request.use(config => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

// 响应拦截器
axiosInstance.interceptors.response.use(
  response => response.data,
  error => {
    if (error.response.status === 401) {
      window.location.href = '/login'
    }
    return Promise.reject(error)
  }
)

export default axiosInstance