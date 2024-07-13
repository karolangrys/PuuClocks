import { useParams } from 'react-router-dom';

export const LobbyPage = () => {
  const { lobbyId } = useParams();

  return <div>LobbyPage: {lobbyId}</div>;
};
