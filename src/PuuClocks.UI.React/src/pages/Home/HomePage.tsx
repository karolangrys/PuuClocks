import { Button } from '~/components/ui/Button/Button';
import card from '~/assets/images/PuuClocksCard.png';
import { useTranslation } from 'react-i18next';
import { LobbyService } from '~/services/LobbyService';
import { useNavigate } from 'react-router-dom';

export const HomePage = () => {
  const { t } = useTranslation('common');
  const navigate = useNavigate();

  const handleCreateLobby = () => {
    LobbyService.createLobby()
      .then((response) => {
        navigate(`lobby/${response.data.lobbyID}`);
      })
      .catch((error) => {
        console.log(error);
      });
  };

  return (
    <div className="flex flex-col gap-8 items-center justify-center my-4">
      <img className="max-w-sm" src={card} alt="PuuClocksLogo" />
      <Button variant="outline" onClick={() => handleCreateLobby()}>
        {t('CreateLobby')}
      </Button>
      <Button variant="outline">{t('JoinLobby')}</Button>
    </div>
  );
};
