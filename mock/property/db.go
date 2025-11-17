package property

import (
	cproperty "github.com/abtransitionit/gocore/mock/property"
)

var PropertyMap = map[string]FnPropertyHandler{
	"path": cproperty.GetPath,
	// "envar":           getEnvar,
	// "cgroupVersion":   getCgroupVersion,
	"cpu":      cproperty.GetCpu,
	"ram":      cproperty.GetRam,
	"uname":    getOsArch2, // code change from original
	"uuid":     getUuid,    // code change from original
	"osDistro": cproperty.GetOsDistro,
	// "osType":          getOsType, // e.g. linux, windows, darwin
	"osFamily":        cproperty.GetOsFamily, // TODO Type vs family
	"osKernelVersion": cproperty.GetOsKernelVersion,
	"osUser":          cproperty.GetOsUser,
	"osVersion":       cproperty.GetOsVersion,
	"osArch":          cproperty.GetOsArch, // e.g. arm64, amd64
	"pathTree":        getPathTree,
	"rcFilePath":      getRcFilePath,
	"selinuxStatus":   getSelinuxStatus,
	"selinuxMode":     getSelinuxMode,
	"serviceStatus":   isServiceActive,
	"serviceEnabled":  isServiceEnabled,
	"infoOs":          cproperty.GetOsInfos,
	"infoSelinux":     getSelinuxInfos,
	"infoService":     getServiceInfos,
	"netip":           getNetIp,      // code change from original
	"netgateway":      getNetGateway, // code change from original

}
