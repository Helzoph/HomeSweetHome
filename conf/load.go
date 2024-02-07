package conf

import (
	"wfhome/game"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var (
	Void     game.StarChart
	Missions []game.Mission
)

func LoadConf() {
	loadfile("yaml/Missions.yaml")
	loadfile("yaml/StarChart.yaml")

	err := viper.UnmarshalKey("Void", &Void)
	if err != nil {
		log.Fatal().Err(err).Str("设置", "Void").Msg("读取配置文件失败")
	}

	err = viper.UnmarshalKey("Missions", &Missions)
	if err != nil {
		log.Fatal().Err(err).Str("设置", "Missions").Msg("读取配置文件失败")
	}
}

func loadfile(filename string) {
	log.Debug().Str("配置文件", filename).Msg("开始读取配置文件")
	viper.SetConfigFile(filename)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Str("配置文件", filename).Msg("读取配置文件失败")
	}

	log.Debug().Str("配置文件", filename).Msg("读取配置文件成功")
}
