# video2audio
go 通过ffmpeg.exe 批量转换视频成音频 
# 使用 
必须输入-videopath
* .\video2audio.exe -videopath=e:\videos
* .\video2audio.exe --help

Usage of C:\**\video2audio.exe:

  -audiopath string
  
        audio out root path  (default "d:\\output_audio")
        
  -filter string
  
        whitch file suffix should to extract audio ,default is mp4;mkv;flv,split by ';'  (default "mp4;mkv;flv")
        
  -keep
  
        keep the src video path,input true or false,default is true (default true)
        
  -suffix string
  
        default is aac (default "aac")
        
  -videopath string
  
        video path
