const whichPage = document.getElementById("whichPage").innerText;
let count = document.getElementById("count").innerText;
let completed = document.getElementById("completed").innerText;
const cpPage = document.getElementById("cpPage").innerText;
console.log(`目前在第${whichPage}頁`);
const record = document.getElementById("record").innerText;
const add = document.getElementById("add").value;
// window.onload = function () {
//   fetch("/onloadData", {
//     method: "DELETE",
//     body: JSON.stringify({
//       WhichPage: whichPage,
//       Number: "1",
//       cpPage: cpPage,
//     }),
//   })
//     .then((response) => {
//       console.log(response);
//       return response.json();
//     })
//     .then((response) => console.log(response));
// };

let pages;
if (cpPage === "false") {
  if (count === "0") {
    count = 1;
  }
  pages = Math.floor((count - 1) / 4) + 1;
} else {
  if (completed === "0") {
    completed = 1;
  }
  pages = Math.floor((completed - 1) / 4) + 1;
}

if (cpPage === "false") {
  for (i = 0; i < pages; i++) {
    const a = document.createElement("a");
    a.innerText = i + 1;
    a.href = `/ncp/${i + 1}/${record}`;
    a.style = "margin:3px";
    document.getElementById("pageDiv").appendChild(a);
  }
} else if (cpPage === "true") {
  for (i = 0; i < pages; i++) {
    const a = document.createElement("a");
    a.innerText = i + 1;
    a.href = `/cp/${i + 1}/${record}`;
    a.style = "margin:3px";
    document.getElementById("pageDiv").appendChild(a);
  }
}

function toggleHandler(x) {
  fetch("/changeState", {
    method: "PUT",
    body: JSON.stringify({
      Id: x,
    }),
    headers: {
      "content-type": "application/json",
    },
  }).then(() => {
    if (cpPage === "false") {
      window.location.href = `/ncp/1/${record}`;
    } else if (cpPage === "true") {
      window.location.href = `/cp/1/${record}`;
    }
  });
}

function updateHandler(x) {
  const newName = document.getElementById(x).value;
  fetch(`/${whichPage}`, {
    method: "PUT",
    body: JSON.stringify({
      WhichPage: whichPage,
      Number: x,
      NewName: newName,
      cpPage: cpPage,
    }),
    headers: {
      "content-type": "application/json",
    },
  }).then(() => {
    if (cpPage === "false") {
      window.location.href = `/${whichPage}`;
    } else {
      window.location.href = `/completed/${whichPage}`;
    }
  });
}

function deleteHandler(x) {
  fetch("/", {
    method: "DELETE",
    body: JSON.stringify({
      WhichPage: whichPage,
      Number: `${x}`,
      cpPage: cpPage,
    }),
    headers: {
      "content-type": "application/json",
    },
  }).then(() => {
    window.location.href = "/1";
  });
}

function addHandler(record) {
  console.log(record)
  fetch("/add", {
    method: "POST",
    body: JSON.stringify({
      Name: "gggg" 
    }),
    headers: {
      "content-type": "application/json",
    },
  }).then(() => {
    window.location.href = `/ncp/1/${record}`;
  });
}

function recordHandler(x) {
  if (typeof x === "number") {
    if (x <= 0 || x % 1 !== 0) {
      alert("資料筆數需為正整數");
      return;
    }
    if (cpPage === "false") {
      window.location.href = `/ncp/1/${x}`;
    } else if (cpPage === "true") {
      window.location.href = `/cp/1/${x}`;
    }
  }
}
