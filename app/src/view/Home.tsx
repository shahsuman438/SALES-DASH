import { useLocation } from "react-router";
import Dashboard from "./Dashboard"
import Products from "./Products"
import { PATHS } from "../routes";

const Home = () => {
  const location = useLocation();

  return (
    <div className="container">
      {location.pathname === PATHS.DASHBOARD_PATH && <Dashboard />}
      {location.pathname === PATHS.ALL_PRODUCTS && <Products />}
    </div>
  )
}

export default Home