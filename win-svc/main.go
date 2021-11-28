package main

import (
	"fmt"
	"log"

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

}
