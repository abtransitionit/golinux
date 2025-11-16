/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package util

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// Name: PropertyHandler
//
// Description: a function that retrieves a property.
type PropertyHandler func(...string) (string, error)

// map a string to a function
var OsPropertyMap = map[string]PropertyHandler{
	// "ostype":          getOsType, // e.g. linux, windows, darwin
	// "osarch":          getOsArch,
	// "cpu":             getCpu,
	// "path":            getPath,
	// "osversion":       getOsVersion,
	// "osuser":          getOsUser,
	// "ram":             getRam,
	// "netip":           getNetIp,
	// "netgateway":      getNetGateway,
	// "oskversion":      getOsKernelVersion,
	// "uuid":     getUuid,
	// "uname":    getUnameM,
	// "osdistro": getOsDistro,
	// "osfamily": getOsFamily,

	"rebootstatus":   getReboot,
	"serviceStatus":  getServiceStatus,
	"serviceEnabled": getServiceEnabled,
	"serviceinfos":   getServiceInfos,
	"cgroup":         getCgroupVersion,

	// "selstatus": getSelinuxStatus,
	"selmode":  getSelinuxMode,
	"selinfos": getSelinuxInfos,

	"clipackage":      getPackage,
	"init":            getInitSystem,
	"host":            getHost,
	"userlinger":      getLinger,
	"osinfos":         getOsInfos,
	"pathext":         getPathExtend,
	"pathtree":        getPathTree,
	"sshreachability": getSshReachability,
}

func GetProperty(property string, params ...string) (string, error) {

	// get function that manages that property
	fnHandler, ok := coreProperties[property]
	if !ok {
		return "", fmt.Errorf("unknown property requested: %s", property)
	}

	// play that function and get it output
	output, err := fnHandler(params...)
	if err != nil {
		return "", fmt.Errorf("error getting %s: %w", property, err)
	}

	return strings.TrimSpace(output), nil
}

// GetProperty handles local (cross-platform or Linux) and remote properties.
func GetProperty(vmName, property string, params ...string) (string, error) {
	// Remote property
	if vmName != "" {
		cmd := fmt.Sprintf("goluc property %s %s", property, strings.Join(params, " "))
		// does CLI goluc exist on the remote
		output, err := run.RunCliSsh(vmName, cmd)
		if err != nil {
			return "", fmt.Errorf("failed to get remote property '%s' from '%s': %w", property, vmName, err)
		}
		return strings.TrimSpace(output), nil
	}

	// Local property: try cross-platform first
	if val, err := coreproperty.GetProperty(property, params...); err == nil {
		return val, nil
	}

	// Local property: try linux-platform then
	if runtime.GOOS == "linux" {
		if handler, ok := linuxProperties[property]; ok {
			val, err := handler(params...)
			if err != nil {
				return "", fmt.Errorf("error getting %s: %w", property, err)
			}
			return strings.TrimSpace(val), nil
		}
	}

	// Local property: unknown property
	return "", fmt.Errorf("unknown property requested: %s", property)
}

func getPackage(params ...string) (string, error) {

	// manage argument
	if len(params) < 1 {
		return "", fmt.Errorf("cli name required")
	}

	// get input
	cliName := params[0]

	// get os:type
	osType, err := getOsType()
	if err != nil {
		return "", err
	}
	// get os:arch
	osfamily, err := getOsFamily()
	if err != nil {
		return "", err
	}

	// manage linux only
	if strings.TrimSpace(strings.ToLower(osType)) != "linux" {
		return "", fmt.Errorf("os:type not manage: [%s]", osType)
	}

	return osfamily + " " + cliName, nil
}
func getLingerStatus(params ...string) (string, error) {
	if len(params) < 1 {
		return "", fmt.Errorf("user name required")
	}

	// get input
	OsUserName := params[0]

	// play test cli - same as testing if cli : loginctl exists
	cli := fmt.Sprintf(`loginctl show-user %s`, OsUserName)
	if _, err := RunCLILocal(cli); err != nil {
		return "", err
	}
	// now grep is safe
	cli = fmt.Sprintf(`loginctl show-user %s | grep -i linger | cut -d= -f2`, OsUserName)
	output, err := RunCLILocal(cli)
	if err != nil {
		return "", err
	}
	// success
	return output, nil
}

func getSshReachability(params ...string) (string, error) {
	// manage argument
	if len(params) < 1 {
		return "", fmt.Errorf("vm name required")
	}

	// get service name
	vm := params[0]

	// play cli
	if _, cli := IsVmSshReachable(vm); cli != nil {
		return "false", nil
	}
	return "true", nil
}

func getRebootStatus(_ ...string) (string, error) {
	// Ensure we're on Linux
	osType, err := getOsType()
	if err != nil {
		return "", fmt.Errorf("could not detect OS type: %w", err)
	}
	if osType != "linux" {
		return "", fmt.Errorf("unsupported OS type: %s (only linux is supported)", osType)
	}

	// Detect the OS family
	osFamily, err := getOsFamily()
	if err != nil {
		return "", fmt.Errorf("could not detect OS family: %w", err)
	}

	// Select appropriate command
	var cli string
	switch strings.TrimSpace(osFamily) {
	case "debian":
		cli = "test -f /var/run/reboot-required && echo true || echo false"
	case "rhel":
		cli = "	command -v needs-restarting >/dev/null && needs-restarting -r | grep -q 'Reboot is required' && echo true || echo false"
	default:
		return "", fmt.Errorf("unsupported OS family: %s", osFamily)
	}

	// Run the command
	output, err := RunCLILocal(cli)
	if err != nil {
		return "", fmt.Errorf("failed to check reboot requirement: %w", err)
	}

	return strings.TrimSpace(output), nil
}

func getHost(_ ...string) (string, error) {
	osType, err := getOsType()
	if err != nil {
		return "", err
	}
	if osType != "linux" {
		return "Unsupported OS", nil
	}
	cmd := "systemd-detect-virt"
	out, err := RunCLILocal(cmd)
	if err != nil {
		return "", err
	}
	return out, nil
}

func getInitSystem(_ ...string) (string, error) {
	output, err := RunCLILocal("ps -p 1 -o comm=")
	if err != nil {
		return "", fmt.Errorf("getting init > %v", err)
	}
	if strings.Contains(output, "systemd") {
		return "systemd (cgroup v2)", nil
	}
	return "initd (likely cgroup v1)", nil
}

func getPathExtend(params ...string) (string, error) {
	// check arg
	if len(params) < 1 {
		return "", fmt.Errorf("semi-colon separated paths required")
	}
	// get property
	path := os.Getenv("PATH")
	if path == "" {
		return "", fmt.Errorf("PATH environment variable is not set")
	}
	// get extension
	extension := params[0]

	pathExtend, err := UpdatePath(extension)
	if err != nil {
		return "", err
	}
	return pathExtend, nil
}

func getCgroupVersion(_ ...string) (string, error) {
	content, err := os.ReadFile("/proc/self/cgroup")
	if err != nil {
		return "", fmt.Errorf("getting cgroup > %w", err)
	}
	if strings.Contains(string(content), "0::/") {
		return "v2", nil
	}
	return "v1", nil
}

// return "", fmt.Errorf("getting net-ip > %s", cErr)

func GetOsPropertyMap() map[string]PropertyHandler {
	return OsPropertyMap
}

// Example Usage:
//
//	props := []string{"cpu", "ram", "osarch", "uuid", "cgroup"}
//
//	for _, prop := range props {
//		value, err := util.GetPropertyLocal(prop)
//		if err != nil {
//			// logx.L.Debugf("%s", err)
//			continue
//		}
//		fmt.Printf("prop: %s value: %s\n", prop, value)
//	}
func GetPropertyLocal(property string, params ...string) (string, error) {
	fnPropertyHandler, ok := OsPropertyMap[property]
	if !ok {
		return "", fmt.Errorf("❌ unknown property requested: %s", property)
	}

	output, err := fnPropertyHandler(params...)
	if err != nil {
		return "", fmt.Errorf("❌ error getting %s: %w", property, err)
	}

	return output, nil
}

func GetPropertyRemote(vm string, property string, params ...string) (string, error) {
	cli := fmt.Sprintf(`luc do getprop %s`, property)

	// Append optional params if any
	if len(params) > 0 {
		cli = fmt.Sprintf(`luc do getprop %s %s`, property, strings.Join(params, " "))
	}

	out, err := RunCLIRemote(vm, cli)
	if err != nil {
		return out, err
	}
	return out, nil
}
