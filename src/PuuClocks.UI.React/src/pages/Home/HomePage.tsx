import { Button } from '~/components/ui/Button';
import card from '~/assets/images/PuuClocksCard.png';
import { useTranslation } from 'react-i18next';
import { LobbyService } from '~/services/LobbyService';
import { useNavigate } from 'react-router-dom';
import { Input } from '~/components/ui/Input';
import { useState } from 'react';
import { StorageKeys } from '~/common/consts/storageKeys';

export const HomePage = () => {
  const { t } = useTranslation('common');
  const navigate = useNavigate();

  const [lobbyId, setLobbyId] = useState<string>('');
  const [nick, setNick] = useState<string>(
    JSON.parse(localStorage.getItem(StorageKeys.Nick) ?? '""')
  );

  const saveNickname = () => {
    localStorage.setItem(StorageKeys.Nick, JSON.stringify(nick));
  };

  const handleCreateLobby = () => {
    saveNickname();
    LobbyService.createLobby()
      .then((response) => {
        navigate(`lobby/${response.data.lobbyID}`);
      })
      .catch((error) => {
        console.log(error);
      });
  };

  const handleJoinLobby = () => {
    saveNickname();
    navigate(`lobby/${lobbyId}`);
  };

  return (
    <div className="flex flex-col gap-8 items-center justify-center my-4">
      <img className="max-w-sm" src={card} alt="PuuClocksLogo" />
      <div>
        <small className="text-sm font-medium leading-none">
          {t('Lobby.Nickname')}
        </small>
        <Input
          value={nick}
          onChange={(e) => setNick(e.currentTarget.value)}
          type="text"
          placeholder={'Grzegorz Floryda'}
        />
      </div>
      <div className="flex gap-2">
        <Input
          value={lobbyId}
          onChange={(e) => setLobbyId(e.currentTarget.value)}
          type="text"
          placeholder={t('Lobby.LobbyId')}
        />
        <Button variant="outline" onClick={() => handleJoinLobby()}>
          {t('Lobby.JoinLobby')}
        </Button>
      </div>
      <Button variant="outline" onClick={() => handleCreateLobby()}>
        {t('Lobby.CreateLobby')}
      </Button>
    </div>
  );
};
