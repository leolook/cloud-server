package flag

import (
	"cloud-server/lib/log"
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
)

var (
	SqliteAddr string
	ServerAddr string
	path       string
)

func init() {
	flag.StringVar(&path, "conf", "../conf/dev_conf.json", "conf")
	flag.StringVar(&SqliteAddr, "sqlite", "tool.db", "sqlite address")
	flag.StringVar(&ServerAddr, "server", "0.0.0.0:2030", "server address")
	flag.Parse()
}

func Start() {
	//读取配置文件内容
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open %s,[err=%v]", path, err)
		return
	}

	//读取配置文件内容
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("failed to read all conf file,[err=%v]", err)
		return
	}

	if buf == nil || len(buf) <= 0 || string(buf) == "" {
		log.Fatalf("conf file is empty")
		return
	}

	var confMp map[string]string
	err = json.Unmarshal(buf, &confMp)
	if err != nil {
		log.Fatalf("failed to unmarshal conf,[err=%v]", err)
		return
	}

	for k, v := range confMp {
		err = flag.Set(k, v)
		if err != nil {
			log.Fatalf("failed to set flag,[err=%v]", err)
			return
		}
	}

}
