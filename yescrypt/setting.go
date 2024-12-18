package yescrypt

import "log"

type setting struct {
	NLog2 int // Memory cost parameter (as a power of 2)
	R     int // Block size
}

func (s *setting) validate() {
	s.validateNLog2()
	s.validateR()
}

func (s *setting) validateNLog2() {
	switch {
	case s.NLog2 <= 1:
		log.Fatalln("N must be greater than 1")
	case s.NLog2 < 9 || s.NLog2 > 17:
		log.Fatalln("N must be in range 9~17")
	}
}

func (s *setting) validateR() {
	if s.R < 1 || s.R > 32 {
		log.Fatalln("R must be in range 1~32")
	}
}
