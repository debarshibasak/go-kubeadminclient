package kubeadmclient_test

import (
	"testing"

	"github.com/debarshibasak/go-kubeadmclient/kubeadmclient"
)

func TestKubeadm_AddNode(t *testing.T) {
	k := kubeadmclient.Kubeadm{
		MasterNodes: []*kubeadmclient.MasterNode{
			kubeadmclient.NewMasterNode(
				"ubuntu",
				"192.168.64.47",
				"/Users/debarshibasak/.ssh/id_rsa",
			),
		},
		WorkerNodes: []*kubeadmclient.WorkerNode{
			kubeadmclient.NewWorkerNode(
				"ubuntu",
				"192.168.64.51",
				"/Users/debarshibasak/.ssh/id_rsa",
			),
			kubeadmclient.NewWorkerNode(
				"ubuntu",
				"192.168.64.52",
				"/Users/debarshibasak/.ssh/id_rsa",
			),
			kubeadmclient.NewWorkerNode(
				"ubuntu",
				"192.168.64.50",
				"/Users/debarshibasak/.ssh/id_rsa",
			),
			kubeadmclient.NewWorkerNode(
				"ubuntu",
				"192.168.64.48",
				"/Users/debarshibasak/.ssh/id_rsa",
			),
			kubeadmclient.NewWorkerNode(
				"ubuntu",
				"192.168.64.49",
				"/Users/debarshibasak/.ssh/id_rsa",
			),
		},
		SkipWorkerFailure: true,
	}

	err := k.AddNode()
	if err != nil {
		t.Fatal(err)
	}
}
