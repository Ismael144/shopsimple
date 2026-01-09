import { useNavigate } from "react-router";
import { useCart } from "../context/CartContext";
import React from "react";

export default function Payment() {
  const navigate = useNavigate()
  const { items, total, clearCart } = useCart()

  function handlePay(e: any) {
    e.preventDefault()
    // pretend to process payment
    clearCart()
    navigate('/order-confirmation', { replace: true })
  }

  return (
    <div className="container my-5">
      <div className="checkout-grid">
        <div>
          <div className="checkout-card">
            <div className="stepper">
              <div className="step">Shipping</div>
              <div className="step active">Payment</div>
              <div className="step">Confirm</div>
            </div>

            <h3 style={{ marginTop: 6 }}>Payment</h3>
            <p className="muted">Secure payment — use test card numbers for demo.</p>

            <form onSubmit={handlePay}>
              <div style={{ marginTop: 8, marginBottom: 10 }}>
                <label className="muted">Name on card</label>
                <input required className="form-control" style={{ padding: 10, borderRadius: 8 }} />
              </div>

              <div style={{ marginBottom: 10 }}>
                <label className="muted">Card number</label>
                <div style={{ display: 'flex', gap: 8 }}>
                  <input required className="form-control" placeholder="4242 4242 4242 4242" style={{ padding: 10, borderRadius: 8 }} />
                </div>
              </div>

              <div style={{ display: 'flex', gap: 10 }}>
                <div style={{ flex: 1 }}>
                  <label className="muted">Expiry</label>
                  <input required className="form-control" placeholder="MM/YY" style={{ padding: 10, borderRadius: 8 }} />
                </div>
                <div style={{ width: 120 }}>
                  <label className="muted">CVC</label>
                  <input required className="form-control" placeholder="123" style={{ padding: 10, borderRadius: 8 }} />
                </div>
              </div>

              <div style={{ marginTop: 14 }}>
                <button className="btn btn-success" type="submit">Pay ${ (total + 10).toFixed(2) }</button>
              </div>
            </form>
          </div>
        </div>

        <aside>
          <div className="checkout-card">
            <h5>Order summary</h5>
            <div style={{ marginTop: 10 }}>
              {items.map(it => (
                <div key={String(it.id)} style={{ display: 'flex', justifyContent: 'space-between', marginBottom: 8 }}>
                  <div style={{ display: 'flex', gap: 10, alignItems: 'center' }}>
                    <img src={it.image || 'https://htmldemo.net/venezo/venezo/assets/img/s-product/product2.jpg'} width={48} alt="" />
                    <div>
                      <div style={{ fontWeight: 600 }}>{it.title}</div>
                      <div className="muted">Qty {it.qty}</div>
                    </div>
                  </div>
                  <div>${(it.price * it.qty).toFixed(2)}</div>
                </div>
              ))}

              <hr />
              <div style={{ display: 'flex', justifyContent: 'space-between', fontWeight: 700 }}>
                <div>Total</div>
                <div>${(total + 10).toFixed(2)}</div>
              </div>
            </div>
          </div>
        </aside>
      </div>
    </div>
  )
}
