//Package examples for deleting cluster
package examples

import "github.com/debarshibasak/go-kubeadmclient/kubeadmclient"

// This is an example for deleting cluster.
// This needs atleast one master to be specified
// This will fetch list of nodes and starting deleting them.
// Alternative you can set ResetOnDeleteCluster = true, this will also try to reset the node with best efforts.
func DeleteClusterExample() {
	k := kubeadmclient.Kubeadm{
		MasterNodes: []*kubeadmclient.MasterNode{
			kubeadmclient.NewMasterNode(
				"ubuntu",
				"192.168.64.7",
				"/Users/debarshibasak/.ssh/id_rsa",
			),
		},
	}

	err := k.DeleteCluster()
	if err != nil {
		t.Fatal(err)
	}
}
