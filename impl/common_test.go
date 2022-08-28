package impl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortNodeByCPU(t *testing.T) {
	nodes := []Node{
		{
			Name:       "Node 1",
			FreeCPU:    5000,
			FreeMemory: 2500,
		},
		{
			Name:       "Node 2",
			FreeCPU:    2500,
			FreeMemory: 3000,
		},
		{
			Name:       "Node 3",
			FreeCPU:    500,
			FreeMemory: 500,
		},
		{
			Name:       "Node 4",
			FreeCPU:    4000,
			FreeMemory: 4000,
		},
	}

	expect := []Node{
		{
			Name:       "Node 3",
			FreeCPU:    500,
			FreeMemory: 500,
		},
		{
			Name:       "Node 2",
			FreeCPU:    2500,
			FreeMemory: 3000,
		},
		{
			Name:       "Node 4",
			FreeCPU:    4000,
			FreeMemory: 4000,
		},
		{
			Name:       "Node 1",
			FreeCPU:    5000,
			FreeMemory: 2500,
		},
	}

	actual := sortNodeByCPU(nodes)
	assert.Equal(t, actual, expect)
}

func TestSortNodeByMem(t *testing.T) {
	nodes := []Node{
		{
			Name:       "Node 1",
			FreeCPU:    5000,
			FreeMemory: 2500,
		},
		{
			Name:       "Node 2",
			FreeCPU:    2500,
			FreeMemory: 3000,
		},
		{
			Name:       "Node 3",
			FreeCPU:    500,
			FreeMemory: 500,
		},
		{
			Name:       "Node 4",
			FreeCPU:    4000,
			FreeMemory: 4000,
		},
	}

	expect := []Node{
		{
			Name:       "Node 3",
			FreeCPU:    500,
			FreeMemory: 500,
		},
		{
			Name:       "Node 1",
			FreeCPU:    5000,
			FreeMemory: 2500,
		},
		{
			Name:       "Node 2",
			FreeCPU:    2500,
			FreeMemory: 3000,
		},
		{
			Name:       "Node 4",
			FreeCPU:    4000,
			FreeMemory: 4000,
		},
	}

	actual := sortNodeByMem(nodes)
	assert.Equal(t, actual, expect)
}

func TestJobByCPU(t *testing.T) {
	jobs := []Job{
		{
			Name:       "Job 1",
			RequestMem: 5000,
			RequestCPU: 2500,
		},
		{
			Name:       "Job 2",
			RequestMem: 2500,
			RequestCPU: 3000,
		},
		{
			Name:       "Job 3",
			RequestMem: 500,
			RequestCPU: 500,
		},
		{
			Name:       "Job 4",
			RequestMem: 4000,
			RequestCPU: 4000,
		},
	}

	expect := []Job{
		{
			Name:       "Job 3",
			RequestMem: 500,
			RequestCPU: 500,
		},
		{
			Name:       "Job 1",
			RequestMem: 5000,
			RequestCPU: 2500,
		},
		{
			Name:       "Job 2",
			RequestMem: 2500,
			RequestCPU: 3000,
		},
		{
			Name:       "Job 4",
			RequestMem: 4000,
			RequestCPU: 4000,
		},
	}

	actual := sortJobByCPU(jobs)
	assert.Equal(t, expect, actual)
}

func TestJobByMem(t *testing.T) {
	jobs := []Job{
		{
			Name:       "Job 1",
			RequestMem: 5000,
			RequestCPU: 2500,
		},
		{
			Name:       "Job 2",
			RequestMem: 2500,
			RequestCPU: 3000,
		},
		{
			Name:       "Job 3",
			RequestMem: 500,
			RequestCPU: 500,
		},
		{
			Name:       "Job 4",
			RequestMem: 4000,
			RequestCPU: 4000,
		},
	}

	expect := []Job{
		{
			Name:       "Job 3",
			RequestMem: 500,
			RequestCPU: 500,
		},
		{
			Name:       "Job 2",
			RequestMem: 2500,
			RequestCPU: 3000,
		},
		{
			Name:       "Job 4",
			RequestMem: 4000,
			RequestCPU: 4000,
		},
		{
			Name:       "Job 1",
			RequestMem: 5000,
			RequestCPU: 2500,
		},
	}

	actual := sortJobByMem(jobs)
	assert.Equal(t, expect, actual)
}
