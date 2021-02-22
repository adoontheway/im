package ctrl

import (
	"fmt"
	"im/util"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

const (
	EndPoint     = ""
	AccessKeyId  = ""
	AccessSecret = ""
	Bucket       = ""
)

func Upload(w http.ResponseWriter, r *http.Request) {
	UploadLocal(w, r)
	// TODO: need oss
	// UploadOss(w,r)
}

// 1. Save directoory ./mnt,need to be insure already exsit
func UploadLocal(w http.ResponseWriter, r *http.Request) {
	// get the uploaded file
	srcfile, header, err := r.FormFile("file")
	if err != nil {
		util.RespFail(w, err.Error())
		return
	}
	// create a new file
	suffix := ".png"
	// if filename contains suffix
	ofilename := header.Filename
	tmp := strings.Split(ofilename, ".")
	if len(tmp) > 1 {
		suffix = "." + tmp[len(tmp)-1]
	}
	// otherwise
	filetype := r.FormValue("filetype")
	if len(filetype) > 0 {
		suffix = filetype
	}
	filename := fmt.Sprintf("%d%04d%s", time.Now().Unix(), rand.Int31(), suffix)
	dstfile, err := os.Create("./mnt/" + filename)
	if err != nil {
		util.RespFail(w, err.Error())
		return
	}
	_, err = io.Copy(dstfile, srcfile)
	if err != nil {
		util.RespFail(w, err.Error())
		return
	}
	url := "/mnt/" + filename
	util.RespOk(w, url, "")
}

// Auth:public
func UploadOss(w http.ResponseWriter, r *http.Request) {
	srcfile, head, err := r.FormFile("file")
	if err != nil {
		util.RespFail(w, err.Error())
		return
	}
	// get file suffix
	suffix := ".png"
	// if name is xx.xx.png
	ofilename := head.Filename
	tmp := strings.Split(ofilename, ".")
	if len(tmp) != 0 {
		suffix = "." + tmp[len(tmp)-1]
	}
	// if filename was specified
	filetype := r.FormValue("filetype")
	if len(filetype) != 0 {
		suffix = filetype
	}
	// initialize ossclient
	client, err := oss.New(EndPoint, AccessKeyId, AccessSecret)
	if err != nil {
		util.RespFail(w, err.Error())
		return
	}
	// get bucket
	bucket, err := client.Bucket(Bucket)
	if err != nil {
		util.RespFail(w, err.Error())
		return
	}
	// set file name
	filename := fmt.Sprintf("mnt/%d%04d%s", time.Now().Unix(), rand.Int31(), suffix)
	// upload to bucket
	err = bucket.PutObject(filename, srcfile)
	if err != nil {
		util.RespFail(w, err.Error())
		return
	}
	// get url
	url := "http://" + Bucket + "." + EndPoint + "/" + filename
	util.RespOk(w, url, "")
}
