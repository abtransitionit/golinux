package dnfapt

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/abtransitionit/gocore/logx"
)

// Name: InstallPackage
//
// Description: install a dnfapt package on a Linux distro
func InstallDaRepository(ctx context.Context, logger logx.Logger, osFamily string, daRepo DaRepo) (string, error) {

	if osFamily != "rhel" && osFamily != "fedora" && osFamily != "debian" {
		return "", fmt.Errorf("this function only supports Linux (rhel, fedora, debian), but found: %s", osFamily)
	}

	// lookup the organization's reference db - get templated URL of the repo
	daRepoRef, ok := MapDaRepoReference[daRepo.Name]
	if !ok {
		return "", fmt.Errorf("found no matches for this package repo: %s", daRepo.Name)
	}
	// lookup the organization's reference db - get CTE OS specific repo data
	daRepoRefCte, ok := MapDaRepoCteReference[osFamily]
	if !ok {
		return "", fmt.Errorf("found no matches for repository CTE for this os family: %s", osFamily)
	}
	// lookup the organization's reference db - get templated Repo file content
	daRepoTplFileContent, ok := MapDaRepoTplFileContent[osFamily]
	if !ok {
		return "", fmt.Errorf("found no matches for repository TPL file for this os family: %s", osFamily)
	}

	// define var from these infos - logic common to all OS families
	urlRepo := ResolveURLRepo(daRepoRef.UrlRepo, daRepo.Version, daRepoRefCte.Pack)
	urlGpg := ResolveURLGpg(daRepoRef.UrlGpg, daRepo.Version, daRepoRefCte.Pack, daRepoRefCte.Gpg)
	repoFilePath := filepath.Join(daRepoRefCte.Folder, daRepo.FileName+daRepoRefCte.Ext)

	// logic specific to OS family
	var gpgFilePath string
	if daRepoRefCte.GpgFolder != "" && daRepoRefCte.GpgExt != "" {
		gpgFilePath = filepath.Join(daRepoRefCte.GpgFolder, daRepo.FileName+daRepoRefCte.GpgExt)
	}

	// define the vars/structure that will be used into a template (the repo file content)
	daRepoTplFileContentVar := RepoFileContentVar{
		RepoName:    daRepo.Name,
		UrlRepo:     urlRepo,
		UrlGpg:      urlGpg,
		GpgFilePath: gpgFilePath,
	}

	_, daRepoFileContent := ResolveRepoFileContent(daRepoTplFileContent, daRepoTplFileContentVar)

	// log
	if gpgFilePath != "" {
		logger.Debugf("🅰️ Gpg file path is: %s", gpgFilePath)
	}
	logger.Debugf("🅰️ Repo file path is: %s", repoFilePath)
	logger.Debugf("🅰️ UrlRepo is: %s", urlRepo)
	logger.Debugf("🅰️ UrlGpg  is: %s", urlGpg)
	fmt.Println(daRepoFileContent)

	return "", nil
}

func ResolveRepoFileContent(tplRepoContent string, setVar RepoFileContentVar) (string, error) {
	return tplRepoContent, nil
}
func ResolveURLRepo(tplUrlRepo string, tag string, pack string) string {
	return substituteUrlRepoPlaceholders(tplUrlRepo, tag, pack)
}
func ResolveURLGpg(tplUrlGpg string, tag string, pack string, gpg string) string {
	return substituteUrlGpgPlaceholders(tplUrlGpg, tag, pack, gpg)
}

func substituteUrlRepoPlaceholders(tplDaRepoUrl string, tag string, pack string) string {

	replacements := map[string]string{
		"$TAG":  tag,
		"$PACK": pack,
	}
	url := tplDaRepoUrl
	for k, v := range replacements {
		url = strings.ReplaceAll(url, k, v)
	}
	return url
}

func substituteUrlGpgPlaceholders(tplDaRepoUrl string, tag string, pack string, gpg string) string {

	replacements := map[string]string{
		"$TAG":  tag,
		"$PACK": pack,
		"$GPG":  gpg,
	}
	url := tplDaRepoUrl
	for k, v := range replacements {
		url = strings.ReplaceAll(url, k, v)
	}
	return url
}
