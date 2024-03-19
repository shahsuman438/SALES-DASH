import { useEffect, useState } from 'react';
import Table from '../components/table/Table';
import { salesType } from '../types/types';
import salesService from '../services/sales.service';

const Sales = () => {
  const [data, setData] = useState<salesType[]>([]);
  useEffect(() => {
    salesService
      .getAllSales()
      .then((res) => {
        setData(res.data);
      })
      .catch((err) => {
        console.log(err);
      });
  }, []);
  return (
    <div>
      <h1 className='text-blue'>ALL Sales</h1>
      <Table data={data} />
    </div>
  );
};

export default Sales;
