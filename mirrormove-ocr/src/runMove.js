const { click, clickMove, keyPress } = require("./service/simulateService");
const { getMove } = require("./service/moveService")

async function executeMove(id, delay) {
    const move = await getMove(id);
    const action = convertMoveToBigAction(move);
    const totalWait = triggerAction(action, delay);
    return totalWait;
}

function convertMoveToBigAction(move) {
    let steps = move.actions
        .map((value) => {
            return [...Array(value.loops)].fill(value.action.steps);
        })
        .flatMap((_) => _)
        .map((action, index, list) => {
            let actionCopy = [...action];
            let hasChanged = false;
            let isFirstAction = index == 0;
            let isFinalAction = !list[index + 1];
            if (isFinalAction == false) {
                actionCopy.pop() // remove the no-op cap in each action end except the last
                hasChanged = true;
            }
            if (!isFirstAction) {
                actionCopy.shift() // remove the inital click in each subsequent action
                hasChanged = true;
            }

            if (hasChanged) return actionCopy;

            return action;
        })
        .reduce(
            (acc, steps, index) => {
                acc.sum += getTotalMoveStepDuration(steps, acc.sum);
                let alteredActions = steps;
                if (index != 0) {
                    alteredActions = steps.map((step) => {
                        return {
                            ...step,
                            time: step.time + acc.sum,
                        };
                    });
                }

                acc.actions.push(alteredActions);
                return acc;
            },
            { actions: [], sum: 0 }
        )
        .actions.flatMap((_) => _)

    return { steps };
}

function getTotalMoveStepDuration(steps, bigSum) {
    return steps.reduce((sum, step, index, list) => {
        const prev = list[index - 1]?.time;
        const curr = step.time;
        if (!prev) return 0;
        const diff = curr - prev;
        sum += diff;
        return sum;
    }, bigSum);
}

function triggerAction(action, milliOffset) {
    let totalWait = milliOffset;
    const timedSteps = action.steps.map((step, index, steps) => {
        const { time, position, act } = getResponseUpdate(
            steps[index - 1],
            step,
            steps[index + 1]
        );
        totalWait += time;
        return {
            step,
            wait: totalWait,
            position,
            act,
        };
    });
    timedSteps.map((ts) => {
        setTimeout(async () => {
            let position = ts?.position;
            let action = ts?.step?.action;
            let key = ts?.step?.key;

            if (action === "key_press") {
                await keyPress(key);
            } else if (action === "click_move") {
                await clickMove(position);
            } else if (action === "click") {
                await click(position);
            }
        }, ts.wait);
    });

    return totalWait;
}

function getResponseUpdate(prevStep, currStep, nextStep) {
    let currTime = currStep.time;
    let prevTime = !prevStep ? currTime : prevStep.time;

    let currPos = { x: currStep.x, y: currStep.y };
    let nextPos = !nextStep ? currPos : { x: nextStep.x, y: nextStep.y };

    let time = currTime - prevTime;
    let position = { x1: currPos.x, y1: currPos.y, x2: nextPos.x, y2: nextPos.y };

    let prevAction = !prevStep ? null : prevStep.action;
    let nextAction = !nextStep ? null : nextStep.action;

    return {
        time,
        position,
        act: {
            prev: prevAction,
            next: nextAction,
        },
    };
}


exports.executeMove = executeMove;