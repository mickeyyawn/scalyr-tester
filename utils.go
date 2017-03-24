package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func Print(line string) {
	fmt.Println(line)
}

func Int32ToString(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}

func UUID() string {

	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out)

	return string(out)
}

func HostName() string {
	name, err := os.Hostname()

	if err != nil {
		panic(err)
	}

	return name
}
