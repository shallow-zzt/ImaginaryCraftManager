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
            document.getElementById("loginResult").innerHTML = `登陆失败，账号或密码错误……`
            return response.text()
        } else {
            window.location.href = "/dashboard"
            document.getElementById("loginResult").innerHTML = `登陆成功，<a href='/dashboard'>点击此处跳转</a>`
            return response.text()
        }
    })
    .then(data => {
       // location.reload();
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
        .then(response => showCmdOutput())
        .catch(error => console.error('Error:', error));
}

function stopProcess() {
    fetch('/control/servercmd/stop')
        .then(response => clearCmdOutput())
        .catch(error => console.error('Error:', error));
}

function restartProcess() {
    fetch('/control/servercmd/restart')
        .then(response => {
            clearCmdOutput();
            startProcess();
        })
        .catch(error => console.error('Error:', error));
}

function logoutFunc() {
    fetch('/auth/logout')
    .then(response => location.reload())
        .catch(error => console.error('Error:', error));

}

function refreshFunc() {
    fetch('/auth/logout/refresh')
    .then(response => location.reload())
        .catch(error => console.error('Error:', error));
}
