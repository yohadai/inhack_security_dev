import React, { useState } from 'react'
import { login } from '../api.js'

export default function Login({ onToken }) {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [message, setMessage] = useState('')

  async function onSubmit(e) {
    e.preventDefault()
    try {
      const data = await login(username, password)
      onToken(data.token)
      setMessage('로그인 성공. 토큰을 받았습니다.')
    } catch (err) {
      setMessage('실패: ' + (err.response?.data?.error || err.message))
    }
  }

  return (
    <form onSubmit={onSubmit}>
      <h2>로그인</h2>
      <input placeholder="username" value={username} onChange={(e) => setUsername(e.target.value)} />
      <br />
      <input placeholder="password" type="password" value={password} onChange={(e) => setPassword(e.target.value)} />
      <br />
      <button type="submit">로그인</button>
      <p>{message}</p>
    </form>
  )
}
