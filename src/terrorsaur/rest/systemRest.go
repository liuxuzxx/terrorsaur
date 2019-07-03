package rest

type SystemInformation struct {
	systemName string
	status     string
}

var systemInformation = SystemInformation{"Go-WEB-Iris系统", "OK"}

func FetchSystemInformation() SystemInformation {
	return systemInformation
}
