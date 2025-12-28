package gopm

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/property"
)

func (i *Cli) Install(hostName string, logger logx.Logger) error {
	// define var
	var osFamily, osDistro string
	var err error
	// 1 - get host:property
	osFamily, err = property.GetProperty(logger, hostName, "osFamily")
	if err != nil {
		return err
	}
	osDistro, err = property.GetProperty(logger, hostName, "osDistro")
	if err != nil {
		return err
	}

	// log
	logger.Infof("%s:%s:%s > install %s:%s from url", hostName, osFamily, osDistro, i.Name, i.Version)
	// get the CLI from the YAML file
	cli, err := GetCliFromYaml(osFamily, osDistro, i.Name)
	if err != nil {
		return fmt.Errorf("%s:%s:%s > %w", hostName, osFamily, osDistro, err)
	}
	// run the install
	fmt.Printf("%v\n", cli)

	// handle success
	return nil
}

func (i *CliSlice) Install(logger logx.Logger) error {
	// log
	logger.Info("Install called")
	// handle success
	return nil
}
