package utils

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
)

//GetFolders .
func GetFolders(folder string) []string {
	var result []string
	files, _ := ioutil.ReadDir(folder) //specify the current dir
	for _, file := range files {
		if file.IsDir() {
			result = append(result, file.Name())
			//GetFiles(folder + "/" + file.Name())
		} else {
			//fmt.Println(folder + "/" + file.Name())
			continue
			//result = append(result, file.Name())
		}
	}
	return result
}

// IsExist -
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

//CopyFile .
func CopyFile(dstName string, srcName string) (written int64, err error) {
	// Open the file
	src, err := os.Open(srcName)
	if err != nil {
		return
	}

	// Close the open file
	defer src.Close()

	// Create a new file
	dst, err := os.Create(dstName)
	if err != nil {
		return
	}

	// Close the newly created file
	defer dst.Close()

	// Copy file
	return io.Copy(dst, src)
}

//Compress .
func Compress(files []*os.File, dest string) error {
	d, _ := os.Create(dest)
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()
	for _, file := range files {
		err := Compresse(file, "", w)
		if err != nil {
			return err
		}
	}
	return nil
}

//Compresse .
func Compresse(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		if len(prefix) == 0 {
			prefix = info.Name()
		} else {
			prefix = prefix + "/" + info.Name()
		}
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = Compresse(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		if len(prefix) == 0 {
			header.Name = header.Name
		} else {
			header.Name = prefix + "/" + header.Name
		}
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
