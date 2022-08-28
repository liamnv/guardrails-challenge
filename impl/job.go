package impl

import (
	"github.com/pkg/errors"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Job struct {
	Name                 string
	Image                string
	RawRequestCPU        string
	RawRequestMem        string
	RequestMem           int64
	RequestCPU           int64
	NodeName             string
	Namespace            string
	BackOffLimit         int32
	ActiveDeadlineSecond int
}

func (j *Job) ParseResource() error {
	if cpu, err := resource.ParseQuantity(j.RawRequestCPU); err != nil {
		return errors.Wrap(err, "parse cpu quantity")
	} else {
		j.RequestCPU = cpu.MilliValue()
	}
	if mem, err := resource.ParseQuantity(j.RawRequestMem); err != nil {
		return errors.Wrap(err, "parse memory quantity")
	} else {
		j.RequestMem = mem.Value() / 1024 / 1024
	}
	return nil
}

func (j *Job) ToK8sJob() (*batchv1.Job, error) {
	var backoffLimit *int32 = &j.BackOffLimit
	cpuQuantity, err := resource.ParseQuantity(j.RawRequestCPU)
	if err != nil {
		return nil, errors.Wrap(err, "parse cpu quantity")
	}
	memQuantity, err := resource.ParseQuantity(j.RawRequestMem)
	if err != nil {
		return nil, errors.Wrap(err, "parse memory quantity")
	}
	k8sJob := batchv1.Job{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Job",
			APIVersion: "batch/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      j.Name,
			Namespace: j.Namespace,
		},

		Spec: batchv1.JobSpec{
			BackoffLimit: backoffLimit,
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "job",
							Image: j.Image,
							Resources: corev1.ResourceRequirements{
								Requests: corev1.ResourceList{
									corev1.ResourceCPU:    cpuQuantity,
									corev1.ResourceMemory: memQuantity,
								},
							},
						},
					},
					RestartPolicy: "Never",
					NodeName:      j.NodeName,
				},
			},
		},
	}
	return &k8sJob, nil
}
