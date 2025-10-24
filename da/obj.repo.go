package da

import (
	"context"
	"fmt"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/run"
)

func NewRepo(name string, url RepoUrl, osType string) (*Repo, error) {
	r := &Repo{Name: name, Url: url}

	mgr, err := r.GetManager(osType)
	if err != nil {
		return nil, err
	}
	r.Mgr = mgr

	return r, nil
}

func (r *Repo) Add(ctx context.Context, local bool, remoteHost string, logger logx.Logger) (string, error) {
	return r.Mgr.CliAdd()
}
func (r *Repo) Delete(ctx context.Context, local bool, remoteHost string, logger logx.Logger) (string, error) {
	return r.Mgr.CliDelete()
}

//	func (r *Repo) List(ctx context.Context, local bool, remoteHost string, logger logx.Logger) (string, error) {
//		return r.Mgr.CliList()
//	}
func (r *Repo) EnableGPG(ctx context.Context) (string, error) {
	return fmt.Sprintf("gpg --import %s.gpg", r.Name), nil
}
func (r *Repo) Update(ctx context.Context) (string, error) {
	return fmt.Sprintf("sudo dnf update -q -y > /dev/null"), nil
}

func (r *Repo) List(ctx context.Context, local bool, remoteHost string, logger logx.Logger) (string, error) {
	// define cli
	cli, err := r.Mgr.CliList()
	if err != nil {
		return "", fmt.Errorf("failed to build CLI list: %w", err)
	}

	// // play cli
	output, err := run.ExecuteCliQuery(cli, logger, local, remoteHost, run.NoOpErrorHandler)
	if err != nil {
		return "", fmt.Errorf("failed to run command: %s: %w", cli, err)
	}

	// return response
	return output, nil

}

// // list repo
// func (r *Repo) List() (string, error) {
// 	if r.config == nil {
// 		return "", fmt.Errorf("config not set for repo listing")
// 	}

// 	files, err := os.ReadDir(r.config.RepoFolder)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to read repo folder: %v", err)
// 	}

// 	if len(files) == 0 {
// 		return "No repositories found.", nil
// 	}

// 	var sb strings.Builder
// 	sb.WriteString("Repositories:\n")

// 	for _, file := range files {
// 		if !file.IsDir() && strings.HasSuffix(file.Name(), ".list") {
// 			sb.WriteString(fmt.Sprintf(" - %s\n", file.Name()))
// 		}
// 	}

// 	return sb.String(), nil
// }

// // adds a repo
// func (r *Repo) Add() (string, error) {
// 	if r.config == nil {
// 		return "", fmt.Errorf("config not set for repo %s", r.Name)
// 	}

// 	listFile := filepath.Join(r.config.RepoFolder, r.Name+".list")
// 	content := fmt.Sprintf("deb %s stable main\n", r.URL)

// 	if err := os.WriteFile(listFile, []byte(content), 0644); err != nil {
// 		return "", fmt.Errorf("failed to write repo file: %v", err)
// 	}

// 	msg := fmt.Sprintf("Repository '%s' added at %s", r.Name, listFile)
// 	return msg, nil
// }

// // deletes a repo
// func (r *Repo) Delete() (string, error) {
// 	if r.config == nil {
// 		return "", fmt.Errorf("config not set for repo %s", r.Name)
// 	}

// 	listFile := filepath.Join(r.config.RepoFolder, r.Name+".list")
// 	if err := os.Remove(listFile); err != nil && !os.IsNotExist(err) {
// 		return "", fmt.Errorf("failed to remove repo file: %v", err)
// 	}

// 	msg := fmt.Sprintf("Repository '%s' removed from %s", r.Name, listFile)
// 	return msg, nil
// }
