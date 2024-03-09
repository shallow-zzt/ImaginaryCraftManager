function startWebSocket(){
    var socket = new WebSocket("ws://localhost:8080/ws/servercmd/status");

    socket.onmessage = function(event) {
        document.getElementById("output").innerHTML += event.data + "<br>";
    };
}