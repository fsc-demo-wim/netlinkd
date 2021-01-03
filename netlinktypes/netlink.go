package netlinktypes

// Link struct
type Link struct {
	Index        int
	Name         string
	HardwareAddr string
	MTU          int
	OperState    string
	ParentIndex  int
	MasterIndex  int
	Vfs          []VfInfo
}

// VfInfo struct
type VfInfo struct {
	ID int
	Vlan int
}
