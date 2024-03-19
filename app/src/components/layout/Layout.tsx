import React, { ReactNode } from 'react';
import Navbar from '../navBar/Navbar';

interface LayoutProps {
  children: ReactNode;
}

const Layout: React.FC<LayoutProps> = ({ children }) => {
  return (
    <div className='d-flex d-flex-col'>
      <Navbar />
      {children}
    </div>
  );
};

export default Layout;
