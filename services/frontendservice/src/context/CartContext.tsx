import React, { createContext, useContext, useEffect, useState } from 'react'

type CartItem = {
  id: string | number
  title: string
  price: number
  qty: number
  image?: string
}

type CartContextValue = {
  items: CartItem[]
  addItem: (item: Omit<CartItem, 'qty'>, qty?: number) => void
  removeItem: (id: string | number) => void
  setQuantity: (id: string | number, qty: number) => void
  clearCart: () => void
  total: number
}

const CartContext = createContext<CartContextValue | undefined>(undefined)

const STORAGE_KEY = 'shop_simple_cart_v1'

export const CartProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [items, setItems] = useState<CartItem[]>(() => {
    try {
      const raw = localStorage.getItem(STORAGE_KEY)
      if (!raw) return []
      return JSON.parse(raw) as CartItem[]
    } catch {
      return []
    }
  })

  useEffect(() => {
    try {
      localStorage.setItem(STORAGE_KEY, JSON.stringify(items))
    } catch {}
  }, [items])

  function addItem(item: Omit<CartItem, 'qty'>, qty = 1) {
    setItems(prev => {
      const existing = prev.find(p => String(p.id) === String(item.id))
      if (existing) {
        return prev.map(p => p.id === existing.id ? { ...p, qty: p.qty + qty } : p)
      }
      return [...prev, { ...item, qty }]
    })
  }

  function removeItem(id: string | number) {
    setItems(prev => prev.filter(p => String(p.id) !== String(id)))
  }

  function setQuantity(id: string | number, qty: number) {
    setItems(prev => prev.map(p => String(p.id) === String(id) ? { ...p, qty: Math.max(0, qty) } : p))
  }

  function clearCart() {
    setItems([])
  }

  const total = items.reduce((s, it) => s + it.price * it.qty, 0)

  return (
    <CartContext.Provider value={{ items, addItem, removeItem, setQuantity, clearCart, total }}>
      {children}
    </CartContext.Provider>
  )
}

export function useCart() {
  const ctx = useContext(CartContext)
  if (!ctx) throw new Error('useCart must be used within CartProvider')
  return ctx
}

export type { CartItem }
