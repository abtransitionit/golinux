# Terminlogy
|Name|type|desc|
|-|-|-|
|PM|acro|**P**ackage **M**anager|
|ONPM|acro|**O**s **N**ative **P**ackage **M**anager|

# Purpose
`onpm` provides a generic abstraction over package managers from two major Linux families:

* **Debian-based** (APT)
* **RHEL-based** (DNF)

It lets you build CLI commands for adding, removing, and listing packages or repositories without caring about which distro you're running on.



# Concepts

## PackageManager (CliBuilder)

An interface that generate **CLI strings**, not execute them, implemented by:

* `AptManager`
* `DnfManager`

Supported `method()` implemented by each **PM**:

| Method        | For Repo               | For Package             |
| ------------- | ---------------------- | ----------------------- |
| `CliList()`   | Lists configured repos | Lists installed package |
| `CliAdd()`    | Adds repo              | Installs package        |
| `CliDelete()` | Removes repo           | Removes package         |


Examples:


| Method               | Debian/Ubuntu                               | RHEL/Fedora/Rocky                    |
| -------------------- | ------------------------------------------- | ------------------------------------ |
| `CliAdd()`           | `add-apt-repository 'URL'`                  | `dnf config-manager --add-repo URL`  |
| `CliAdd()` (pkg)     | `apt-get install -y pkg`                    | `dnf install -y pkg`                 |
| `CliDelete()`        | `apt-get remove -y pkg`                     | `dnf remove -y pkg`                  |
| `CliDelete()` (repo) | *remove file in* `/etc/apt/sources.list.d/` | *remove file in* `/etc/yum.repos.d/` |



## Factory: `CliBuilderFactory.get()`

Creates the appropriate manager based on:

* `osFamily` → `"debian"` or `"rhel"` (normalized)
* `osDistro` → `"ubuntu"`, `"debian"`, `"rocky"`, `"fedora"`, etc.
* Either a `*Repo` or a `*Package` (the other must be `nil`)





# Usage 

## Build a CLI for a package

```go
pkg := &onpm.Package{Name: "nginx"}
err := pkg.SetCliBuilder("debian", "ubuntu", nil)
if err != nil {
    panic(err)
}

cmd, _ := pkg.Cbd.CliAdd()
fmt.Println(cmd) // apt-get install -y nginx
```

## Build a CLI for a repository

```go
repo := &onpm.Repo{
    Name: "myrepo",
    Url:  "http://example.com/repo",
}

err := repo.GetCliBuilder("rhel", "rocky")
if err != nil {
    panic(err)
}

cmd, _ := repo.Cbd.CliAdd()
fmt.Println(cmd) // dnf config-manager --add-repo http://example.com/repo
```


## Repo
### formal definition
```go
type Repo struct {
	Name string
	Url  Url
	Cbd  CliBuilder
}
```
### definition
a `dnf` or `apt` package repository

## Url
### formal definition
```go
type Url struct {
	Repo string
	Gpg  string
}
```
### definition
a container for the URLs of a `dnf` or `apt` package repository. eg.

```json
"crio": {
    Repo: "https://download.opensuse.org/repositories/isv:/cri-o:/stable:/v$TAG/$PACK/",
    Gpg:  "https://download.opensuse.org/repositories/isv:/cri-o:/stable:/v$TAG/$PACK/$GPG",
},
"k8s": {
    Repo: "https://pkgs.k8s.io/core:/stable:/v$TAG/$PACK/",
    Gpg:  "https://pkgs.k8s.io/core:/stable:/v$TAG/$PACK/$GPG",
}

```
## Package
### formal definition
```go
type Package struct {
	Name    string
	Version string
	Cbd     CliBuilder
}
```
### definition
a `dnf` or `apt` package of a repository

# Directory layout
```
onpm/
  obj.pkg.go       # methods to manage package - eg. pkg.add(...)
  obj.repo.go      # methods to manage repo    - eg. repo.delete(...)
  pm.go           # return the manager to be used according to the OS type 
  pm.apt.go       # return the apt CLIs to manage repo and package for ubuntu-based linux OS
  pm.dnf.go       # return the dnf CLIs to manage repo and package for rhel/fedora-based linux OS
  type.go          # define structure and interface
  README.md.       # this doc
```

# How it works
1. OS repository and package are **modeled** by a structure (ie. `Repo` and `Package`)
1. Each structure to model the object define members among which
    - a method `GetManager` that detect the os type (rhel, fedora, ubuntu, ...)
    - has access a method (`GetManager`) that detect the os type (rhel, fedora, ubuntu, ...)
1. according to the os type, `GetPM` create another object (the manager) inside the repo or the package
   - `AptManager` if it's a `ubuntu` like OS
   - `DnfManager` if it's a `rhel` or `fedora` like OS
1. Now the oject (repo or package) knows on which OS it resides
1. the rest is only a matter of wrapping `dnf` or `apt` CLI to manage repositories and packages
1. for example calling `repo.add(...)` does the following
   - call the manager
   - the manager (Apt... or Dnf...) get the CLI to add a repo onto the OS
   - the manager execute the CLI localy or remotely acccording to the parameters passed to the cmde.






1. **`da/pkg.go`**

   * Defines `RepositoryManager` interface
   * Implements `APTRepositoryManager` and `DNFRepositoryManager`
   * Handles adding, removing, listing repos

2. **`pkg/pkg`**

   * Defines `PackageManager` interface
   * Implements `APTPackageManager` and `DNFPackageManager`
   * Handles installing, updating, removing, listing packages
   * Can include logic for repo-dependent packages later

3. **`internal/osdetect`**

   * Detects OS type at runtime
   * Returns constants like `OS_APT` or `OS_DNF`
   * Ensures library chooses the right implementation automatically

4. **`config/packages.yaml`**

   * Optional mapping table for standard packages between distros
   * Example:

     ```yaml
     vim:
       apt: vim
       dnf: vim-enhanced
     gnupg:
       apt: gnupg
       dnf: gnupg2
     ```

---

## **3. Incremental Development Path**

1. Implement **OS detection**
2. Implement **APTRepositoryManager**
3. Implement **APTPackageManager** for standard packages
4. Implement **DNFRepositoryManager**
5. Implement **DNFPackageManager** for standard packages
6. Add **package mapping logic** for cross-distro
7. Implement **repo-aware package installation** if needed
8. Add **logging, config, and tests** at each step

---

This layout allows you to **grow the library incrementally** without touching existing code.

If you want, the next step could be a **detailed interface definition for RepositoryManager and PackageManager**, including OS detection and optional mapping support. This will be the blueprint before writing any code.

Do you want me to do that next?


# Example usage
```go
func main() {
  // Load manager.yaml
  config, err := LoadConfig("manager.yaml")
  if err != nil {
      log.Fatal(err)
  }

  r := Repo{
      Name:   "kubernetes",
      URL:    "https://apt.kubernetes.io/",
      Enabled: true,
      config: &config.APT, // set APT paths
  }

  msg, err := repo.Add()
  if err != nil {
      log.Fatal(err)
  }
  fmt.Println(msg)

  msg, err = repo.Remove()
  if err != nil {
      log.Fatal(err)
  }
  fmt.Println(msg)

      // r.Remove() can be called later
  }

```






requirements for a software projects written in go
- manage dnf and apt package repositories and packages. 
- the code must 
  - be production grade
  - written incrementally

and 
- never write code before i ask
- go slowly




idea:

type ManagerConfig struct {
    RepoFolder string `yaml:"repoFolder"`
    KeyFolder  string `yaml:"keyFolder,omitempty"`
    GPGCheck   bool   `yaml:"gpgCheck,omitempty"`
}

type Config struct {
    APT ManagerConfig `yaml:"apt"`
    DNF ManagerConfig `yaml:"dnf"`
}

// LoadConfig loads manager.yaml
func LoadConfig(path string) (*Config, error) {
    // 1. Read YAML file
    // 2. Unmarshal into Config struct
    return nil, nil
}



type Repo struct {
    Name    string
    URL     string
    GPGKey  string
    Enabled bool
    config  *ManagerConfig // pointer to config for paths
}

// Add creates the repository in the system
func (r *Repo) Add() error {
    if r.config == nil {
        return fmt.Errorf("config not set for repo %s", r.Name)
    }

    listFile := filepath.Join(r.config.RepoFolder, r.Name+".list")

    // Example: write basic repo URL (APT only for now)
    content := fmt.Sprintf("deb %s stable main\n", r.URL)
    if err := os.WriteFile(listFile, []byte(content), 0644); err != nil {
        return fmt.Errorf("failed to write repo file: %v", err)
    }

    // GPG key handling can be added later
    return nil
}

// Remove deletes the repository
func (r *Repo) Remove() error {
    if r.config == nil {
        return fmt.Errorf("config not set for repo %s", r.Name)
    }

    listFile := filepath.Join(r.config.RepoFolder, r.Name+".list")
    if err := os.Remove(listFile); err != nil && !os.IsNotExist(err) {
        return fmt.Errorf("failed to remove repo file: %v", err)
    }

    return nil
}

func main() {
    // Load manager.yaml
    config, err := LoadConfig("manager.yaml")
    if err != nil {
        log.Fatal(err)
    }

    r := Repo{
        Name:   "kubernetes",
        URL:    "https://apt.kubernetes.io/",
        Enabled: true,
        config: &config.APT, // set APT paths
    }

    if err := r.Add(); err != nil {
        log.Fatal(err)
    }

    // r.Remove() can be called later
}
