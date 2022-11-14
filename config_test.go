package config

import (
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	if err := Init("./template"); err != nil {
		panic(err)
	}
	m.Run()
}

func TestLog(t *testing.T) {
	if Log["error"].LogFile != "log/error.log-*-*-*" ||
		Log["run"].Level != 0x1 ||
		Log["access"].Expire != 30 ||
		Log["access"].Trace != 0 {
		t.Fatal(Log)
	}
}

func TestMysql(t *testing.T) {
	if Mysql["default-slave"].Addr != "127.0.0.1:3306" ||
		Mysql["default-master"].Username != "root" ||
		Mysql["user-slave"].Password != "123456" ||
		Mysql["default-master"].Database != "db_test" ||
		Mysql["user-slave"].DsnOptions != "charset=utf8mb4&parseTime=True" ||
		Mysql["default-master"].MaxIdle != 8 ||
		Mysql["user-master"].MaxOpen != 32 ||
		Mysql["default-master"].MaxLifeTime != 300*time.Second ||
		Mysql["user-master"].SlowTime != 200*time.Millisecond ||
		Mysql["default-master"].SlowLogger != "run" ||
		Mysql["user-slave"].MaxIdle != 16 ||
		Mysql["user-slave"].MaxOpen != 128 ||
		Mysql["user-slave"].MaxLifeTime != 600*time.Second ||
		Mysql["user-slave"].SlowTime != 100*time.Millisecond ||
		Mysql["user-slave"].SlowLogger != "slow" {
		t.Fatal(Mysql)
	}
}
