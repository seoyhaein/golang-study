package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
)

type Config struct {
	S1             string //test server https address
	Silent         bool
	ConfigFilePath string

	// fetching data from config file
	Filename string
	// 버전 관련 Makefile 관련 내용 다루기
	// Version string
}

// 10/16
/*
	이함수의 개선점은?  현재 이 struct 는 간단한 struct 지만 커지면??
	이함수가 개선되면 당연히 이 함수를 호출하는 부분도 수정되어야함.
*/
func DefaultConfig() Config {
	return Config{
		S1:             "https://daum.net",
		Silent:         true,
		ConfigFilePath: "./config/config.json",
	}
}

/*
	아래 메서드에서 리시버( (c *Config) )가 포인터냐 포인터가 아니냐에 따라 메서드를 호출하는 방식이 달라진다.
	지금은 포인터로 리시버를 만들어 주었기 때문에 server.go 에서 아래와 같이 작성해줄 수 있다.

	c.RegisterFlags(fs)

	만약 리시버를 value 로 만들어 줄 경우( (c Config) ) server.go 에서는 아래와 같이 작성해줄 것이다.
    (*c).RegisterFlags(fs)
*/
func (c *Config) RegisterConfig(fs *flag.FlagSet) (*Config, error) {

	// 일반적으로 리턴값에서 error 는 가장 오른쪽에 둔다.
	// error 의 인터페이스를 활용해서 error 메세지 표시를 바꿀수 있다.
	// 함수 만들어서 처리, 개선할 필요가 있다.
	var conf = c

	/*
		default value 가 세팅된다.
		형태론적으로 혼동이 올 수 있는데 세번째 파라미터는 디폴트 값이다.
	*/
	fs.StringVar(&c.S1, "u", c.S1, "server address")
	fs.BoolVar(&c.Silent, "silent", c.Silent, "Log nothing to stdout/stderr")
	fs.StringVar(&c.ConfigFilePath, "path", c.ConfigFilePath, "config file path")

	// 최소 디폴트 값의 url 주소를 가지고 온다.
	file, err := ioutil.ReadFile(c.ConfigFilePath)

	if err != nil {
		return nil, err
	}

	// '=' 기호로 교체되었다.
	err = json.Unmarshal(file, conf)

	if err != nil {
		return nil, err
	}

	/*
		리턴을 표현할 때

		// 명시적으로 str2 를 리턴값의 이름을 지어줌.
		func FunctionA(str1 string) (str2 string)  {

		str2 = str1
		return
		}

		이렇게 두는 것이 return str2 라고 명시하는 것보다 더 효과적이라는 test 결과를 예전에 웹 상에서 본적이 있다.
		참고 삼아 한번 적어 보았다.
	*/
	return conf, nil
}
