package loggerFactory

import (
	"log"
	"log/syslog"
	"os"
	"strings"
)

type Logger struct {
	Context  string
	ReportId string
	testMode bool
	writer   *syslog.Writer
}

//
//
func (l *Logger) buildMessage(prio string, msg string) (logMsg string) {

	msg = strings.ReplaceAll(msg, "\n", "")
	msg = strings.ReplaceAll(msg, "\t", "")
	msg = strings.ReplaceAll(msg, "  ", " ")

	logMsg = "[" + prio + "] [context:" + l.Context + ", reportId:" + l.ReportId + "] " + msg

	return
}

//
//
func (l *Logger) SetTestMode() {

	l.testMode = true
}

//
//
func (l *Logger) Debug(message string) {

	if l.testMode {
		log.Print(l.buildMessage("DEBUG", message))
		return
	}

	err := l.writer.Debug(l.buildMessage("DEBUG", message))
	if err != nil {
		log.Fatal(err)
	}
}

//
//
func (l *Logger) Error(message string) {

	if l.testMode {
		log.Print(l.buildMessage("ERROR", message))
		return
	}

	err := l.writer.Err(l.buildMessage("ERROR", message))
	if err != nil {
		log.Fatal(err)
	}
}

//
//
func (l *Logger) Warning(message string) {

	if l.testMode {
		log.Fatal(l.buildMessage("WARNING", message))
		return
	}
	err := l.writer.Err(l.buildMessage("WARNING", message))
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(3)
}

//
//
func (l *Logger) Emerg(message string) {

	if l.testMode {
		log.Fatal(l.buildMessage("ERROR", message))
		return
	}
	err := l.writer.Err(l.buildMessage("ERROR", message))
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(3)
}

//
//
func New(rsyslogAddr string, app string) *Logger {

	var err error
	logger := Logger{}

	log.Printf("setting up the logger (%s, %s)", rsyslogAddr, app)

	if rsyslogAddr != "" {
		logger.writer, err = syslog.Dial("tcp", rsyslogAddr, syslog.LOG_WARNING|syslog.LOG_DAEMON, app)

		if err != nil {
			log.Fatalf("syslog.Dial(%s, %s) failed: %s", rsyslogAddr, app, err.Error())
		}

		if logger.writer == nil {
			log.Fatalf("syslog.Dial(%s, %s) failed", rsyslogAddr, app)
		}
	} else {
		logger.SetTestMode()
		log.SetOutput(os.Stdout)
		log.SetOutput(os.Stderr)
	}

	return &logger
}
