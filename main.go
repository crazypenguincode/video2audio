package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"video2audio/convert"
)

var video_path = flag.String("videopath", "", "video path")
var audio_out_path = flag.String("audiopath", `d:\output_audio`, "audio out root path ")
var filter = flag.String("filter", "mp4;mkv;flv", "whitch file suffix should to extract audio ,default is mp4;mkv;flv,split by ';' ")
var keep_dir = flag.Bool("keep", true, "keep the src video path,input true or false,default is true")
var audio_suffix = flag.String("suffix", "aac", "default is aac")

// todo https://github.com/giorgisio/goav 用go的接口调用可以节省一个exe 也更灵活
func main() {
	fmt.Println("if output file in C: ,please run as administrator")
	flag.Parse()
	fmt.Println("vpath", *video_path)
	if *video_path == "" {
		fmt.Println("please input videopath")
		return
	}
	err := convert.SetLogFile("run")
	if err != nil {
		fmt.Println(err.Error())
	}
	flt := strings.Split(*filter, ";")
	filterx := flt[:0]
	for _, f := range flt {
		filterx = append(filterx, "."+f)
	}
	fmt.Println("suffix:", flt)
	files, err := convert.GetAllFiles(*video_path, filterx)
	if err != nil {
		fmt.Println(err.Error())
		log.Println(err.Error())
		return
	}
	convert.Run(files, *audio_out_path, *audio_suffix, *video_path, *keep_dir)
	fmt.Println("----------finish---------")
}
