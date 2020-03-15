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

		ClusterName: "testcluster",

		MasterNodes: []*kubeadmclient.MasterNode{
			kubeadmclient.NewMasterNode(
				"ubuntu",
				"192.168.64.23",
				"/Users/debarshibasak/.ssh/id_rsa",
			),
		},
		WorkerNodes: []*kubeadmclient.WorkerNode{
			kubeadmclient.NewWorkerNode(
				"ubuntu",
				"192.168.64.24",
				"/Users/debarshibasak/.ssh/id_rsa",
			),
			kubeadmclient.NewWorkerNode(
				"ubuntu",
				"192.168.64.25",
				"/Users/debarshibasak/.ssh/id_rsa",
			),
		},
		Netorking:   networking.Flannel,
		VerboseMode: false,
	}

	err := k.CreateCluster()
	if err != nil {
		log.Fatal(err)
	}
}
