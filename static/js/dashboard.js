var showFileMode = 'none';

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