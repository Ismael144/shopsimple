import React from "react"

export default function Footer() {
    return (
        <footer className="footer bg-white" style={{ borderTop: '1px solid #e6e7eb' }}>
            <div className="container" style={{ display: 'flex', gap: 12, alignItems: 'center', justifyContent: 'space-between', padding: '10px 0' }}>
                <div style={{ display: 'flex', gap: 12, alignItems: 'center' }}>
                    <div style={{ fontWeight: 700 }}>ShopSimple</div>
                    <div className="muted">© {new Date().getFullYear()}</div>
                </div>

                <div className="muted" style={{ fontSize: 14 }}>
                    Built with care — simple e-commerce demo
                </div>
            </div>
        </footer>
    )
}