package conf

import (
	"homesweethome/game"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var (
	StarChartList []game.StarChart
	StarChartMap  map[string]game.StarChart
	Missions      []game.Mission
)

func LoadConf(yamlPath string) {
	yamlMap := map[string]any{
		"StarChart": &StarChartList,
		"Missions":  &Missions,
	}

	// 读取配置文件
	for k, v := range yamlMap {
		loadfile(yamlPath + "/" + k + ".yaml")
		err := viper.UnmarshalKey(k, v)
		if err != nil {
			log.Fatal().Err(err).Str("设置", k).Msg("读取配置文件失败")
		}
		log.Debug().Str("设置", k).Msg("读取配置文件成功")
	}

	// 将星图列表转换为map
	scMap := make(map[string]game.StarChart)
	for _, starChart := range StarChartList {
		log.Debug().Str("星图", starChart.Name).Msg("开始转换星图列表为map")
		scMap[starChart.Name] = starChart
		log.Debug().Interface("星图", scMap).Msg("转换星图列表为map成功")
	}
	StarChartMap = scMap
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
