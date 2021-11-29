package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

func main() {

	const name = "postgresql-x64-13"

	m, err := mgr.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer m.Disconnect()

	s, err := m.OpenService(name)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	status, err := s.Query()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s status: %v\n", s.Name, status.State)

	if status.State != svc.Running {
		fmt.Printf("starting...\n")
		if err := s.Start(); err != nil {
			log.Fatal(err)
		}

		if err := wait(s, svc.Running); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("running\n")
	}

	fmt.Printf("stopping...\n")
	if err := controlService(s, svc.Stop, svc.Stopped); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("stopped\n")
}

func controlService(s *mgr.Service, c svc.Cmd, to svc.State) error {

	_, err := s.Control(c)
	if err != nil {
		return fmt.Errorf("could not send control=%d: %v", c, err)
	}

	return wait(s, to)
}

func wait(s *mgr.Service, to svc.State) error {

	status, err := s.Query()
	if err != nil {
		return fmt.Errorf("could not retrieve service status: %v", err)
	}

	timeout := time.Now().Add(10 * time.Second)
	for status.State != to {
		if timeout.Before(time.Now()) {
			return fmt.Errorf("timeout waiting for service to go to state=%d", to)
		}
		time.Sleep(300 * time.Millisecond)
		status, err = s.Query()
		if err != nil {
			return fmt.Errorf("could not retrieve service status: %v", err)
		}
	}
	return nil
}
