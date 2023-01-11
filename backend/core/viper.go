package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
	"timedeliver/global"
)

func Viper() *viper.Viper {
	v := viper.New()

	prompt := promptui.Select{
		Label: "Where to find your config file",
		Items: []string{"Path", "Consul", "Etcd", "Env"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		panic(fmt.Errorf("Prompt run error: %s \n", err))
	}

	switch result {
	case "Path":
		fmt.Print("Input your file path:")

		//confPath := util.ReciveCmd()

		//写死记得上线要关掉
		v.AddConfigPath("C:\\Users\\86153\\OneDrive\\桌面\\develop\\TimeDeliver")

		v.SetConfigType("yaml")
	case "Consul":
		//Todo
		fmt.Println("Todoing")
	case "Etcd":
		//Todo
		fmt.Println("Todoing")
	case "Env":
		//Todo
		fmt.Println("Todoing")
	}

	err = v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	if err = v.Unmarshal(&global.Config); err != nil {
		fmt.Println(err)
	}

	v.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})

	return v
}
