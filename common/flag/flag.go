package flag

import "flag"

var (
	SqliteAddr string
	ServerAddr string
)

func init() {
	flag.StringVar(&SqliteAddr, "sqlite", "tool.db", "sqlite address")
	flag.StringVar(&ServerAddr, "server", "0.0.0.0:2030", "server address")
	flag.Parse()
}
