package onpm

func (d *DnfRepoManager) List() (string, error) {
	return "dnf list repos", nil
}

func (d *DnfRepoManager) Add() (string, error) {
	return "dnf config-manager --add-repo <repo>", nil
}

func (d *DnfRepoManager) Remove() (string, error) {
	return "dnf config-manager --remove-repo <repo>", nil
}
