package kubeadmclient_test

import (
	"testing"

	"github.com/debarshibasak/go-kubeadmclient/kubeadmclient"
)

func TestKubeadm_DeleteCluster(t *testing.T) {
	k := kubeadmclient.Kubeadm{
		MasterNodes: []*kubeadmclient.MasterNode{
			kubeadmclient.NewMasterNode(
				"ubuntu",
				"192.168.64.7",
				"/Users/debarshibasak/.ssh/id_rsa",
			),
		},
		SkipWorkerFailure: false,
	}

	err := k.DeleteCluster()
	if err != nil {
		t.Fatal(err)
	}
}
