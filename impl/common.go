package impl

import (
	"sort"
)

func sortNodeByCPU(nodes []Node) []Node {
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].FreeCPU < nodes[j].FreeCPU
	})
	return nodes
}

func sortNodeByMem(nodes []Node) []Node {
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].FreeMemory < nodes[j].FreeMemory
	})
	return nodes
}

func sortJobByCPU(jobs []Job) []Job {
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].RequestCPU < jobs[j].RequestCPU
	})
	return jobs
}

func sortJobByMem(jobs []Job) []Job {
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].RequestMem < jobs[j].RequestMem
	})
	return jobs
}
