package property

import (
	cproperty "github.com/abtransitionit/gocore/mock/property"
)

var PropertyMap = map[string]FnPropertyHandler{
	"path":  cproperty.GetPath,
	"envar": getEnvar,
	// "cgroupVersion":   getCgroupVersion,
	"cpu":             cproperty.GetCpu,
	"ram":             cproperty.GetRam,
	"uuid":            getUuid, // code change from original
	"osDistro":        cproperty.GetOsDistro,
	"osType":          cproperty.GetOsType,   // e.g. linux, windows, darwin
	"osFamily":        cproperty.GetOsFamily, // linux:rhel, debian; windows:windows; mac:Worstation
	"osKernelVersion": cproperty.GetOsKernelVersion,
	"osUser":          cproperty.GetOsUser,
	"osVersion":       cproperty.GetOsVersion,
	"osArch":          cproperty.GetOsArch, // e.g. arm64, amd64
	"uname":           getOsArch2,          // code change from original
	"pathTree":        getPathTree,
	"rcFilePath":      getRcFilePath,
	"selinuxStatus":   getSelinuxStatus,
	"selinuxMode":     getSelinuxMode,
	"serviceStatus":   isServiceActive,
	"serviceEnabled":  isServiceEnabled,
	"infoOs":          cproperty.GetOsInfos,
	"infoSelinux":     getSelinuxInfos,
	"infoService":     getServiceInfos,
	"netip":           getNetIp,
	"netgateway":      getNetGateway,
	"needReboot":      getNeedReboot,
}
