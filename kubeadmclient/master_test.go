package kubeadmclient

import (
	"testing"
)

func TestCreateCluster(t *testing.T) {
	//
	//log.Println("starting master node creation")
	//masterNode := NewMasterNode("ubuntu", "192.168.64.16", "/Users//.ssh/id_rsa")
	//if err := masterNode.install(); err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println("fetching join command")
	//
	//joinCommand, err := masterNode.getJoinCommand()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(joinCommand)
	//
	//err = masterNode.taintAsMaster()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println("applying flannel")
	//
	//err = masterNode.applyFile("https://raw.githubusercontent.com/coreos/flannel/2140ac876ef134e0ed5af15c65e414cf26827915/Documentation/kube-flannel.yml")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println("creating worker node")
	//
	//workerNode := NewWorkerNode("ubuntu", "192.168.64.15", "/Users//.ssh/id_rsa")
	//
	//if err := workerNode.install(joinCommand); err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println("created worker node")
}
