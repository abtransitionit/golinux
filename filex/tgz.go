package filex

import (
	"fmt"
	"strings"
)

func DetectTgzFolder(filePath string) string {
	return fmt.Sprintf(`tar -tzf %q | awk -F/ 'NF>1 {print $1"/"}' | sort -u | wc -l`, filePath)
}

func DetectTgzFile(filePath string) string {
	return fmt.Sprintf("tar -tzf %q | wc -l", filePath)
}

func CpTgzFileLevel1(tgzFilePath, goCliName string) string {
	var cmds = []string{
		fmt.Sprintf("sudo mkdir -p /usr/local/bin/%s ", goCliName),
		fmt.Sprintf("sudo tar -xvzf %q -C /usr/local/bin/%s --strip-components=1", tgzFilePath, goCliName),
	}

	cli := strings.Join(cmds, " && ")
	return cli

}

func CpTgzFileLevel0(tgzFilePath, goCliName string) string {
	var cmds = []string{
		fmt.Sprintf("sudo mkdir -p /usr/local/bin/%s ", goCliName),
		fmt.Sprintf("sudo tar -xvzf %q -C /usr/local/bin/%q", tgzFilePath, goCliName),
	}

	cli := strings.Join(cmds, " && ")
	return cli

}

cmds := []string{
	fmt.Sprintf(`nbFolder=$(tar -tzf %q | awk -F/ 'NF>1 {print $1"/"}' | sort -u | wc -l)`, filePath),
	fmt.Sprintf("sudo mkdir -p /usr/local/bin/%s", goCliName),
	fmt.Sprintf(`if [ "$nbFolder" -eq 1 ]; then sudo tar -xvzf %q -C /usr/local/bin/%s --strip-components=1; fi`, filePath, goCliName),
	fmt.Sprintf(`sudo tar -xvzf %q -C /usr/local/bin/%s`, filePath, goCliName),
}

cli := strings.Join(cmds, " && ")

func tgzTest(filePath, goCliName string) string {
	var cmds = []string{
		fmt.Sprintf(`nbFolder=$(tar -tzf %q | awk -F/ 'NF>1 {print $1"/"}' | sort -u | wc -l)`, filePath),
		fmt.Sprintf(`sudo mkdir -p /usr/local/bin/%s`, goCliName),
		fmt.Sprintf(`if [ "$nbFolder" -eq 0 ]; then sudo tar -xvzf %q -C /usr/local/bin/%s;fi`, filePath, goCliName),
		fmt.Sprintf(`if [ "$nbFolder" -eq 1 ]; then sudo tar -xvzf %q -C /usr/local/bin/%s --strip-components=1; fi`, filePath, goCliName),
	}

	cli := strings.Join(cmds, " && ")
	return cli
}
