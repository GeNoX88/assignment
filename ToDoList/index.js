var whichPage = document.getElementById("whichPage").innerText;
var count = document.getElementById("count").innerText;
console.log(`目前在第${whichPage}頁`);

var pages = Math.floor(((count || 1) - 1) / 4) + 1;
for (i = 0; i < pages; i++) {
  document
    .getElementById("pageDiv")
    .appendChild(
      document.createElement(
        "button",
        (innerText = i + 1),
        (onclick = `javascript:location.href="/${i + 1}"`)
      )
    );
}

function updateHandler(x) {
  var newName = document.getElementById(x).value;
  fetch(`/${whichPage}`, {
    method: "PUT",
    body: JSON.stringify({ WhichPage: whichPage, Number: x, NewName: newName }),
    headers: {
      "content-type": "application/json",
    },
  }).then((response) => {
    window.location.href = "/1";
    console.log(response.status);
  });
}

function delete1Handler() {
  fetch(`/${whichPage}`, {
    method: "DELETE",
    body: JSON.stringify({
      WhichPage: whichPage,
      Number: "1",
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
  fetch(`/${whichPage}`, {
    method: "DELETE",
    body: JSON.stringify({
      WhichPage: whichPage,
      Number: "2",
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
  var newName = document.getElementById(3).value;
  fetch(`/${whichPage}`, {
    method: "DELETE",
    body: JSON.stringify({
      WhichPage: whichPage,
      Number: "3",
      NewName: newName,
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
  var newName = document.getElementById(4).value;
  fetch(`/${whichPage}`, {
    method: "DELETE",
    body: JSON.stringify({
      WhichPage: whichPage,
      Number: "4",
      NewName: newName,
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
