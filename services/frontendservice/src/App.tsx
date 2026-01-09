import { BrowserRouter, Routes, Route } from "react-router-dom";
import SignIn from './pages/SignIn';
import SignUp from './pages/SignUp';
import Navbar from './components/Navbar';
import Footer from './components/Footer';
import Toast from './components/Toast';
import ErrorBoundary from './components/ErrorBoundary';
import ProductsListing from './pages/ProductsListing'
import ShoppingCart from "./pages/ShoppingCart";
import ProductsDetails from "./pages/ProductDetails";
import Payment from './pages/Payment';
import OrderConfirmation from './pages/OrderConfirmation';
import Shipping from './pages/Shipping';
import { CartProvider } from './context/CartContext';
import { ToastProvider } from './context/ToastContext';
// bootstrap CSS is imported in main.tsx so app-specific styles can load after it

function App() {
  return (
    <>
      <ErrorBoundary>
        <BrowserRouter>
          <ToastProvider>
            <CartProvider>
              <Navbar />
              <Routes>
              <Route path="/" element={<ProductsListing />} />
              <Route path="/cart" element={<ShoppingCart />} />
              <Route path="/shipping" element={<Shipping />} />
              <Route path="/checkout" element={<Payment />} />
              <Route path="/order-confirmation" element={<OrderConfirmation />} />
              <Route path="/:id/details" element={<ProductsDetails />} />
              <Route path="/signin" element={<SignIn />} />
              <Route path="/signup" element={<SignUp />} />
            </Routes>
              <Footer />
            </CartProvider>
            <Toast />
          </ToastProvider>
        </BrowserRouter>
      </ErrorBoundary>
    </>
  )
}

export default App
