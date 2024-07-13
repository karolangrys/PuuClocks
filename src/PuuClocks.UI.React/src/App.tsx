import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { HomePage } from './pages/Home/HomePage';
import { Layout } from './components/ui/Layout/Layout';
import { Suspense } from 'react';
import { LobbyPage } from './pages/Lobby/LobbyPage';

export const App = () => {
  return (
    <>
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
    </>
  );
};
