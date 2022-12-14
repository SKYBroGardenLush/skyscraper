package command

import (
	"fmt"
	"github.com/SKYBroGardenLush/skyscraper/framework/cobra"
	"github.com/SKYBroGardenLush/skyscraper/framework/contract"
	"github.com/erikdubbelboer/gspt"
	"github.com/sevlyar/go-daemon"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

var cronDaemon = false

func initCronCommand() *cobra.Command {

	cronStartCommand.Flags().BoolVarP(&cronDaemon, "daemon", "d", false, "start serve daemon")
	cronCommand.AddCommand(cronStartCommand)

	return cronCommand
}

var cronCommand = &cobra.Command{
	Use:   "cron",
	Short: "定时任务相关命令",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()

		}
		return nil
	},
}

var cronStartCommand = &cobra.Command{
	Use:   "start",
	Short: "启动cron常驻进程",
	RunE: func(c *cobra.Command, args []string) error {
		// 获取容器
		container := c.Root().GetContainer()
		// 获取容器中app服务
		appService := container.MustMake(contract.AppKey).(contract.App)

		//设置cron的日志地址和进程id地址
		pidFolder := appService.RuntimeFolder()
		serverPidFile := filepath.Join(pidFolder, "cron.pid")
		logFolder := appService.LogFolder()
		serverLogfile := filepath.Join(logFolder, "cron.log")
		currentFolder := appService.BaseFolder()

		// daemon 模式
		if cronDaemon {
			//创键一个Context
			cntxt := &daemon.Context{
				// 设置pid文件
				PidFileName: serverPidFile,
				PidFilePerm: 0664,
				//设置日志文件
				LogFileName: serverLogfile,
				LogFilePerm: 0640,
				//设置工作路径
				WorkDir: currentFolder,
				//设置所有设置文件的mask，默认为750
				Umask: 072,
				//子进程的参数，按照这个参数设置，子进程命令为./hase cron start  --daemon=ture
				Args: []string{"", "cron", "start", "--daemon=true"},
			}
			//启动子进程，d不空表示当前是父进程，d为空表示当前是子进程
			d, err := cntxt.Reborn()
			if err != nil {
				return err
			}
			if d != nil {
				//父进程直接打印成功信息，不做任何操作
				fmt.Println("cron serve started,pid :", d.Pid)
				fmt.Println("log file", serverLogfile)
				return nil
			}

			//子进程执行Cron.Run
			defer cntxt.Release()
			fmt.Println("daemon started")
			gspt.SetProcTitle("hade cron")
			c.Root().Cron.Run()
			return nil

		}
		//not daemon mode
		fmt.Println("start cron job")
		content := strconv.Itoa(os.Getpid())
		fmt.Println("[PID]", content)
		err := ioutil.WriteFile(serverPidFile, []byte(content), 0664)
		if err != nil {
			return err
		}
		gspt.SetProcTitle("hade cron")
		c.Root().Cron.Run()
		return nil

	},
}
