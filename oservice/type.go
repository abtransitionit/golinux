package oservice

import (
	"context"
	"fmt"

	"github.com/abtransitionit/gocore/logx"
)

type OsService struct {
	Name    string // logical name
	CName   string // canonical name
	Path    string // path to the unit file
	Content string // content of the unit file
}

type SliceOsService []OsService
type MapOsService map[string]OsService

// Usage:
//
//	cname, err = s2.GetCName()
//	if err != nil {
//	  fmt.Println("Error:", err) // Error: OsService with Name "nginx" not found in OsServiceReference
//	}
func (m MapOsService) GetCName(s OsService) (string, error) {
	if result, ok := m[s.Name]; ok {
		return result.CName, nil
	}
	return "", fmt.Errorf("OsService with Name %q not found in MapOsService", s.Name)
}

func (osService OsService) GetName() string {
	return osService.Name
}

func SetContent(ctx context.Context, logger logx.Logger, service OsService) (string, error) {
	logger.Infof("To implement: set the content of the unit file %s", service.Path)
	return "", nil
}

func (list SliceOsService) GetListName() []string {
	names := make([]string, 0, len(list))
	for _, s := range list {
		names = append(names, s.Name)
	}
	return names
}
