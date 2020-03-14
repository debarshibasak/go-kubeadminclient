package kubeadmclient

import "errors"

var (
	errMasterNotSpecified = errors.New("master node not specified")
	errWorkerNotSpecified = errors.New("worker not specified")
)

func (k *Kubeadm) AddNode() error {

	if len(k.MasterNodes) == 0 {
		return errMasterNotSpecified
	}

	if len(k.WorkerNodes) == 0 {
		return errWorkerNotSpecified
	}

	joinCommand, err := k.MasterNodes[0].GetJoinCommand()
	if err != nil {
		return err
	}

	if err := k.setupWorkers(joinCommand); err != nil {
		return err
	}

	return err
}
