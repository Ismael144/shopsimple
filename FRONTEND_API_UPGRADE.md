# Frontend API Integration - Products Listing Upgrade

## Overview
The frontend has been upgraded to integrate with the new products API that returns enhanced product data with proper pricing and categorization.

## Changes Made

### 1. **ProductsListing Component** (`src/frontendservice/src/pages/ProductsListing.tsx`)

#### Updated Interfaces:
- **Price Interface**: Changed from simple `cents: string` to structured pricing:
  ```typescript
  interface Price {
      currencyCode: string;
      units: string;
      nanos: number;
  }
  ```
- **Product Interface**: Now includes:
  - `id: string` (changed from number)
  - `categories: string[]` (new field)
  - Removed `imageUrl` (not in new API)
  - Changed `unitPrice` to use new Price interface

- **Pagination Interface**: Updated to match API response:
  ```typescript
  interface Pagination {
      totalPages: string;
      currentPage: string;
      totalItems: string;
  }
  ```

#### Price Parsing:
- Implemented `parsePrice(p: Price)` function that:
  - Converts units (major currency units) and nanos (nano units) to decimal value
  - Formula: `units + (nanos / 1_000_000_000)`
  - Example: units=100, nanos=32 → 100.000000032

#### Pagination Handling:
- Updated to parse string values from API: `parseInt(pagination.totalPages || '1')`
- Maintains same pagination UI with Previous/Next buttons

### 2. **ProductCard Component** (`src/frontendservice/src/components/ProductCard.tsx`)

#### New Props:
- `price: Price | string` - Now accepts Price object or string
- `categories?: string[]` - Display product categories

#### New Price Formatting:
- Added `formatPrice()` function using Intl.NumberFormat for proper currency formatting
- Displays currency code from API (e.g., "USD")
- Handles both object and string price formats for backwards compatibility

#### UI Enhancements:
- Categories Display: Shows comma-separated categories below product title
- Price Display: Properly formatted currency with currency code
- Maintains all existing functionality (cart integration, stock status, ratings)

### 3. **API Data Flow**

The component now expects API responses in this format:
```json
{
  "products": [
    {
      "id": "3da2f930-cb0e-42b5-80c1-1f4cbfa49913",
      "name": "some product",
      "description": "some description",
      "unitPrice": {
        "currencyCode": "USD",
        "units": "100",
        "nanos": 32
      },
      "stock": "100",
      "categories": ["gaming", "pc", "rgb"],
      "createdAt": "2026-01-19T17:50:07.030Z"
    }
  ],
  "pagination": {
    "totalPages": "10",
    "currentPage": "1",
    "totalItems": "1"
  }
}
```

## Key Features

✅ **Type-Safe**: Full TypeScript support with proper interfaces
✅ **Currency Formatting**: Proper internationalization with currency codes
✅ **Nano Unit Support**: Handles fractional currency amounts (1/1,000,000,000)
✅ **Category Display**: Products show their categories for better filtering
✅ **Pagination**: Supports string-based pagination from API
✅ **Backwards Compatible**: ProductCard still accepts string prices
✅ **Stock Management**: Shows availability status
✅ **Cart Integration**: Full add-to-cart functionality preserved

## Testing

To test the integration:

1. Start the product service API on `http://localhost:8080`
2. Run the frontend: `npm run dev` in `src/frontendservice/`
3. The products page should display:
   - Product names
   - Categories (gaming, pc, rgb)
   - Properly formatted prices (e.g., $100.00)
   - Stock status
   - Add to cart functionality
   - Pagination controls

## Notes

- The API endpoint: `http://localhost:8080/products/v1/list?page={page}&page_size=10`
- Images are currently using placeholder URLs (imageUrl removed from API)
- Ratings default to 4 stars (can be enhanced with actual API data if provided)
- Category filtering can be implemented in FilterSection component for future enhancement
