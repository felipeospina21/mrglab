package logger

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// Logs to debug.log file
//
//	logger.Debug("log", func() {
//		log.Println(strconv.Itoa(msg.Width))
//		log.Println("tw " + strconv.Itoa(m.table.Width()))
//	})
func Debug(logPrefix string, cb func()) {
	f, err := tea.LogToFile("debug.log", logPrefix)
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()
	cb()
}

func Error(error error) {
	f, err := tea.LogToFile(".log", "Error")
	if err != nil {
		log.Fatal("Opening error \n", err)
	}
	defer f.Close()
	log.Fatalf("%s\n", error)
}
