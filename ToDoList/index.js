const whichPage = document.getElementById("whichPage").innerText;
let count = document.getElementById("count").innerText;
let completed = document.getElementById("completed").innerText;
const cpPage = document.getElementById("cpPage").innerText;
console.log(`目前在第${whichPage}頁`);

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
if (cpPage == "false") {
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

if (cpPage == "false") {
  for (i = 0; i < pages; i++) {
    const a = document.createElement("a");
    a.innerText = i + 1;
    a.href = `/${i + 1}`;
    a.style = "margin:3px";
    document.getElementById("pageDiv").appendChild(a);
  }
} else if (cpPage == "true") {
  for (i = 0; i < pages; i++) {
    const a = document.createElement("a");
    a.innerText = i + 1;
    a.href = `/completed/${i + 1}`;
    a.style = "margin:3px";
    document.getElementById("pageDiv").appendChild(a);
  }
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
  }).then((response) => {
    window.location.href = "/1";
    console.log(response.status);
  });
}

let deleteURL = (function () {
  if (cpPage === "false") {
    return `/${whichPage}`;
  } else if (cpPage === "true") {
    return `/completed/${whichPage}`;
  }
})();

console.log(deleteURL);
function delete1Handler() {
  fetch(deleteURL, {
    method: "DELETE",
    body: JSON.stringify({
      WhichPage: whichPage,
      Number: "1",
      cpPage: cpPage,
    }),
    headers: {
      "content-type": "application/json",
    },
  }).then((response) => {
    window.location.href = "/1";
    console.log(response.status);
  });
}
function delete2Handler() {
  fetch(deleteURL, {
    method: "DELETE",
    body: JSON.stringify({
      WhichPage: whichPage,
      Number: "2",
      cpPage: cpPage,
    }),
    redirect: "manual",
    headers: {
      "content-type": "application/json",
    },
  }).then((response) => {
    window.location.href = "/1";
    console.log(response.status);
  });
}
function delete3Handler() {
  fetch(deleteURL, {
    method: "DELETE",
    body: JSON.stringify({
      WhichPage: whichPage,
      Number: "3",
      NewName: newName,
      cpPage: cpPage,
    }),
    redirect: "manual",
    headers: {
      "content-type": "application/json",
    },
  }).then((response) => {
    window.location.href = "/1";
    console.log(response.status);
  });
}
function delete4Handler() {
  fetch(deleteURL, {
    method: "DELETE",
    body: JSON.stringify({
      WhichPage: whichPage,
      Number: "4",
      NewName: newName,
      cpPage: cpPage,
    }),
    redirect: "manual",
    headers: {
      "content-type": "application/json",
    },
  }).then((response) => {
    window.location.href = "/1";
    console.log(response.status);
  });
}
