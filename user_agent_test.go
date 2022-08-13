package UserAgent_test

import (
	"log"
	"testing"

	ua "github.com/msadg/UserAgent"
)

func TestRand(t *testing.T) {
	d := ua.Rand()
	log.Print(d)
}
