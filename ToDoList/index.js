let count = document.getElementById("count").innerText;
let completed = document.getElementById("completed").innerText;
let cpPage = document.getElementById("cpPage").innerText;
let record = document.getElementById("record").innerText;
let whichPage = document.getElementById("whichPage").innerText;
let pages;
if (cpPage === "all") {
  if (parseInt(count) + parseInt(completed) == 0) {
    pages = 1;
  } else pages = Math.floor((parseInt(count) + parseInt(completed)  - 1) / record) + 1;
} else if (cpPage === "ncp") {
  if (count === "0") {
    pages = 1;
  } else pages = Math.floor((count - 1) / record) + 1;
} else if (cpPage === "cp") {
  if (completed === "0") {
    completed = 1;
  }
  pages = Math.floor((completed - 1) / record) + 1;
}

for (i = 0; i < pages; i++) {
  const a = document.createElement("a");
  a.innerText = i + 1;
  a.href = `/${cpPage}/${i + 1}/${record}`;
  a.style = "margin:3px";
  document.getElementById("pageDiv").appendChild(a);
}

function addHandler(x) {
  let myForm = document.getElementById("myForm");
  let formData = new FormData(myForm);
  fetch("/add", {
    method: "POST",
    body: formData,
  }).then(() => {
    window.location.href = `/${cpPage}/1/${x}`;
  });
}

function toggleHandler(x) {
  fetch("/changeState", {
    method: "PUT",
    body: JSON.stringify(x),
    headers: {
      "content-type": "application/json",
    },
  }).then(() => {
    window.location.href = `/${cpPage}/1/${record}`;
  });
}

function updateHandler(x) {
  const newName = document.getElementById(x).value;
  fetch("/changeName", {
    method: "PUT",
    body: JSON.stringify({
      Id: x,
      NewName: newName,
    }),
    headers: {
      "content-type": "application/json",
    },
  }).then(() => {
    window.location.href = `/${cpPage}/${whichPage}/${record}`;
  });
}

function deleteHandler(x) {
  fetch("/deleteTodo", {
    method: "DELETE",
    body: JSON.stringify(x),
    headers: {
      "content-type": "application/json",
    },
  }).then(() => {
    window.location.href = `/${cpPage}/1/${record}`;
  });
}

function recordHandler(x) {
  if (x > 0 && x % 1 === 0) {
    window.location.href = `/${cpPage}/1/${x}`;
  }
}
