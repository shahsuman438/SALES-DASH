import { useEffect, useState, useCallback } from 'react';
import Table from '../components/table/Table';
import productService from '../services/product.service';
import config from '../config/config';
import { toast } from 'react-toastify';
import { productType } from '../types/types';

const Products = () => {
  const [data, setData] = useState<productType[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  const fetchAllProducts = useCallback(async () => {
    setLoading(true);
    try {
      const res = await productService.getAllProducts();
      setData(res.data);
    } catch (err) {
      console.error(err);
    } finally {
      setLoading(false);
    }
  }, []);

  const handleSync = () => {
    fetchAllProducts();
  };

  useEffect(() => {
    const fetchDataAndSetupSSE = async () => {
      await fetchAllProducts();
      const sse = new EventSource(`${config.apiEndpoint}/events`);

      sse.onmessage = (e) => {
        const message = JSON.parse(e.data).message;
        if (message) toast.info(message);
      };

      sse.onerror = () => {
        console.error('SSE error');
        sse.close();
      };

      return () => {
        sse.close();
      };
    };

    fetchDataAndSetupSSE();
  }, [fetchAllProducts]);

  return (
    <div>
      <div className='flex justify-between items-center'>
        <h1 className='text-blue'>ALL Products</h1>
        <button
          onClick={handleSync}
          disabled={loading}
          className='btn btn--primary'
        >
          {loading ? (
            <span className='mr-2'>Syncing...</span>
          ) : (
            <span className='mr-2'>Sync</span>
          )}
        </button>
      </div>
      {data && <Table data={data} />}
    </div>
  );
};

export default Products;
