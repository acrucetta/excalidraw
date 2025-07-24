window.addEventListener('load', () => {
      const canvas = document.getElementById('draw-canvas');
      const ctx    = canvas.getContext('2d');
      const proto = location.protocol === "https:" ? "wss" : "ws";
      const ws = new WebSocket(`${proto}://${location.host}/ws`);

      let drawing  = false;
      let lastX = 0, lastY = 0;

      ws.addEventListener("open", (event) => {
        console.log()
      })

      // Draw incoming strokes
      ws.addEventListener('message', ev => {
        const stroke = JSON.parse(ev.data);
        for (let index = 1; index < stroke.Points.length; index++) {
          const p0 = stroke.Points[index-1];
          const p1 = stroke.Points[index];
          drawLine(p0.X, p0.Y, p1.X, p1.Y, stroke.Color, stroke.Width);
        }
      })

      // Start drawing
      canvas.addEventListener('mousedown', e => {
        drawing = true;
        [lastX, lastY] = [e.offsetX, e.offsetY];
        currentStroke = {
        Points: [{X: lastX, Y:lastY}],
        Color: '#000',
        Widht: 2,
        PlayerID: 'user123'
      }
      });

      // Draw as the mouse moves
      canvas.addEventListener('mousemove', e => {
          if (!drawing) return;
          const x = e.offsetX, y = e.offsetY;

          // draw the new segment locally
          drawLine(lastX, lastY, x, y, currentStroke.Color, currentStroke.Width);

          // append to our stroke and send the updated array
          currentStroke.Points.push({ X: x, Y: y });
          ws.send(JSON.stringify(currentStroke));

          lastX = x;
          lastY = y;
        });

      // Stop drawing
      ['mouseup','mouseout'].forEach(evt =>
        canvas.addEventListener(evt, () => {drawing = false, currentStroke = null})
      );

      // The actual drawing function
      function drawLine(x0, y0, x1, y1, color, width) {
        ctx.beginPath();
        ctx.moveTo(x0, y0);
        ctx.lineTo(x1, y1);
        ctx.strokeStyle = color;
        ctx.lineWidth   = width;
        ctx.stroke();
        ctx.closePath();
      }
    });
