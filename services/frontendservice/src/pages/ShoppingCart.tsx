import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faTrashAlt } from '@fortawesome/free-solid-svg-icons';
import { Link, useNavigate } from "react-router";
import { useCart } from "../context/CartContext";

export default function ShoppingCart() {
    const navigate = useNavigate()
    const { items, removeItem, setQuantity, total } = useCart()

    return (
        <div className="container my-5">
            <div style={{ display: "grid", gridTemplateColumns: "1fr 320px", gap: 20 }}>
                <div>
                    {items.length === 0 ? (
                        <div className="card" style={{ padding: 20 }}>
                            <div className="text-center muted">Your cart is empty</div>
                            <div className="text-center mt-2"><Link to="/" className="btn btn-success">Browse products</Link></div>
                        </div>
                    ) : (
                        items.map(it => (
                            <div key={String(it.id)} className="card" style={{ padding: 12, marginBottom: 12 }}>
                                <div style={{ display: "flex", gap: 12, alignItems: "center" }}>
                                    <img src={it.image || 'https://htmldemo.net/venezo/venezo/assets/img/s-product/product2.jpg'} width={88} alt="" />
                                    <div>
                                        <div style={{ fontWeight: 600 }}>{it.title}</div>
                                        <div className="muted">${it.price.toFixed(2)}</div>
                                    </div>
                                    <div style={{ marginLeft: "auto", display: "flex", gap: 8, alignItems: "center" }}>
                                        <input type="number" value={it.qty} onChange={(e) => setQuantity(it.id, Number(e.target.value || 0))} style={{ width: 72, padding: 6, borderRadius: 6, border: "1px solid #e5e7eb" }} />
                                        <button style={{ background: "transparent", border: "none", color: "#ef4444", cursor: "pointer" }} aria-label="remove" onClick={() => removeItem(it.id)}>
                                            <FontAwesomeIcon icon={faTrashAlt} />
                                        </button>
                                    </div>
                                </div>
                            </div>
                        ))
                    )}
                </div>

                <aside>
                    <div className="card" style={{ padding: 14 }}>
                        <div style={{ fontWeight: 700, marginBottom: 8 }}>Order Summary</div>
                        <div style={{ display: "flex", justifyContent: "space-between", marginBottom: 6 }}>
                            <div className="muted">Subtotal</div>
                            <div>${total.toFixed(2)}</div>
                        </div>
                        <div style={{ display: "flex", justifyContent: "space-between", marginBottom: 12 }}>
                            <div className="muted">Shipping</div>
                            <div>$10.00</div>
                        </div>
                        <hr />
                        <div style={{ display: "flex", justifyContent: "space-between", fontWeight: 700, marginTop: 12 }}>
                            <div>Total</div>
                            <div>${(total + 10).toFixed(2)}</div>
                        </div>
                        <div style={{ marginTop: 14 }}>
                            <button className="btn btn-success" onClick={() => navigate('/shipping')}>Proceed to Checkout</button>
                        </div>
                    </div>
                </aside>
            </div>
        </div>
    )
}