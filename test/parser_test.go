package test

import (
	"testing"

	"github.com/arsharaj/project-logger/parser"
)

var source string = "/var/log/testlog"

func TestParseSyslogLineInfo(t *testing.T) {
	line := "Aug 9 10:12:01 arsharaj project-logger-test[123]: Service started successfully"

	entry := parser.ParseSyslogLine(line, source)

	if entry.Level != "INFO" {
		t.Errorf("expected level INFO, got %s", entry.Level)
	}

	if entry.Source != source {
		t.Errorf("expected source %s, got %s", source, entry.Source)
	}

	if entry.Message == "" {
		t.Errorf("expected non-empty message")
	}

	if entry.Timestamp.Month() != 8 || entry.Timestamp.Day() != 9 {
		t.Errorf("expected timestamp month - day to match log, got %v", entry.Timestamp)
	}
}

func TestParseSyslogLineError(t *testing.T) {
	line := "Aug 9 10:12:01 arsharaj project-logger-test[123]: Error starting service"

	entry := parser.ParseSyslogLine(line, source)

	if entry.Level != "ERROR" {
		t.Errorf("expected level ERROR, got %s", entry.Level)
	}
}

func TestParseSyslogLineWarn(t *testing.T) {
	line := "Aug 9 10:12:01 arsharaj project-logger-test[123]: Warning low memory consumption"

	entry := parser.ParseSyslogLine(line, source)

	if entry.Level != "WARN" {
		t.Errorf("expected level WARN, got %s", entry.Level)
	}
}

func TestParseSyslogLineUnrecognized(t *testing.T) {
	line := "This is not a syslog line"

	entry := parser.ParseSyslogLine(line, source)

	if entry.Message != line {
		t.Errorf("expected raw message to match input, got %s", entry.Message)
	}

	if entry.Level != "INFO" {
		t.Errorf("expected default level INFO, got %s", entry.Level)
	}
}
