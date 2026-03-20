"use strict";
function indexToPoint(point, gridSize) {
    return point.row * gridSize + point.col;
}
function pointToIndex(index, gridSize) {
    return {
        row: Math.floor(index / gridSize),
        col: index % gridSize,
    };
}
function rotate90degrees(point, gridSize) {
    return {
        row: point.col,
        col: gridSize - 1 - point.row,
    };
}
function getAllRotations(point, gridSize) {
    const r1 = rotate90degrees(point, gridSize);
    const r2 = rotate90degrees(r1, gridSize);
    const r3 = rotate90degrees(r2, gridSize);
    return [r1, r2, r3];
}
function getCenterIndex(gridSize) {
    return Math.floor(gridSize / 2) * gridSize + Math.floor(gridSize / 2);
}
function isOdd(n) {
    return n % 2 !== 0;
}
function getNodeByIndex(index) {
    return table.querySelector(`[data-index="${index}"]`);
}
function applyToRotations(index, fn) {
    const point = pointToIndex(index, gridSize);
    getAllRotations(point, gridSize).forEach(rotation => {
        const node = getNodeByIndex(indexToPoint(rotation, gridSize));
        if (node)
            fn(node);
    });
}
const params = new URLSearchParams(window.location.search);
const nParam = params.get("n");
document.querySelectorAll("#number").forEach(x => x.innerHTML = nParam !== null && nParam !== void 0 ? nParam : "");
const gridSize = Number(nParam);
const required = (gridSize * gridSize - gridSize % 2) / 4;
const centerIndex = getCenterIndex(gridSize);
document.documentElement.style.setProperty("--grid-size", String(gridSize));
let key = [];
const table = document.getElementById("matrix");
for (let i = 0; i < gridSize; i++) {
    const tr = document.createElement("tr");
    for (let j = 0; j < gridSize; j++) {
        const td = document.createElement("td");
        td.classList.add("node");
        td.dataset.index = `${i * gridSize + j}`;
        td.textContent = "X";
        if (isOdd(gridSize) && i === Math.floor(gridSize / 2) && j === Math.floor(gridSize / 2)) {
            td.classList.add("blocked");
        }
        tr.appendChild(td);
    }
    table.appendChild(tr);
}
table.addEventListener("click", (event) => {
    const node = event.target.closest(".node");
    if (!node || node.classList.contains("blocked"))
        return;
    const index = Number(node.dataset.index);
    if (node.classList.toggle("selected")) {
        key.push(index);
        applyToRotations(index, n => n.classList.add("blocked"));
    }
    else {
        key = key.filter(i => i !== index);
        applyToRotations(index, n => n.classList.remove("blocked"));
    }
});
document.getElementById("clear").addEventListener("click", () => {
    key = [];
    table.querySelectorAll(".node").forEach(node => {
        if (isOdd(gridSize) && Number(node.dataset.index) === centerIndex)
            return;
        node.classList.remove("selected", "blocked");
    });
});
document.getElementById("submit").addEventListener("click", async () => {
    const submitBtn = document.getElementById("submit");
    if (key.length !== required) {
        submitBtn.textContent = `Need ${required} cells (${key.length}/${required})`;
        setTimeout(() => submitBtn.textContent = "Submit", 2000);
        return;
    }
    await fetch("/submit", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ key }),
    }).then(() => document.body.innerHTML = '<p>Key submitted! You can close this tab.</p>');
});
