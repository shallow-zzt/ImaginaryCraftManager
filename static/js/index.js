
function showMods() {
    fetch('/api/mods')
        .then(response => response.json())
        .then(data => {
            var mods = data["mods"];
            for(var i=0;i<mods.length;i++){
                document.getElementById('modsList').innerHTML += `${mods[i]}<br>`;              
            }
        })
        .catch(error => console.error('Error:', error));
}

function startProcess() {
    fetch('/control/servercmd/start')
        .then(response => response.text())
        .then(data => {
            document.getElementById('serverCmd').innerHTML += data;
        })
        .catch(error => console.error('Error:', error));
}

function stopProcess() {
    fetch('/control/servercmd/stop')
        .then(response => response.text())
        .then(data => {
            document.getElementById('serverCmd').innerHTML += data;
        })
        .catch(error => console.error('Error:', error));
}

function restartProcess() {
    fetch('/control/servercmd/restart')
        .then(response => response.text())
        .then(data => {
            document.getElementById('serverCmd').innerHTML += data;
        })
        .catch(error => console.error('Error:', error));
}

showMods();