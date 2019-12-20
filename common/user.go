package common

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
)

const RegionPrefix = "region"
const AccessKeyPrefix = "access_key"
const SecretKeyPrefix = "secret_key"

var UserInfoPath string

type User struct {
	Region    string
	AccessKey string
	SecretKey string
}

func GetUser() *User {
	var region, accessKey, secretKey string

	f, err := os.Open(UserInfoPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		line = strings.TrimRight(line, "\n")
		userInfo := strings.Split(line, "=")
		if len(userInfo) == 2 {
			value := userInfo[1]
			switch userInfo[0] {
			case RegionPrefix:
				region = value
			case AccessKeyPrefix:
				accessKey = value
			case SecretKeyPrefix:
				secretKey = value
			}
		}
	}
	return &User{
		Region:    region,
		AccessKey: accessKey,
		SecretKey: secretKey,
	}
}

func init() {
	home, _ := Home()
	path := home + "/.sca/"
	UserInfoPath = home + "/.sca/config"
	os.MkdirAll(path, 0777)
}

func Home() (string, error) {
	user, err := user.Current()
	if nil == err {
		return user.HomeDir, nil
	}

	if "windows" == runtime.GOOS {
		return homeWindows()
	}

	return homeUnix()
}

func homeUnix() (string, error) {
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}
