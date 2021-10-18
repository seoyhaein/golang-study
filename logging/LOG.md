# 로그에 대해서
0. 기초 설명은 http://golang.site/go/article/114-Logging 잘 되어 있음.

1. 로그의 인스턴스 관리에 대해서 시작할 것.

2. 로그는 다양한 패키지가 존재하고 있다. 하지만 내가 오픈소스등에서 경험한 것은 https://github.com/sirupsen/logrus 이 대표적이다. 이 패키지를 중심으로 살펴보도록 한다. hook 관련해서 잊어버렸던 것들 찾아서 자료 보강 하자.

3. 로그를 파일로 남길때 이를 효과적(파일사이즈, 백업등 - > log rotate) 다루는 방법들 및 기타 

4. github.com/sirupsen/logrus 다룰때 log 와 error 를 통합해서 다룬다.
