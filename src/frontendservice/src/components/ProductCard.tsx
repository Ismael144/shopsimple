import React, { useState } from "react";
import { Link } from "react-router";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faStar } from "@fortawesome/free-solid-svg-icons";
import { useCart } from "../context/CartContext";
import { useToast } from "../context/ToastContext";

interface Price {
  currencyCode: string;
  units: string;
  nanos: number;
}

type Props = {
  id?: string | number;
  title: string;
  price: Price | string;
  image?: string;
  rating?: number;
  stock?: number;
  categories?: string[];
};

export default function ProductCard({ id, title, price, image, rating = 4, stock = 0, categories = [] }: Props) {
  const [imageLoaded, setImageLoaded] = useState(false)
  const { addItem, items } = useCart()
  const { addToast } = useToast()

  function parsePrice(p: Price | string): number {
    if (typeof p === 'string') {
      const n = Number(String(p).replace(/[^0-9.-]+/g, ''))
      return Number.isFinite(n) ? n : 0
    }
    const units = Number(p.units) || 0;
    const nanos = p.nanos || 0;
    return units + (nanos / 1_000_000_000);
  }

  function formatPrice(p: Price | string): string {
    const amount = parsePrice(p);
    const currency = typeof p === 'object' ? p.currencyCode : 'USD';
    return new Intl.NumberFormat('en-US', {
      style: 'currency',
      currency: currency,
      minimumFractionDigits: 2,
      maximumFractionDigits: 2
    }).format(amount);
  }

  const cartItem = items.find(item => String(item.id) === String(id || title))
  const cartQuantity = cartItem?.qty || 0
  const isOutOfStock = stock === 0

  return (
    <Link to={`/${id}/details`} className="product-card" style={{ textDecoration: "none" }}>
      <div className="card" style={{ position: 'relative', display: 'flex', flexDirection: 'column', height: '100%' }}>
        <div className="card-image" style={{ position: 'relative', backgroundColor: '#f3f4f6', height: '250px' }}>
          <img 
            src={image || "https://htmldemo.net/venezo/venezo/assets/img/product/product1-big.jpg"} 
            alt={title}
            loading="lazy"
            onLoad={() => setImageLoaded(true)}
            onError={() => setImageLoaded(true)}
            style={{ width: '100%', height: '100%', objectFit: 'cover', opacity: imageLoaded ? 1 : 0, transition: 'opacity 0.3s ease-in' }}
          />
          {!imageLoaded && (
            <div style={{ position: 'absolute', top: '50%', left: '50%', transform: 'translate(-50%, -50%)', textAlign: 'center', color: '#9ca3af', fontSize: '12px' }}>Loading image...</div>
          )}
          {cartQuantity > 0 && (
            <div style={{
              position: 'absolute',
              top: '8px',
              right: '8px',
              backgroundColor: '#16a34a',
              color: 'white',
              borderRadius: '50%',
              width: '32px',
              height: '32px',
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'center',
              fontSize: '14px',
              fontWeight: 'bold',
              boxShadow: '0 2px 8px rgba(0,0,0,0.15)'
            }}>
              {cartQuantity}
            </div>
          )}
        </div>
        <div className="card-body" style={{ display: 'flex', flexDirection: 'column', flex: 1 }}>
          <div className="title">{title}</div>
          <div className="muted" style={{ fontSize: '12px', marginBottom: '8px' }}>
            {categories.length > 0 ? categories.join(', ') : 'Uncategorized'}
          </div>
          <div className="price mt-2">{formatPrice(price)}</div>
          <div className="mt-2" style={{ display: "flex", gap: 6, alignItems: "center" }}>
            {Array.from({ length: Math.max(0, Math.min(5, Math.round(rating))) }).map((_, i) => (
              <FontAwesomeIcon key={i} icon={faStar} style={{ color: "#f59e0b" }} />
            ))}
          </div>
          <div style={{ marginTop: '8px', fontSize: '13px', fontWeight: '600' }}>
            {isOutOfStock ? (
              <span style={{ color: '#dc2626' }}>Out of Stock</span>
            ) : (
              <span style={{ color: '#16a34a' }}>In Stock: {stock}</span>
            )}
          </div>
          {cartQuantity > 0 && (
            <div style={{
              marginTop: '8px',
              padding: '8px 12px',
              backgroundColor: '#16a34a',
              color: '#fff',
              borderRadius: '6px',
              fontSize: '13px',
              fontWeight: '700',
              textAlign: 'center',
              border: '2px solid #16a34a'
            }}>
              ✓ {cartQuantity} in cart
            </div>
          )}
          <div className="mt-2" style={{ marginTop: 'auto' }}>
            <button 
              className="btn btn-success" 
              onClick={(e) => { 
                e.preventDefault()
                if (!isOutOfStock) {
                  addItem({ id: id || title, title, price: parsePrice(price), image }, 1)
                  addToast(`${title} added to cart!`, 'success', 2500)
                }
              }}
              disabled={isOutOfStock}
              style={{ opacity: isOutOfStock ? 0.6 : 1, cursor: isOutOfStock ? 'not-allowed' : 'pointer' }}
            >
              {isOutOfStock ? 'Out of Stock' : 'Add To Cart'}
            </button>
          </div>
        </div>
      </div>
    </Link>
  );
}
