package kubeadmclient

import "errors"

func (k *Kubeadm) DeleteCluster() error {

	if len(k.MasterNodes) == 0 {
		return errors.New("no master specified")
	}

	k.MasterNodes[0].getListOfNodes()

	//k.MasterNodes[0].getJoinCommand()
	return nil
}
