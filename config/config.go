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
/*
	아래 메서드에서 리시버( (c *Config) )가 포인터냐 포인터가 아니냐에 따라 메서드를 호출하는 방식이 달라진다.
	지금은 포인터로 리시버를 만들어 주었기 때문에 server.go 에서 아래와 같이 작성해줄 수 있다.

	c.RegisterFlags(fs)

	만약 리시버를 value 로 만들어 줄 경우( (c Config) ) server.go 에서는 아래와 같이 작성해줄 것이다.
    (*c).RegisterFlags(fs)
*/
func (c *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.StringVar(&c.S1, "u", c.S1, "server address")
	fs.BoolVar(&c.Silent, "silent", c.Silent, "Log nothing to stdout/stderr")
}
