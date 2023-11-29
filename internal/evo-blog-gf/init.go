package evo_blog_gf

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

const (
	defaultConfigName  = "evo-blog-gf.yaml.yaml"
	recommendedHomeDir = "evo-blog-bf"
)

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(filepath.Join(home, recommendedHomeDir))
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(defaultConfigName)
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
	}
	_, err := fmt.Fprintln(os.Stdout, "Using config file:", viper.ConfigFileUsed())
	if err != nil {
		return
	}
}
