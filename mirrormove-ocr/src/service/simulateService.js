const { mouse, straightTo, keyboard, Point, Button } = require("@nut-tree-fork/nut-js");

async function click(position) {

	const point = new Point(position.x1, position.y1);
	await mouse.setPosition(point);
	await mouse.click(Button.LEFT);
}

async function clickMove(position) {
    const startPoint = { x: position.x1, y: position.y1 };
    const endPoint = { x: position.x2, y: position.y2 }; 
  
    await mouse.move(straightTo(startPoint));
    await mouse.pressButton(Button.LEFT);
    await mouse.move(straightTo(endPoint));
    await mouse.releaseButton(Button.LEFT);
}

async function keyPress(key) {
    await keyboard.type(key);
}

exports.click = click;
exports.clickMove = clickMove;
exports.keyPress = keyPress;