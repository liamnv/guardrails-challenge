package impl

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"scheduler/k8s"
	"testing"
)

func TestProcessScenario1(t *testing.T) {
	jobs := []Job{
		{
			Name:       "job1",
			Image:      "docker_img_1",
			RequestMem: 500,
			RequestCPU: 200,
		},
		{
			Name:       "job2",
			Image:      "docker_img_2",
			RequestMem: 1000,
			RequestCPU: 100,
		},
		{
			Name:       "job3",
			Image:      "docker_img_3",
			RequestMem: 2000,
			RequestCPU: 200,
		},
	}
	nodes := []Node{
		{
			Name: "node1",
			Labels: map[string]string{
				"kubernetes.io/hostname": "ip-172-100-1-1.ec2.internal",
			},
			AllocatableCPU:    1000,
			AllocatableMemory: 4000,
			AllocatedCPU:      1000,
			AllocatedMemory:   3000,
			FreeCPU:           0,
			FreeMemory:        1000,
		},
		{
			Name: "node2",
			Labels: map[string]string{
				"kubernetes.io/hostname": "ip-172-100-1-2.ec2.internal",
			},
			AllocatableCPU:    1000,
			AllocatableMemory: 4000,
			AllocatedCPU:      0,
			AllocatedMemory:   0,
			FreeCPU:           1000,
			FreeMemory:        4000,
		},
	}

	expect := []Job{
		{
			Name:       "job1",
			Image:      "docker_img_1",
			RequestMem: 500,
			RequestCPU: 200,
			NodeName:   "node2",
		},
		{
			Name:       "job2",
			Image:      "docker_img_2",
			RequestMem: 1000,
			RequestCPU: 100,
			NodeName:   "node2",
		},
		{
			Name:       "job3",
			Image:      "docker_img_3",
			RequestMem: 2000,
			RequestCPU: 200,
			NodeName:   "node2",
		},
	}
	actual := chooseNode(jobs, nodes)
	assert.Equal(t, expect, actual)
}

func TestProcessScenario2(t *testing.T) {
	jobs := []Job{
		{
			Name:       "job1",
			Image:      "docker_img_1",
			RequestMem: 500,
			RequestCPU: 200,
		},
		{
			Name:       "job2",
			Image:      "docker_img_2",
			RequestMem: 1000,
			RequestCPU: 100,
		},
		{
			Name:       "job3",
			Image:      "docker_img_3",
			RequestMem: 2000,
			RequestCPU: 200,
		},
	}
	nodes := []Node{
		{
			Name: "node1",
			Labels: map[string]string{
				"kubernetes.io/hostname": "ip-172-100-1-1.ec2.internal",
			},
			AllocatableCPU:    1000,
			AllocatableMemory: 4000,
			AllocatedCPU:      0,
			AllocatedMemory:   0,
			FreeCPU:           1000,
			FreeMemory:        4000,
		},
		{
			Name: "node2",
			Labels: map[string]string{
				"kubernetes.io/hostname": "ip-172-100-1-2.ec2.internal",
			},
			AllocatableCPU:    1000,
			AllocatableMemory: 4000,
			AllocatedCPU:      0,
			AllocatedMemory:   0,
			FreeCPU:           1000,
			FreeMemory:        4000,
		},
	}
	expect := []Job{
		{
			Name:       "job1",
			Image:      "docker_img_1",
			RequestMem: 500,
			RequestCPU: 200,
			NodeName:   "node2",
		},
		{
			Name:       "job2",
			Image:      "docker_img_2",
			RequestMem: 1000,
			RequestCPU: 100,
			NodeName:   "node2",
		},
		{
			Name:       "job3",
			Image:      "docker_img_3",
			RequestMem: 2000,
			RequestCPU: 200,
			NodeName:   "node2",
		},
	}
	actual := chooseNode(jobs, nodes)
	assert.Equal(t, expect, actual)
}

func TestProcessScenario3(t *testing.T) {
	jobs := []Job{
		{
			Name:       "job1",
			Image:      "docker_img_1",
			RequestMem: 500,
			RequestCPU: 200,
		},
		{
			Name:       "job2",
			Image:      "docker_img_2",
			RequestMem: 1000,
			RequestCPU: 100,
		},
		{
			Name:       "job3",
			Image:      "docker_img_3",
			RequestMem: 2000,
			RequestCPU: 200,
		},
	}
	nodes := []Node{
		{
			Name: "node1",
			Labels: map[string]string{
				"kubernetes.io/hostname": "ip-172-100-1-1.ec2.internal",
			},
			AllocatableCPU:    1000,
			AllocatableMemory: 4000,
			AllocatedCPU:      100,
			AllocatedMemory:   2500,
			FreeCPU:           900,
			FreeMemory:        1500,
		},
		{
			Name: "node2",
			Labels: map[string]string{
				"kubernetes.io/hostname": "ip-172-100-1-2.ec2.internal",
			},
			AllocatableCPU:    1000,
			AllocatableMemory: 3000,
			AllocatedCPU:      0,
			AllocatedMemory:   2000,
			FreeCPU:           1000,
			FreeMemory:        1000,
		},
		{
			Name: "node3",
			Labels: map[string]string{
				"kubernetes.io/hostname": "ip-172-100-1-2.ec2.internal",
			},
			AllocatableCPU:    1000,
			AllocatableMemory: 3000,
			AllocatedCPU:      0,
			AllocatedMemory:   0,
			FreeCPU:           1000,
			FreeMemory:        3000,
		},
	}
	expect := []Job{
		{
			Name:       "job1",
			Image:      "docker_img_1",
			RequestMem: 500,
			RequestCPU: 200,
			NodeName:   "node3",
		},
		{
			Name:       "job2",
			Image:      "docker_img_2",
			RequestMem: 1000,
			RequestCPU: 100,
			NodeName:   "node1",
		},
		{
			Name:       "job3",
			Image:      "docker_img_3",
			RequestMem: 2000,
			RequestCPU: 200,
			NodeName:   "node3",
		},
	}
	actual := chooseNode(jobs, nodes)
	assert.Equal(t, expect, actual)
}

func TestImplement_mapNodeResource(t *testing.T) {
	client, err := k8s.NewClient("/Users/liamnv/.kube/config")
	assert.Nil(t, err)
	implement := NewImplement(client)
	nodes, err := implement.GetNode(context.Background())
	fmt.Println(nodes)
}
