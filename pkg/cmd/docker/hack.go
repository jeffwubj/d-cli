package docker

import (
	api "k8s.io/api/core/v1"
	"os"
	"strconv"
)

func SetDockerHost(port int) {
	os.Setenv("DOCKER_HOST", DefaultHost+ strconv.Itoa(port))
}

func IsPodTerminating(pod *api.Pod) bool{
	return pod.ObjectMeta.DeletionTimestamp != nil
}
