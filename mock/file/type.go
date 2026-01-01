package file

import (
	"path/filepath"
)

// defrine types
type FileProperty struct {
	Src   string // filePath
	Dst   string // filePath or folderPath
	Chmod string // optional
}
type File struct {
	Name     string // filename
	Path     string // filePath
	FullPath string // full path = Path + Name
}

// description: creates a File either from (name + path) or from fullFilePath
func GetFile(fileName, filePath, fullFilePath string) *File {

	// Case 1: name + path provided
	if fileName != "" && filePath != "" {
		full := filepath.Join(filePath, fileName)

		return &File{
			Name:     fileName,
			Path:     filePath,
			FullPath: full,
		}
	}

	// Case 2: full path provided
	if fullFilePath != "" {
		return &File{
			Name:     filepath.Base(fullFilePath),
			Path:     filepath.Dir(fullFilePath),
			FullPath: fullFilePath,
		}
	}

	return nil
}

// // decription: returns a File pointing to "$HOME/.profile"
// func GetRcFile() *File {
// 	home := os.Getenv("HOME")                   // get user's home directory
// 	fullPath := filepath.Join(home, ".profile") // $HOME/.profile

// 	return &File{
// 		Name:     ".profile",
// 		Path:     home,
// 		FullPath: fullPath,
// 	}
// }
