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
)

func Upload(w http.ResponseWriter, r *http.Request) {
	UploadLocal(w, r)
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
