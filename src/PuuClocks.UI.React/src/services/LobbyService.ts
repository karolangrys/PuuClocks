import { BackendApi } from './BackendApi';

interface CreatedLobby {
  lobbyID: string;
}

export const LobbyService = {
  createLobby: async () => {
    return await BackendApi.post<CreatedLobby>('create-lobby');
  },
};
