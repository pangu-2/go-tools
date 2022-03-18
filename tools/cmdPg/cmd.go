package cmdPg

import (
	"io/ioutil"
	"log"
	"os/exec"
)

func Cmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	stdout2, err := cmd.StdoutPipe()
	if err != nil { //获取输出对象，可以从该对象中读取输出结果
		log.Fatal(err)
	}
	defer stdout2.Close()               // 保证关闭输出流
	if err := cmd.Start(); err != nil { // 运行命令
		log.Fatal(err)
	}
	if opBytes, err := ioutil.ReadAll(stdout2); err != nil { // 读取输出结果
		log.Fatal(err)
	} else {
		log.Println(string(opBytes))
	}
}
