package convert

import (
	"errors"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetAllFiles(path string, filter []string) ([]string, error) {
	out := make([]string, 0)
	if _, err := os.Stat(path); err != nil {
		return out, errors.New(path + "not exist")
	}
	err := filepath.Walk(path, func(file_path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			for _, f := range filter {
				if filepath.Ext(file_path) == f {
					out = append(out, file_path)
					break
				}
			}
		}
		return err
	})
	return out, err
}

func ExtractAudio(input_video, output_audio string) {

	log.Println(input_video, output_audio)
	// "\"" + input_video + "\""
	// 这里浪费了很多时间，一个小时调试为何不成功，总是返回参数错误，原来 变量无需多此一举去添加引号，添加了引号反而会失败，因为go底层调用cmd之前已经做过处理了

	//cmdstr := "E:\\src-test\\go\\ffmpeg.exe" + " " + "-i" + " " + "\"" + input_video + "\"" + " " + "-vn" + " " + "-y" + " " + "-acodec" + " " + "copy" + " " + "\"" + output_audio + "\""
	//args := []string{"E:\\src-test\\go\\ffmpeg.exe", "-i", "\"" + input_video + "\"", "-vn", "-y", "-acodec", "copy", "\"" + output_audio + "\""}
	//cmdX := exec.Command(`C:\Program Files\Git\git-bash.exe`, "-c", cmdstr)
	cmd := exec.Command("E:\\src-test\\go\\ffmpeg.exe", "-i", input_video, "-vn", "-y", "-acodec", "copy", output_audio)
	buf, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error(), " the log file show details")
		log.Println(err.Error())
	}
	log.Println(string(buf))

	//fmt.Println(`./ffmpeg.exe`, "-i", "\""+input_video+"\"", "-vn", "-y", "-acodec", "copy", "\""+output_audio+"\"")

}

func Run(files []string, outpath, audio_suffix, src_root string, keep_dir bool) error {
	for index, file := range files {

		audio := outpath
		if keep_dir {
			audio = filepath.Join(outpath, strings.TrimSuffix(file[len(src_root):], filepath.Ext(file)))
		} else {
			audio = filepath.Join(outpath, strings.TrimSuffix(filepath.Base(file), filepath.Ext(file)))
		}
		if err := os.MkdirAll(filepath.Dir(audio), os.ModePerm); err != nil {
			log.Println(err.Error())
		}
		audio += "." + audio_suffix
		fmt.Println()
		fmt.Println("start convert:", file, "to", audio)
		ExtractAudio(file, audio)
		fmt.Println("current:", index+1, "total:", len(files), "percentage:", 100*float64(index+1)/float64(len(files)), "%")
	}
	return nil
}

func ConvertToString(text string) string {
	var decodeBytes, _ = simplifiedchinese.GBK.NewDecoder().Bytes([]byte(text))
	return string(decodeBytes)
}

func SetLogFile(name string) (err error) {
	appLog := filepath.Join(name + ".log")
	logFile, err := os.OpenFile(appLog, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(logFile)
	return nil
}
