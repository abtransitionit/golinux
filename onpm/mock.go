package onpm

func UpgradePkg() (string, error) {
	return "apt-get update -y", nil
}

func UpgradeOs() string {
	return "apt-get upgrade -y"
}
