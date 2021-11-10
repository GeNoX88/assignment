var whichPage = document.getElementById("whichPage").innerText;
console.log(`目前在第${whichPage}頁`);

function updateHandler(x) {
  var newName = document.getElementById(x).value;
  fetch(`/${whichPage}`, {
    method: "PUT",
    body: JSON.stringify({ WhichPage: whichPage, Number: x, NewName: newName }),
    headers: {
      "content-type": "application/json",
    },
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

  }).then((response) => window.location.href="/1");
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
    }.then((response) => window.location.href="/1")
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
    }.then((response) => window.location.href="/1")
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
    }.then((response) => window.location.href="/1")
  });
}
