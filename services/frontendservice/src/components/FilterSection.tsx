import React, { useEffect, useState } from 'react'

type Filters = {
    query?: string
    categories: string[]
    minPrice?: number
    maxPrice?: number
}

const STORAGE_KEY = 'shop_filters_v1'

const ALL_CATEGORIES = ['Phones', 'Tablets', 'Clothings', 'Machinery']

export default function FilterSection() {
    const [query, setQuery] = useState('')
    const [categories, setCategories] = useState<string[]>([])
    const [minPrice, setMinPrice] = useState<string>('')
    const [maxPrice, setMaxPrice] = useState<string>('')

    useEffect(() => {
        try {
            const raw = localStorage.getItem(STORAGE_KEY)
            if (!raw) return
            const parsed = JSON.parse(raw) as Filters
            setQuery(parsed.query || '')
            setCategories(parsed.categories || [])
            setMinPrice(parsed.minPrice ? String(parsed.minPrice) : '')
            setMaxPrice(parsed.maxPrice ? String(parsed.maxPrice) : '')
        } catch {}
    }, [])

    function toggleCategory(cat: string) {
        setCategories(prev => prev.includes(cat) ? prev.filter(c => c !== cat) : [...prev, cat])
    }

    function applyFilters() {
        const payload: Filters = { query: query || undefined, categories, minPrice: minPrice ? Number(minPrice) : undefined, maxPrice: maxPrice ? Number(maxPrice) : undefined }
        localStorage.setItem(STORAGE_KEY, JSON.stringify(payload))
        window.dispatchEvent(new Event('filters-changed'))
    }

    function clearFilters() {
        setQuery('')
        setCategories([])
        setMinPrice('')
        setMaxPrice('')
        localStorage.removeItem(STORAGE_KEY)
        window.dispatchEvent(new Event('filters-changed'))
    }

    return (
        <div className="filter-section">
            <div className="product-categories-section bg-white p-4">
                <h5>Search</h5>
                <input value={query} onChange={e => setQuery(e.target.value)} placeholder="Search products" className="form-control mb-3" />

                <h5 className="mt-2">Categories</h5>
                <div className="product-categories">
                    {ALL_CATEGORIES.map(cat => (
                        <div key={cat} className="product-category d-flex align-items-center gap-3 border-bottom py-2">
                            <input type="checkbox" checked={categories.includes(cat)} onChange={() => toggleCategory(cat)} className="form-check-input" />
                            <label className="form-check-label">{cat}</label>
                        </div>
                    ))}
                </div>

                <h5 className="mt-3">Price</h5>
                <div style={{ display: 'flex', gap: 8 }}>
                    <input value={minPrice} onChange={e => setMinPrice(e.target.value)} placeholder="Min" className="form-control" />
                    <input value={maxPrice} onChange={e => setMaxPrice(e.target.value)} placeholder="Max" className="form-control" />
                </div>

                <div style={{ display: 'flex', gap: 8, marginTop: 12 }}>
                    <button className="btn btn-success" onClick={applyFilters}>Apply</button>
                    <button className="btn btn-light" onClick={clearFilters}>Clear</button>
                </div>
            </div>
        </div>
    )
}