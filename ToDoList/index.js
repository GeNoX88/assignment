let count = document.getElementById("count").innerText;
let completed = document.getElementById("completed").innerText;
const cpPage = document.getElementById("cpPage").innerText;
const record = document.getElementById("record").innerText;

function addHandler(x) {
  let myForm = document.getElementById("myForm");
  let formData = new FormData(myForm);
  fetch("/add", {
    method: "POST",
    body: formData,
  }).then(() => {
    window.location.href = `/ncp/1/${x}`;
  });
}

let pages;
let pageURL;
if (cpPage === "false") {
  if (count === "0") {
    count = 1;
  }
  pages = Math.floor((count - 1) / record) + 1;
  pageURL = "ncp";
} else {
  if (completed === "0") {
    completed = 1;
  }
  pages = Math.floor((completed - 1) / record) + 1;
  pageURL = "cp";
}

for (i = 0; i < pages; i++) {
  const a = document.createElement("a");
  a.innerText = i + 1;
  a.href = `/${pageURL}/${i + 1}/${record}`;
  a.style = "margin:3px";
  document.getElementById("pageDiv").appendChild(a);
}

function toggleHandler(x) {
  fetch("/changeState", {
    method: "PUT",
    body: JSON.stringify({
      Id: x,
      cpPage,
    }),
    headers: {
      "content-type": "application/json",
    },
  }).then(() => {
    window.location.href = `/${pageURL}/1/${record}`;
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
    window.location.href = `/${pageURL}/1/${record}`;
  });
}

function deleteHandler(x) {
  fetch("/deleteTodo", {
    method: "DELETE",
    body: JSON.stringify({
      Id: x,
    }),
    headers: {
      "content-type": "application/json",
    },
  }).then(() => {
    window.location.href = `/${pageURL}/1/${record}`;
  });
}

function recordHandler(x) {
  if (x > 0 && x % 1 === 0) {
    window.location.href = `/${pageURL}/1/${x}`;
  }
}
