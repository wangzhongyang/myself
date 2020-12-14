package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"
)

var (
	pathRoot string
)

func main() {
	path := "/Users/wangzhongyang/go/src/myself/forza4-backup"
	_ = Compress(path)
	delete()
}

// Compress 压缩文件，返回压缩文件地址
func Compress(path string) string {
	pathArr := strings.Split(path, "/")
	pathRoot = strings.Join(pathArr[:len(pathArr)-1], "/")
	dirName := pathArr[len(pathArr)-1]
	compressName := fmt.Sprintf("%s-%d.zip", dirName, time.Now().Unix())

	outFile, err := os.Create(fmt.Sprintf("%s/%s", pathRoot, compressName))
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	// Create a new zip archive.
	w := zip.NewWriter(outFile)
	defer w.Close()

	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Errorf("open dir failed, path:%s, error:%w", path, err))
	}
	defer file.Close()

	if err := compress(file, pathRoot, w); err != nil {
		panic(fmt.Errorf("compress failed,error:%w", err))
	}

	return fmt.Sprintf("%s/%s", pathRoot, compressName)
}

func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	defer file.Close()
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			fileName := file.Name() + "/" + fi.Name()
			f, err := os.Open(fileName)
			if err != nil {
				panic(fmt.Errorf("open file failed, file name:%s, %w", fileName, err))
			}
			defer f.Close()
			if err = compress(f, prefix, zw); err != nil {
				panic(fmt.Errorf("compress failed,%w", err))
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			panic(fmt.Errorf("zip.FileInfoHeader failed,%w", err))
		}
		header.Name = opposite(prefix + "/" + header.Name)
		writer, err := zw.CreateHeader(header)
		if err != nil {
			panic(fmt.Errorf("zw.CreateHeader failed,%w", err))
		}
		if _, err = io.Copy(writer, file); err != nil {
			panic(fmt.Errorf("io.Copy failed,%w", err))
		}
	}
	return nil
}

// opposite 计算相对目录
func opposite(longPath string) string {
	a1 := strings.Split(pathRoot, "/")
	a2 := strings.Split(longPath, "/")
	a3 := a2[len(a1):]
	return strings.Join(a3, "/")
}

func delete() {
	funcName := "delete file,"
	root, err := os.Open(pathRoot)
	if err != nil {
		panic(fmt.Errorf("%s Open file failed, error:%w", funcName, err))
	}
	rootStat, err := root.Stat()
	if err != nil {
		panic(fmt.Errorf("%s get file stat failed, error:%w", funcName, err))
	}
	if !rootStat.IsDir() {
		panic(fmt.Errorf("%s this path is not dir", funcName))
	}
	filesInfo, err := root.Readdir(-1)
	if err != nil {
		panic(fmt.Errorf("%s get child node failed, error:%w", funcName, err))
	}
	sort.Sort(filesInfoType(filesInfo))
	count := 0
	for _, file := range filesInfo {
		if strings.Contains(file.Name(), "forza4-backup-") {
			fmt.Println("file name: ", file.Name())
			count += 1
		}
		if count > 5 {
			removeName := pathRoot + "/" + file.Name()
			if err := os.Remove(removeName); err != nil {
				panic(fmt.Errorf("%s remove file failed, file name:%s, error:%w", funcName, removeName, err))
			}
		}
	}
}

type filesInfoType []os.FileInfo

func (s filesInfoType) Len() int           { return len(s) }
func (s filesInfoType) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s filesInfoType) Less(i, j int) bool { return s[i].Name() < s[j].Name() }
