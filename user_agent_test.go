package userAgent_test

import (
	"log"
	"testing"

	ua "github.com/msadg/userAgent"
)

func TestRand(t *testing.T) {
	d := ua.Rand()
	log.Print(d)
}
