## 간단히 기획하는 app

1. client 는 server 에게 shell script 를 전송하고, server 는 그 전송을 처리하여 그 결과 값을 client 에 전송한다.
2. client 는 복수일 수 있다. 여기서는 일단 server 는 한개로 한정한다. 

3. client 가 shell script 를 전송할 경우, server 에서 처리한다는 의미는 독립적으로 script 가 실행된다는 의미 이다. 
   이와 같은 경우는, 처리 시간이 문제가 발생할 수 있다. 즉, 서버에서 해당 shell script 가 실행되어서 그 결과를 받는데 까지 얼마 만큼의
   시간이 걸릴지 알 수 없다. 또한 shell script 를 실행하는 것은 별도의 프로세스가 담당한다.
   
   이러한 문제를 해결할때 두가지 조건을 생각할 수 있다. client 가 접속을 유지할 경우와 client 가 접속을 끊을 경우 이다.
   이러한 두가지 사항도 server 는 정상적으로 그 결과 값을 client 에게 전송해주어야 한다.

4. polling 방식은 구현 시나리오로써 제외한다.

### 구현 방식(아이디어-쉽게 바뀔 수 있음)

server 는 long-lived 방식의 stream 을 리턴하는 방식을 가진 method 를 하나 구성한다.
여러 client 가 접속할 수 있으므로 client_id int32 를 두고 sync.map 을 사용해서 safe 하게 client_id 를 저장할 수 있도록 한다. 
context 와 별도로, client 에서 접속을 끊을 경우 이러한 상태를 전송할 수 있는 make(chan bool), finished <-chan bool 채널을 하나 만들어 준다.

server 는 client 에에게 ping 을 지속적으로 보내서 client 가 접속 하고 있는지 파악해야 한다. 
만약 해당 ping 에 client 가 응답을 보내지 않는다면, client 가 disconnected 라고 판단한다.

server 는 전송된 shell script 의 처리 결과 상태를 기록으로 남긴다. 이때, 해당 기록은 파일로 기록한다. 
향후 저장되는 방식은 여러 툴을 사용할 수 있다. (zookeeper, sql, ldap, ...)

shell script 의 처리 상태가 저장된 후, 그 결과 값은 client 에 전송된다. 이때, 생각해볼 문제가 발생한다.

1. client 가 접속하고 있으며, shell script 의 결과 나왔을때 (가장 간단한 처리) 
   - 그냥 전송하고 받고 끝.
   
2. client 가 접속을 하지 않을때, shell script 의 결과 완료 및 미전송 기록 - 기본적으로 해당 shell script 가 수행된다.
   이후, client 가 재 접속 할 경우, 해당 결과를 전송하여야 한다.
   
   재접속할때 제일 처음 server 에서 미전송 기록이 있는지 확인하고 미전송 결과를 송수신 하는 method 가 필요. 
   만약 이때, 미전송 기록이 없으면 없다라는 사실을 리턴함. 
   
3. shell script 의 오류로 무한 루프 도는 것처럼 절대로 종료 하지 않는 오류 코드
   client 에서 timeout 을 설정할 수 있도록 하고 해당 timeout 을 넘어서면 server 에서 해당 서비스 또는 method 를 중단 시킨다.
    

