package utils

import (
	"log"
	"os"
	"runtime"
	"time"
)

const LOG_PATH = "./logs/"

/*
	To turn off all output from a log.Logger, set the output destination to ioutil.Discard,
	a writer on which all calls succeed without doing anything.
	log.SetOutput(ioutil.Discard)
*/

func Debug(e ...interface{}) {
	_, fn, line, _ := runtime.Caller(1)

	f, err := os.OpenFile(LOG_PATH+time.Now().Format("20060102")+"_debug.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "DEBUG |", log.LstdFlags)
	logger.Printf("| %s:%d | %v\n", fn, line, e)
}
