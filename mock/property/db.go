package property

var PropertyMap = map[string]FnPropertyHandler{
	"path": getPath,
	// "envar":           getEnvar,
	// "cgroupVersion":   getCgroupVersion,
	"cpu":      getCpu,
	"ram01":    getRam01,
	"ram02":    getRam02,
	"uname":    getUnameM, // code change from original
	"uuid":     getUuid,   // code change from original
	"osDistro": getOsDistro,
	// "osType":          getOsType, // e.g. linux, windows, darwin
	"osFamily":        getOsFamily, // TODO Type vs family
	"osKernelVersion": getOsKernelVersion,
	"osUser":          getOsUser,
	"osVersion":       getOsVersion,
	"osArch":          getOsArch, // e.g. arm64, amd64
	"pathTree":        getPathTree,
	"rcFilePath":      getRcFilePath,
	"selinuxStatus":   getSelinuxStatus,
	"selinuxMode":     getSelinuxMode,
	"serviceStatus":   isServiceActive,
	"serviceEnabled":  isServiceEnabled,

	"netip":      getNetIp,      // code change from original
	"netgateway": getNetGateway, // code change from original

}
