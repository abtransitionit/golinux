package onpm

import (
	"context"
	"fmt"

	"github.com/abtransitionit/gocore/logx"
)

// func NewPackage(name, version, osType string, repo *Repo) (*Package, error) {
// 	p := &Package{Name: name, Version: version}

// 	cbd, err := p.GetCliBuilder(osType, repo)
// 	if err != nil {
// 		return nil, err
// 	}
// 	p.Cbd = cbd

// 	return p, nil
// }

// Package
func (p *Package) Add(ctx context.Context, local bool, remoteHost string, logger logx.Logger) (string, error) {
	return p.Cbd.CliAdd()
}
func (p *Package) Delete(ctx context.Context, local bool, remoteHost string, logger logx.Logger) (string, error) {
	return p.Cbd.CliDelete()
}
func (p *Package) List(ctx context.Context, local bool, remoteHost string, logger logx.Logger) (string, error) {
	return p.Cbd.CliList()
}
func (p *Package) Upgrade(ctx context.Context) (string, error) {
	return fmt.Sprintf("apt-get upgrade -y %s", p.Name), nil
}
