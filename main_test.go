package logger

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestError(t *testing.T) {
	file := directory + "/error_" + time.Now().Format("2006_01_02") + ".log"
	if _, err := os.Stat(file); !os.IsNotExist(err) {
		os.Remove(file)
	}
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := errors.New("Is Error")
	tag := "The tag"
	msg := fmt.Sprintf("[%s] ERROR/%s: %s\n", time.Now().Format("2006/01/02 15:04:05"), tag, err.Error())
	Error(tag, err)

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = old

	got := []byte(`

========

` + msg + `

========

`)
	if !reflect.DeepEqual(got, out) {
		t.Errorf("console returned unexpected: got %s want %s",
			got, out)
	}

	if b, _ := ioutil.ReadFile(file); !reflect.DeepEqual(b, []byte(msg)) {
		t.Errorf("file content unexpected: got %s want %s",
			b, msg)
	}
}
