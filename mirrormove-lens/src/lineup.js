const { makeFile } = require("./service/fileManager");
const canvasContainer = document.getElementById("canvasContainer");
const canvas = document.createElement("canvas");

const screenWidth = window.innerWidth;
const screenHeight = window.innerHeight;

canvas.width = screenWidth - 50;
canvas.height = screenHeight - 40;
const ctx = canvas.getContext("2d");

canvasContainer.appendChild(canvas);

window.addEventListener("keydown", function (event) {
  if (event.keyCode === 27 || event.key === "Escape") {
    process.exit(1);
  }
});

let isDrawing = false;
let rect = { x: 0, y: 0, width: 0, height: 0, color: "red" };

canvas.addEventListener("mousedown", (event) => {
  isDrawing = true;
  const modal = document.getElementById("modal");
  if (modal.style == "block") isDrawing = false;

  rect.x = event.clientX - canvas.offsetLeft;
  rect.y = event.clientY - canvas.offsetTop;
  rect.width = 0;
  rect.height = 0;
});

canvas.addEventListener("mousemove", (event) => {
  if (isDrawing) {
    rect.width = event.clientX - canvas.offsetLeft - rect.x;
    rect.height = event.clientY - canvas.offsetTop - rect.y;
    redrawCanvas();
  }
});

canvas.addEventListener("mouseup", () => {
  isDrawing = false;
  const modal = document.getElementById("modal");
  modal.style.display = "block";
});

function redrawCanvas() {
  ctx.clearRect(0, 0, canvas.width, canvas.height);

  ctx.strokeStyle = rect.color;
  ctx.strokeRect(rect.x, rect.y, rect.width, rect.height);
}

document.getElementById("saveTitle").addEventListener("click", async () => {
  const title = document.getElementById("selectTitle").value.replace(/ /g, '_');
  const updatedRect = {
    ...rect,
    title,
  };

  const modal = document.getElementById("modal");
  modal.style.display = "none";
  console.log(updatedRect)
  makeFile(`./savedPositions/${title}.json`, JSON.stringify(updatedRect));
  document.getElementById("selectTitle").value = ""

});
