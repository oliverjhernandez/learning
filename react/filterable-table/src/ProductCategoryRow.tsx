interface ProductCategoryRowProps {
  category: string;
}

const ProductCategoryRow: React.FC<ProductCategoryRowProps> = ({
  category,
}) => {
  return (
    <tr>
      <th className="category-row" colSpan={2}>
        {category}
      </th>
    </tr>
  );
};

export default ProductCategoryRow;
