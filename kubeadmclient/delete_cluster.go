package kubeadmclient

import (
	"errors"
	"log"
)

func (k *Kubeadm) DeleteCluster() error {

	if len(k.MasterNodes) == 0 {
		return errors.New("no master specified")
	}

	nodelist, err := k.MasterNodes[0].getAllWorkerNodeNames()
	if err != nil {
		return err
	}

	if k.ResetOnDeleteCluster && len(k.WorkerNodes) < len(nodelist) {
		return errors.New("will not be able to reset nodes as the nodelist is greater than worker nodes. This hints that some node details are missing")
	}

	masterNodeList, err := k.MasterNodes[0].getAllMasterNodeNames()
	if err != nil {
		return err
	}

	if k.ResetOnDeleteCluster && len(k.MasterNodes) < len(masterNodeList) {
		return errors.New("will not be able to reset nodes as the nodelist is greater than master nodes. This hints that some node details are missing")
	}

	if len(masterNodeList) == 0 && len(nodelist) == 0 {
		return errors.New("looks like an empty cluster")
	}

	if k.ResetOnDeleteCluster {
		err := k.RemoveNode()
		if err != nil {
			return err
		}
	}

	if len(nodelist) > 0 {
		var errC = make(chan error, 1)
		for i, node := range nodelist {
			go func(node string, index int) {
				if err := k.MasterNodes[0].deleteNode(node); err != nil {
					errC <- err
				}

				if index == len(nodelist)-1 {
					close(errC)
				}

			}(node, i)
		}

		for e := range errC {
			if e != nil {
				log.Println("error - " + e.Error())
				if !k.SkipWorkerFailure {
					return e
				}
			}
		}
	}

	if len(masterNodeList) > 0 {
		var errMasterDeletion = make(chan error, 1)
		defer close(errMasterDeletion)

		for i, node := range masterNodeList {
			go func(node string, index int) {
				if err := k.MasterNodes[0].deleteNode(node); err != nil {
					errMasterDeletion <- err
				}

				log.Println(index)
				if index == len(masterNodeList)-1 {
					close(errMasterDeletion)
				}

			}(node, i)
		}

		for e := range errMasterDeletion {
			if e != nil {
				return e
			}
		}
	}

	return nil
}
