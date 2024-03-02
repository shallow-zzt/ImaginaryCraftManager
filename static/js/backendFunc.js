function authLogin(formData){
    fetch('/auth/login', {
        method: 'POST',
        body: JSON.stringify(Object.fromEntries(formData)),
        headers: {
            'Content-Type': 'application/json'
        }
    })
    .then(response => {
        if (!response.ok) {
            alert("登陆失败")
        }
        return response.text();
    })
    .then(data => {
        location.reload();
    })
    .catch(error => {
        alert(error.message);
    });
}

function showMods() {      
    fetch('/api/mods')
        .then(response => response.json())
        .then(data => {
            var mods = data["mods"];
            for(var i=0;i<mods.length;i++){
                document.getElementById('fileList').innerHTML += `${mods[i]}<br>`;              
            }
        })
        .catch(error => console.error('Error:', error));
}

function showModConfigs() {     
    fetch('/api/mods/configs')
        .then(response => response.json())
        .then(data => {
            var config = data["configs"];
            for(var i=0;i<config.length;i++){
                document.getElementById('fileList').innerHTML += `${config[i]}<br>`;              
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
