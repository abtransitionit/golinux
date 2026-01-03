package property

import (
	"fmt"
	"os"
	"strings"

	cproperty "github.com/abtransitionit/gocore/mock/property"
	"github.com/abtransitionit/golinux/mock/run"
	"github.com/opencontainers/selinux/go-selinux"
)

func getEnvar(params ...string) (string, error) {
	// get parameter
	if len(params) < 1 {
		return "", fmt.Errorf("envar name required")
	}
	envarName := params[0]
	// get
	result := os.Getenv(envarName)
	if strings.TrimSpace(result) == "" {
		return "", fmt.Errorf("environment variable %s is not set", envarName)
	}
	// handle success
	return result, nil
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
func getOsArch2(_ ...string) (string, error) {
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

func getNeedReboot(_ ...string) (string, error) {
	// 1 - get property
	osType, err := cproperty.GetOsType()
	if err != nil {
		return "", fmt.Errorf("could not detect OS type: %w", err)
	}
	osFamily, err := cproperty.GetOsFamily()
	if err != nil {
		return "", err
	}
	// 2 - manage only linux
	if osType != "linux" {
		return "", err
	}

	// 3 - define CLI
	var cli string
	switch strings.TrimSpace(osFamily) {
	case "debian":
		cli = "test -f /var/run/reboot-required && echo true || echo false"
	case "rhel", "fedora":
		cli = "	command -v needs-restarting >/dev/null && needs-restarting -r | grep -q 'Reboot is required' && echo true || echo false"
	default:
		return "", fmt.Errorf("unsupported OS family: %s", osFamily)
	}
	// 4 - run cli
	output, err := run.RunCli("local", cli, nil)
	// handle system error
	if err != nil {
		return "", fmt.Errorf("getting reboot status: %w", err)
	}
	// handle success
	return strings.TrimSpace(output), nil
}
