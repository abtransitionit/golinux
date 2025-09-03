package dnfapt

import (
	"context"
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
)

// Name: InstallPackage
//
// Description: install a dnfapt package on a Linux distro
func InstallDaRepository(ctx context.Context, logger logx.Logger, osFamily string, daRepository DaRepository) (string, error) {

	if osFamily != "rhel" && osFamily != "fedora" && osFamily != "debian" {
		return "", fmt.Errorf("this function only supports Linux (rhel, fedora, debian), but found: %s", osFamily)
	}

	// lookup and get in the reference database that daRepository object
	DaRepositoryRef, ok := DaRepositoryReference[daRepository.Name]
	if !ok {
		return "", fmt.Errorf("failed to fetch in reference db : %s", daRepository.Name)
	}
	// check the field is not empty
	if DaRepositoryRef.UrlRepo == "" {
		return "", fmt.Errorf("no URL template defined for %s", daRepository.Name)
	}
	// check the field is not empty
	if DaRepositoryRef.UrlGpg == "" {
		return "", fmt.Errorf("no URL template defined for %s", daRepository.Name)
	}

	// // logic for installtion
	// switch osFamily {
	// case "rhel", "fedora":
	// case "debian":
	// }

	logger.Debugf("🅰️ resolving UrlRepo and UrlGpg for %s:%s", osFamily, daRepository.Name)
	logger.Debugf("🅰️ templated UrlRepo is: %s", DaRepositoryRef.UrlRepo)
	logger.Debugf("🅰️ templated UrlGpg  is: %s", DaRepositoryRef.UrlGpg)

	return "", nil
}

func ResolveURLRepo(logger logx.Logger, osType string, osArch string, uname string) (string, error) {
	return substituteUrlRepoPlaceholders(goCliRef.Url, cli, tag, osType, osArch, uname), nil(goCliRef.Url, cli, tag, osType, osArch, uname), nil
	return "", nil
}
func ResolveURLGpg(logger logx.Logger, osType string, osArch string, uname string) (string, error) {
	return substituteUrlGpgPlaceholders(goCliRef.Url, cli, tag, osType, osArch, uname), nil(goCliRef.Url, cli, tag, osType, osArch, uname), nil
	return "", nil
}

func substituteUrlRepoPlaceholders(tplDaRepoUrl string, tag string, pack string) string {

	replacements := map[string]string{
		"$TAG":  tag,
		"$PACK": osType,
	}
	url := tplDaRepoUrl
	for k, v := range replacements {
		url = strings.ReplaceAll(url, k, v)
	}
	return url
}

func substituteUrlGpgPlaceholders(tplDaRepoUrl string, cli GoCli, tag string, osType string, osArch string, uname string) string {

	replacements := map[string]string{
		"$TAG":  tag,
		"$PACK": osType,
		"$GPG":  osArch,
	}
	url := tplDaRepoUrl
	for k, v := range replacements {
		url = strings.ReplaceAll(url, k, v)
	}
	return url
}
