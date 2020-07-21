package config

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/NYTimes/logrotate"
	"github.com/labstack/echo/v4"
)

const (
	accessLog    = "echo-blog-access.log"
	apiLog       = "echo-blog-api.log"
	eventLog     = "echo-blog-event.log"
	massEventLog = "echo-blog-mass-event.log"
)

// AccessLogSkipper アクセスログのスキップ条件
func AccessLogSkipper(c echo.Context) bool {
	return false
}

// AccessLog log writter
func AccessLog() io.Writer {
	return outputOf(accessLog)
}

func logPath() string {
	if env := Env(); env == DEV || env == PRD {
		return "/var/log/echo-blog/"
	}
	return "./"
}

func outputOf(file string) io.Writer {
	// TODO ECS移行が完了したら常にos.Stdoutを返す

	// write to stdout in docker
	if os.Getenv("DOCKER_LOG_STDOUT") == "1" {
		return os.Stdout
	}

	// otherwise, write to each path
	logfile, err := logrotate.NewFile(logPath() + addHostname(addPrefix(file)))
	if err != nil {
		log.Fatal(err)
	}
	return logfile
}

// addPrefix returns log file name w/ prefix if neccesarry.
func addPrefix(file string) string {
	prefix := os.Getenv("DOCKER_LOG_PREFIX")
	if prefix == "" {
		return file
	}

	return prefix + "-" + file
}

// addHostname returns log file name w/ hostname if neccessary.
func addHostname(file string) string {
	// 環境変数がセットされている場合のみログファイルにホスト名付与
	if os.Getenv("DOCKER_LOG_HOSTNAME_SUFFIX") != "1" {
		return file
	}

	hostname, err := os.Hostname()
	if err != nil {
		return file
	}

	return strings.Replace(file, ".log", "-"+hostname+".log", -1)
}
