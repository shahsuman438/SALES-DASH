// @ts-ignore
import React from 'react';
import { Link, useLocation } from 'react-router-dom';
import { PATHS } from '../../routes';

const Navbar = () => {
  const location = useLocation();

  // Array of navigation items
  const navItems = [
    { path: PATHS.DASHBOARD_PATH, label: 'Home' },
    { path: PATHS.ALL_PRODUCTS, label: 'Products' },
    { path: PATHS.ALL_SALES, label: 'Sales' },
    { path: PATHS.SALES_BY_PRODUCTS, label: 'Sales by Products' },
    { path: PATHS.SALES_BY_BRANDS, label: 'Sales by Brand' },
  ];

  return (
    <nav>
      <ul>
        {navItems.map((item) => (
          <li
            key={item.path}
            className={location.pathname === item.path ? 'selected' : ''}
          >
            <Link to={item.path}>{item.label}</Link>
          </li>
        ))}
      </ul>
    </nav>
  );
};

export default Navbar;
