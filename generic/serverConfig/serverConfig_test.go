package serverConfig

import (
	"fmt"
	"testing"
)

func TestWriteProperty(t *testing.T) {
	WriteProperty("../../fabric-server/server.properties", "max-players", "80")
}
func TestLoadServerConfig(t *testing.T) {
	properties, err := ReadServerConfig("../../fabric-server/server.properties")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//fmt.Println(properties)
	fmt.Println(WriteServerConfig2Json(properties, "../../fabric-server"))
}
