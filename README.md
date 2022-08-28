# GuardRails Coding Challenge Submission

## Part 1:
### Question 1
#### How should do to improve current design
##### In this design, I realized that `DB` and `Storage` could be the bottleneck of system, so I try to reduce workloads to them 
- `API` should have its own database (serve some logic independent from `Storage`), storage should store only scan data
- In (2), we could add detail about job, `workers` will have enough information and avoid sending request to `Storage`
- Sharding/ Autoscaling `DB` could be good options
- For historical data, we can use replicas of `DB`

##### We should maintain the failed job
- Implement alert when `Scan Job` failed with unexpected exception
- Set `backoffLimit` on Kubernetes Jobs to auto retry failed `Scan Job`, after that, resend job into another "dead" queue to manual debug &  retry

##### Some other things we could do
- `Scan Job` send result into `Data Processing` through queue or REST API/ gRPC instead of use stdout because job/ pod could be deleted
- Consider to use some CDC tool to guarantee atomic between save data into database and send into queue
- Implement push-based metrics collector for `Scan Job` because some job run in only 5 second, too short for metrics scraping
- Distributed tracing will help us to debug/ trace problems easier
- Choose resource requests for `Scan Job` base on repository's size to reduce OOM
- Use multiple clusters, I have some painful experience when maintain Kubernetes cluster with a lot of Jobs/ Events
- Download source code from customer's repository to volume and mount into `Scan Jobs` use same source code to reduce traffic cost and avoid node storage full
- Implement graceful shutdown for all components
- Use strategy `RollingUpdate` for APIs and `Recreate` for jobs/ workers
- Carefully implement `Liveness` and `Readiness` for all components
- Use some pre-defined alert for Kubernetes cluster to manage clusters' health (for example: [Prometheus mixin](https://github.com/kubernetes-monitoring/kubernetes-mixin))
### Question 2
#### Strategy would use to save cost:
- Use horizontal pod autoscaling for `API`, `Storage`, `Worker`, `Data Processing` and `Notification`
- Use database solution support autoscaling (example: AuroraDB)
- Use [EKS Cluster Autoscaling](https://docs.aws.amazon.com/eks/latest/userguide/autoscaling.html), [EKS Fargate](https://docs.aws.amazon.com/eks/latest/userguide/fargate.html) or manage EKS nodes programmatically
## Part 2:
### To run this source:
1. Install [k3d](https://k3d.io/v5.4.5/)
2. Create k3d cluster `k3d cluster create scheduler`
3. Run command `go run main.go`