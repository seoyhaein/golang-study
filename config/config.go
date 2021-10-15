package config

import "flag"

type Config struct {
	S1     string //test server https address
	Silent bool
}

func DefaultConfig() Config {
	return Config{
		S1:     "https://daum.net",
		Silent: true,
	}
}

/* default value 가 세팅된다.*/
func (c *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.StringVar(&c.S1, "u", c.S1, "server address")
	fs.BoolVar(&c.Silent, "silent", c.Silent, "Log nothing to stdout/stderr")
}
