package impl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJob_ParseResource(t *testing.T) {
	job := Job{
		Name:          "job1",
		RawRequestCPU: "500m",
		RawRequestMem: "200Mi",
	}
	err := job.ParseResource()
	assert.Nil(t, err)
	assert.Equal(t, int64(500), job.RequestCPU)
	assert.Equal(t, int64(200), job.RequestMem)
}
