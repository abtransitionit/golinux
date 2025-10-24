package da

import (
	"context"
	"fmt"

	"github.com/abtransitionit/gocore/logx"
)

func NewPackage(name, version, osType string, repo *Repo) (*Package, error) {
	p := &Package{Name: name, Version: version}

	mgr, err := p.GetManager(osType, repo)
	if err != nil {
		return nil, err
	}
	p.Mgr = mgr

	return p, nil
}

// Package
func (p *Package) Add(ctx context.Context, local bool, remoteHost string, logger logx.Logger) (string, error) {
	return p.Mgr.CliAdd()
}
func (p *Package) Delete(ctx context.Context, local bool, remoteHost string, logger logx.Logger) (string, error) {
	return p.Mgr.CliDelete()
}
func (p *Package) List(ctx context.Context, local bool, remoteHost string, logger logx.Logger) (string, error) {
	return p.Mgr.CliList()
}
func (p *Package) Upgrade(ctx context.Context) (string, error) {
	return fmt.Sprintf("apt-get upgrade -y %s", p.Name), nil
}
