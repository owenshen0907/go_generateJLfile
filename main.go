// generationJLfile project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	//	"path"
	"path/filepath"
	"strconv"
	"strings"
	//"github.com/qiniu/iconv"
)

var fileName = [10]string{""} //文件名
var tmpContent []string       //读取分离开的文件内容

var body, currentDirectory string //文件内容

func main() {
	getFilename()
	//generateJL()
}

//使用时将，sq文件放入send文件夹，执行完，hf文件将会自动生成。
//C:\GoPath\src\generateJLfile\send
func getFilename() {
	//读取文件名
	fmt.Println("正在读取文件。。。")
	currentDirectory = getCurrentDirectory() + "\\send"
	fmt.Println(currentDirectory)
	//dir, _ := os.OpenFile("C:\\GoPath\\src\\generateJLfile\\send", os.O_RDONLY, os.ModeDir)
	dir, _ := os.OpenFile(currentDirectory, os.O_RDONLY, os.ModeDir)

	fileinfo, _ := dir.Stat()
	fmt.Println(fileinfo.IsDir())

	fileNameStress, _ := dir.Readdir(-1)
	for i, fileNames := range fileNameStress {
		//fmt.Println(i)
		//fmt.Println(fileNames.Name())
		fileName[i] = fileNames.Name()

	}
	if fileName[0] == "" {

		fmt.Println("文件夹未找到相关内容，请核实申请文件已经放入指定文件夹")
		fmt.Println("程序已退出。请检查，文件重新执行")
	} else {
		fmt.Println("-------------------文件名已获取------------------")
		generateJL()
	}
	//	for j, _ := range fileName {
	//		fmt.Println(fileName[j])
	//		if fileName[j] == "0" {
	//			break
	//		}
	//		defer dir.Close()
	//		//判断是否为目录
	//		//	for _, name := range names {
	//		//		fmt.Println(name.Name(), "目录?", name.IsDir())
	//		//	}
	//	}
}
func generateJL() {
	fmt.Println("----------------开始获取相关数据-----------------")
	s := "^?"
	tmp1 := "40216061678"
	tmp2 := "FFT160729021399"

	var bottom, count string
	for _, names := range fileName {
		//fmt.Println(fileName[i])
		if names == "" {
			break
		}
		fmt.Println("---开始获取" + names + "的数据---")
		sourceFile := "send/" + names
		//fmt.Println(sourceFile)

		file, _ := os.Open(sourceFile)
		hfFilePath := strings.Replace(names, "sq", "hf", -1)
		hfFilePath = "received/" + hfFilePath
		//fmt.Println(hfFilePath)

		fileBody, _ := ioutil.ReadAll(file)
		tmpContent = strings.Split(string(fileBody), "^?")
		fmt.Println("-------------------开始写入数据-----------------")
		for i := 0; i < (len(tmpContent)-1)/139; i++ {
			body = tmpContent[139*i+2] + s + tmpContent[139*i+3] + s + tmpContent[139*i+5] + s + tmpContent[139*i+6] + s + "1" + s
			body = body + tmp1 + s + tmpContent[139*i+51] + s + tmp2 + s + s + "\r\n"
			//			if i == len(tmpContent)-1)/139 {
			//				bottom := "--------------------" + "\r\n"
			//				fmt.Println(i)
			//				body = body + bottom + "i"
			//			}
			count = strconv.Itoa(i + 1)
			bottom = "--------------------" + "\r\n" + count
			//			cd, err := iconv.Open("gbk", "utf-8")
			//			erro(err)
			//			defer cd.Close()
			//			gbk := cd.ConvString(body)
			writeToFile(hfFilePath, body)
		}
		writeToFile(hfFilePath, bottom)
		fmt.Println(hfFilePath + "文件已生成。")
	}
	fmt.Println("执行完毕，请查看文件，继续下一步。")
}

func writeToFile(fileName, body string) {
	dstfile, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE, 0777)
	erro(err)
	defer dstfile.Close()
	dstfile.WriteString(body)
}
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	erro(err)
	//return strings.Replace(dir, "\\", "/", -1)
	return dir
}
func erro(err error) {
	if err != nil {
		fmt.Println("出错了", err)
	}
}

//1姓名，2合同号，5
//	fullFilename := "C:/Users/shent/Desktop/与mac共享/操作手册-Owen/6联贷反馈成功"
//	fmt.Println("fullFilename =", fullFilename)
//	var filenameWithSuffix string
//	filenameWithSuffix = path.Base(fullFilename) //获取文件名带后缀
//	fmt.Println("filenameWithSuffix =", filenameWithSuffix)
//	var fileSuffix string
//	fileSuffix = path.Ext(filenameWithSuffix) //获取文件后缀
//	fmt.Println("fileSuffix =", fileSuffix)

//	var filenameOnly string
//	filenameOnly = strings.TrimSuffix(filenameWithSuffix, fileSuffix) //获取文件名
//	fmt.Println("filenameOnly =", filenameOnly)
