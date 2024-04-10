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
    })
    .catch(error => {
        alert(error.message);
    });
}

function setServerConfigs(formData){
    fetch('/setting/modify/servercmd/configs', {
        method: 'POST',
        body: JSON.stringify(Object.fromEntries(formData)),
        headers: {
            'Content-Type': 'application/json'
        }
    })
    .then(response => {
        if (!response.ok) {
            alert("修改失败");
        } else {
            alert("修改成功");
        }
    })
    .then(data => {
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
            var modnums = data["modnums"];
            document.getElementById('fileListTitle').innerHTML += `(${modnums}个)`;
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
            var confignums = data["confignums"];
            document.getElementById('fileListTitle').innerHTML += `(${confignums}个)`;
            for(var i=0;i<config.length;i++){
                document.getElementById('fileList').innerHTML += `${config[i]}<br>`;              
            }
        })
        .catch(error => console.error('Error:', error));
}
function showServerConfigs(){
    fetch('/api/server/setting')
        .then(response => response.json())
        .then(data => {
            var settingClass = Object.keys(data);
            for(var i=0;i<settingClass.length;i++){
                subData = data[settingClass[i]];
                var settingKeys = Object.keys(subData);
                for(var j=0;j<settingKeys.length;j++){
                    contentData = subData[settingKeys[j]];
                    document.getElementById(settingKeys[j]).value = contentData;
                }
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

