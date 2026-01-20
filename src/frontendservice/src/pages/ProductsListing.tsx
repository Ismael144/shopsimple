import { useEffect, useState } from "react";
import FilterSection from "../components/FilterSection";
import ProductCard from "../components/ProductCard";
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
    stock: string;
    categories: string[];
    imageUrl: string;
    createdAt: string;
}

interface Pagination {
    totalPages: string;
    currentPage: string;
    totalItems: string;
}

interface Products {
    products: Product[];
    pagination: Pagination;
}

interface ProductFilters {
    query?: string;
    categories?: string[];
    minPrice?: number;
    maxPrice?: number;
}

export default function ProductsListing() {
    const [loading, setLoading] = useState(false)
    const [error, setError] = useState<any>(null)
    const [data, setData] = useState<Products>({} as Products)
    const [pagination, setPagination] = useState<any>({})
    const [filters, setFilters] = useState<ProductFilters | null>(null)
    const [page, setPage] = useState<number>(1)
    const PAGE_SIZE = 10

    useEffect(() => {
        const controller = new AbortController();
        const fetchData = async () => {
            setLoading(true)
            setError(null)

            try {
                let url: string;
                
                // Use Filter API if filters are applied, otherwise use List API
                if (filters && (filters.query || (filters.categories && filters.categories.length > 0) || filters.minPrice !== undefined || filters.maxPrice !== undefined)) {
                    // Build filter query with FilterRequest parameters
                    const params = new URLSearchParams({
                        page: String(page),
                        page_size: String(PAGE_SIZE)
                    })
                    
                    if (filters.query) {
                        params.append('search_string', filters.query)
                    }
                    
                    if (filters.categories && filters.categories.length > 0) {
                        filters.categories.forEach((cat: string) => {
                            params.append('categories', cat)
                        })
                    }
                    
                    if (filters.minPrice !== undefined || filters.maxPrice !== undefined) {
                        if (filters.minPrice !== undefined) {
                            params.append('price_ranges.min.units', String(Math.floor(filters.minPrice)))
                            params.append('price_ranges.min.currency_code', 'USD')
                        }
                        if (filters.maxPrice !== undefined) {
                            params.append('price_ranges.max.units', String(Math.floor(filters.maxPrice)))
                            params.append('price_ranges.max.currency_code', 'USD')
                        }
                    }
                    
                    url = `http://localhost:8080/products/v1/filter?${params.toString()}`
                } else {
                    // Use regular List API
                    url = `http://localhost:8080/products/v1/list?page=${page}&page_size=${PAGE_SIZE}`
                }
                
                const res = await axios.get<any>(url, {
                    signal: controller.signal,
                    headers: {
                        "Access-Control-Allow-Origin": "*"
                    }
                })
                setData(res.data)
                setPagination(res.data.pagination || {})
            } catch (err) {
                if (axios.isAxiosError(err) && err.message !== 'canceled') {
                    setError(err.message)
                } else if (!axios.isAxiosError(err)) {
                    setError("An unexpected error occured...")
                }
            } finally {
                setLoading(false)
            }
        }

        fetchData()

        return () => controller.abort()
    }, [page, filters])

    // Load product filters from localStorage
    useEffect(() => {
        function loadFilters() {
            try {
                const raw = localStorage.getItem('shop_filters_v1')
                if (!raw) { 
                    setFilters(null)
                    setPage(1)
                    return 
                }
                const parsed = JSON.parse(raw) as ProductFilters
                setFilters(parsed)
                setPage(1) // Reset to first page when filters change
            } catch { 
                setFilters(null) 
            }
        }

        loadFilters()
        window.addEventListener('filters-changed', loadFilters)
        return () => window.removeEventListener('filters-changed', loadFilters)
    }, [])


    const totalPages = parseInt(pagination.totalPages || '1') || 1

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
                        {loading && <p>Loading products...</p>}
                        {error && <p style={{ color: 'red' }}>Error: {error}</p>}
                        {!loading && (
                            <>
                                <div className="products-listings">
                                    {data.products?.map(p => (
                                        <ProductCard 
                                            key={p.id} 
                                            id={p.id} 
                                            title={p.name} 
                                            price={p.unitPrice}
                                            image={p.imageUrl}
                                            categories={p.categories}
                                            rating={4} 
                                            stock={parseInt(p.stock)} 
                                        />
                                    ))}
                                </div>
                                <div style={{ marginTop: 18 }}>
                                    <nav aria-label="Pagination">
                                        <ul className="pagination-custom">
                                            <li>
                                                <button onClick={() => setPage(Math.max(1, page - 1))} disabled={page === 1} className="btn btn-light">Prev</button>
                                            </li>
                                            {Array.from({ length: totalPages }).map((_, i) => (
                                                <li key={i}>
                                                    <button onClick={() => setPage(i + 1)} className={i + 1 === page ? 'active' : ''}>{i + 1}</button>
                                                </li>
                                            ))}
                                            <li>
                                                <button onClick={() => setPage(Math.min(totalPages, page + 1))} disabled={page === totalPages} className="btn btn-light">Next</button>
                                            </li>
                                        </ul>
                                    </nav>
                                </div>
                            </>
                        )}
                    </div>
                </div>
            </div>
        </div>
    )
}