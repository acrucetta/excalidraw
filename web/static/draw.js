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
        const stroke.JSON.parse(ev.)
      })

      // Start drawing
      canvas.addEventListener('mousedown', e => {
        drawing = true;
        [lastX, lastY] = [e.offsetX, e.offsetY];
      });

      // Draw as the mouse moves
      canvas.addEventListener('mousemove', e => {
        if (!drawing) return;
        drawLine(lastX, lastY, e.offsetX, e.offsetY);
        [lastX, lastY] = [e.offsetX, e.offsetY];
      });

      // Stop drawing
      ['mouseup','mouseout'].forEach(evt =>
        canvas.addEventListener(evt, () => drawing = false)
      );

      // The actual drawing function
      function drawLine(x0, y0, x1, y1) {
        ctx.beginPath();            // start a new drawing path
        ctx.moveTo(x0, y0);         // move “pen” to last point
        ctx.lineTo(x1, y1);         // draw a line to the new point
        ctx.strokeStyle = '#000';   // stroke color
        ctx.lineWidth   = 2;        // line thickness
        ctx.stroke();               // actually paint the line
        ctx.closePath();            // finish this path
      }
    });
