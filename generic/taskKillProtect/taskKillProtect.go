package taskKillProtect

import (
	"ImaginaryCraftManager/generic/serverCmd"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func InterruptFunction(cm *serverCmd.CommandManager) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT)

	go func() {
		<-sigCh
		serverCmd.CloseProcessAndPipe(cm)
		fmt.Println("关闭运行的java服务器进程中……")

	}()

	select {}
}
