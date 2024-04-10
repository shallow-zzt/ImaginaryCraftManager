var showFileMode = 'none';

cmdOutput.scrollTop = cmdOutput.scrollHeight;

function setDefaultServerConfig(){
  showServerConfigs();
}

document.getElementById('serverGeneralSettingForm').addEventListener('submit', function(event) {
  event.preventDefault();

  var formData = new FormData(event.target);
  console.log(formData);
  setServerConfigs(formData);
})
document.getElementById('serverWorldSettingForm').addEventListener('submit', function(event) {
  event.preventDefault();

  var formData = new FormData(event.target);
  console.log(formData);
  setServerConfigs(formData);
})
document.getElementById('serverNetworkingSettingForm').addEventListener('submit', function(event) {
  event.preventDefault();

  var formData = new FormData(event.target);
  console.log(formData);
  setServerConfigs(formData);
})
document.getElementById('serverPlayerSettingForm').addEventListener('submit', function(event) {
  event.preventDefault();

  var formData = new FormData(event.target);
  console.log(formData);
  setServerConfigs(formData);
})
document.getElementById('serverResourcesPackSettingForm').addEventListener('submit', function(event) {
  event.preventDefault();

  var formData = new FormData(event.target);
  console.log(formData);
  setServerConfigs(formData);
})
document.getElementById('serverAdditionalSettingForm').addEventListener('submit', function(event) {
  event.preventDefault();

  var formData = new FormData(event.target);
  console.log(formData);
  setServerConfigs(formData);
})

function showFormContent(formId) {
    var form = document.getElementById(formId);
    var formData = new FormData(form);
    for (var pair of formData.entries()) {
      console.log(pair[0] + ": " + pair[1]);
    }
}

function switchFileMode2ModsConfig(){ 
    document.getElementById('fileListTitle').innerHTML = `Mods配置文件列表`; 
    document.getElementById('fileList').innerHTML = ``;    
    document.getElementById('fileControl').innerHTML = `			
    <form action="file/mods/config/upload" method="post" enctype="multipart/form-data">
    <input type="file" name="file" />
    <input type="submit" value="Upload" />
    </form>`    
    showModConfigs();
    showFileMode = 'modConfigs'
}

function switchFileMode2Mods(){
    document.getElementById('fileListTitle').innerHTML = `Mod文件列表`;  
    document.getElementById('fileList').innerHTML = ``; 
    document.getElementById('fileControl').innerHTML = `			
    <form action="file/mods/upload" method="post" enctype="multipart/form-data">
    <input type="file" name="file" />
    <input type="submit" value="Upload" />
    </form>`              
    showMods();
    showFileMode = 'mods'
}

function switchFileMode2None(){
    document.getElementById('fileListTitle').innerHTML = ``;  
    document.getElementById('fileList').innerHTML = ``;      
    document.getElementById('fileControl').innerHTML = ``;
    showFileMode = 'none'
}

function showCmdOutput(){
    startWebSocket();
}

function clearCmdOutput(){
    document.getElementById('cmdOutput').innerHTML= ``;
}

setDefaultServerConfig();