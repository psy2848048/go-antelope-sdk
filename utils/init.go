package utils

import (
	"log"
	"os"
	"sync"
)

func init() {
	logger = log.New(os.Stderr, "utils: ", log.Ldate|log.Ltime|log.Lshortfile)

	NodeRESTResponseStatus = sync.Map{}
}
