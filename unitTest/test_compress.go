package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
)

func CompressedFile(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil || info.IsDir() {
		return err
	}
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	header.Name = prefix + "/" + header.Name
	writer, err := zw.CreateHeader(header)
	if err != nil {
		return err
	}
	if _, err = io.Copy(writer, file); err != nil {
		return err
	}
	return nil
}

func TestCompress() {
	f, _ := os.Open("Unit_Test.txt")
	// 压缩文件
	dst, _ := os.Create("Unit_Test.zip")
	zipWriter := zip.NewWriter(dst)
	if err := CompressedFile(f, "", zipWriter); err != nil {
		log.Fatalln(err.Error())
	}
	// Make sure to check the error on Close.
	if err := zipWriter.Close(); err != nil {
		log.Fatalln(err.Error())
	}
	return
}
