import React from 'react'
import { useNavigate } from 'react-router'

export default function SignUp(){
  const navigate = useNavigate()

  function handleSubmit(e: React.FormEvent<HTMLFormElement>){
    e.preventDefault()
    const fd = new FormData(e.currentTarget)
    const payload = {
      username: fd.get('username'),
      email: fd.get('email'),
      password: fd.get('password'),
      phone: fd.get('phone')
    }
    // In a real app: send payload to server
    console.log('signup', payload)
    navigate('/signin')
  }

  return (
    <div className="container d-flex align-center justify-content-center" style={{ minHeight: '80vh' }}>
      <div className="auth-form bg-white p-4" style={{ width: 480 }}>
        <h2 className="text-secondary text-center my-3">Create account</h2>
        <form onSubmit={handleSubmit}>
          <div className="mb-3">
            <label className="form-label">Username</label>
            <input name="username" required className="form-control" placeholder="Choose a username" />
          </div>

          <div className="mb-3">
            <label className="form-label">Email</label>
            <input name="email" type="email" required className="form-control" placeholder="you@example.com" />
          </div>

          <div className="mb-3">
            <label className="form-label">Phone</label>
            <input name="phone" type="tel" className="form-control" placeholder="Optional phone number" />
          </div>

          <div className="mb-3">
            <label className="form-label">Password</label>
            <input name="password" type="password" required className="form-control" placeholder="Create a password" />
          </div>

          <div className="d-grid">
            <button className="btn btn-success">Create account</button>
          </div>

          <div className="text-center mt-3 muted">Already have an account? <a href="/signin">Sign in</a></div>
        </form>
      </div>
    </div>
  )
}
