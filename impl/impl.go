package impl

import (
	"context"
	"github.com/pkg/errors"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
)

type Client interface {
	GetNodes(context.Context) ([]corev1.Node, error)
	AllPods(context.Context) ([]corev1.Pod, error)
	ScheduleJobs(context.Context, []*batchv1.Job) error
}

type Implement struct {
	Client Client
}

func NewImplement(client Client) *Implement {
	return &Implement{
		Client: client,
	}
}

func (impl *Implement) Process(ctx context.Context, jobs []Job) error {
	for i := range jobs {
		if err := jobs[i].ParseResource(); err != nil {
			return errors.Wrap(err, "parse job resoure")
		}
	}
	nodes, err := impl.GetNode(ctx)
	if err != nil {
		return errors.Wrap(err, "get all nodes")
	}
	jobs = chooseNode(jobs, nodes)
	k8sJobs := make([]*batchv1.Job, 0, len(jobs))
	for _, job := range jobs {
		if k8sJob, err := job.ToK8sJob(); err != nil {
			return errors.Wrap(err, "parse job to k8s job")
		} else {
			k8sJobs = append(k8sJobs, k8sJob)
		}
	}
	return impl.Client.ScheduleJobs(ctx, k8sJobs)
}

func (impl *Implement) GetNode(ctx context.Context) ([]Node, error) {
	k8sNodes, err := impl.Client.GetNodes(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get k8s nodes")
	}
	nodes := make([]Node, 0, len(k8sNodes))
	for _, r := range k8sNodes {
		n := Node{
			Name:              r.Name,
			Labels:            r.Labels,
			AllocatableCPU:    r.Status.Allocatable.Cpu().MilliValue(),
			AllocatableMemory: r.Status.Allocatable.Memory().Value() / 1024 / 1024,
		}
		nodes = append(nodes, n)
	}

	k8sPods, err := impl.Client.AllPods(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get k8s pods")
	}
	nodes = mapNodeResource(nodes, k8sPods)
	return nodes, nil

}
func chooseNode(jobs []Job, nodes []Node) []Job {
	//In example Jobs, 3 jobs are very different at Memory => I will use memory for schedule job into node
	jobs = sortJobByMem(jobs)
	nodes = sortNodeByMem(nodes)
	for i := len(jobs) - 1; i >= 0; i-- {
		if node := scheduleJob(jobs[i], nodes); node != nil {
			jobs[i].NodeName = node.Name
			node.FreeCPU -= jobs[i].RequestCPU
			node.FreeMemory -= jobs[i].RequestMem
		}
	}
	return jobs
}

func scheduleJob(job Job, nodes []Node) *Node {
	for i := len(nodes) - 1; i >= 0; i-- {
		if nodes[i].FreeCPU > job.RequestCPU && nodes[i].FreeMemory > job.RequestMem {
			return &nodes[i]
		}
	}
	return nil
}

func mapNodeResource(nodes []Node, pods []corev1.Pod) []Node {
	type Resource struct {
		CPU    int64
		Memory int64
	}
	nodeResourceMap := make(map[string]Resource)
	for _, pod := range pods {
		for _, container := range pod.Spec.Containers {
			if _, ok := nodeResourceMap[pod.Spec.NodeName]; !ok {
				nodeResourceMap[pod.Spec.NodeName] = Resource{}
			}
			resource := nodeResourceMap[pod.Spec.NodeName]
			resource.CPU += container.Resources.Requests.Cpu().MilliValue()
			resource.Memory += container.Resources.Requests.Memory().Value() / 1024 / 1024
			nodeResourceMap[pod.Spec.NodeName] = resource
		}
	}
	for i := range nodes {
		nodes[i].AllocatedCPU = nodeResourceMap[nodes[i].Name].CPU
		nodes[i].AllocatedMemory = nodeResourceMap[nodes[i].Name].Memory
	}
	return nodes
}
