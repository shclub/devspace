package kubectl

import (
	"context"
	"io/ioutil"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

// Logs prints the container logs
func Logs(client *kubernetes.Clientset, namespace, podName, containerName string, lastContainerLog bool, tail *int64) (string, error) {
	lines := int64(100)
	if tail != nil {
		lines = *tail
	}

	request := client.Core().Pods(namespace).GetLogs(podName, &v1.PodLogOptions{
		Container: containerName,
		TailLines: &lines,
		Previous:  lastContainerLog,
	})

	reader, err := request.Context(context.Background()).Stream()
	if err != nil {
		return "", err
	}

	logs, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}

	return string(logs), nil
}
