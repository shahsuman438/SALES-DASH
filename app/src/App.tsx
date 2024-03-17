import './App.css'
import { Routes, Route } from 'react-router-dom'
import Home from './view/Home'
import { PATHS } from './routes'
import Products from './view/Products'
import Layout from './components/layout/Layout'
import NotFound from './view/NotFound'

function App() {
  return (
    <Layout>
      <Routes>
        <Route path={PATHS.DASHBOARD_PATH} element={<Home />} />
        <Route path={PATHS.ALL_PRODUCTS} element={<Products />} />
        <Route path={PATHS.NOT_FOUND} element={<NotFound />} />
      </Routes>
    </Layout>
  )
}

export default App