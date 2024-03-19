import { useEffect, useState } from 'react';
import Card from '../components/card/Card';
import reportService from '../services/report.service';

type reportType = {
  mostProfitableProduct: string;
  leastProfitableProduct: string;
  dateOfHighestSales: string;
  dateOfLeastSales: string;
};

const Dashboard = () => {
  const [data, setData] = useState<reportType>();

  useEffect(() => {
    reportService
      .getSummeryReport()
      .then((res) => {
        setData(res.data);
      })
      .catch((err) => {
        console.log(err);
      });
  }, []);

  return (
    <div className='d-flex justify-content-space-between pt-5 p-5 '>
      <Card
        title='Most Profitable Product'
        value={data?.mostProfitableProduct || 'N/A'}
      />
      <Card
        title='Least Profitable Product'
        value={data?.leastProfitableProduct || 'N/A'}
      />
      <Card
        title='Date of Highest Sales'
        value={data?.dateOfHighestSales || 'N/A'}
      />
      <Card
        title='Date of Least Sales'
        value={data?.dateOfLeastSales || 'N/A'}
      />
    </div>
  );
};

export default Dashboard;
