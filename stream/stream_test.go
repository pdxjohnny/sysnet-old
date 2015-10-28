package stream

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
	"testing"
)

func TestStream(t *testing.T) {
	counter := 0
	cmd := exec.Command("../test/long_running")
	reader, writer := io.Pipe()
	PrepCmd(writer, cmd)
	go OnReadString(reader, func(text string) {
		counter++
		fmt.Println("Output number", counter, text)
	})
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	if counter != 11 {
		panic(errors.New("Did not reach 10"))
	}
}
