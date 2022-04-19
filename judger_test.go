package judger

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_Case1(t *testing.T) {
	// test in linux 
	ExecShell("gcc -o file/main file/main.c")

	r, err := Run(1000, 2000, 128*1024*1024, 32*1024*1024, 10000, 200, 0, 0, 0, []string{}, []string{}, "file/main", "file/1.in", "file/1.out", "file/1.out", "file/judger.log", "c_cpp")
	if err != nil {
		t.Error(err)
	}

	b, _ := json.Marshal(r)
	fmt.Println(string(b))

}
