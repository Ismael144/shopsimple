import { useEffect, useMemo, useState } from "react";
import { Link } from "react-router";
import FilterSection from "../components/FilterSection";
import ProductCard from "../components/ProductCard";

export default function ProductsListing() {
    const SAMPLE = useMemo(() => ([
        { id: '1', title: 'High Frequency Earphones', price: '$170.00', image: 'https://htmldemo.net/venezo/venezo/assets/img/product/product1-big.jpg', category: 'Phones', rating: 4, stock: 15 },
        { id: '2', title: 'Wireless Headset', price: '$120.00', image: 'https://htmldemo.net/venezo/venezo/assets/img/product/product2-big.jpg', category: 'Phones', rating: 4, stock: 8 },
        { id: '3', title: 'Bluetooth Speaker', price: '$90.00', image: 'https://htmldemo.net/venezo/venezo/assets/img/product/product3-big.jpg', category: 'Speakers', rating: 5, stock: 0 },
        { id: '4', title: 'Running Shoes', price: '$85.00', image: 'https://htmldemo.net/venezo/venezo/assets/img/product/product4-big.jpg', category: 'Clothings', rating: 4, stock: 25 }
    ]), [])

    const [filters, setFilters] = useState<any>(null)

    useEffect(() => {
        function loadFilters() {
            try {
                const raw = localStorage.getItem('shop_filters_v1')
                if (!raw) { setFilters(null); return }
                setFilters(JSON.parse(raw))
            } catch { setFilters(null) }
        }

        loadFilters()
        window.addEventListener('filters-changed', loadFilters)
        return () => window.removeEventListener('filters-changed', loadFilters)
    }, [])

    function parsePrice(p: string) {
        const n = Number(String(p).replace(/[^0-9.-]+/g, ''))
        return Number.isFinite(n) ? n : 0
    }

    const filtered = useMemo(() => {
        if (!filters) return SAMPLE
        const q = (filters.query || '').toLowerCase()
        return SAMPLE.filter(p => {
            if (filters.categories && filters.categories.length) {
                if (!filters.categories.includes(p.category)) return false
            }
            const price = parsePrice(p.price)
            if (filters.minPrice != null && price < filters.minPrice) return false
            if (filters.maxPrice != null && price > filters.maxPrice) return false
            if (q) {
                return p.title.toLowerCase().includes(q) || (p.category || '').toLowerCase().includes(q)
            }
            return true
        })
    }, [SAMPLE, filters])

    const [page, setPage] = useState<number>(1)
    const perPage = 6
    const totalPages = Math.max(1, Math.ceil(filtered.length / perPage))

    useEffect(() => {
        setPage(1)
    }, [filtered])

    return (
        <div>
            <div style={{ background: 'linear-gradient(90deg,#eef2ff,#f5f7fb)', padding: '40px 0' }}>
                <div className="container" style={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', gap: 20 }}>
                    <div>
                        <h1 style={{ margin: 0 }}>Discover simple, beautiful products</h1>
                        <p className="muted" style={{ marginTop: 8 }}>Curated selection, easy checkout — shop with confidence.</p>
                    </div>
                    <div>
                        <img src="https://htmldemo.net/venezo/venezo/assets/img/product/product1-big.jpg" alt="hero" style={{ width: 220, borderRadius: 12, boxShadow: '0 8px 24px rgba(15,23,42,0.08)' }} />
                    </div>
                </div>
            </div>

            <div className="container my-5">
                <div className="products-listing-grid" style={{ marginBottom: "100px" }}>
                    <FilterSection />
                    <div className="products-listing-section mx-3">
                        <div className="products-listings">
                            {filtered.slice((page-1)*perPage, page*perPage).map(p => (
                                <ProductCard key={p.id} id={p.id} title={p.title} price={p.price} image={p.image} rating={p.rating} stock={p.stock} />
                            ))}
                        </div>
                        <div style={{ marginTop: 18 }}>
                            <nav aria-label="Pagination">
                                <ul className="pagination-custom">
                                    <li>
                                        <button onClick={() => setPage(Math.max(1, page-1))} disabled={page===1} className="btn btn-light">Prev</button>
                                    </li>
                                    {Array.from({ length: totalPages }).map((_, i) => (
                                        <li key={i}>
                                            <button onClick={() => setPage(i+1)} className={i+1===page ? 'active' : ''}>{i+1}</button>
                                        </li>
                                    ))}
                                    <li>
                                        <button onClick={() => setPage(Math.min(totalPages, page+1))} disabled={page===totalPages} className="btn btn-light">Next</button>
                                    </li>
                                </ul>
                            </nav>
                        </div>

                    </div>
                </div>
            </div>
        </div>
    )
}