package main

import (
	"context"
	"log"
	"scheduler/impl"
	"scheduler/k8s"
)

func main() {
	k8sClient, err := k8s.NewClient("/Users/liamnv/.kube/config")
	if err != nil {
		log.Fatal(err)
	}
	imp := impl.NewImplement(k8sClient)

	jobs := []impl.Job{
		{
			Name:          "job1",
			Image:         "nginx",
			RawRequestMem: "500Mi",
			RawRequestCPU: "200m",
			Namespace:     "default",
		},
		{
			Name:          "job2",
			Image:         "nginx",
			RawRequestMem: "1Gi",
			RawRequestCPU: "100m",
			Namespace:     "default",
		},
		{
			Name:          "job3",
			Image:         "nginx",
			RawRequestMem: "2Gi",
			RawRequestCPU: "200m",
			Namespace:     "default",
		},
	}
	if err := imp.Process(context.TODO(), jobs); err != nil {
		log.Fatal(err)
	}

}
