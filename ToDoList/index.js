let count = document.getElementById("count").innerText;
let completed = document.getElementById("completed").innerText;
// let cpPage = document.getElementById("cpPage").innerText;
let record = document.getElementById("record").innerText;
let whichPage = document.getElementById("whichPage").innerText;

window.onload = async function () {
  const res = await fetch("/onload");
  const data = await res.json();
  const cpPage = data.cpPage;
  let pages;
  switch (cpPage) {
    case "all":
      if (parseInt(count) + parseInt(completed) == 0) {
        pages = 1;
      } else {
        pages =
          Math.floor((parseInt(count) + parseInt(completed) - 1) / record) + 1;
      }
      break;
    case "ncp":
      if (count === "0") {
        pages = 1;
      } else {
        pages = Math.floor((count - 1) / record) + 1;
      }
      break;
    case "cp":
      if (completed === "0") {
        completed = 1;
      } else {
        pages = Math.floor((completed - 1) / record) + 1;
      }
  }
  for (i = 0; i < pages; i++) {
    const a = document.createElement("a");
    a.innerText = i + 1;
    a.href = `/${cpPage}/${i + 1}/${record}`;
    a.style = "margin:3px";
    document.getElementById("pageDiv").appendChild(a);
  }
};
function addHandler() {
  let myForm = document.getElementById("myForm");
  let formData = new FormData(myForm);
  fetch("/add", {
    method: "POST",
    body: formData,
    redirect: "manual",
  }).then(() => {
    window.location.href = window.location.href;
  });
}
async function toggleHandler(x) {
  const res = await fetch("/onload");
  const data = await res.json();
  fetch("/changeState", {
    method: "PUT",
    body: JSON.stringify(x),
    headers: {
      "content-type": "application/json",
    },
    redirect: "manual",
  }).then(() => {
    window.location.href = `/${data.cpPage}/1/${data.record}`;
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
    redirect: "manual",
  }).then(() => {
    window.location.href = window.location.href;
  });
}

async function deleteHandler(x) {
  const res = await fetch("/onload");
  const data = await res.json();
  fetch("/deleteTodo", {
    method: "DELETE",
    body: JSON.stringify(x),
    headers: {
      "content-type": "application/json",
    },
  }).then(() => {
    window.location.href = `/${data.cpPage}/1/${record}`;
  });
}

async function recordHandler(x) {
  const res = await fetch("/onload");
  const data = await res.json();
  if (x > 0 && x % 1 === 0) {
    window.location.href = `/${data.cpPage}/1/${x}`;
  }
}
