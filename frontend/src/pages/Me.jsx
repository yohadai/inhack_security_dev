import React, { useState } from 'react'
import { getMe } from '../api.js'

export default function Me({ token }) {
  const [info, setInfo] = useState(null)
  const [message, setMessage] = useState('')

  async function onLoad() {
    try {
      const data = await getMe(token)
      setInfo(data)
      setMessage('')
    } catch (err) {
      setInfo(null)
      setMessage('실패: ' + (err.response?.data?.error || err.message))
    }
  }

  return (
    <div>
      <h2>내 정보</h2>
      <button onClick={onLoad}>불러오기</button>
      {info && <pre>{JSON.stringify(info, null, 2)}</pre>}
      <p>{message}</p>
    </div>
  )
}
