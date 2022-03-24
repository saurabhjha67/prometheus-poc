package contracts

type CPUQuota struct {
	Modules []Module `json:"module"`
}

type Module struct {
	CPUs []CPU `json:"cpu"`
}

type CPU struct {
	CPUNames     []string `json:"cpu_names"`
	Override     string   `json:"override"`
	OverrideDisk string   `json:"overridedisk"`
	ProjectId    string   `json:"projectid"`
	Region       string   `json:"region"`
	Source       string   `json:"source"`
}
