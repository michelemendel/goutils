package tests

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/michelemendel/goutils/log"
	"go.uber.org/zap"
)

var lg *zap.SugaredLogger

const LOG_LEVEL = "DEBUG"

func init() {
	lg = log.InitWithConsole(LOG_LEVEL)
}

const (
	DEBUG = "DEBUG"
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
)

type logtests struct {
	out   func()
	check func(t *testing.T, res string)
}

var logTests = []logtests{
	{
		out: func() {
			lg = log.SetDebugLevel()
			lg.Debug(DEBUG)
			lg.Info(INFO)
		},
		check: func(t *testing.T, res string) {
			if !(strings.Contains(res, DEBUG) && strings.Contains(res, INFO)) {
				t.Errorf("Expected %s and %s", DEBUG, INFO)
			}
		},
	},
	{
		out: func() {
			lg = log.SetInfoLevel()
			lg.Debug(DEBUG)
			lg.Info(INFO)
		},
		check: func(t *testing.T, res string) {
			if strings.Contains(res, DEBUG) || !strings.Contains(res, INFO) {
				t.Errorf("Expected only %s, but also got %s", INFO, DEBUG)
			}
		},
	},
	{
		out: func() {
			lg = log.SetWarnLevel()
			lg.Debug(DEBUG)
			lg.Info(INFO)
			lg.Warn(WARN)
			lg.Error(ERROR)
		},
		check: func(t *testing.T, res string) {
			if (strings.Contains(res, DEBUG) || strings.Contains(res, INFO)) || !(strings.Contains(res, WARN) && strings.Contains(res, ERROR)) {
				t.Errorf("Expected %s and %s, but also got %s and %s", WARN, ERROR, INFO, DEBUG)
			}
		},
	},
}

func TestLog(t *testing.T) {
	for _, lt := range logTests {
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		lt.out() //call the function that writes to stdout

		w.Close()
		os.Stdout = oldStdout

		var out bytes.Buffer
		io.Copy(&out, r)

		lt.check(t, out.String()) //check the output
	}
}
