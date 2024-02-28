
var name = "fabric-server/mods";
var xhr = new XMLHttpRequest();
xhr.open("POST", "/api/message", true);
xhr.setRequestHeader("Content-Type", "application/json");
xhr.onreadystatechange = function () {
    if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
        var response = JSON.parse(xhr.responseText);
        document.getElementById("response").innerText = response.text;
    }
};
xhr.send(JSON.stringify({ text: name }));
