import React, { useMemo, useState } from "react";
import { useParams, useNavigate } from "react-router";
import FilterSection from "../components/FilterSection";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faStar, faChevronLeft } from "@fortawesome/free-solid-svg-icons";
import { useCart } from "../context/CartContext";
import { useToast } from "../context/ToastContext";

type Product = {
    id: string
    title: string
    price: number
    image: string
    rating?: number
    description?: string
    category?: string
    stock: number
}

const SAMPLE: Record<string, Product> = {
    '1': { id: '1', title: 'High Empeded High Tech Headphones', price: 170, image: 'https://htmldemo.net/venezo/venezo/assets/img/product/product1-big.jpg', rating: 4.3, description: 'Clear sound, comfortable fit and long battery life.' , category: 'Audio', stock: 15 },
    '2': { id: '2', title: 'Wireless Headset', price: 120, image: 'https://htmldemo.net/venezo/venezo/assets/img/product/product2-big.jpg', rating: 4, description: 'Lightweight and powerful.' , category: 'Audio', stock: 8 },
    '3': { id: '3', title: 'Bluetooth Speaker', price: 90, image: 'https://htmldemo.net/venezo/venezo/assets/img/product/product3-big.jpg', rating: 4.5, description: 'Portable speaker with rich bass.' , category: 'Speakers', stock: 0 }
}

export default function ProductDetails() {
    const { id } = useParams()
    const product = useMemo(() => SAMPLE[String(id)] ?? SAMPLE['1'], [id])
    const [qty, setQty] = useState<number>(1)
    const navigate = useNavigate()
    const { addItem, items, setQuantity } = useCart()
    const { addToast } = useToast()
    const cartItem = items.find(item => String(item.id) === String(product.id))
    const cartQuantity = cartItem?.qty || 0

    function handleAdd() {
        if (cartQuantity + qty > product.stock) {
            addToast(`Only ${product.stock - cartQuantity} items available!`, 'error', 2500)
            return
        }
        addItem({ id: product.id, title: product.title, price: product.price, image: product.image }, qty)
        addToast(`${qty} ${title} added to cart!`, 'success', 2500)
    }

    function handleReduce() {
        if (cartQuantity > 0) {
            const newQty = cartQuantity - qty
            if (newQty <= 0) {
                setQuantity(product.id, 0)
            } else {
                setQuantity(product.id, newQty)
            }
        }
    }

    return (
        <div className="container my-5">
            <div className="products-listing-grid">
                <FilterSection />
                <div style={{ padding: 12 }}>
                    <div className="card" style={{ padding: 18 }}>
                        <div style={{ display: 'grid', gridTemplateColumns: '1fr 380px', gap: 20 }}>
                            <div>
                                <img src={product.image} alt={product.title} style={{ width: '100%', borderRadius: 8 }} />
                            </div>
                            <div>
                                <div style={{ display: 'flex', gap: 12, alignItems: 'center', marginBottom: 8 }}>
                                    <button onClick={() => navigate(-1)} className="btn btn-light" style={{ padding: '6px 10px' }}><FontAwesomeIcon icon={faChevronLeft} />&nbsp;Back</button>
                                </div>

                                <h2 style={{ margin: '6px 0' }}>{product.title}</h2>
                                <div style={{ fontSize: 20, color: '#16a34a', fontWeight: 700 }}>${product.price.toFixed(2)}</div>

                                <div style={{ display: 'flex', gap: 6, alignItems: 'center', marginTop: 8 }}>
                                    {Array.from({ length: Math.round(product.rating ?? 4) }).map((_, i) => (
                                        <FontAwesomeIcon key={i} icon={faStar} style={{ color: '#f59e0b' }} />
                                    ))}
                                    <div className="muted" style={{ marginLeft: 8 }}>{(product.rating ?? 0).toFixed(1)}</div>
                                </div>

                                <div style={{ marginTop: 12, fontSize: '15px', fontWeight: '600' }}>
                                    {product.stock === 0 ? (
                                        <span style={{ color: '#dc2626' }}>❌ Out of Stock</span>
                                    ) : (
                                        <span style={{ color: '#16a34a' }}>✓ In Stock: {product.stock} available</span>
                                    )}
                                </div>

                                <p className="muted" style={{ marginTop: 12 }}>{product.description}</p>

                                {cartQuantity > 0 && (
                                    <div style={{
                                        marginTop: 12,
                                        padding: '12px 16px',
                                        backgroundColor: '#dcfce7',
                                        border: '2px solid #16a34a',
                                        borderRadius: '8px',
                                        color: '#166534',
                                        fontWeight: '700',
                                        fontSize: '16px',
                                        textAlign: 'center'
                                    }}>
                                        ✓ You have {cartQuantity} of this item in cart
                                    </div>
                                )}

                                <div style={{ display: 'flex', gap: 10, alignItems: 'center', marginTop: 14 }}>
                                    <label className="muted">Quantity</label>
                                    <input 
                                        type="number" 
                                        min={1} 
                                        value={qty} 
                                        onChange={(e) => setQty(Math.max(1, Number(e.target.value || 1)))} 
                                        style={{ width: 80, padding: 8, borderRadius: 8, border: '1px solid #e5e7eb' }} 
                                        disabled={product.stock === 0}
                                    />
                                    {cartQuantity > 0 ? (
                                        <>
                                            <button 
                                                className="btn btn-success" 
                                                onClick={handleAdd}
                                                disabled={product.stock === 0 || cartQuantity + qty > product.stock}
                                                style={{ opacity: (product.stock === 0 || cartQuantity + qty > product.stock) ? 0.6 : 1, cursor: (product.stock === 0 || cartQuantity + qty > product.stock) ? 'not-allowed' : 'pointer' }}
                                            >
                                                Add More
                                            </button>
                                            <button className="btn btn-danger" onClick={handleReduce}>Reduce</button>
                                        </>
                                    ) : (
                                        <button 
                                            className="btn btn-success" 
                                            onClick={handleAdd}
                                            disabled={product.stock === 0}
                                            style={{ opacity: product.stock === 0 ? 0.6 : 1, cursor: product.stock === 0 ? 'not-allowed' : 'pointer' }}
                                        >
                                            {product.stock === 0 ? 'Out of Stock' : 'Add to cart'}
                                        </button>
                                    )}
                                </div>

                                <div className="muted" style={{ marginTop: 14 }}>Category: <strong>{product.category}</strong></div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}