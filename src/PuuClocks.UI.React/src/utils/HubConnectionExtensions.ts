// import { HubConnection } from '@microsoft/signalr';

// // Declare the Extension
// declare module '@microsoft/signalr' {
//   interface HubConnection {
//     startConnection(): Promise<void>;
//   }
// }

// async function startConnection(this: HubConnection) {
//   try {
//     await this.start();
//     console.log('connected');
//   } catch (err) {
//     setTimeout(() => this.startConnection(), 5000);
//   }
// }

// // Implement the Extension
// HubConnection.prototype.startConnection = startConnection;
