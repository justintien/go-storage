package imagesrv

import (
	"os"
	"util"
)

type config struct {
	Root    string
	Cache   string
	URI     string
	MaxSize int64
}

var cfg = config{"/upload/file", "/cache/image", "/image", (1 << (10 * 2)) * 10}

func init() {
	if err := util.LoadConfig(&cfg, "imagesrv"); err != nil {
		return
	}

	os.MkdirAll(cfg.Root, os.ModePerm)
	os.MkdirAll(cfg.Cache, os.ModePerm)
}
