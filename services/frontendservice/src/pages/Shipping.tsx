import { useNavigate } from "react-router";
import { useCart } from "../context/CartContext";
import React from "react";

export default function Shipping(){
  const navigate = useNavigate()
  const { items, total } = useCart()

  function handleContinue(e:any){
    e.preventDefault()
    // save shipping if needed
    navigate('/checkout')
  }

  return (
    <div className="container my-5">
      <div className="checkout-grid">
        <div>
          <div className="checkout-card">
            <div className="stepper">
              <div className="step active">Shipping</div>
              <div className="step">Payment</div>
              <div className="step">Confirm</div>
            </div>

            <h3 style={{ marginTop: 6 }}>Shipping</h3>
            <p className="muted">Enter shipping address for your order.</p>

            <form onSubmit={handleContinue}>
              <div style={{ marginTop: 8, marginBottom: 10 }}>
                <label className="muted">Full name</label>
                <input required className="form-control" style={{ padding: 10, borderRadius: 8 }} />
              </div>

              <div style={{ marginBottom: 10 }}>
                <label className="muted">Address</label>
                <input required className="form-control" style={{ padding: 10, borderRadius: 8 }} />
              </div>

              <div style={{ display: 'flex', gap: 10 }}>
                <div style={{ flex: 1 }}>
                  <label className="muted">City</label>
                  <input required className="form-control" style={{ padding: 10, borderRadius: 8 }} />
                </div>
                <div style={{ width: 120 }}>
                  <label className="muted">ZIP</label>
                  <input required className="form-control" style={{ padding: 10, borderRadius: 8 }} />
                </div>
              </div>

              <div style={{ marginTop: 14 }}>
                <button className="btn btn-success" type="submit">Continue to Payment</button>
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
                <div>Subtotal</div>
                <div>${total.toFixed(2)}</div>
              </div>
            </div>
          </div>
        </aside>
      </div>
    </div>
  )
}
