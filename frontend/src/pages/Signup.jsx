import React, { useState } from 'react'
import trim from 'lodash/trim'
import { signup } from '../api.js'

export default function Signup() {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [message, setMessage] = useState('')

  async function onSubmit(e) {
    e.preventDefault()
    try {
      await signup(trim(username), password)
      setMessage('회원가입 성공')
    } catch (err) {
      setMessage('실패: ' + (err.response?.data?.error || err.message))
    }
  }

  return (
    <form onSubmit={onSubmit}>
      <h2>회원가입</h2>
      <input placeholder="username" value={username} onChange={(e) => setUsername(e.target.value)} />
      <br />
      <input placeholder="password" type="password" value={password} onChange={(e) => setPassword(e.target.value)} />
      <br />
      <button type="submit">가입</button>
      <p>{message}</p>
    </form>
  )
}
