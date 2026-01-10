import React from "react";
import { Link } from "react-router";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faStar } from "@fortawesome/free-solid-svg-icons";
import { useCart } from "../context/CartContext";
import { useToast } from "../context/ToastContext";

type Props = {
  id?: string | number;
  title: string;
  price: string;
  image?: string;
  rating?: number;
  stock?: number;
};

export default function ProductCard({ id, title, price, image, rating = 4, stock = 0 }: Props) {
  const { addItem, items } = useCart()
  const { addToast } = useToast()

  function parsePrice(p: string) {
    const n = Number(String(p).replace(/[^0-9.-]+/g, ''))
    return Number.isFinite(n) ? n : 0
  }

  const cartItem = items.find(item => String(item.id) === String(id || title))
  const cartQuantity = cartItem?.qty || 0
  const isOutOfStock = stock === 0

  return (
    <Link to={`/${id}/details`} className="product-card" style={{ textDecoration: "none" }}>
      <div className="card" style={{ position: 'relative' }}>
        <div className="card-image" style={{ position: 'relative' }}>
          <img src={image || "https://htmldemo.net/venezo/venezo/assets/img/product/product1-big.jpg"} alt={title} />
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
        <div className="card-body">
          <div className="title">{title}</div>
          <div className="muted">Simple short description</div>
          <div className="price mt-2">{price}</div>
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
          <div className="mt-2">
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
