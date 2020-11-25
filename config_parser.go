package config_parser

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const dot = "."

type Config struct {
	Mymap  map[string]string
	strcet string
}

func (cfg *Config) InitConfig(filepath string) {
	cfg.Mymap = make(map[string]string)

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		str := strings.TrimSpace(string(line))
		fmt.Println(str) //打印每行

		if strings.Index(str, "#") == 0 || strings.Index(str, ";") == 0 {
			continue
		}

		idx1 := strings.Index(str, "[")
		idx2 := strings.LastIndex(str, "]")

		if idx1 > -1 && idx2 > -1 && idx2 > idx1+1 {
			cfg.strcet = strings.TrimSpace(str[idx1+1 : idx2])
			continue
		}

		if len(cfg.strcet) == 0 {
			continue
		}

		idx := strings.Index(str, "=")
		if idx < 0 {
			continue
		}

		key := strings.TrimSpace(str[:idx])
		if len(key) == 0 {
			continue
		}

		val := strings.TrimSpace(str[idx+1:])
		//防止 值的后面跟有#或//的注释符合
		pos := strings.Index(val, "\t#")
		if pos > -1 {
			val = val[0:pos]
		}

		pos = strings.Index(val, " #")
		if pos > -1 {
			val = val[0:pos]
		}

		pos = strings.Index(val, "\t//")
		if pos > -1 {
			val = val[0:pos]
		}

		pos = strings.Index(val, " //")
		if pos > -1 {
			val = val[0:pos]
		}

		if len(val) == 0 {
			continue
		}

		//cfg.Mymap[key] = strings.TrimSpace(val)

		newkey := cfg.strcet + dot + key

		cfg.Mymap[newkey] = strings.TrimSpace(val)
	}

}

func (cfg Config) Read(node, key string) string {
	key = node + dot + key
	v, found := cfg.Mymap[key]
	if !found {
		return ""
	}
	return v
}
