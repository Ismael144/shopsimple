import { Link } from "react-router";

export default function OrderConfirmation(){
  const orderId = Math.random().toString(36).slice(2,9).toUpperCase()

  return (
    <div className="container my-5 text-center">
      <div className="card" style={{ padding: 28, maxWidth: 640, margin: '0 auto' }}>
        <h2>Thank you — your order is confirmed</h2>
        <p className="muted">Order ID: <strong>{orderId}</strong></p>
        <p className="muted">We emailed the receipt and will notify you when your items ship.</p>
        <div style={{ marginTop: 16 }}>
          <Link to="/" className="btn btn-success">Continue shopping</Link>
        </div>
      </div>
    </div>
  )
}
