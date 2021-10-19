# golang-study

1. flag 를 통해서 commandline 에서 사용자의 입력을 받고 app 실행. (사용자가 입력된 데이터를 통해서 app 설정 및 기능 동작)
   일반적으로 flag 를 통해서 가능하나 보다 더 추가적인 기능들이 필요하다면 cli(https://github.com/mitchellh/cli) 나 cobra (https://github.com/spf13/cobra) 를 이용한다.
   flag 를 통해서 app 이 동작시 log 기록을 나타나도록 하거나 아니면 log 기록을 나타나지 않게 샘플로 제작한다. 백엔드의 경우 실시간으로 해당 app 의 실시간 상태를 나타나게 해줄 수 있다.
   
2. 1에서와 같이 부가적인 패키지를 사용하게 되면 디펜던시 이슈가 발생함으로 go mod 나 go dep 을 활용해야 한다. 둘다 사용하기가 편한 툴이라 편한 것을 사용하면 되지만, hashicorp 에서 만든 패키지    를 이용할때는 go dep 을 사용해야 한다. 물론 go mod 와 관련된 이슈를 해결했다고는 하나 go dep 으로 패키지를 진행하는 것이 정신건강에 좋다. 대표적으로 vault 같은 것들이 있다.

3. go mod 적용. 처음 go mod init 으로 초기화 하고 go mod tidy 로 디펜던시를 추가해준다. 그후 go mod vendor 를 적용해서 현재 적용한 패키지들을 vendor 로 관리 한다. vendor 를 적용하는 이유는 현재 내가 쓰고 있는 패키지가(현재 사용중인 버전이) 향후 새롭게 업데이트 된 패키지의 버전이 하위 종속성을 지원하지 않게 업데이트가 되거나 또는 패키지가 더이상 업데이트가 되지 않을 경우에 내가 적용한 패키지를 보다 안정적으로 관리 하기 위해서 vendor 를 적용하는 것이 바람직 하다.

4. 10/16 까지 간단하게 내부에서 json 파일에 세팅 정보등을 저장해서 app 에서 활용하는 방식과 flag 를 통해서 app 사용자가 초기 세팅 이나 명령을 주어 app 을 실행 시킬때 필요한 간단한 소스를 만들었다. 부가적으로 몇가지 필요 기능들을 추가(필요함수들, makefile)하고 로그로 넘어가도록 하겠다. 오늘 다룬자료 중에 언마샬해주는 함수를 가져다 썼는데 (encoding/json) 예전에 읽은 자료가 있어서 링크를 걸어둔다. 시리얼라이즈와 마샬링을 잘 설명해주고 있다. (참고 : https://hyesun03.github.io/2019/09/08/marshalling-vs-serialization/)
.
5. Version 정보 입력 하기. app 의 버전 정보를 코드상에서 상수로 집어넣을 수도 있다. 하지만 빌드 타임에서 넣어줌으로써 보다 버전관리를 보다 편하게 해줄 수 있다. 간단한 예를 소스상에 했다. 자동화된 방식으로 진행하기 위해서는 몇가지 개선점이 필요하다. go build 시 -ldflags 를 적용하면 여기서 적용한 데이터를 빌드 시 해당 변수에 값을 적용할 수 있다. 즉 Version 이라는 변수에 값을 빌드 타임에 적용한다는 것이다.

예시) go build -ldflags=-X 'main.Version=0.0.2'

하지만, 보다 더 자동화 하기 위해선 Makefile 을 이용하여 go build 를 적용하여 제작할 수도 있다.

6. errors (참고 : https://mingrammer.com/gobyexample/errors/ ) 소스가 좀 복잡하게 작성되었지만 어렵지 않을듯 type assertion 에 대한 내용도 있긴 하지만 이 부분도 추가적으로 다루도록 함. 추후 소스에 반영하도록 함.

7. Makefile 사용법 
```
make standard

make gogofast
```

## 다루고자 하는 내용들 순서대로 (일단 이렇게 예상) - 중간에 기타 내용들은 추가적으로 다룸.
 log -> context -> 인터페이스 -> grpc -> 고루틴/채널
 
 grpc 를 다루면서 context, 인터페이스, 메서드 및 고루틴/채널 등을 다루도록 하겠다.
 
 
 ## 스터디 주제 (여러분들이 한번 주제를 추가하셔도 되고, 제가 리스트 해놓은 것들에 대해서 정리해주셔도 됩니다.)
 1. cli, cobra 사용법. (저는 주로 cli 를 사용하지만, 혹시 각각 자세한 사용법을 정리 해주시면 좋을듯 합니다.)
 2. github 통한 코드 리뷰 방식등. 
 3. panic, recover, defer 설명.
 4. go build 시 -ldflags 를 이용하여 빌드시 자동으로 Version 을 만들어 주기.
 5. https://www.sohamkamani.com/golang/type-assertions-vs-type-conversions/ 읽어보고 한번 정리해보기.
 6. proto 문법 정리해보기.

