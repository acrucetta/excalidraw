# Dummy Excalidraw

Game Logic:

- Shared canvas with drawings (stored in the server)
- Player A draws
  - Sends message to server with drawing
  - The server broadcasts the change to all clients
  - Player B, C... updates its canvas with the messages

Architecture

- Hub processes messages from all registered clients
- Broadcast service sends the message to all the client browsers

## Tasks

- [] Update the chat / message hub to display small drawings in the shared HTML
  - Send points to the HTML, then have the HTML draw them
- [] When a user draws in the HTML, we need to create calls in the Golang Websocket that
  share that JSON to the rest of the clients connected to it
- [] Split the messages into smaller chunks, if we keep sending one large message with all
  points it will fail
