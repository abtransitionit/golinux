package property

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
	"strings"

	"github.com/abtransitionit/golinux/mock/run"
	"github.com/opencontainers/selinux/go-selinux"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

func getPath(_ ...string) (string, error) {
	path := os.Getenv("PATH")
	if path == "" {
		return "", fmt.Errorf("PATH environment variable is not set")
	}
	return path, nil
}

func isServiceEnabled(params ...string) (string, error) {
	// get parameter
	if len(params) < 1 {
		return "", fmt.Errorf("service name required")
	}
	service := params[0]
	// define CLI
	cli := fmt.Sprintf("systemctl is-enabled %s", service)
	// run cli
	output, err := run.RunCli("local", cli, nil)
	// handle system error
	if err != nil {
		return "", fmt.Errorf("getting service status: %w", err)
	}
	// handle success
	return output, nil
}

func isServiceActive(params ...string) (string, error) {
	// get parameter
	if len(params) < 1 {
		return "", fmt.Errorf("service name required")
	}
	service := params[0]
	// define CLI
	cli := fmt.Sprintf("systemctl is-active %s", service)
	// run cli
	output, err := run.RunCli("local", cli, nil)

	// in theses case the cli returns err = nil
	if output == "active" || output == "inactive" || output == "failed" {
		return output, nil
	}

	// manage other real errors
	if err != nil {
		return "", fmt.Errorf("getting service status: %w", err)
	}

	// handle other success
	return strings.TrimSpace(string(output)), nil
}

func getSelinuxMode(_ ...string) (string, error) {
	switch selinux.EnforceMode() {
	case selinux.Enforcing:
		return "enforcing", nil
	case selinux.Permissive:
		return "permissive", nil
	case selinux.Disabled:
		return "disabled", nil
	default:
		return "unknown", nil
	}
}

func getSelinuxStatus(_ ...string) (string, error) {
	if selinux.GetEnabled() {
		return "enabled", nil
	}
	return "disabled", nil
}

func getUuid(_ ...string) (string, error) {
	// define CLI
	cli := "sudo cat /sys/class/dmi/id/product_uuid"
	// run cli
	output, err := run.RunCli("local", cli, nil)
	// handle system error
	if err != nil {
		return "", fmt.Errorf("getting uuid: %w", err)
	}
	// handle success
	return output, nil
}
func getUnameM(_ ...string) (string, error) {
	// define CLI
	cli := "uname -m"
	// run cli
	output, err := run.RunCli("local", cli, nil)
	// handle system error
	if err != nil {
		return "", fmt.Errorf("getting uname: %w", err)
	}
	// handle success
	return output, nil
}

func getRcFilePath(params ...string) (string, error) {
	// define CLI
	cli := "$HOME/.profile"
	// run cli
	output, err := run.RunCli("local", cli, nil)
	// handle system error
	if err != nil {
		return "", fmt.Errorf("getting rc file path: %w", err)
	}
	// handle success
	return output, nil
}

func getPathTree(params ...string) (string, error) {
	// check params
	if len(params) < 1 {
		return "", fmt.Errorf("base path name required. Eg. `/`")
	}
	// get input
	basePath := params[0]
	// define CLI
	cli := fmt.Sprintf(`find %s -type d | sort | paste -sd\:`, basePath)
	// run cli
	output, err := run.RunCli("local", cli, nil)
	// handle system error
	if err != nil {
		return "", fmt.Errorf("getting path tree: %w", err)
	}
	// handle success
	return output, nil
}

func getOsDistro(_ ...string) (string, error) {
	info, err := host.Info()
	if err != nil {
		return "", err
	}
	return info.Platform, nil
}

func getOsFamily(_ ...string) (string, error) {
	info, err := host.Info()
	if err != nil {
		return "", fmt.Errorf("getting os family > %v", err)
	}
	return info.PlatformFamily, nil
}

func getOsVersion(_ ...string) (string, error) {
	info, err := host.Info()
	if err != nil {
		return "", fmt.Errorf("getting os version > %v", err)
	}
	return info.PlatformVersion, nil
}

func getOsUser(_ ...string) (string, error) {
	output, err := user.Current()
	if err != nil {
		return "", err
	}
	return output.Username, nil
}

func getOsKernelVersion(_ ...string) (string, error) {
	info, err := host.Info()
	if err != nil {
		return "", fmt.Errorf("getting os kernel version > %v", err)
	}
	return info.KernelVersion, nil
}

func getCpu(_ ...string) (string, error) {
	output, err := cpu.Info()
	if err != nil {
		return "", fmt.Errorf("getting cpu info > %v", err)
	}
	return fmt.Sprintf("%v", output[0].Cores), nil
}

func getRam01(_ ...string) (string, error) {
	output, err := cpu.Info()
	if err != nil {
		return "", fmt.Errorf("getting cpu info > %v", err)
	}
	return fmt.Sprintf("%v", output[0].Cores), nil
}

func getRam02(_ ...string) (string, error) {
	output, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", output.Total/(1024*1024*1024)), nil
}

func getOsArch(_ ...string) (string, error) {
	return runtime.GOARCH, nil // go env GOARCH
}

// ---------------- TODO -------------------------------- TODO -------------------------------- TODO ----------------
// ---------------- TODO -------------------------------- TODO -------------------------------- TODO ----------------
// ---------------- TODO -------------------------------- TODO -------------------------------- TODO ----------------

func getOsType(_ ...string) (string, error) {
	return runtime.GOOS, nil // go env GOOS
}

func getOsInfos(_ ...string) (string, error) {
	family, err := getOsFamily()
	if err != nil {
		return "", err
	}

	distro, err := getOsDistro()
	if err != nil {
		return "", err
	}

	version, err := getOsVersion()
	if err != nil {
		return "", err
	}

	kernel, err := getOsKernelVersion()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("family: %-6s :: distro: %-10s :: OsVersion: %-6s :: OsKernelVersion: %s", family, distro, version, kernel), nil
}

func getServiceInfos(params ...string) (string, error) {
	if len(params) < 1 {
		return "", fmt.Errorf("service name required")
	}
	// get service
	serviceName := params[0]

	// get
	isActive, err := isServiceActive(serviceName)
	if err != nil {
		return "", fmt.Errorf("serviceStatus: %v", err)
	}

	// get
	isEnabled, err := isServiceEnabled(serviceName)
	if err != nil {
		return "", fmt.Errorf("serviceEnabled: %v", err)
	}

	// return
	return fmt.Sprintf("%-6s / %-6s", isActive, isEnabled), nil

}

func getSelinuxInfos(_ ...string) (string, error) {
	status, err := getSelinuxStatus()
	if err != nil {
		return "", fmt.Errorf("selstatus: %v", err)
	}

	mode, err := getSelinuxMode()
	if err != nil {
		return "", fmt.Errorf("selmode: %v", err)
	}

	return fmt.Sprintf("status: %-10s :: mode: %s", status, mode), nil
}

func getNetIp(_ ...string) (string, error) {
	// define CLI
	cli := "curl -s ifconfig.me -4"
	// run cli
	output, err := run.RunCli("local", cli, nil)
	// handle system error
	if err != nil {
		return "", fmt.Errorf("getting net-ip > %v", err)
	}
	// handle success
	return output, nil
}

func getNetGateway(_ ...string) (string, error) {
	// define CLI
	cli := "ip route get 2.2.2.2"
	// run cli
	output, err := run.RunCli("local", cli, nil)
	// handle system error
	if err != nil {
		return "", fmt.Errorf("getting net-ip > %v", err)
	}
	// First line only
	line := strings.Split(output, "\n")[0]
	// handle success
	return strings.TrimSpace(line), nil
}
