package signal

import (
	"os"
	"os/signal"
	"port/pkg/db"
	"syscall"
)

func Handler(q chan bool, dbHandler db.Handler) {
	repeat := true

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	for signal := range c {
		switch signal {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL:
			repeat = false
		}

		if !repeat {
			dbHandler.Close()
		}
		q <- repeat
	}
}
