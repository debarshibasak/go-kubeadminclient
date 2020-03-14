package kubeadmclient_test

import (
	"log"
	"testing"

	"github.com/debarshibasak/go-kubeadmclient/kubeadmclient/networking"

	"github.com/debarshibasak/go-kubeadmclient/kubeadmclient"
)

func TestKubeadm_CreateCluster2(t *testing.T) {

}

func TestKubeadm_CreateCluster(t *testing.T) {

	k := kubeadmclient.Kubeadm{
		ClusterName: "test",
		MasterNodes: []*kubeadmclient.MasterNode{
			kubeadmclient.NewMasterNode(
				"ubuntu",
				"192.168.64.47",
				"USER_HOME/.ssh/id_rsa",
			),
		},

		WorkerNodes: []*kubeadmclient.WorkerNode{
			kubeadmclient.NewWorkerNode(
				"ubuntu",
				"192.168.64.48",
				"USER_HOME/.ssh/id_rsa",
			),
			kubeadmclient.NewWorkerNode(
				"ubuntu",
				"192.168.64.49",
				"USER_HOME/.ssh/id_rsa",
			),
		},

		SkipWorkerFailure: false,
		Netorking:         networking.Flannel,
	}

	err := k.CreateCluster()
	if err != nil {
		log.Fatal(err)
	}
}
