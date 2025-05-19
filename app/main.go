package main

import (
	"runtime"
	"os"

	"golang.org/x/sys/windows/svc"
	"github.com/apernet/hysteria/app/v2/cmd"
)

type hysteriaService struct{}

func (m *hysteriaService) Execute(args []string, r <-chan svc.ChangeRequest, s chan<- svc.Status) (bool, uint32) {
	s <- svc.Status{State: svc.StartPending}

	go func() {
		cmd.Execute()
	}()

	s <- svc.Status{State: svc.Running, Accepts: svc.AcceptStop | svc.AcceptShutdown}

	for c := range r {
		switch c.Cmd {
		case svc.Stop, svc.Shutdown:
			s <- svc.Status{State: svc.Stopped}
			os.Exit(0)
			break
		}
	}

	s <- svc.Status{State: svc.Stopped}
	return false, 0
}

func main() {
	if runtime.GOOS == "windows" {
		isWinService, err := svc.IsWindowsService()
		if err == nil && isWinService {
			svc.Run("Hysteria", &hysteriaService{})
			return
		}
	}

	cmd.Execute()
}
