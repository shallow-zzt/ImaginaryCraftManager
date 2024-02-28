package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	// 设置路由
	http.HandleFunc("/api/message", handleMessage)
	http.HandleFunc("/api/file", handleMessage)

	http.Handle("/", http.FileServer(http.Dir("static")))

	// 启动Web服务器
	http.ListenAndServe(":8080", nil)
	fmt.Println("running……")
}

func getAllFileNames(dirPath string) ([]string, error) {
	var fileNames []string

	// 读取目录
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	// 遍历目录下的文件和子目录
	for _, file := range files {
		if file.IsDir() {
			subDirPath := filepath.Join(dirPath, file.Name())
			subFileNames, err := getAllFileNames(subDirPath)
			if err != nil {
				return nil, err
			}
			fileNames = append(fileNames, subFileNames...)
		} else {
			fileNames = append(fileNames, file.Name())
		}
	}

	return fileNames, nil
}

// handleMessage 处理从前端发送的消息
func handleMessage(w http.ResponseWriter, r *http.Request) {
	// 从请求中解析JSON数据
	decoder := json.NewDecoder(r.Body)
	var message Message
	var filePaths string

	err := decoder.Decode(&message)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	fileNames, err := getAllFileNames(message.Text)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, fileName := range fileNames {
		filePaths = filePaths + "\n" + fileName
	}

	response := Message{Text: filePaths}

	// 将响应数据转换为JSON格式并写回到客户端
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
