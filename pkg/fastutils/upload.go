package fastutils

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"time"

	"emperror.dev/errors"
	"github.com/valyala/fasthttp"
)

// UploadFast upload
func UploadFast(url, uploadFile, field, tid string, timeOut time.Duration) ([]byte, error) {
	// 新建一个缓冲，用于存放文件内容
	bodyBufer := &bytes.Buffer{}
	// 创建一个multipart文件写入器，方便按照http规定格式写入内容
	bodyWriter := multipart.NewWriter(bodyBufer)
	// 从bodyWriter生成fileWriter,并将文件内容写入fileWriter,多个文件可进行多次
	fileWriter, err := bodyWriter.CreateFormFile(field, path.Base(uploadFile)) // "file"
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	file, er1 := os.Open(uploadFile)
	if er1 != nil {
		fmt.Println(er1)
		return nil, er1
	}
	// 不要忘记关闭打开的文件
	defer file.Close()
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	// 关闭bodyWriter停止写入数据
	_ = bodyWriter.Close()
	contentType := bodyWriter.FormDataContentType()
	// 构建request，发送请求
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	req.Header.SetContentType(contentType)
	// 直接将构建好的数据放入post的body中
	req.SetBody(bodyBufer.Bytes())
	req.Header.SetMethod("POST")
	req.Header.Set("tid", tid)
	req.Header.Set("sid", tid)
	req.SetRequestURI(url)

	err = fasthttp.DoTimeout(req, resp, timeOut)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		return nil, errors.New("uplodd fail")
	}
	return resp.Body(), nil
}
