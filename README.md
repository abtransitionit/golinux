# Golinux

A foundational Go library (i.e. no `main()`) containing low level reusable universal linux functions, for managing any Linux platform. 

It provides a set of primitives that handle cross-distribution differences automatically. no need to know/worry if the linux host is `Rhel`, `debian`, `alpine`, `ubuntu`, ... .

----

[![Dev CI](https://github.com/abtransitionit/golinux/actions/workflows/ci-dev.yaml/badge.svg?branch=dev)](https://github.com/abtransitionit/golinux/actions/workflows/ci-dev.yaml)
[![Main CI](https://github.com/abtransitionit/golinux/actions/workflows/ci-main.yaml/badge.svg?branch=main)](https://github.com/abtransitionit/golinux/actions/workflows/ci-main.yaml)
[![LICENSE](https://img.shields.io/badge/license-Apache_2.0-blue.svg)](https://choosealicense.com/licenses/apache-2.0/)

---


# Features  
This project template includes the following components:  


|Component|Description|
|-|-|
|Licensing|Predefined open-source license (Apache 2.0) for legal compliance.|
|Code of Conduct| Ensures a welcoming and inclusive environment for all contributors.|  
|README|Structured documentation template for clear project onboarding.|  

This library offers a high-level API for common Linux administration tasks, uning the same universal primitive** including:

- `dnfapt`: A package for managing software packages on RHEL and Debian systems.
- `oservice`: A package for starting, stopping, and checking the status of OS services.
- `runCliLocal()`: A primitive to play any complex linux CLI on the local VM
- `runCliSsh()`: A primitive to play any complex linux CLI on a remote VM

---

## Installation

To use this library in your project, run:

```bash
go get [github.com/abtransitionit/golinux](https://github.com/abtransitionit/golinux)
```

---

# Contributing  

We welcome contributions! Before participating, please review:  
- **[Code of Conduct](.github/CODE_OF_CONDUCT.md)** – Our community guidelines.  
- **[Contributing Guide](.github/CONTRIBUTING.md)** – How to submit issues, PRs, and more.  

----


# Release History & Changelog  

Track version updates and changes:  
- **📦 Latest Release**: `vX.X.X` ([GitHub Releases](#))  
- **📄 Full Changelog**: See [CHANGELOG.md](CHANGELOG.md) for detailed version history.  

---


# Howtos
## get code from repo
**purpose**
How to use code from another `go` project that is also in active development.
```shell
# in golinux/go.mod (need code form gocore)
require github.com/abtransitionit/gocore v0.0.0
replace github.com/abtransitionit/gocore => ../gocore
```
- then `go mod tidy`


