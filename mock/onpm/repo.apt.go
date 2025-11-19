package onpm

func (a *AptRepoManager) List() (string, error) {
	return "apt list repos", nil
}

func (a *AptRepoManager) Add() (string, error) {
	return "add-apt-repo <repo>", nil
}

func (a *AptRepoManager) Remove() (string, error) {
	return "remove-apt-repo <repo>", nil
}
