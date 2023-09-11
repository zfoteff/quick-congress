package bin

import (
	"fmt"
	"log"
	"os"

	"path/filepath"
)

type QuickCongressLogger struct {
	Log *log.Logger
}

func NewLogger(name string, logfilePath string) *QuickCongressLogger {
	var customLogger *log.Logger

	path, err := filepath.Abs("./logs")
	if err != nil {
		// Exit on error
		fmt.Println("Could not find log output path:", err)
		os.Exit(1)
	}

	logger, err := os.OpenFile(path+"/"+logfilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		// Exit on error
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	customLogger = log.New(logger, "["+name+"] ", log.Ldate|log.Ltime|log.Lshortfile)
	return &QuickCongressLogger{
		customLogger,
	}
}

func (q *QuickCongressLogger) Debug(text string) {
	q.Log.Printf("[DEBUG]\t%s", text)
}

func (q *QuickCongressLogger) Info(text string) {
	q.Log.Printf("[INFO]\t%s", text)
}

func (q *QuickCongressLogger) Warning(text string) {
	q.Log.Printf("[WARNING]\t%s", text)
}

func (q *QuickCongressLogger) Error(text string, err error) {
	q.Log.Fatalf("[ERROR] %s:\n%s", text, err)
}
