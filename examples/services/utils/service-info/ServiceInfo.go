package serviceinfo

type (
	ServiceInfo struct {
		Name string `json:"name"`
		IP   string `json:"ip"`
		Port int32  `json:"port"`
	}
)
