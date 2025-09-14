package dnfapt

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
	"strings"
)

func GetPackCName(daPack DaPack) (string, error) {

	// lookup the organization's reference db - get obj thet contains the property of the package
	packName := daPack.Name
	daPackRef, ok := MapDaPackReference[packName]
	if !ok {
		return "", fmt.Errorf("found no matches for dnfapt package '%s' in reference DB", packName)
	}
	return daPackRef.CName, nil
}
func GetRepoFilePath(osFamily string, daRepo DaRepo) (string, error) {
	// lookup the organization's reference db - get obj thet contains the OS specific CTE of the repo
	daRepoRefCte, ok := MapDaRepoCteReference[osFamily]
	if !ok {
		return "", fmt.Errorf("found no matches for OS family '%s' in repository CTE reference DB", osFamily)
	}
	return filepath.Join(daRepoRefCte.Folder, daRepo.FileName+daRepoRefCte.Ext), nil
}
func GetGpgFilePath(osFamily string, daRepo DaRepo) (string, error) {
	// lookup the organization's reference db - get obj thet contains the OS specific CTE of the repo
	daRepoRefCte, ok := MapDaRepoCteReference[osFamily]
	if !ok {
		return "", fmt.Errorf("found no matches for repository CTE for this os family: %s", osFamily)
	}
	return filepath.Join(daRepoRefCte.GpgFolder, daRepo.FileName+daRepoRefCte.GpgExt), nil
}

func GetUrlGpgResolved(osFamily string, daRepo DaRepo) (string, error) {
	// lookup the organization's reference db - get obj that contains the templated URL of the repo
	daRepoRef, ok := MapDaRepoReference[daRepo.Name]
	if !ok {
		return "", fmt.Errorf("found no matches for this package repo: %s", daRepo.Name)
	}

	// lookup the organization's reference db - get obj thet contains the OS specific CTE of the repo
	daRepoRefCte, ok := MapDaRepoCteReference[osFamily]
	if !ok {
		return "", fmt.Errorf("found no matches for repository CTE for this os family: %s", osFamily)
	}
	return ResolveURLGpg(daRepoRef.UrlGpg, daRepo.Version, daRepoRefCte.Pack, daRepoRefCte.Gpg), nil
}

func GetRepoFileContent(osFamily string, daRepo DaRepo) (string, error) {

	// lookup the organization's reference db - get obj that contains the templated URL of the repo
	daRepoRef, ok := MapDaRepoReference[daRepo.Name]
	if !ok {
		return "", fmt.Errorf("found no matches for this package repo: %s", daRepo.Name)
	}

	// lookup the organization's reference db - get obj thet contains the OS specific CTE of the repo
	daRepoRefCte, ok := MapDaRepoCteReference[osFamily]
	if !ok {
		return "", fmt.Errorf("found no matches for repository CTE for this os family: %s", osFamily)
	}

	// lookup the organization's reference db - get obj that contains the templated file of the repo
	daRepoTplFileContent, ok := MapDaRepoTplFileContent[osFamily]
	if !ok {
		return "", fmt.Errorf("found no matches for repository TPL file for this os family: %s", osFamily)
	}

	// define var from these infos - logic common to all OS families
	urlRepo := ResolveURLRepo(daRepoRef.UrlRepo, daRepo.Version, daRepoRefCte.Pack)
	urlGpg := ResolveURLGpg(daRepoRef.UrlGpg, daRepo.Version, daRepoRefCte.Pack, daRepoRefCte.Gpg)

	// logic specific to OS family
	var gpgFilePath string
	if daRepoRefCte.GpgFolder != "" && daRepoRefCte.GpgExt != "" {
		gpgFilePath = filepath.Join(daRepoRefCte.GpgFolder, daRepo.FileName+daRepoRefCte.GpgExt)
	}

	// define the structure that holds the vars that will be used to resolve the templated file
	daRepoTplFileContentVar := RepoFileContentVar{
		RepoName:    daRepo.Name,
		UrlRepo:     urlRepo,
		UrlGpg:      urlGpg,
		GpgFilePath: gpgFilePath,
	}

	// resolve the templated file
	daRepoFileContent, err := ResolveRepoFileContent(daRepoTplFileContent, daRepoTplFileContentVar)
	if err != nil {
		return "", fmt.Errorf("failde to resolve templated repo file, for the file: %s", daRepoTplFileContent)
	}
	return daRepoFileContent, nil
}

func ResolveRepoFileContent(tplRepoContent string, vars RepoFileContentVar) (string, error) {
	tpl, err := template.New("repo").Parse(tplRepoContent)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, vars); err != nil {
		return "", err
	}

	return buf.String(), nil
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

func DaRepoUpdate(osFamily string) (string, error) {
	var cmds []string
	switch osFamily {

	case "rhel", "fedora":
		cmds = []string{
			"sudo dnf update -q -y > /dev/null",
		}

	case "debian":
		cmds = []string{
			"sudo apt update -qq -y > /dev/null",
		}

	default:
		return "", fmt.Errorf("this function only supports Linux (rhel, fedora, debian), but found: %s", osFamily)
	}
	cli := strings.Join(cmds, " && ")
	return cli, nil

}

// func ResolveFileContent(tplContent string, data RepoData) (string, error) {
//     // 1. Create a new template and parse the provided string.
//     tpl, err := template.New("repo_file").Parse(tplContent)
//     if err != nil {
//         return "", fmt.Errorf("failed to parse template: %w", err)
//     }

//     // 2. Create a buffer to write the output to.
//     var buf bytes.Buffer

//     // 3. Execute the template with the provided data.
//     if err := tpl.Execute(&buf, data); err != nil {
//         return "", fmt.Errorf("failed to execute template: %w", err)
//     }

//     // Return the resolved string from the buffer.
//     return buf.String(), nil
// }

// Name: InstallPackage
//
// Description: install a dnfapt package on a Linux distro
// func InstallDaRepository(ctx context.Context, logger logx.Logger, osFamily string, daRepo DaRepo) (string, error) {

// 	if osFamily != "rhel" && osFamily != "fedora" && osFamily != "debian" {
// 		return "", fmt.Errorf("this function only supports Linux (rhel, fedora, debian), but found: %s", osFamily)
// 	}

// 	// lookup the organization's reference db - get obj that contains the templated URL of the repo
// 	daRepoRef, ok := MapDaRepoReference[daRepo.Name]
// 	if !ok {
// 		return "", fmt.Errorf("found no matches for this package repo: %s", daRepo.Name)
// 	}
// 	// lookup the organization's reference db - get obj thet contains the OS specific CTE of the repo
// 	daRepoRefCte, ok := MapDaRepoCteReference[osFamily]
// 	if !ok {
// 		return "", fmt.Errorf("found no matches for repository CTE for this os family: %s", osFamily)
// 	}
// 	// lookup the organization's reference db - get obj that contains the templated file of the repo
// 	daRepoTplFileContent, ok := MapDaRepoTplFileContent[osFamily]
// 	if !ok {
// 		return "", fmt.Errorf("found no matches for repository TPL file for this os family: %s", osFamily)
// 	}

// 	// define var from these infos - logic common to all OS families
// 	urlRepo := ResolveURLRepo(daRepoRef.UrlRepo, daRepo.Version, daRepoRefCte.Pack)
// 	urlGpg := ResolveURLGpg(daRepoRef.UrlGpg, daRepo.Version, daRepoRefCte.Pack, daRepoRefCte.Gpg)
// 	repoFilePath := filepath.Join(daRepoRefCte.Folder, daRepo.FileName+daRepoRefCte.Ext)

// 	// logic specific to OS family
// 	var gpgFilePath string
// 	if daRepoRefCte.GpgFolder != "" && daRepoRefCte.GpgExt != "" {
// 		gpgFilePath = filepath.Join(daRepoRefCte.GpgFolder, daRepo.FileName+daRepoRefCte.GpgExt)
// 	}

// 	// define the structure that holds the vars that will be used to resolve the templated file
// 	daRepoTplFileContentVar := RepoFileContentVar{
// 		RepoName:    daRepo.Name,
// 		UrlRepo:     urlRepo,
// 		UrlGpg:      urlGpg,
// 		GpgFilePath: gpgFilePath,
// 	}

// 	// resolve the templated file
// 	daRepoFileContent, err := ResolveRepoFileContent(daRepoTplFileContent, daRepoTplFileContentVar)
// 	if err != nil {
// 		return "", fmt.Errorf("failde to resolve templated repo file, for the file: %s", daRepoTplFileContent)
// 	}

// 	// // write the file
// 	// if err := WriteFile(repoFilePath, daRepoFileContent); err != nil {
// 	// 	return "", fmt.Errorf("failde to write repo file: %s", err)
// 	// }

// 	// log
// 	// if gpgFilePath != "" {
// 	// 	logger.Debugf("🅰️ Gpg file path is: %s", gpgFilePath)
// 	// }
// 	// logger.Debugf("🅰️ Repo file path is: %s", repoFilePath)
// 	// logger.Debugf("🅰️ UrlRepo is: %s", urlRepo)
// 	// logger.Debugf("🅰️ UrlGpg  is: %s", urlGpg)
// 	// logger.Debugf("🅰️ Resolved repo file content:\n%s", daRepoFileContent)
// 	// fmt.Println(daRepoFileContent)

// 	return "", nil
// }
