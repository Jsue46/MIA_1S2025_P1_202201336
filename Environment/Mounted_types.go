package Environment

// MountedPartition almacena información de particiones montadas
type MountedPartition struct {
	MountPath   string
	MountName   string
	MountID     string
	MountStatus byte
	LoggedIn    bool
	MountStart  int32
}

func GetMountedPartitions() [][]MountedPartition {
	// Example data: Replace this with actual logic to retrieve mounted partitions.
	return [][]MountedPartition{
		{
			{MountID: "partition1", MountPath: "/dev/sda1"},
			{MountID: "partition2", MountPath: "/dev/sda2"},
		},
		{
			{MountID: "partition3", MountPath: "/dev/sdb1"},
		},
	}
}
