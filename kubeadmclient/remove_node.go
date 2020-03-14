package kubeadmclient

import (
	"errors"
	"log"
	"sync"
	"time"
)

var (
	errNoWorkerForRemoveNode = errors.New("no worker information is set while removing node")
)

// RemoveNode will take the incoming Kubeadm struct.
// For each worker, it just reset the configuration.
// By default it will fail if any of the worker fails to reset.
// However, you can skip that with the field SkipWorkerFailure in Kubeadm
func (k *Kubeadm) RemoveNode() error {
	startTime := time.Now()

	if len(k.WorkerNodes) == 0 {
		return errNoWorkerForRemoveNode
	}

	var workerWG sync.WaitGroup
	errc := make(chan *WorkerError, 1)

	for _, worker := range k.WorkerNodes {

		workerWG.Add(1)

		go func(wg *sync.WaitGroup, worker *WorkerNode) {
			if err := worker.Reset(); err != nil {
				errc <- &WorkerError{
					worker: worker,
					err:    err,
				}
			}

			wg.Done()
		}(&workerWG, worker)
	}

	log.Println("time taken = " + time.Since(startTime).String())

	return nil
}
