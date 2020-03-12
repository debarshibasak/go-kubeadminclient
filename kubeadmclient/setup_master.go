package kubeadmclient

import (
	"github.com/debarshibasak/go-kubeadmclient/kubeadmclient/common"
	"github.com/pkg/errors"
)

func (k *Kubeadm) setupMaster(setup Setup) (string, error) {

	var joinCommand string
	var err error
	if setup == NONHA {
		joinCommand, err = k.setupNonHAMaster()
		if err != nil {
			return "", err
		}
	}
	if setup == HA {

		if k.HaProxyNode == nil {
			return "", errors.New("haproxy node is not set")
		}

		err := k.setupHAPRoxy()
		if err != nil {
			return "", err
		}

		joinCommand, err = k.setupHAMaster(k.HaProxyNode.ipOrHost)
		if err != nil {
			return "", err
		}
	}
	return joinCommand, nil
}

func (k *Kubeadm) setupHAMaster(vip string) (string, error) {

	var joinCommand string

	primaryMaster := k.MasterNodes[0]
	primaryMaster.verboseMode = k.VerboseMode

	masterJoinCommand, err := primaryMaster.installAndFetchCommand(k.ClusterName, vip)
	if err != nil {
		return "", err
	}

	for _, master := range k.MasterNodes[1:len(k.MasterNodes)] {
		err := master.Install(&common.HighAvailability{JoinCommand: masterJoinCommand})
		if err != nil {
			return "", err
		}
	}

	err = primaryMaster.ChangePermissionKubeconfig()
	if err != nil {
		return "", err
	}

	err = primaryMaster.TaintAsMaster()
	if err != nil {
		return "", err
	}

	joinCommand, err = primaryMaster.GetJoinCommand()
	if err != nil {
		return "", err
	}

	return joinCommand, nil
}

func (k *Kubeadm) setupNonHAMaster() (string, error) {

	var joinCommand string
	//nonha setup
	masterNode := k.MasterNodes[0]
	masterNode.verboseMode = k.VerboseMode

	err := masterNode.Install(nil)
	if err != nil {
		return "", err
	}

	err = masterNode.ChangePermissionKubeconfig()
	if err != nil {
		return "", err
	}

	err = masterNode.TaintAsMaster()
	if err != nil {
		return "", err
	}

	joinCommand, err = masterNode.GetJoinCommand()
	if err != nil {
		return "", err
	}

	return joinCommand, nil
}