package ctrl

import "os"

func init() {
	os.MkdirAll("./mnt", os.ModePerm)
}
