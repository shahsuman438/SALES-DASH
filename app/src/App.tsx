import './App.css';
import { Routes, Route } from 'react-router-dom';
import { PATHS } from './routes';
import Products from './view/Products';
import Layout from './components/layout/Layout';
import NotFound from './view/NotFound';
import Sales from './view/Sales';
import Dashboard from './view/Dashboard';
import SalesByProduct from './view/reports/SalesByProduct';
import SalesByBrand from './view/reports/SalesByBrand';

function App() {
  return (
    <Layout>
      <Routes>
        <Route path={PATHS.DASHBOARD_PATH} element={<Dashboard />} />
        <Route path={PATHS.ALL_PRODUCTS} element={<Products />} />
        <Route path={PATHS.ALL_SALES} element={<Sales />} />
        <Route path={PATHS.SALES_BY_PRODUCTS} element={<SalesByProduct />} />
        <Route path={PATHS.SALES_BY_BRANDS} element={<SalesByBrand />} />
        <Route path={PATHS.NOT_FOUND} element={<NotFound />} />
      </Routes>
    </Layout>
  );
}

export default App;
