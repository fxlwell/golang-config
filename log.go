package config

import (
	log "github.com/fxlwell/golang-log"
	"github.com/go-ini/ini"
)

var Log map[string]*log.Conf

const (
	CFG_KEY_LOG         = "Log"
	CFG_KEY_LOG_LOGFILE = "LogFile"
	CFG_KEY_LOG_LEVEL   = "Level"
	CFG_KEY_LOG_EXPIRE  = "Expire"
	CFG_KEY_LOG_TRACE   = "Trace"
)

func init() {
	Log = make(map[string]*log.Conf)
	RegisterParser(parseLog)
}

func parseLog(fp *ini.File) error {
	psec, err := fp.GetSection(CFG_KEY_LOG)
	if err != nil {
		return ErrSectionNotExists
	}

	for _, sec := range psec.ChildSections() {
		nodeName := sec.Name()
		secName := nodeName[len(CFG_KEY_LOG)+1:]

		Log[secName] = &log.Conf{
			LogFile: GetKeyMust(sec, nodeName, CFG_KEY_LOG_LOGFILE).MustString(""),
			Level:   GetKeyParentMust(sec, psec, nodeName, CFG_KEY_LOG_LEVEL).MustInt(0x0),
			Expire:  GetKeyParentMust(sec, psec, nodeName, CFG_KEY_LOG_EXPIRE).MustInt(-1),
			Trace:   GetKeyParentMust(sec, psec, nodeName, CFG_KEY_LOG_TRACE).MustInt(-1),
		}
	}

	return nil
}
