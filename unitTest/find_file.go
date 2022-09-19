package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 查找目录及子目录下的文件
// 添加参数判断
// path : 目标目录
// files : 容纳所有文件路径的结果数组
// targetType : 目标文件类型
// ignoreFile : 忽略文件（文件名，包括扩展名）
// ignorePath : 忽略目录
// ignoreType : 忽略文件类型
func GetAllFile(path string, files *[]string, targetType *[]string, ignoreFile *[]string, ignorePath *[]string, ignoreType *[]string) (err error) {

	if !isAllEmpty(targetType) && !isAllEmpty(ignoreType) {

		fmt.Printf("WARNGING: 目标文件类型已指定, 忽略文件类型无须指定。后续处理中忽略文件类型作为空处理\n")
	}

	err = getAllFileRecursion(path, files, targetType, ignoreFile, ignorePath, ignoreType)
	return err
}

// 递归查找目录及子目录下的文件
// path : 目标目录
// files : 容纳所有文件路径的结果数组
// targetType : 目标文件类型
// ignoreFile : 忽略文件（文件名，包括扩展名）
// ignorePath : 忽略目录
// ignoreType : 忽略文件类型
func getAllFileRecursion(path string, files *[]string, targetType *[]string, ignoreFile *[]string, ignorePath *[]string, ignoreType *[]string) (err error) {
	l, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	separator := string(os.PathSeparator)
	for _, f := range l {
		tmp := string(path + separator + f.Name())

		if f.IsDir() {

			// 过滤被忽略的文件夹（文件夹名完全相同）
			if !isInArray(ignorePath, f.Name()) {

				err = getAllFileRecursion(tmp, files, targetType, ignoreFile, ignorePath, ignoreType)
				if err != nil {
					return err
				}
			}
		} else {
			// 目标文件类型被指定
			if !isAllEmpty(targetType) {

				// 属于目标文件类型
				if isInSuffix(targetType, f.Name()) {

					// 忽略文件为空 或者 目标文件中不含有指定忽略文件
					if isAllEmpty(ignoreFile) || !isInArray(ignoreFile, f.Name()) {

						*files = append(*files, tmp)
					}
				}
			} else { // 目标文件类型为空

				// 忽略文件类型被指定
				if !isAllEmpty(ignoreType) {

					// 不属于忽略文件类型
					if !isInSuffix(ignoreType, f.Name()) {

						// 忽略文件为空 或者 目标文件中不含有指定忽略文件
						if isAllEmpty(ignoreFile) || !isInArray(ignoreFile, f.Name()) {

							*files = append(*files, tmp)
						}
					}
				} else { // 忽略文件类型为空

					// 忽略文件为空 或者 目标文件中不含有指定忽略文件
					if isAllEmpty(ignoreFile) || !isInArray(ignoreFile, f.Name()) {

						*files = append(*files, tmp)
					}
				}
			}
		}
	}

	return nil
}

// 判断目标字符串是否是在数组中
func isInArray(list *[]string, s string) (isIn bool) {

	if len(*list) == 0 {
		return false
	}

	isIn = false
	for _, f := range *list {

		if f == s {
			isIn = true
			break
		}
	}

	return isIn
}

// 判断目标字符串的末尾是否含有数组中指定的字符串
func isInSuffix(list *[]string, s string) (isIn bool) {

	isIn = false
	for _, f := range *list {

		if strings.TrimSpace(f) != "" && strings.HasSuffix(s, f) {
			isIn = true
			break
		}
	}

	return isIn
}

// 判断数组各元素是否是空字符串或空格
func isAllEmpty(list *[]string) (isEmpty bool) {

	if len(*list) == 0 {
		return true
	}

	isEmpty = true
	for _, f := range *list {

		if strings.TrimSpace(f) != "" {
			isEmpty = false
			break
		}
	}

	return isEmpty
}

func FindFilePath() {
	fmt.Printf("Info: Start\n")

	targetType := []string{".js", "", ""}
	ignoreFile := []string{"index.js"}
	ignorePath := []string{".git"}
	ignoreType := []string{".gitignore", ".exe"}

	var files []string
	var path = "C:\\Users\\Administrator\\go\\src\\github.com\\bettersun\\hellogo"
	err := GetAllFile(path, &files, &targetType, &ignoreFile, &ignorePath, &ignoreType)
	if err != nil {
		fmt.Printf(err.Error() + "\n")
	}

	fmt.Printf("文件名（全路径）列表:\n")
	for _, file := range files {
		fmt.Printf("  [%s]\n", file)
	}

	fmt.Printf("Info: End\n")
}
