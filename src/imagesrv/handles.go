package imagesrv

import (
	"fmt"
	"httpsrv"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/nfnt/resize"
)

func Image(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()
	fileOfPath := filepath.Join(cfg.Root, r.URL.Path)
	info, err := os.Stat(fileOfPath)

	if err != nil || info.IsDir() {
		httpsrv.SendFailure(w, r, 404, httpsrv.Failure{Message: "Not Found"})
		return
	}

	suffix := filepath.Ext(fileOfPath)

	if suffix != ".jpg" && suffix != ".jpeg" && suffix != ".png" {
		log.Printf("not a jpeg|png file %s\n", fileOfPath)

		httpsrv.SendFailure(w, r, 404, httpsrv.Failure{Message: "Not Found"})
		return
	}

	file, _ := os.Open(fileOfPath)
	defer file.Close()

	image, _ := decode(suffix, file)

	s, q := getScaleQuality(qs)

	log.Printf("%s: [s=%d:q=%d]", r.URL.Path, s, q)

	if s == 100 && q == 100 {
		log.Printf("serveFile: %s", fileOfPath)
		http.ServeFile(w, r, fileOfPath)
		return
	}

	fileOfPath = strings.Replace(fileOfPath, suffix, fmt.Sprintf(".%ds.%dq%s", s, q, suffix), -1)
	fileOfPath = strings.Replace(fileOfPath, cfg.Root, cfg.Cache, -1)

	os.MkdirAll(fileOfPath[:strings.LastIndex(fileOfPath, "/")], os.ModePerm)

	out, err := os.Create(fileOfPath)
	if err != nil {
		log.Printf("fail to create scale|quality file %s %q\n", fileOfPath, err)

		httpsrv.SendFailure(w, r, 404, httpsrv.Failure{Message: "Not Found"})
		return
	}

	image = resize.Resize(uint(image.Bounds().Dx()*s/100), 0, image, resize.Lanczos3)

	if err := encode(suffix, out, image, q); err != nil {
		log.Printf("fail to encode scale[%d]|quality[%d] jpeg file %s %q\n", s, q, fileOfPath, err)

		httpsrv.SendFailure(w, r, 404, httpsrv.Failure{Message: "Not Found"})
		return
	}

	http.ServeFile(w, r, fileOfPath)
	return
}

func getScaleQuality(qs url.Values) (s int, q int) {
	s, _ = strconv.Atoi(qs.Get("s"))
	q, _ = strconv.Atoi(qs.Get("q"))

	if s == 0 {
		s = 100
	}
	if q == 0 {
		q = 100
	}
	return
}

func decode(suffix string, file *os.File) (image.Image, error) {
	if suffix == ".png" {
		return png.Decode(file)
	}

	return jpeg.Decode(file)
}

func encode(suffix string, out *os.File, image image.Image, quality int) error {
	if suffix == ".png" {
		return png.Encode(out, image)
	}

	return jpeg.Encode(out, image, &jpeg.Options{Quality: quality})
}
