import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { HomePage } from './pages/Home/HomePage';
import { Layout } from './components/ui/Layout';
import { Suspense } from 'react';
import { LobbyPage } from './pages/Lobby/LobbyPage';
import { Slide, ToastContainer } from 'react-toastify';

import 'react-toastify/dist/ReactToastify.css';
import { ThemeProvider } from '~/context/ThemeProvider';

export const App = () => {
  return (
    <>
      <ThemeProvider defaultTheme="light" storageKey="color-theme">
        <Suspense fallback="loading">
          <BrowserRouter basename="/">
            <Routes>
              <Route element={<Layout />}>
                <Route path="/" element={<HomePage />} />
                <Route path="/lobby/:lobbyId" element={<LobbyPage />} />
              </Route>
            </Routes>
          </BrowserRouter>
        </Suspense>
        <ToastContainer
          position="top-right"
          autoClose={5000}
          hideProgressBar
          newestOnTop={false}
          closeOnClick
          rtl={false}
          pauseOnFocusLoss={false}
          draggable
          pauseOnHover
          theme="colored"
          transition={Slide}
        />
      </ThemeProvider>
    </>
  );
};
