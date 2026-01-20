# Product Filtering Implementation

## Overview
Implemented product filtering using the `Filter` RPC from `product.proto`. The filtering system supports search, category filtering, and price range filtering.

## Changes Made

### 1. FilterSection Component (`src/frontendservice/src/components/FilterSection.tsx`)

**Features:**
- **Search Input**: Filter by product name/description
  - Supports Enter key to apply filters
- **Category Checkboxes**: Select multiple categories to filter
  - Updated with common category options: gaming, electronics, audio, pc, rgb, phones, accessories
  - Formatted labels with proper capitalization
- **Price Range**: Filter by minimum and max price
  - Uses number input fields with currency support
  - Both fields are optional

**Functions:**
- `applyFilters()`: Saves filters to localStorage and triggers update via custom event
- `clearFilters()`: Resets all filters and localStorage
- `toggleCategory()`: Add/remove categories from selection

### 2. ProductsListing Component (`src/frontendservice/src/pages/ProductsListing.tsx`)

**Filter Logic:**
```typescript
// When filters are applied:
if (filters && (filters.query || filters.categories?.length > 0 || filters.minPrice !== undefined || filters.maxPrice !== undefined)) {
    // Use /products/v1/filter endpoint
} else {
    // Use /products/v1/list endpoint
}
```

**API Integration:**
- **List API**: `GET /products/v1/list?page={page}&page_size={page_size}`
  - Used when no filters are applied
  
- **Filter API**: `GET /products/v1/filter?...`
  - Used when any filter is active
  - Supports these parameters:
    - `page`: Current page number
    - `page_size`: Items per page
    - `search_string`: Search query
    - `categories`: Category filters (repeated/array)
    - `price_ranges.min.units`: Minimum price
    - `price_ranges.min.currency_code`: Currency for min price
    - `price_ranges.max.units`: Maximum price
    - `price_ranges.max.currency_code`: Currency for max price

**Features:**
- Resets pagination to page 1 when filters change
- Listens for `filters-changed` event from FilterSection
- Handles both filtered and unfiltered requests seamlessly
- Error handling for canceled and unexpected errors

## Proto Alignment

Aligns with `product.proto` definitions:

```proto
rpc Filter(FilterRequest) returns (ListResponse) {
  option (google.api.http) = {
    get: "/products/v1/filter"
  };
};

message FilterRequest {
  uint32          page          = 1;
  uint32          page_size     = 2;
  string          search_string = 3;
  repeated string categories    = 4;
  PriceRanges     price_ranges  = 5;
}

message PriceRanges {
  common.v1.Money min = 1;
  common.v1.Money max = 2;
}
```

## User Flow

1. **Initial Load**: Shows all products using List API
2. **Apply Filters**: 
   - User selects search query, categories, and/or price range
   - Click "Apply Filters" button
   - ProductsListing detects filter change
   - Switches to Filter API endpoint
   - Displays filtered results
3. **Clear Filters**:
   - Click "Clear All" button
   - All filters reset
   - Back to showing all products

## Example Filter Scenarios

### Search Only
```
GET /products/v1/filter?page=1&page_size=10&search_string=airpods
```

### Category Filter Only
```
GET /products/v1/filter?page=1&page_size=10&categories=audio&categories=electronics
```

### Price Range Only
```
GET /products/v1/filter?page=1&page_size=10&price_ranges.min.units=50&price_ranges.max.units=300&price_ranges.min.currency_code=USD&price_ranges.max.currency_code=USD
```

### Combined Filters
```
GET /products/v1/filter?page=1&page_size=10&search_string=wireless&categories=audio&categories=electronics&price_ranges.min.units=50&price_ranges.max.units=300&price_ranges.min.currency_code=USD&price_ranges.max.currency_code=USD
```

## Storage & State Management

- Filters are persisted in localStorage under key `shop_filters_v1`
- Custom `filters-changed` event triggers filter reload across components
- Pagination resets to page 1 when filters change
- All state is managed locally via React useState hooks

## Testing

To test the filtering:

1. Start the product service
2. Navigate to products listing page
3. Try filtering by:
   - Searching for a product name
   - Selecting categories
   - Setting price ranges
   - Combining multiple filters
4. Verify products update correctly
5. Clear filters to see all products again
