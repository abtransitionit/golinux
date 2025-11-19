package onpm

// --- Package Managers ---

func (a *AptPkgManager) List() (string, error) {
	return "apt list --installed", nil
}

func (a *AptPkgManager) Install() (string, error) {
	return "apt install <pkg>", nil
}

func (a *AptPkgManager) Remove() (string, error) {
	return "apt remove <pkg>", nil
}
