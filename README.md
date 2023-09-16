# Hello AWS Secret Generator! 👋
안녕하세요! 해당 도구는 Secret Manager에서 저장중인 Secret을 가져와서 파일의 형태로 저장하는 도구입니다. 

Hello! This tool is a tool that retrieves the secret stored in Secret Manager and saves it in the form of a file.

# Why AWS Secret Generator? ⚙️ 
AWS Secrets Manager는 AWS에서 제공하는 Secret 관리 도구입니다. 해당 서비스는 데이터베이스 자격 증명, API 키 및 기타 Secret을 수명 주기 동안 쉽게 로테이션하거나, 관리 및 검색할 수 있도록 지원합니다.

AWS Secrets Manager는 is a secrets management service that helps you protect access to your applications, services, and IT resources. This service enables you to easily rotate, manage, and retrieve database credentials, API keys, and other secrets throughout their lifecycle.

해당 도구를 사용하면 Secret을 관리하고, Secret을 사용하는 서비스에서는 Secret을 가져와서 파일의 형태로 로컬 PC에 저장할 수 있습니다.

if you use this tool, you can manage your secret and get secret from AWS Secret Manager and save it as a file in your local PC.

보안 상의 이유로 AWS 콘솔을 개발자에게 제공하지 않거나, Slack과 같은 메신저를 이용해 Secret을 공유하는 형태를 방지하고자 해당 도구를 만들었습니다.

For security reasons, we do not provide AWS console to developers or prevent sharing secrets using messengers such as Slack.



# Install with Go
[혹시 Golang이 설치되어 있지 않나요?](https://golang.org/doc/install)

[Do you have Golang installed?](https://golang.org/doc/install) 

```bash
$ go install github.com/YangTaeyoung/aws-secret-gen@v1.0.0
```

# Usage 🛠️

## Interface
```bash
$ aws-secret-gen -c <config path> -o <output path>
```

| 플래그 | 설명                                                              | 기본 값                     |
|-----|-----------------------------------------------------------------|--------------------------|
| `-c` | AWS Secret Generator의 Config를 저장하거나, 불러올 주소                     | `~/.aws-secret-gen/config` |
| `-o` | AWS Secrets Manager에서 가져온 Secrets를 저장할 파일 경로 (확장자를 포함할 수 있습니다.) | ✅ 기본값 없음 (required)      |

| Flag | Description                                                                                          | Default                     |
|-----|------------------------------------------------------------------------------------------------------|--------------------------|
| `-c` | Save or load the Config of AWS Secret Generator                                                      | `~/.aws-secret-gen/config` |
| `-o` | The file path to save the Secrets retrieved from AWS Secrets Manager (can include the extension.) | ✅ No default value (required)      |

## Trouble Shooting 👊

```bash
$ aws-secret-gen
> zsh: command not found: aims-cli
```

go로 설치한 프로그램을 실행할 때 발생하는 에러입니다. ~/.zshrc 파일(혹은 ~/.bashrc)의 하단에 다음과 같이 환경변수를 추가합니다.

error that occurs when running a program installed with go. Add the environment variable as follows at the bottom of the ~/.zshrc file (or ~/.bashrc).

```bash
# ...
export PATH="$HOME/go/bin:$PATH"
```
