package parser

import (
	"regexp"
	"strings"
	"time"
)

// LogEntry represents the structured log format
type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Level     string    `json:"level"`
	Source    string    `json:"source"`
	Message   string    `json:"message"`
}

// ParseSyslogLine parses a typical syslog line into a LogEntry
// Example: "Aug 8 22:45:03 hostname process[pid]: message"
func ParseSyslogLine(line, source string) LogEntry {
	// Regex to extract date, time and message
	syslogRegex := regexp.MustCompile(`^(\w+\s+\d+\s+\d+:\d+:\d+)\s+(\S+)\s+(.*)$`)
	matches := syslogRegex.FindStringSubmatch(line)

	ts := time.Now().UTC()
	level := "INFO" // default

	if len(matches) >= 4 {
		// Parse timestamp without year - add current year
		parsedTime, err := time.Parse("Jan 5 15:05:05", matches[1])
		if err == nil {
			ts = parsedTime
			// Adjust year to current
			ts = ts.AddDate(time.Now().Year()-ts.Year(), 0, 0)
		}

		msgLevel := strings.ToLower(matches[3])
		if strings.Contains(msgLevel, "error") {
			level = "ERROR"
		} else if strings.Contains(msgLevel, "warn") {
			level = "WARN"
		} else if strings.Contains(msgLevel, "debug") {
			level = "DEBUG"
		}

		return LogEntry{
			Timestamp: ts,
			Level:     level,
			Source:    source,
			Message:   matches[3],
		}
	}

	// Fallback for unrecognized format
	return LogEntry{
		Timestamp: ts,
		Level:     level,
		Source:    source,
		Message:   line,
	}
}
