package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
  // RDWR = 読み書き \ CREATE = 無ければ作成 \ APPEND = 追記            //パーミッション
  logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
  if err != nil {
    log.Fatalln(err)
  }

  // io.MultiWriter(os.Stdout = ログの標準出力 \ lofileに出力)
  multiLogFile := io.MultiWriter(os.Stdout, logfile)
  //log.SetFlagsでログのフォーマットを指定
  log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
  //ログの出力先を指定
  log.SetOutput(multiLogFile)

}
