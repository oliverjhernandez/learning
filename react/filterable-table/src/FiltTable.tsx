import React, { useState } from "react";
import "./index.css";
import ProductTable from "./ProductTable";
import SearchBar from "./SearchBar";

export type Product = {
  name: string;
  stocked: boolean;
  price: string;
  category: string;
};

interface FilterableProductTableProps {
  products: Product[];
}

const FilterableProductTable: React.FC<FilterableProductTableProps> = ({
  products,
}) => {
  const [filterText, setFilterText] = useState("");
  const [inStockOnly, setInStockOnly] = useState(false);

  return (
    <div className="container">
      <SearchBar
        filterText={filterText}
        inStockOnly={inStockOnly}
        onFilterTextChange={(newText: string) => {
          setFilterText(newText);
        }}
        onInStockOnlyChange={(newInStock: boolean) => {
          setInStockOnly(newInStock);
        }}
      />
      <ProductTable
        products={products}
        filterText={filterText}
        inStockOnly={inStockOnly}
      />
    </div>
  );
};

export default FilterableProductTable;
