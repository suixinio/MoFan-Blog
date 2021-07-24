package gsky

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"mofan-blog/utils"
	"mofan-blog/utils/errmsg"
	"net/http"
	"time"
)

type Gsky struct {
}

func (g Gsky) UploadFile(file multipart.File, fileHandler *multipart.FileHeader) (string, int) {
	//var params map[string]string
	//params["file"]
	if fileHandler == nil {
		fmt.Println("ninininininin")
	}
	fmt.Println("filesize:", fileHandler.Size)
	content, err := UploadGskyFile(utils.GskyServer, map[string]string{}, "file", fileHandler.Filename, file)
	if err != nil {
		fmt.Println(err)
		return "", errmsg.ERROR
	}
	var data map[string]interface{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		fmt.Println(err)
		return "", errmsg.ERROR
	}

	if data["status"] != 200.0 {
		return "", errmsg.ERROR
	}

	return fmt.Sprintf("%s%s", utils.BaseUrl, data["url"]), 0
}

//注意client 本身是连接池，不要每次请求时创建client
var (
	HttpClient = &http.Client{
		Timeout: 3 * time.Second,
	}
)

// 上传文件
// url                请求地址
// params        post form里数据
// nameField  请求地址上传文件对应field
// fileName     文件名
// file               文件
func UploadGskyFile(url string, params map[string]string, nameField, fileName string, file io.Reader) ([]byte, error) {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)

	formFile, err := writer.CreateFormFile(nameField, fileName)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(formFile, file)
	if err != nil {
		return nil, err
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	//req.Header.Set("Content-Type","multipart/form-data")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	resp, err := HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return content, nil
}
