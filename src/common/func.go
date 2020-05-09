package common

import (
	"log"
)

func ChkErr(err error) {
	if err != nil {
        log.Fatalf("error: %v", err)
	}
}