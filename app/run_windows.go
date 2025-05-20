//go:build windows
// +build windows

package main

import (
	"time"
	"os"
	"github.com/apernet/hysteria/app/v2/cmd"
    "golang.org/x/sys/windows/svc"
)

type hysteriaService struct{}

func (m *hysteriaService) Execute(args []string, r <-chan svc.ChangeRequest, s chan<- svc.Status) (bool, uint32) {
	s <- svc.Status{State: svc.StartPending}

	go func() {
		cmd.Execute()
	}()

	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown
	s <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}

	for c := range r {
		switch c.Cmd {
		case svc.Stop, svc.Shutdown:
			s <- svc.Status{State: svc.StopPending}
			time.AfterFunc(1*time.Second, func() {
				s <- svc.Status{State: svc.Stopped}
				os.Exit(0)
			})
			break
		}
	}

	s <- svc.Status{State: svc.Stopped}
	return false, 0
}

func run() {
	isWinService, err := svc.IsWindowsService()
	if err == nil && isWinService {
		svc.Run("Hysteria", &hysteriaService{})
		return
	}

	cmd.Execute()
}
