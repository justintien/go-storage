package filesrv

import (
	"os"
	"util"
)

type config struct {
	Root    string
	URI     string
	MaxSize int64
}

var cfg = config{"/upload/file", "/file", (1 << (10 * 2)) * 10}

func init() {
	if err := util.LoadConfig(&cfg, "filesrv"); err != nil {
		return
	}

	os.MkdirAll(cfg.Root, os.ModePerm)
}
