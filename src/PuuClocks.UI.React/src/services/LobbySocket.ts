// import {
//   HttpTransportType,
//   HubConnectionBuilder,
//   LogLevel,
// } from '@microsoft/signalr';

// export const LobbySocket = {
//   connect: (lobbyId: string) => {
//     const connection = new HubConnectionBuilder()
//       .withUrl(import.meta.env.VITE_API_URL + `/join-lobby/${lobbyId}`, {
//         skipNegotiation: true,
//         transport: HttpTransportType.WebSockets,
//       })
//       .configureLogging(LogLevel.Debug)
//       .build();

//     // Autoreconnect
//     connection.onclose(async () => {
//       await connection.startConnection();
//     });

//     return connection;
//   },
// };
