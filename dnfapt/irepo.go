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

	// log
	if gpgFilePath != "" {
		logger.Debugf("🅰️ Gpg file path is: %s", gpgFilePath)
	}
	logger.Debugf("🅰️ Repo file path is: %s", repoFilePath)
	logger.Debugf("🅰️ UrlRepo is: %s", urlRepo)
	logger.Debugf("🅰️ UrlGpg  is: %s", urlGpg)
	fmt.Println(daRepoTplFileContent)

	return "", nil
}

func ResolveURLRepo(tplUrlRepo string, tag string, pack string) string {
	return substituteUrlRepoPlaceholders(tplUrlRepo, tag, pack)
}
func ResolveURLGpg(tplUrlGpg string, tag string, pack string, gpg string) string {
	return substituteUrlGpgPlaceholders(tplUrlGpg, tag, pack, gpg)
}

// func ResolveURLGpg(logger logx.Logger, osType string, osArch string, uname string) (string, error) {
// 	return substituteUrlGpgPlaceholders(goCliRef.Url, cli, tag, osType, osArch, uname), nil(goCliRef.Url, cli, tag, osType, osArch, uname), nil
// 	return "", nil
// }

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

// func ResolveFileContent(tplContent string, data RepoData) (string, error) {
// 	// 1. Create a new template and parse the provided string.
// 	tpl, err := template.New("repo_file").Parse(tplContent)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to parse template: %w", err)
// 	}

// 	// 2. Create a buffer to write the output to.
// 	var buf bytes.Buffer

// 	// 3. Execute the template with the provided data.
// 	if err := tpl.Execute(&buf, data); err != nil {
// 		return "", fmt.Errorf("failed to execute template: %w", err)
// 	}

// 	// Return the resolved string from the buffer.
// 	return buf.String(), nil
// }
