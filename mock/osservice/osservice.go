package osservice

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

// Description: Starts the service for the current session (at runtime)
func (i *Service) Start(hostName string, logger logx.Logger) error {
	// 1 - get and play cli
	if _, err := run.RunCli(hostName, i.cliToStart(), logger); err != nil {
		return err
	}
	// handle success
	logger.Debugf("%s:%s > started OS service for the current session", hostName, i.Name)
	return nil
}
func (i *Service) Stop(hostName string, logger logx.Logger) error {
	// 1 - get and play cli
	if _, err := run.RunCli(hostName, i.cliToStop(), logger); err != nil {
		return err
	}
	// handle success
	logger.Debugf("%s:%s > stopped OS service for the current session", hostName, i.Name)
	return nil
}

// Description: enables a service to start after at host reboot (at startup)
func (i *Service) Enable(hostName string, logger logx.Logger) error {
	// 1 - get and play cli
	if _, err := run.RunCli(hostName, i.cliToEnable(), logger); err != nil {
		return err
	}
	// handle success
	logger.Debugf("%s:%s > enabled OS service at startup", hostName, i.Name)
	return nil
}

func (i *Service) cliToEnable() string {
	var cmds = []string{
		"sudo systemctl daemon-reload",
		fmt.Sprintf("sudo systemctl enable %s", i.Name),
	}
	cli := strings.Join(cmds, " && ")
	return cli
}
func (i *Service) cliToStart() string {
	var cmds = []string{
		"sudo systemctl daemon-reload",
		fmt.Sprintf("sudo systemctl start %s", i.Name),
	}
	cli := strings.Join(cmds, " && ")
	return cli
}
func (i *Service) cliToStop() string {
	var cmds = []string{
		"sudo systemctl daemon-reload",
		fmt.Sprintf("sudo systemctl stop %s", i.Name),
	}
	cli := strings.Join(cmds, " && ")
	return cli
}

// Description: installs a service using its configuration file
func (i *Service) Install(hostName string, logger logx.Logger) error {

	// log
	logger.Infof("%s > installing OS service: %s", hostName, i.Name)

	// handle success
	return nil
}

// linux cli to check a service status
// s=crio;for a in o1u o2a o3r o4f o5d;    do ssh "$a" "printf '%s - ' '$a'; systemctl status $s  2>/dev/null  || echo disabled"; done
// s=kubelet;for a in o1u o2a o3r o4f o5d; do ssh "$a" "printf '%s - ' '$a'; systemctl status $s  2>/dev/null  || echo disabled"; done
// s=sshd;for a in o1u o2a o3r o4f o5d;    do ssh "$a" "printf '%s - ' '$a'; systemctl status $s  2>/dev/null  || echo disabled"; done

// s=crio;for a in o1u o2a o3r o4f o5d;    do ssh "$a" "printf '%s - ' '$a'; systemctl status $s  2>/dev/null  | head -1"; done
// s=kubelet;for a in o1u o2a o3r o4f o5d; do ssh "$a" "printf '%s - ' '$a'; systemctl status $s  2>/dev/null  | head -1"; done
// s=sshd;for a in o1u o2a o3r o4f o5d;    do ssh "$a" "printf '%s - ' '$a'; systemctl status $s  2>/dev/null  | head -1"; done

// s=crio;for a in o1u o2a o3r o4f o5d;    do ssh "$a" "printf '%s - ' '$a'; systemctl is-enabled $s  2>/dev/null  | head -1"; done
// s=kubelet;for a in o1u o2a o3r o4f o5d; do ssh "$a" "printf '%s - ' '$a'; systemctl is-enabled $s  2>/dev/null  | head -1"; done
// s=sshd;for a in o1u o2a o3r o4f o5d;    do ssh "$a" "printf '%s - ' '$a'; systemctl is-enabled $s  2>/dev/null  | head -1"; done

// s=crio;for a in o1u o2a o3r o4f o5d;    do ssh "$a" "printf '%s - ' '$a'; systemctl is-active $s  2>/dev/null  | head -1"; done
// s=kubelet;for a in o1u o2a o3r o4f o5d; do ssh "$a" "printf '%s - ' '$a'; systemctl is-active $s  2>/dev/null  | head -1"; done
// s=sshd;for a in o1u o2a o3r o4f o5d;    do ssh "$a" "printf '%s - ' '$a'; systemctl is-active $s  2>/dev/null  | head -1"; done

// s=crio;for a in o1u o2a o3r o4f o5d;    do ssh "$a" "printf '%s - ' '$a'; systemctl status $s | egrep -i act 2>/dev/null | head -1 || echo disabled"; done
// s=crio;for a in o1u o2a o3r o4f o5d;    do ssh "$a" "printf '%s - ' '$a'; systemctl status $s | egrep -i loa 2>/dev/null | head -1 || echo disabled"; done
// s=crio;for a in o1u o2a o3r o4f o5d;    do ssh "$a" "printf '%s - ' '$a'; systemctl status $s | egrep -i 'act|loa' 2>/dev/null  || echo disabled"; done
// s=kubelet;for a in o1u o2a o3r o4f o5d; do ssh "$a" "printf '%s - ' '$a'; systemctl status $s | grep -i active 2>/dev/null | head -1 || echo disabled"; done
// s=sshd;for a in o1u o2a o3r o4f o5d;    do ssh "$a" "printf '%s - ' '$a'; systemctl status $s | grep -i active 2>/dev/null | head -1 || echo disabled"; done

// s=sshd;for a in o1u o2a o3r o4f o5d; do ssh "$a" "printf '%s - ' '$a'; systemctl is-enabled $s 2>/dev/null || echo disabled"; done
// s=sshd;for a in o1u o2a o3r o4f o5d; do echo $a; ssh "$a" "systemctl is-enabled $s 2>/dev/null || echo disabled"; done
// s=sshd;for a in o1u o2a o3r o4f o5d; do echo $a; ssh "$a" "systemctl status $s 2>/dev/null || echo disabled"; done
// s=sshd;for a in o1u o2a o3r o4f o5d; do echo $a; ssh "$a" "systemctl status $s 2>/dev/null | head -1 || echo disabled"; done
// s=sshd;for a in o1u o2a o3r o4f o5d; do ssh "$a" "printf '%s - ' '$a'; systemctl status $s 2>/dev/null | head -1 || echo disabled"; done
// s=sshd;for a in o1u o2a o3r o4f o5d; do echo $a; ssh "$a" "systemctl status $s | grep Active 2>/dev/null || echo disabled"; done
