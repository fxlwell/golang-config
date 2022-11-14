package config

import (
	mysql "github.com/fxlwell/golang-mysql"
	"github.com/go-ini/ini"
)

var Mysql map[string]*mysql.Conf

const (
	CFG_KEY_MYSQL             = "Mysql"
	CFG_KEY_MYSQL_ADDR        = "Addr"
	CFG_KEY_MYSQL_USERNAME    = "Username"
	CFG_KEY_MYSQL_PASSWORD    = "Password"
	CFG_KEY_MYSQL_DATABASE    = "Database"
	CFG_KEY_MYSQL_DSNOPTIONS  = "Dsnoptions"
	CFG_KEY_MYSQL_MAXIDEL     = "Maxidle"
	CFG_KEY_MYSQL_MAXOPEN     = "Maxopen"
	CFG_KEY_MYSQL_MAXLIFETIME = "Maxlifetime"
	CFG_KEY_MYSQL_SLOWTIME    = "Slowtime"
	CFG_KEY_MYSQL_SLOWLOGGER  = "Slowlogger"
)

func init() {
	Mysql = make(map[string]*mysql.Conf)
	RegisterParser(parseMysql)
}

func parseMysql(fp *ini.File) error {
	psec, err := fp.GetSection(CFG_KEY_MYSQL)
	if err != nil {
		return ErrSectionNotExists
	}

	for _, sec := range psec.ChildSections() {
		nodeName := sec.Name()
		secName := nodeName[len(CFG_KEY_MYSQL)+1:]

		Mysql[secName] = &mysql.Conf{
			Addr:        GetKeyMust(sec, nodeName, CFG_KEY_MYSQL_ADDR).MustString(""),
			Username:    GetKeyMust(sec, nodeName, CFG_KEY_MYSQL_USERNAME).MustString(""),
			Password:    GetKeyMust(sec, nodeName, CFG_KEY_MYSQL_PASSWORD).MustString(""),
			Database:    GetKeyMust(sec, nodeName, CFG_KEY_MYSQL_DATABASE).MustString(""),
			DsnOptions:  GetKeyMust(sec, nodeName, CFG_KEY_MYSQL_DSNOPTIONS).MustString(""),
			MaxIdle:     GetKeyParentMust(sec, psec, nodeName, CFG_KEY_MYSQL_MAXIDEL).MustInt(1),
			MaxOpen:     GetKeyParentMust(sec, psec, nodeName, CFG_KEY_MYSQL_MAXOPEN).MustInt(1),
			MaxLifeTime: GetKeyParentMust(sec, psec, nodeName, CFG_KEY_MYSQL_MAXLIFETIME).MustDuration(),
			SlowTime:    GetKeyParentMust(sec, psec, nodeName, CFG_KEY_MYSQL_SLOWTIME).MustDuration(),
			SlowLogger:  GetKeyParentMust(sec, psec, nodeName, CFG_KEY_MYSQL_SLOWLOGGER).MustString(""),
		}
	}

	return nil
}
