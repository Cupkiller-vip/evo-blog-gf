package evo_blog_gf

import (
	"encoding/json"
	"evo-blog-gf/internal/log"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func GetBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		// 指定命令的名字，该名字会出现在帮助信息中
		Use: "evo-blog-gf.yaml",
		// 命令的简短描述
		Short: "A good Go practical project",
		// 命令出错时，不打印帮助信息。不需要打印帮助信息，设置为 true 可以保持命令出错时一眼就能看到错误信息
		SilenceUsage: true,
		// 指定调用 cmd.Execute() 时，执行的 Run 函数，函数执行失败会返回错误信息
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Init(logOptions())
			defer log.Sync()
			return run()
		},
		// 这里设置命令运行时，不需要指定命令行参数
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}
	cobra.OnInitialize(initConfig)
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "The path to the blog configuration file. Empty string for no configuration file.")
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	return cmd
}

func run() error {
	settings, _ := json.Marshal(viper.AllSettings())
	log.Infow(string(settings))
	log.Infow(viper.GetString("db.username"))
	return nil
}
