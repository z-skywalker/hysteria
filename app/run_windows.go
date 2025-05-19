//go:build windows
// +build windows

package main

import (
    "log"
	"github.com/apernet/hysteria/app/v2/cmd"
    "golang.org/x/sys/windows/svc"
)

type hysteriaService struct{}

func (m *hysteriaService) Execute(args []string, r <-chan svc.ChangeRequest, s chan<- svc.Status) (bool, uint32) {
	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown
	s <- svc.Status{State: svc.StartPending}
	log.Print("start pending")

	go func() {
		log.Print("begin to execute")
		cmd.Execute()
		log.Print("end of execute")
	}()

	s <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}
	log.Print("running")

	for c := range r {
		switch c.Cmd {
		// case svc.Interrogate:
		// 	log.Print("Interrogate")
		// 	s <- c.CurrentStatus
		case svc.Stop, svc.Shutdown:
			s <- svc.Status{State: svc.StopPending}
			log.Print("StopPending")
			// time.AfterFunc(1*time.Second, func() {
			// 	s <- svc.Status{State: svc.Stopped}
			// 	log.Print("Stopped")
			// 	os.Exit(0)
			// })
			break
		}
	}

	s <- svc.Status{State: svc.Stopped}
	log.Print("exit")
	return false, 0
}

func run() {
    log.Println("running as Windows service")

	isWinService, err := svc.IsWindowsService()
	if err == nil && isWinService {
		err := svc.Run("HysteriaClientService", &hysteriaService{})
		if err != nil {
			log.Fatalf("failed to run service: %v", err)
		}
		return
	}

	cmd.Execute()
}
