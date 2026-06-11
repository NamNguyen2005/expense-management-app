package utils

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	// "github.com/google/uuid"
)
var maxSizeFile int = 5 << 20
var alowExtension = map[string]bool{
    ".img" : true,  
    ".jpg" : true, 
    ".png" : true,
	".jpeg" : true,
}
var alowMineType = map[string]bool{
	"image/jpeg" : true,
	"image/png" : true,	
}
func ValidateAndSaveFile(fileHeader *multipart.FileHeader , uploadDr string )(string , error){
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if !alowExtension[ext] {
		return "" , errors.New("unsupported file extension")
	}
	if fileHeader.Size > int64(maxSizeFile) {
		return "" , errors.New("file is too large(100 MB)")
	}
	file , err := fileHeader.Open()
	if err != nil{
		return "" , errors.New("khong the mo file")
	}
	defer file.Close()
	buffer := make([]byte , 512)
	_ , err = file.Read(buffer)
	if err != nil {
		return "" , errors.New("can not read file")
	}
	// create file
	err = os.MkdirAll("./upload", os.ModePerm )
	if err != nil{
		return "" , errors.New("can not create file upload")
	}
	alowMine := http.DetectContentType(buffer)
	if !alowMineType[alowMine] {
		return "" , errors.New("file khong hop le")
	}

	//change file name not repeat
	baseName := fileHeader.Filename[:len(fileHeader.Filename)-len(ext)]
	fileName := fileHeader.Filename
	savePath := filepath.Join(uploadDr, fileName)
	
	if _, err := os.Stat(savePath); err == nil {
		fileName = fmt.Sprintf("%s_%d_copy%s", baseName, time.Now().Unix(), ext)
		savePath = filepath.Join(uploadDr, fileName)
	}

	err = SaveFile(fileHeader, savePath)
	if err != nil{
		return "" , err
	}
	
	return fileName ,nil
}


// save file function with no ctx
func SaveFile(fileHeader *multipart.FileHeader , destination string)(error){
	src , err := fileHeader.Open()
	if err != nil{
		return  err
	}
	defer src.Close()

	out , err := os.Create(destination)
	if err != nil {
		return err
	}
	defer out.Close()
	_ , err = io.Copy(out , src)
	if err != nil {
		return err
	}
	return err
}