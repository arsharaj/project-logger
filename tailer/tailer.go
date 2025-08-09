package tailer

import (
	"fmt"
	"log"

	"github.com/hpcloud/tail"
)

// TailFile continuously reads new lines from the given log file
func TailFile(filePath string, lineHandler func(string)) error {
	t, err := tail.TailFile(filePath, tail.Config{
		Follow:    true, // Follow the file as it grows
		ReOpen:    true, // Reopen file if rotated
		MustExist: true, // Fail if file doesn't exist
		Poll:      true, // Use polling instead of inotify for compatibility
	})

	if err != nil {
		return fmt.Errorf("failed to tail file %s: %w", filePath, err)
	}

	log.Printf("started tailing: %s\n", filePath)

	for line := range t.Lines {
		if line.Err != nil {
			log.Printf("error reading line from %s: %v\n", filePath, line.Err)
			continue
		}
		lineHandler(line.Text)
	}

	return nil
}
