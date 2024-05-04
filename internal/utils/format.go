package utils

import (
	"errors"
	"strconv"
	"strings"
)

type Input struct {
	Host     string
	Wordlist string
	Robots   bool
	// Return codes
	StatShow []int
	StatHide []int

	Workers int
}

func Args(args []string) (Input, error) {
	scan := Input{}

	var err error

	for index, element := range args {
		switch element {
		case "-w":
			scan.Wordlist = args[index+1]

		case "-r":
			scan.Robots = true

		case "-S":
			for i := 1; len(args) > index+i && !strings.Contains(args[index+i], "-"); i++ {
				code, _ := strconv.Atoi(args[index+i])
				scan.StatShow = append(scan.StatShow, code)
			}

			if len(scan.StatShow) == 0 {
				scan.StatShow = []int{}
			}

		case "-f":
			for i := 1; len(args) > index+i && !strings.Contains(args[index+i], "-"); i++ {
				code, _ := strconv.Atoi(args[index+i])
				scan.StatHide = append(scan.StatHide, code)

			}
		case "-a":
			scan.Workers, _ = strconv.Atoi(args[index+1])

		}
	}

	scan.Host = args[1]

	if InSclice(args, "--help") || InSclice(args, "help") || InSclice(args, "-h") {
		scan.Host = "help"
	}

	if scan.Host == "" {
		err = errors.New("host")
	}

	if scan.Workers == 0 {
		scan.Workers = 16
	}

	if !InSclice(args, "-S") {
		scan.StatHide = append(scan.StatHide, 403, 404)
	}

	return scan, err
}

func Host(scan Input) Input {
	if strings.Contains(scan.Host, "http") {
		return scan
	}

	scan.Host = "http://" + scan.Host
	return scan
}
