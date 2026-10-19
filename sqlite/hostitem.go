package sqlite

type 服务器数据库条目 struct {
	Hostname  string `json:"hostname"`
	IP        string `json:"ip"`
	CPUMax    string `json:"cpu_max"`
	MemMax    string `json:"mem_max"`
	Appsystem string `json:"appsystem"`
	Dept      string `json:"dept"`
	Count     string `json:"count"`
	Modified  string `json:"modified"`
}

type 数据库中的表 struct {
	Table []服务器数据库条目 `json:"table"`
}
