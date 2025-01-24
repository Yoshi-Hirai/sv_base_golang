const socket = new WebSocket("ws:localhost:8080/ws");

socket.onopen = () => {
  console.log("Connected to chat room");
};

socket.onmessage = (event) => {
  console.log("Message from server: ", event.data);
  const messageElement = document.createElement("p");
  messageElement.textContent = event.data;
  document.body.appendChild(messageElement);
};

document.getElementById("sendButton").onclick = () => {
  const message = document.getElementById("messageInput").value;
  socket.send(message);
};
