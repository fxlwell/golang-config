package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-ini/ini"
)

type ParseFunc func(*ini.File) error

var parseFuncs []ParseFunc
var ErrSectionNotExists = errors.New("section not exists")

func RegisterParser(f ParseFunc) {
	parseFuncs = append(parseFuncs, f)
}

func Init(conf_path string) error {
	fp, err := init_ini_file(conf_path)
	if err != nil {
		return err
	}

	for _, f := range parseFuncs {
		if err := f(fp); err != nil && err != ErrSectionNotExists {
			return err
		}
	}

	return nil
}

func init_ini_file(dirname string) (*ini.File, error) {
	fd, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}

	fnames, err := fd.Readdirnames(0)
	if err != nil {
		return nil, err
	}

	fp := ini.Empty()

	for _, fname := range fnames {
		if !strings.HasSuffix(fname, ".ini") {
			continue
		}

		fullname := filepath.Join(dirname, fname)

		if content, err := ioutil.ReadFile(fullname); err != nil {
			return nil, err
		} else if err := fp.Append([]byte(string(content))); err != nil {
			return nil, err
		}
	}

	return fp, nil
}

func GetKeyMust(sec *ini.Section, node, key string) *ini.Key {
	if k, err := sec.GetKey(key); err == nil {
		return k
	}
	panic(fmt.Errorf("config: [%s#%s] not exists", node, key))
	return nil
}

func GetKeyParentMust(sec, psec *ini.Section, node, key string) *ini.Key {
	if k, err := sec.GetKey(key); err == nil {
		return k
	}
	if k, err := psec.GetKey(key); err == nil {
		return k
	}
	panic(fmt.Errorf("config: [%s#%s] not exists", node, key))
	return nil
}
