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

- [] Drawing Rooms
  - Have the user elect a room with some code, then enter to that drawing room
- [] Send all canvas code upon first registering
  - Upon first signing in, the user will get the latest canvas code
  - He will then be able to add his updates to the canvas
- [] Select color and width as some toolbar
- [] Clear canvas command
