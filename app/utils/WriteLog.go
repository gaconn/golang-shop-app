package utils

import (
	"os"
	"runtime"
)

func WriteErrorLog(message error) {
	_, file, line, _ := runtime.Caller(1)
	f, err := os.OpenFile("./data/logs/logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}

	f.WriteString("File: " + file + "----Line: " + string(line) + "---message: " + message.Error() + "\n")

	f.Close()
}
