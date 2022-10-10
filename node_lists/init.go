package nodeLists

import (
	"log"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stderr, "node list fetcher: ", log.Ldate|log.Ltime|log.Lshortfile)
}
