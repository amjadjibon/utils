package k8s

import (
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func Client() (kubernetes.Interface, error) {
	var masterUrl = os.Getenv("MASTER_URL")
	var kubeConfigPath = os.Getenv("KUBE_CONFIG_PATH")

	var kubeConfig *rest.Config
	var client kubernetes.Interface
	var err error

	kubeConfig, err = clientcmd.BuildConfigFromFlags(masterUrl, kubeConfigPath)
	if err != nil {
		return nil, err
	}

	client, err = kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return nil, err
	}

	return client, nil
}
