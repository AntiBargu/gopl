package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // 画板中的第一种颜色
	blackIndex = 1 // 画板中的第二种颜色
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w, r)
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))

	return
}

func lissajous(out io.Writer, r *http.Request) {
	cycles := 5 // 完整的x振荡器变化次数

	const (
		res     = 0.001 // 角度分辨率
		size    = 100   // 图像画布大小
		nframes = 64    // 动画的帧数
		delay   = 8     // 以10ms为单位的延迟，总计80ms延迟
	)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		if k == "cycles" && v[0] != "" {
			var err error

			if cycles, err = strconv.Atoi(v[0]); err != nil {
				fmt.Fprintf(out, "cycles parameter wrong")
				return
			}
		}
	}

	freq := rand.Float64() * 3.0 // y振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 初相位差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1) // 创建画板
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay) // 指定动画的延迟
		anim.Image = append(anim.Image, img)   // 追加帧
	}
	gif.EncodeAll(out, &anim) // 注意：忽略编码错误
}
