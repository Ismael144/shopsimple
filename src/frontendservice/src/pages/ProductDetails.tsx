import React, { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router";
import FilterSection from "../components/FilterSection";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faStar, faChevronLeft } from "@fortawesome/free-solid-svg-icons";
import { useCart } from "../context/CartContext";
import { useToast } from "../context/ToastContext";
import axios from "axios";

interface Price {
  currencyCode: string;
  units: string;
  nanos: number;
}

interface Product {
    id: string;
    name: string;
    description: string;
    unitPrice: Price;
    rating: number;
    imageUrl: string;
    stock: string;
    categories: string[];
    createdAt: string;
}

export default function ProductDetails() {
    const { id } = useParams()
    const [product, setProduct] = useState<Product | null>(null)
    const [loading, setLoading] = useState(true)
    const [qty, setQty] = useState<number>(1)
    const navigate = useNavigate()
    const { addItem, items, setQuantity } = useCart()
    const { addToast } = useToast()
    
    const cartItem = items.find(item => String(item.id) === String(product?.id))
    const cartQuantity = cartItem?.qty || 0
    const stock = product ? parseInt(product.stock) : 0

    useEffect(() => {
        const fetchProduct = async () => {
            try {
                setLoading(true)
                const res = await axios.get<Product>(`http://localhost:8080/products/v1?id=${id}`)
                setProduct(res.data)
            } catch (err) {
                addToast('Failed to load product', 'error', 2500)
                navigate('/')
            } finally {
                setLoading(false)
            }
        }

        if (id) {
            fetchProduct()
        }
    }, [id, navigate, addToast])

    function parsePrice(p: Price): number {
        const units = Number(p.units) || 0;
        const nanos = p.nanos || 0;
        return units + (nanos / 1_000_000_000);
    }

    function formatPrice(p: Price): string {
        const amount = parsePrice(p);
        return new Intl.NumberFormat('en-US', {
            style: 'currency',
            currency: p.currencyCode,
            minimumFractionDigits: 2,
            maximumFractionDigits: 2
        }).format(amount);
    }

    function handleAdd() {
        if (!product) return
        if (cartQuantity + qty > stock) {
            addToast(`Only ${stock - cartQuantity} items available!`, 'error', 2500)
            return
        }
        addItem({ id: product.id, title: product.name, price: parsePrice(product.unitPrice), image: product.imageUrl }, qty)
        addToast(`${qty} ${product.name} added to cart!`, 'success', 2500)
    }

    function handleReduce() {
        if (!product) return
        if (cartQuantity > 0) {
            const newQty = cartQuantity - qty
            if (newQty <= 0) {
                setQuantity(product.id, 0)
            } else {
                setQuantity(product.id, newQty)
            }
        }
    }

    if (loading) {
        return (
            <div className="container my-5">
                <div style={{ textAlign: 'center', padding: '40px 0' }}>
                    <p>Loading product...</p>
                </div>
            </div>
        )
    }

    if (!product) {
        return (
            <div className="container my-5">
                <div style={{ textAlign: 'center', padding: '40px 0' }}>
                    <p>Product not found</p>
                </div>
            </div>
        )
    }

    return (
        <div className="container my-5">
            <div className="products-listing-grid">
                <FilterSection />
                <div style={{ padding: 12 }}>
                    <div className="card" style={{ padding: 18 }}>
                        <div style={{ display: 'grid', gridTemplateColumns: '1fr 380px', gap: 20 }}>
                            <div>
                                <img src={product.imageUrl} alt={product.name} style={{ width: '100%', borderRadius: 8 }} />
                            </div>
                            <div>
                                <div style={{ display: 'flex', gap: 12, alignItems: 'center', marginBottom: 8 }}>
                                    <button onClick={() => navigate(-1)} className="btn btn-light" style={{ padding: '6px 10px' }}><FontAwesomeIcon icon={faChevronLeft} />&nbsp;Back</button>
                                </div>

                                <h2 style={{ margin: '6px 0' }}>{product.name}</h2>
                                <div style={{ fontSize: 20, color: '#16a34a', fontWeight: 700 }}>{formatPrice(product.unitPrice)}</div>

                                <div style={{ display: 'flex', gap: 6, alignItems: 'center', marginTop: 8 }}>
                                    {Array.from({ length: Math.max(0, Math.min(5, product.rating)) }).map((_, i) => (
                                        <FontAwesomeIcon key={i} icon={faStar} style={{ color: "#f59e0b" }} />
                                    ))}
                                    <div className="muted" style={{ marginLeft: 8 }}>{product.rating.toFixed(1)}</div>
                                </div>

                                <div style={{ marginTop: 12, fontSize: '15px', fontWeight: '600' }}>
                                    {stock === 0 ? (
                                        <span style={{ color: '#dc2626' }}>❌ Out of Stock</span>
                                    ) : (
                                        <span style={{ color: '#16a34a' }}>✓ In Stock: {stock} available</span>
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
                                        disabled={stock === 0}
                                    />
                                    {cartQuantity > 0 ? (
                                        <>
                                            <button 
                                                className="btn btn-success" 
                                                onClick={handleAdd}
                                                disabled={stock === 0 || cartQuantity + qty > stock}
                                                style={{ opacity: (stock === 0 || cartQuantity + qty > stock) ? 0.6 : 1, cursor: (stock === 0 || cartQuantity + qty > stock) ? 'not-allowed' : 'pointer' }}
                                            >
                                                Add More
                                            </button>
                                            <button className="btn btn-danger" onClick={handleReduce}>Reduce</button>
                                        </>
                                    ) : (
                                        <button 
                                            className="btn btn-success" 
                                            onClick={handleAdd}
                                            disabled={stock === 0}
                                            style={{ opacity: stock === 0 ? 0.6 : 1, cursor: stock === 0 ? 'not-allowed' : 'pointer' }}
                                        >
                                            {stock === 0 ? 'Out of Stock' : 'Add to cart'}
                                        </button>
                                    )}
                                </div>

                                <div className="muted" style={{ marginTop: 14 }}>
                                    Categories: <strong>{product.categories.length > 0 ? product.categories.join(', ') : 'Uncategorized'}</strong>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}