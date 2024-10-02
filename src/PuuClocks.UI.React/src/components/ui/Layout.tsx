import { Outlet } from 'react-router-dom';
import { Navbar } from './Navbar';
import { Footer } from './Footer';

export const Layout = () => {
  return (
    <div className="max-w-2xl mx-auto">
      <Navbar />
      <Outlet />
      <Footer />
    </div>
  );
};
