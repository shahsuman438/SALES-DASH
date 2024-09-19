import { useEffect, useState, useCallback } from 'react';
import Table from '../components/table/Table';
import { salesType } from '../types/types';
import salesService from '../services/sales.service';
import config from '../config/config';
import { toast } from 'react-toastify';

const Sales = () => {
  const [data, setData] = useState<salesType[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  const fetchAllSales = useCallback(async () => {
    setLoading(true);
    try {
      const res = await salesService.getAllSales();
      setData(res.data);
    } catch (err) {
      console.error(err);
    } finally {
      setLoading(false);
    }
  }, []);

  const handleSync = () => {
    fetchAllSales();
  };

  useEffect(() => {
    fetchAllSales();

    const sse = new EventSource(`${config.apiEndpoint}/events`);

    sse.onmessage = (e) => {
      const message = JSON.parse(e.data).message;
      if (message) toast.info(message);
    };

    sse.onerror = () => {
      console.error('error');
      sse.close();
    };

    return () => {
      sse.close();
    };
  }, []);

  return (
    <div>
      <div className='flex justify-between items-center'>
        <h1 className='text-blue'>ALL Sales</h1>
        <button
          onClick={handleSync}
          disabled={loading}
          className='btn btn--primary ml-1'
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

export default Sales;
