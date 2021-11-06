let socket = new WebSocket("ws://127.0.0.1:8080/ws");
console.log("Attempting Connection...");

socket.onopen = (capture) => {
  console.log("Successfully Connected");
  socket.send("Hi from the client!");
};

socket.onclose = (event) => {
  console.log("Socket Closed Connection: ", event);
  socket.send("Client Closed!");
};

socket.onmessage = (e) => {
  const capture = document.getElementById("capture");
  const imgdom = document.getElementById("img");
  let data = e.data;
  console.log(data)
  if (data.size > 1000) {
    let blob = new Blob([data], { type: "image/jpeg" });
    console.log(blob);
    function createImageFromBlob(img, b) {
      imgdom.src = URL.createObjectURL(b);
      imgdom.onload = () => {
        URL.revokeObjectURL(imgdom.src);
      };
    }
    const reader = new FileReader();
    reader.readAsDataURL(blob);
    // console.log(reader.result)
    capture.addEventListener(
      "click",
      createImageFromBlob(imgdom, blob),
      false
    );
  }
};

socket.onerror = (error) => {
  console.log("Socket Error: ", error);
};
function captureHandler() {
  socket.send("Taking selfie");
}
function saveHandler() {
  socket.send("Saving selfie");
}