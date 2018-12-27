package main

import (
	"io"
	"os"
)

//复制文件
func copyFile(DistName string, SrcName string) (written int64, err error) {
	src, err := os.Open(SrcName)
	if err != nil {
		return
	}
	defer src.Close()
	Dist, err := os.OpenFile(DistName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer Dist.Close()
	return io.Copy(Dist, src)
}

//文件内容展示到终端
func copyout(SrcName string) (written int64, err error) {
	src, err := os.Open(SrcName)
	if err != nil {
		return
	}
	defer src.Close()

	return io.Copy(os.Stdout, src)
}

func main() {
	//copyFile("dist.txt", "src.txt")
	copyout("src.txt")
}
