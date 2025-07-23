# Dummy Excalidraw

Game Logic:

- Shared canvas with drawings (stored in the server)
- Player A draws
  - Sends message to server with drawing
  - The server broadcasts the change to all clients
  - Player B, C... updates its canvas with the messages

Messages

- 'draw' -> point (x,y) to point(x,y)
{
  "type" : "draw"
  "points" : [
    [1,1], [1,2] ...
  ]
}
- 'erase' -> point (x,y) to point(x,y)

Architecture
Hub processes messages from all registered clients
Broadcast service sends the message to all the client browsers

Tasks:

- [] Update the chat / message hub to display small drawings in the shared HTML
  - Send points to the HTML, then have the HTML draw them

References:

- Websocket Chat Example: <https://github.com/gorilla/websocket/tree/main/examples/chat>
