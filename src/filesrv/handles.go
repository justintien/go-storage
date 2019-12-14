package filesrv

import (
	"httpsrv"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"util"
)

type File struct {
	Size int64  `json:"size"`
	Path string `json:"path"`
}

func Upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(cfg.MaxSize)

	file, handler, err := r.FormFile("file")

	if err != nil {
		httpsrv.SendFailure(w, r, 400, httpsrv.Failure{Message: "上传错误"})
		return
	}

	extname := filepath.Ext(handler.Filename)

	if isAllowedExt(extname) == false {
		httpsrv.SendFailure(w, r, 400, httpsrv.Failure{Message: "不允许的上传类型"})
		return
	}

	filename := strconv.FormatInt(time.Now().Unix(), 10) + extname

	if filename, err = util.Md5FromReader(file); err != nil {
		defer file.Close()

		log.Printf("fail to sum md5: %q\n", err)
		httpsrv.SendFailure(w, r, 400, httpsrv.Failure{Message: "fail to sum md5"})
		return
	}

	filename = util.Md5PathName(filename) + extname

	fileOfPath := filepath.Join(cfg.Root, filename)

	os.MkdirAll(fileOfPath[:strings.LastIndex(fileOfPath, "/")], os.ModePerm)

	f, err := os.OpenFile(fileOfPath, os.O_CREATE|os.O_WRONLY, 0660)
	if err != nil {
		httpsrv.SendFailure(w, r, 400, httpsrv.Failure{Message: "上传失败"})
		return
	}

	file, _, _ = r.FormFile("file")
	defer file.Close()
	defer f.Close()

	_, err = io.Copy(f, file)

	if err != nil {
		httpsrv.SendFailure(w, r, 400, httpsrv.Failure{Message: "上传失败"})
		return
	}

	httpsrv.SendSuccess(w, r, File{
		Size: handler.Size,
		Path: filepath.Join(cfg.URI, filename),
	})
}

func isAllowedExt(extname string) (isAllowedExt bool) {
	disallowed := []string{".exe", ".js"}
	isAllowedExt = true

	for _, v := range disallowed {
		if v == extname {
			isAllowedExt = false
			return
		}
	}

	return
}
