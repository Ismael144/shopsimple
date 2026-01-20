import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faTrashAlt } from '@fortawesome/free-solid-svg-icons';
import { Link, useNavigate } from "react-router";
import { useCart } from "../context/CartContext";
import { useToast } from "../context/ToastContext";
import { useState, useEffect } from "react";
import axios from "axios";

interface Money {
    currencyCode: string;
    units: string;
    nanos: number;
}

interface CartItem {
    product_id: string;
    quantity: number;
    product_name: string;
    unit_price: Money;
}

interface Cart {
    user_id: string;
    cart_items: CartItem[];
    cart_price_total: Money;
    currency: string;
}

export default function ShoppingCart() {
    const navigate = useNavigate()
    const { items, removeItem, setQuantity, total } = useCart()
    const { addToast } = useToast()
    const [syncLoading, setSyncLoading] = useState(false)

    // Format price from Money object
    function formatPrice(price: number): string {
        return new Intl.NumberFormat('en-US', {
            style: 'currency',
            currency: 'USD',
            minimumFractionDigits: 2,
            maximumFractionDigits: 2
        }).format(price);
    }

    // Sync cart with backend when items change
    useEffect(() => {
        const syncCart = async () => {
            if (items.length === 0) return;
            
            try {
                setSyncLoading(true)
                // In a real app, you would send cart items to backend
                // For now, we're using local storage via CartContext
                // When user is authenticated, you can swap to auth user cart via AssignToAuthUser RPC
            } catch (err) {
                console.error('Error syncing cart:', err)
            } finally {
                setSyncLoading(false)
            }
        }

        syncCart()
    }, [items])

    const handleRemoveItem = (id: string | number) => {
        removeItem(id)
        addToast('Item removed from cart', 'success', 2000)
    }

    const handleCheckout = () => {
        if (items.length === 0) {
            addToast('Cart is empty', 'error', 2000)
            return
        }
        navigate('/shipping')
    }

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
                                    <img src={it.image || 'https://htmldemo.net/venezo/venezo/assets/img/s-product/product2.jpg'} width={88} alt={it.title} />
                                    <div>
                                        <div style={{ fontWeight: 600 }}>{it.title}</div>
                                        <div className="muted">{formatPrice(it.price)}</div>
                                    </div>
                                    <div style={{ marginLeft: "auto", display: "flex", gap: 8, alignItems: "center" }}>
                                        <input 
                                            type="number" 
                                            value={it.qty} 
                                            onChange={(e) => {
                                                const newQty = Number(e.target.value || 0);
                                                if (newQty === 0) {
                                                    handleRemoveItem(it.id);
                                                } else {
                                                    setQuantity(it.id, newQty);
                                                }
                                            }} 
                                            style={{ width: 72, padding: 6, borderRadius: 6, border: "1px solid #e5e7eb" }} 
                                        />
                                        <button 
                                            style={{ background: "transparent", border: "none", color: "#ef4444", cursor: "pointer" }} 
                                            aria-label="remove" 
                                            onClick={() => handleRemoveItem(it.id)}
                                        >
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
                            <div>{formatPrice(total)}</div>
                        </div>
                        <div style={{ display: "flex", justifyContent: "space-between", marginBottom: 12 }}>
                            <div className="muted">Shipping</div>
                            <div>{formatPrice(10)}</div>
                        </div>
                        <hr />
                        <div style={{ display: "flex", justifyContent: "space-between", fontWeight: 700, marginTop: 12 }}>
                            <div>Total</div>
                            <div>{formatPrice(total + 10)}</div>
                        </div>
                        <div style={{ marginTop: 14 }}>
                            <button 
                                className="btn btn-success" 
                                onClick={handleCheckout}
                                disabled={items.length === 0 || syncLoading}
                                style={{ opacity: (items.length === 0 || syncLoading) ? 0.6 : 1, cursor: (items.length === 0 || syncLoading) ? 'not-allowed' : 'pointer' }}
                            >
                                {syncLoading ? 'Syncing...' : 'Proceed to Checkout'}
                            </button>
                        </div>
                    </div>
                </aside>
            </div>
        </div>
    )
}