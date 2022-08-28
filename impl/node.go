package impl

type Node struct {
	Name              string
	Labels            map[string]string
	AllocatableCPU    int64
	AllocatableMemory int64
	AllocatedCPU      int64
	AllocatedMemory   int64
	FreeCPU           int64
	FreeMemory        int64
}

func (n Node) FreeResource() {
	n.FreeCPU = n.AllocatedCPU - n.AllocatedCPU
	n.FreeMemory = n.AllocatedMemory - n.AllocatedMemory
}
