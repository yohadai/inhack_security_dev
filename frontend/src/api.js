import axios from 'axios'

// auth-service 와 api-service 의 주소. 개발 중에는 localhost 를 쓴다.
const AUTH_URL = import.meta.env.VITE_AUTH_URL || 'http://localhost:8081'
const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8082'

export async function signup(username, password) {
  const res = await axios.post(`${AUTH_URL}/signup`, { username, password })
  return res.data
}

export async function login(username, password) {
  const res = await axios.post(`${AUTH_URL}/login`, { username, password })
  return res.data
}

export async function getMe(token) {
  const res = await axios.get(`${API_URL}/me`, {
    headers: { Authorization: `Bearer ${token}` },
  })
  return res.data
}
