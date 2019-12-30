package local

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

const (
	host    = "unix:///var/run/docker.sock"
	version = "1.24"

	IMAGE           = "jdcloudchina/sca:%s"
	TARGET_PORT     = "9090"
	TARGET_VOL_PATH = "/function/code"

	RUNTIME_Python27 = "python27"
	RUNTIME_Python36 = "python36"
	RUNTIME_Python37 = "python37"
)

type DockerClient interface {
	PullImage(runtime string) error
	CreateDocker(hostVolPath, runtime string, hostPort string) error
	StartDocker() error
	StopDocker() error
	RemoveDocker() error
	Close()
}

type dockerClient struct {
	*client.Client
	ContainerId string
}

func NewDockerClient() (DockerClient, error) {
	cli, err := client.NewClient(host, version, nil, nil)
	if err != nil {
		return nil, err
	}

	return &dockerClient{
		Client: cli,
	}, nil
}

func (d *dockerClient) PullImage(runtime string) error {
	cmdStr := fmt.Sprintf("docker pull %s", getImageName(runtime))
	cmd := GenExecCommand(cmdStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()

	//	auth := types.AuthConfig{
	//		Username: "",
	//		Password: "",
	//	}
	//	authBytes, err := json.Marshal(auth)
	//	if err != nil {
	//		return err
	//	}
	//	authBase64 := base64.URLEncoding.EncodeToString(authBytes)
	//	option := types.ImagePullOptions{
	//		RegistryAuth: authBase64,
	//	}
	//	if _, err := d.ImagePull(context.Background(), getImageName(runtime), option); err != nil {
	//		return err
	//	}
	//	return nil
}

func (d *dockerClient) CreateDocker(hostVolPath, runtime string, hostPort string) error {
	exports := make(nat.PortSet)
	port, err := nat.NewPort("tcp", TARGET_PORT)
	if err != nil {
		return err
	}
	exports[port] = struct{}{}
	config := container.Config{
		Image:        getImageName(runtime),
		ExposedPorts: exports,
		Env: []string{
			"HANDLER=index.handler",
			"CHECKSUM=jsa-functionmd5",
			"CODEPATH=/function/code",
		},
	}

	// -p 9090:9090
	pBindMap := make(nat.PortMap)
	npBind := nat.PortBinding{
		HostPort: hostPort,
	}
	pBindMap[port] = []nat.PortBinding{npBind}

	//-v hostVolPath:TARGET_VOL_PATH
	mTmp := mount.Mount{
		Type:   mount.TypeBind,
		Source: hostVolPath,
		Target: TARGET_VOL_PATH,
	}

	hostConfig := container.HostConfig{
		PortBindings: pBindMap,
		Mounts:       []mount.Mount{mTmp},
	}
	nwConfig := network.NetworkingConfig{}

	//create docker
	containerName := fmt.Sprintf("%s-%s-%d", "sca", runtime, time.Now().Unix())
	resp, err := d.ContainerCreate(context.Background(), &config, &hostConfig, &nwConfig, containerName)
	if err != nil {
		return err
	}

	d.ContainerId = resp.ID

	return nil
}

func (d *dockerClient) StartDocker() error {
	option := types.ContainerStartOptions{}
	return d.ContainerStart(context.Background(), d.ContainerId, option)
}

func (d *dockerClient) StopDocker() error {
	return d.ContainerStop(context.Background(), d.ContainerId, nil)
}

func (d *dockerClient) RemoveDocker() error {
	option := types.ContainerRemoveOptions{
		Force: true,
	}
	return d.ContainerRemove(context.Background(), d.ContainerId, option)
}

func (d *dockerClient) Close() {
	d.Close()
	d.ContainerId = ""
}

func getImageName(runtime string) string {
	return fmt.Sprintf(IMAGE, runtime)
}

func GenExecCommand(commandStr string) *exec.Cmd {
	if runtime.GOOS == "windows" {
		return exec.Command("cmd.exe", "/C", commandStr)
	}
	return exec.Command("/bin/bash", "-c", commandStr)
}
