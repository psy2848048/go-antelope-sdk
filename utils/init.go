package utils

import (
	"log"
	"os"
	"sync"
)

func init() {
	logger = log.New(os.Stderr, "xxx: ", log.Ldate|log.Ltime|log.Lshortfile)

	NodeRESTResponseStatus = sync.Map{}
}
