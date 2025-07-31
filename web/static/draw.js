window.addEventListener('load', () => {
  const canvas = document.getElementById('draw-canvas');
  const ctx = canvas.getContext('2d');
  const proto = location.protocol === "https:" ? "wss" : "ws";

  const pathParts = window.location.pathname.split("/");
  const roomCode = pathParts[2] | "";
  const ws = new WebSocket(`${proto}://${location.host}/ws?room=${roomCode}`); 


  let drawing = false;
  let lastX = 0, lastY = 0;
  let currentStroke = null;

  ws.addEventListener("open", (event) => {
    console.log()
  })

  // Draw incoming strokes
  ws.addEventListener('message', ev => {
    const segment = JSON.parse(ev.data);
    drawLine(segment.P0.X, segment.P0.Y, segment.P1.X, segment.P1.Y, segment.Color, segment.Width);
  }
  )

  // Start drawing
  canvas.addEventListener('mousedown', e => {
    drawing = true;
    [lastX, lastY] = [e.offsetX, e.offsetY];
    currentStroke = {
      Color: '#000',
      Width: 2,
      PlayerID: 'user123'
    }
  });

  // Draw as the mouse moves
  canvas.addEventListener('mousemove', e => {
    if (!drawing || !currentStroke) return;
    const x = e.offsetX, y = e.offsetY;

    // draw the new segment locally
    drawLine(lastX, lastY, x, y, currentStroke.Color, currentStroke.Width);

    // append to our stroke and send the updated array
    const segment = {
      P0: { X: lastX, Y: lastY },
      P1: { X: x, Y: y},
      Color: currentStroke.Color,
      Width: currentStroke.Width,
      PlayerID: currentStroke.PlayerID
    }
    console.log("sending segment: ", segment)
    ws.send(JSON.stringify(segment));

    lastX = x;
    lastY = y;
  });

  // Stop drawing
  ['mouseup', 'mouseout'].forEach(evt =>
    canvas.addEventListener(evt, () => { drawing = false, currentStroke = null })
  );

  // The actual drawing function
  function drawLine(x0, y0, x1, y1, color, width) {
    ctx.beginPath();
    ctx.moveTo(x0, y0);
    ctx.lineTo(x1, y1);
    ctx.strokeStyle = color;
    ctx.lineWidth = width;
    ctx.stroke();
    ctx.closePath();
  }
});
