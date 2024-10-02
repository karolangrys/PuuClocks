import { useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { toast } from 'react-toastify';
import useWebSocket, { ReadyState } from 'react-use-websocket';
import { StorageKeys } from '~/common/consts/storageKeys';
import { Urls } from '~/common/consts/urls';

export const LobbyPage = () => {
  const { lobbyId } = useParams();
  const navigate = useNavigate();

  const [messageHistory, setMessageHistory] = useState<MessageEvent<any>[]>([]);

  const [nick, setNick] = useState<string>(
    JSON.parse(localStorage.getItem(StorageKeys.Nick) ?? '""')
  );

  console.log(nick);
  console.log('dupa');

  const { sendMessage, lastMessage, readyState } = useWebSocket(
    Urls.connectLobby(lobbyId as string, nick)
  );

  useEffect(() => {
    if (lastMessage !== null) {
      setMessageHistory((prev) => prev.concat(lastMessage));
    }
  }, [lastMessage]);

  const connectionStatus = {
    [ReadyState.CONNECTING]: 'Connecting',
    [ReadyState.OPEN]: 'Open',
    [ReadyState.CLOSING]: 'Closing',
    [ReadyState.CLOSED]: 'Closed',
    [ReadyState.UNINSTANTIATED]: 'Uninstantiated',
  }[readyState];

  useEffect(() => {
    if (!lobbyId) {
      toast.error('No lobbyID!');
      navigate('/');
      return;
    }
  }, []);

  return (
    <div className="flex flex-col">
      <span>LobbyPage: {lobbyId}</span>
      <span>The WebSocket is currently {connectionStatus}</span>
      {lastMessage ? <span>Last message: {lastMessage.data}</span> : null}
      <ul>
        {messageHistory.map((message, idx) => (
          <span key={idx}>{message ? message.data : null}</span>
        ))}
      </ul>
    </div>
  );
};
