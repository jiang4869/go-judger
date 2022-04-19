package judger

import (
	"io/ioutil"
	"os/exec"
)

func ExecShell(s_cmd string) ([]byte, error) {

	cmd := exec.Command("/bin/bash", "-c", s_cmd)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	
	cmd.Start()

	data, err := ioutil.ReadAll(stdout)
	if err != nil {
		return nil, err
	}
	return data, nil
}