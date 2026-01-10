import React from "react"
import { Link } from "react-router";
import { useCart } from "../context/CartContext";

export default function Navbar() {
    const { items } = useCart()
    const cartCount = items.reduce((sum, item) => sum + item.qty, 0)

    return (
        <nav className="navbar">
            <div className="container nav-inner">
                <div className="brand"><Link to="/">ShopSimple</Link></div>
                <div className="nav-links">
                    <Link to="/">Home</Link>
                    <Link to="/">Products</Link>
                    <Link to="/cart" style={{ position: 'relative', display: 'inline-block' }}>
                        Cart
                        {cartCount > 0 && (
                            <span style={{
                                position: 'absolute',
                                top: '-8px',
                                right: '-12px',
                                backgroundColor: '#dc3545',
                                color: 'white',
                                borderRadius: '50%',
                                width: '20px',
                                height: '20px',
                                display: 'flex',
                                alignItems: 'center',
                                justifyContent: 'center',
                                fontSize: '12px',
                                fontWeight: 'bold'
                            }}>
                                {cartCount}
                            </span>
                        )}
                    </Link>
                    <Link to="/signin">Sign In</Link>
                </div>
            </div>
        </nav>
    );
}