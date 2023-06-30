package internal

import "time"

type FCVObject struct {
	FilePath    string
	MD5CheckSum string
	TimeStamp   time.Time
}
