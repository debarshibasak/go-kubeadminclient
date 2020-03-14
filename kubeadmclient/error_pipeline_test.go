package kubeadmclient_test

import (
	"log"
	"testing"
	"time"

	"errors"
)

func TestError(t *testing.T) {

	errc := make(chan error, 100)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			errc <- errors.New("data" + time.Now().String())
		}
	}()

	for c := range errc {
		if c != nil {
			log.Println(c)
		}
	}
}
