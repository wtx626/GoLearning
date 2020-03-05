package main

import (
	"fmt"
	"github.com/robfig/cron"
	"os/exec"
)

func main() {
	i := 0
	c := cron.New()
	spec := "* * * * * ?"
	c.AddFunc(spec, func() {
		i++
		fmt.Println("execute per second", i)
		cmd := exec.Command("/home/william/workspace/crontest/echo.sh")
		//cmd := exec.Command("/home/tianxiongwu/echo.sh")
		//cmd := exec.Command("/Users/wutianxiong/GolandProjects/GoLearning/src/crontab/echo.sh")
		err := cmd.Start()
		fmt.Println(err)
		if err != nil {
			fmt.Println("cmd err:", err)
		}
		pid := cmd.Process.Pid
		fmt.Println("子进程PID", pid)
		//err = cmd.Wait()
		//if err != nil {
		//	fmt.Println(err)
		//}
	})
	c.Start()
	select {}
	//for {
	//	//	cmd := exec.Command("/home/tianxiongwu/echo.sh")
	//	//	//cmd := exec.Command("/Users/wutianxiong/GolandProjects/GoLearning/src/crontab/echo.sh")
	//	//	err := cmd.Start()
	//	//	fmt.Println(err)
	//	//	if err != nil {
	//	//		fmt.Println("cmd err:", err)
	//	//	}
	//	//	pid := cmd.Process.Pid
	//	//	fmt.Printf("父进程ID%d--》子进程PID%d", os.Getpid(), pid)
	//	//	err = cmd.Wait()
	//	//	if err != nil {
	//	//		fmt.Println(err)
	//	//	}
	//	//}
}
