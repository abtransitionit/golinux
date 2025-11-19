package onpm

func (d *DnfPkgManager) List() (string, error) {
	return "dnf list installed", nil
}

func (d *DnfPkgManager) Install() (string, error) {
	return "dnf install <pkg>", nil
}

func (d *DnfPkgManager) Remove() (string, error) {
	return "dnf remove <pkg>", nil
}
