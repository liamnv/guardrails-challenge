package k8s

import (
	"context"
	"github.com/pkg/errors"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Client struct {
	clientSet *kubernetes.Clientset
}

func NewClient(kubeConfig string) (*Client, error) {

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		return nil, errors.Wrap(err, "Init config")
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errors.Wrap(err, "Init Kubernetes Config")
	}
	return &Client{
		clientSet: clientSet,
	}, nil
}

func (c *Client) GetNodes(ctx context.Context) ([]corev1.Node, error) {
	result, err := c.clientSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "get list nodes")
	}
	return result.Items, nil
}

func (c *Client) AllPods(ctx context.Context) ([]corev1.Pod, error) {
	pods, err := c.clientSet.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "get all pods")
	}
	return pods.Items, nil
}

func (c *Client) ScheduleJobs(ctx context.Context, jobs []*batchv1.Job) error {
	//TODO: use goroutines
	for _, job := range jobs {
		if _, err := c.clientSet.BatchV1().Jobs(job.Namespace).Create(ctx, job, metav1.CreateOptions{}); err != nil {
			return errors.Wrap(err, "create new job")
		}
	}
	return nil
}
