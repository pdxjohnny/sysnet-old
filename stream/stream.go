package stream

import (
	"bufio"
	"io"
	"net/http"
	"os/exec"
)

type flushWriter struct {
	f http.Flusher
	w io.Writer
}

// Write writes to the flushWriter and flushes if it can
func (fw *flushWriter) Write(p []byte) (n int, err error) {
	n, err = fw.w.Write(p)
	if fw.f != nil {
		fw.f.Flush()
	}
	return
}

// PrepCmd takes io.Writer and sets the output of the comands
// stdout and err to the writer
func PrepCmd(w io.Writer, cmd *exec.Cmd) {
	fw := flushWriter{w: w}
	if f, ok := w.(http.Flusher); ok {
		fw.f = f
	}
	cmd.Stdout = &fw
	cmd.Stderr = &fw
}

// OnReadBytes calls a function evertime scanner sees new bytes
func OnReadBytes(reader io.Reader, callOnRead func([]byte)) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		callOnRead(scanner.Bytes())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

// OnReadString wraps OnReadBytes with a conversion from bytes to string
func OnReadString(reader io.Reader, callOnRead func(string)) error {
	var callAndConvertToString = func(bytes []byte) {
		callOnRead(string(bytes))
	}
	return OnReadBytes(reader, callAndConvertToString)
}
