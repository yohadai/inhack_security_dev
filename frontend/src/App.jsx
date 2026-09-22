import React, { useState } from 'react'
import Signup from './pages/Signup.jsx'
import Login from './pages/Login.jsx'
import Me from './pages/Me.jsx'

// 화면 세 개(회원가입 / 로그인 / 내 정보)를 버튼으로 오가는 아주 단순한 앱
export default function App() {
  const [page, setPage] = useState('login')
  const [token, setToken] = useState('')

  return (
    <div style={{ fontFamily: 'sans-serif', maxWidth: 420, margin: '40px auto', padding: 16 }}>
      <h1>inhack demo</h1>
      <nav style={{ display: 'flex', gap: 8, marginBottom: 16 }}>
        <button onClick={() => setPage('signup')}>회원가입</button>
        <button onClick={() => setPage('login')}>로그인</button>
        <button onClick={() => setPage('me')}>내 정보</button>
      </nav>

      {page === 'signup' && <Signup />}
      {page === 'login' && <Login onToken={setToken} />}
      {page === 'me' && <Me token={token} />}
    </div>
  )
}
