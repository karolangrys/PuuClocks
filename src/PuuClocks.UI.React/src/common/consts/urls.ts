export const WS_BASE_URL = import.meta.env.VITE_WS_URL;

export const Urls = {
  connectLobby: (lobbyId: string, nickname: string) =>
    `${WS_BASE_URL}/join-lobby/${lobbyId}?nickname=${nickname}`,
};
