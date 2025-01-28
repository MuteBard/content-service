const axios = require("axios");
const dayjs = require('dayjs')
const jwt = require("jsonwebtoken");
const { JWT } = require("../../settings")

const contentServiceHost = "http://localhost:8080";


async function getMove(id) {
    const response = await axios.get(`${contentServiceHost}/move/${id}`);
    const result = parseData(response.data);
    return result[0];
}

function parseActionData(data) {
    return data.map((d) => {
        const { Token } = d;
        const { steps } = jwt.verify(Token, JWT.secretKey);
        return {
            id: d.Id,
            name: d.Name,
            createdAt: formatDate(d.CreatedAt),
            updatedAt: formatDate(d.UpdatedAt),
            isHidden: d.IsHidden,
            description: d.Description,
            seconds: d.Seconds,
            steps
        };
    });
}

function parseData(data) {
    return data.map((d, i) => {
        try {
            const actions = d.Actions.map((al) => {
                return {
                    loops: al.Loops,
                    action: parseActionData([al.Action])[0]
                }
            })
            return {
                id: d.Id,
                name: d.Name,
                createdAt: formatDate(d.CreatedAt),
                updatedAt: formatDate(d.UpdatedAt),
                isHidden: d.IsHidden,
                description: d.Description,
                seconds: d.Seconds,
                actions
            };
        } catch (err) {
            return {
                id: "hidden",
                name: "hidden",
                createdAt: "hidden",
                updatedAt: "hidden",
                isHidden: "hidden",
                description: "hidden",
                seconds: "hidden",
                actions: "hidden",
            };
        }
    });
}

function formatDate(date) {    
    return dayjs(date).format('ddd, MM-DD-YYYY h:mm A');
}

exports.getMove = getMove;

