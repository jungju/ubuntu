package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"
)

const (
	apiurl = "https://vagrantcloud.com/api/v1"
)

// VagrantCloudBox .
type VagrantCloudBox struct {
	Username     string
	Name         string
	Version      string
	Provider     string
	Token        string
	Filename     string
	ReturnUpload *ReturnUpload
}

type ReturnUpload struct {
	UploadPath string `json:"upload_path"`
}

func (v *VagrantCloudBox) boxurl(url string) string {
	return fmt.Sprintf("%s/box/%s/%s/version/%s/provider/%s/%s?access_token=%s", apiurl, v.Username, v.Name, v.Version, v.Provider, url, v.Token)
}

func (v *VagrantCloudBox) Upload() error {
	if err := v.checkError(); err != nil {
		return err
	}

	if err := v.create(); err != nil {
		return err
	}

	if err := v.upload(); err != nil {
		return err
	}

	return nil
}

func (v *VagrantCloudBox) checkError() error {
	if _, err := os.Stat(v.Filename); os.IsNotExist(err) {
		return err
	}
	return nil
}

func (v *VagrantCloudBox) create() error {
	v.ReturnUpload = &ReturnUpload{}
	createURL := v.boxurl("upload")
	println("URL Call : ", createURL)
	output, err := Executer("create", "curl", []string{createURL}, false, ".", nil)
	if err != nil {
		return err
	}
	println(output)
	return nil
}

func (v *VagrantCloudBox) upload() error {
	v.ReturnUpload = &ReturnUpload{}
	output, err := Executer("create", "curl", []string{"-X", "PUT", "--upload-file", v.Filename, v.ReturnUpload.UploadPath}, false, ".", nil)
	if err != nil {
		return err
	}
	println(output)
	return nil
}

func Executer(cmdType, cmdName string, cmdArgs []string, cmdWait bool, executeDir string, addEnv []string) (string, error) {
	cmd := exec.Command(cmdName, cmdArgs...)

	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, addEnv...)
	if executeDir != "" {
		cmd.Dir = executeDir
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}
	logrus.WithField("cmdType", cmdType).WithField("output", output).Debug("Output of executed")
	return string(output), nil
}
