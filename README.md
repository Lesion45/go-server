### WebSocket
This server can be used as template for your project that using client-server architecture. 
# Usage
The server application defines two types, Client and Hub. The server creates an instance of the Client type for each websocket connection. A Client acts as an intermediary between the websocket connection and a single instance of the Hub type. The Hub maintains a set of registered clients and broadcasts messages to the clients.

The application runs one goroutine for the Hub and two goroutines for each Client. The goroutines communicate with each other using channels. The Hub has channels for registering clients, unregistering clients and broadcasting messages. A Client has a buffered channel of outbound messages. One of the client's goroutines reads messages from this channel and writes the messages to the websocket. The other client goroutine reads messages from the websocket and sends them to the hub.
