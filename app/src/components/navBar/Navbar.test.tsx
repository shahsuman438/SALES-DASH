// @ts-ignore
import React from 'react';
import '@testing-library/jest-dom';
import { render, screen } from '@testing-library/react';
import { BrowserRouter } from 'react-router-dom';
import Navbar from './Navbar';
import { PATHS } from '../../routes';

describe('Navbar Component', () => {
  test('renders navigation links correctly', () => {
    render(
      <BrowserRouter>
        <Navbar />
      </BrowserRouter>,
    );

    const navItems = [
      { path: PATHS.DASHBOARD_PATH, label: 'Home' },
      { path: PATHS.ALL_PRODUCTS, label: 'Products' },
      { path: PATHS.ALL_SALES, label: 'Sales' },
      { path: PATHS.SALES_BY_PRODUCTS, label: 'Sales by Products' },
      { path: PATHS.SALES_BY_BRANDS, label: 'Sales by Brand' },
    ];

    navItems.forEach((item) => {
      const link = screen.getByRole('link', { name: item.label });
      expect(link).toHaveAttribute('href', item.path);
    });
  });
});
