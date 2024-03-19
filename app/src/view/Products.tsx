import { useEffect, useState } from 'react';
import Table from '../components/table/Table';
import productService from '../services/product.service';

type productType = {
  productId: number;
  productName: string;
  brandName: string;
  costPrice: number;
  sellingPrice: number;
  category: string;
  ExpiryDate: string;
};
const Products = () => {
  const [data, setData] = useState<productType[]>([]);

  useEffect(() => {
    productService
      .getAllProducts()
      .then((res) => {
        console.log(res);
        setData(res.data);
      })
      .catch((err) => {
        console.log(err);
      });
  }, []);
  return (
    <div>
      <h1 className='text-blue'>ALL Products</h1>
      <Table data={data} />
    </div>
  );
};

export default Products;
