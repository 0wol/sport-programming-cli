package file

import (
	"bufio"
	"os"
	"strings"
)

type File struct {
	Name         string
	IsDirectory  bool
	AbsolutePath string
	files        []*File
}

func NewFile(name string, absolutePath string) (*File, error) {
	if !isFile(absolutePath) {
		files, err := getFiles(absolutePath)

		if err != nil {
			return nil, err
		}

		return &File{
			Name:         name,
			IsDirectory:  true,
			AbsolutePath: absolutePath,
			files:        files,
		}, nil
	}

	return &File{
		Name:         name,
		IsDirectory:  false,
		AbsolutePath: absolutePath,
		files:        []*File{},
	}, nil
}

func (f *File) GetTree() string {
	return f.getTree(0)
}

func (f *File) getTree(h int) string {
	res := ""
	for i := 0; i < h; i++ {
		res += " "
	}
	res += f.Name + "\n"

	if f.IsDirectory {
		for _, file := range f.files {
			res += file.getTree(h + 4)
		}
	} else {
		names, err := f.getFunctionsNames()
		if err == nil {
			for _, name := range names {
				for i := 0; i < h+4; i++ {
					res += " "
				}
				res += name + "\n"
			}
		}
	}

	return res
}

func (f *File) FindFunctionsByFileName(fileName string) (string, error) {
	if strings.Contains(strings.ToLower(f.Name), fileName) {
		bodies, err := f.getAllFunctions()
		if err != nil {
			return "", err
		}
		return bodies, nil
	}

	res := ""

	for _, inF := range f.files {
		bodies, err := inF.FindFunctionsByFileName(fileName)

		if err != nil {
			return "", err
		}

		res += bodies
	}

	return res, nil
}

func (f *File) FindFunctionsByName(name string) (string, error) {
	bodies, err := f.getFunctionsBodiesByNames([]string{name})
	if err != nil {
		return "", err
	}
	return bodies, nil
}

func (f *File) getAllFunctions() (string, error) {
	names, err := f.getFunctionsNames()

	if err != nil {
		return "", err
	}

	res, err := f.getFunctionsBodiesByNames(names)

	if err != nil {
		return "", err
	}

	return res, nil
}

func (f *File) getFunctionsNames() ([]string, error) {
	var res []string

	if f.IsDirectory {
		for _, inF := range f.files {
			names, err := inF.getFunctionsNames()
			if err != nil {
				return nil, err
			}
			res = append(res, names...)
		}
		return res, nil
	}

	file, err := os.Open(f.AbsolutePath)

	if err != nil {
		return nil, err
	}

	var (
		scanner = bufio.NewScanner(file)
		level   = 0
	)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "{") {
			if level == 0 {
				res = append(res, getSignatureName(line))
			}
			level++
		}

		if strings.Contains(line, "}") {
			level--
		}
	}

	return res, nil
}

func (f *File) getFunctionsBodiesByNames(functionsNames []string) (string, error) {
	if f.IsDirectory {
		res := ""
		for _, inF := range f.files {
			bodies, err := inF.getFunctionsBodiesByNames(functionsNames)
			if err != nil {
				return "", err
			}
			res += bodies
		}
		return res, nil
	}

	file, err := os.Open(f.AbsolutePath)

	if err != nil {
		return "", err
	}

	var (
		scanner = bufio.NewScanner(file)
		res     = ""
		level   = 0
		prev    = ""
	)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "{") {
			if level == 0 {
				if SliceContainsSimilar(functionsNames, strings.ToLower(getSignatureName(line))) {
					if prev != "" {
						res += prev + "\n"
					}
					res += line + "\n"
					level++
					for level > 0 && scanner.Scan() {
						line := scanner.Text()

						if strings.Contains(line, "{") {
							level++
						} else if strings.Contains(line, "}") {
							level--
						}
						res += line + "\n"
					}
					res += "\n"
				} else {
					level++
				}
			} else {
				level++
			}
		}

		if strings.Contains(line, "}") {
			level--
		}

		prev = line
	}

	return res, nil
}
