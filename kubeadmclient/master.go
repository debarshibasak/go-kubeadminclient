package kubeadmclient

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/debarshibasak/go-kubeadmclient/sshclient"
	"github.com/google/uuid"
)

type MasterNode struct {
	*Node
}

func NewMasterNode(username string, ipOrHost string, privateKeyLocation string) *MasterNode {
	return &MasterNode{
		&Node{
			username:           username,
			ipOrHost:           ipOrHost,
			privateKeyLocation: privateKeyLocation,
			clientID:           uuid.New().String(),
		},
	}
}

func (n *MasterNode) changePermissionKubeconfig() error {
	return n.run("sudo chown $USER:$USER /etc/kubernetes/admin.conf")
}

func (n *MasterNode) taintAsMaster() error {
	return n.run("KUBECONFIG=/etc/kubernetes/admin.conf kubectl taint nodes --selector=kubernetes.io/hostname=`hostname` node-role.kubernetes.io/master-")
}

func (n *MasterNode) applyFile(file string) error {
	return n.run("KUBECONFIG=/etc/kubernetes/admin.conf kubectl apply -f " + file)
}

func (n *MasterNode) getToken() (string, error) {

	sh := sshclient.SSHConnection{
		Username:    n.username,
		IP:          n.ipOrHost,
		KeyLocation: n.privateKeyLocation,
	}

	out, err := sh.Collect("sudo kubeadm token list -o json")
	if err != nil {
		return "", err
	}

	c := make(map[string]interface{})

	err = json.Unmarshal([]byte(out), &c)
	if err != nil {
		return "", err
	}

	return c["token"].(string), nil
}

func (n *MasterNode) run(shell string) error {
	return n.sshClient().Run([]string{shell})
}

func (n *MasterNode) ctlCommand(cmd string) error {
	return n.run("KUBECONFIG=/etc/kubernetes/admin.conf " + cmd)
}

func (n *MasterNode) getKubeConfig() (string, error) {
	return n.sshClient().Collect("sudo cat /etc/kubernetes/admin.conf")
}

func (n *MasterNode) getJoinCommand() (string, error) {
	return n.sshClient().Collect("sudo kubeadm token create --print-join-command")
}

func (n *MasterNode) installAndFetchCommand(kubeadm Kubeadm, vip string) (string, error) {

	osType := n.determineOS()

	if osType == nil {
		return "", errors.New("ostpye not found")
	}

	err := n.sshClient().Run(osType.Commands())
	if err != nil {
		return "", err
	}

	err = n.sshClient().ScpToWithData([]byte(generateKubeadmConfig(vip, kubeadm)), "/tmp/kubeadm-config.yaml")
	if err != nil {
		return "", err
	}

	out, err := n.sshClientWithTimeout(30 * time.Minute).Collect("sudo kubeadm init --config /tmp/kubeadm-config.yaml --upload-certs")
	if err != nil {
		log.Println(out)
		return "", err
	}

	return getControlPlaneJoinCommand(out), nil
}

func (n *MasterNode) install(kubeadm Kubeadm, availability *highAvailability) error {

	osType := n.determineOS()

	err := n.sshClientWithTimeout(30 * time.Minute).Run(osType.Commands())
	if err != nil {
		return err
	}

	var s string

	if availability != nil {
		s = "sudo " + availability.JoinCommand
	} else {
		s = "sudo kubeadm init --pod-network-cidr=" + kubeadm.PodNetwork + " --service-cidr=" + kubeadm.ServiceNetwork + " --service-dns-domain=" + kubeadm.DNSDomain
	}

	return n.sshClientWithTimeout(30 * time.Minute).Run([]string{s})
}

func getControlPlaneJoinCommand(data string) string {
	var cmd string

	for _, line := range strings.Split(data, "\n") {

		if strings.HasPrefix(strings.TrimSpace(line), "kubeadm") {
			cmd = cmd + strings.ReplaceAll(line, "\\", "")
		}

		if strings.HasPrefix(strings.TrimSpace(line), "--discovery") {
			cmd = cmd + strings.ReplaceAll(line, "\\", "")
		}

		if strings.HasPrefix(strings.TrimSpace(line), "--control-plane") {
			cmd = cmd + strings.ReplaceAll(line, "\\", "")
			return cmd
		}
	}

	return cmd
}
