package file

import (
	"os"

	"github.com/rs/zerolog/log"
)

// ReadLoad函数从指定的日志文件中读取内容并进行处理。
// 参数logFile是日志文件的路径。
// 函数会打开日志文件并读取文件内容，然后根据特定的规则进行处理。
// 如果打开日志文件失败，函数会记录错误并终止程序。
// 函数会查找任务开始标志，并根据任务所在星图进行处理。
// 最后，函数会查找目标地形并记录结果。
func ReadLoad(logFile string) {
	file, err := os.Open(logFile)
	if err != nil {
		log.Fatal().Err(err).Msg("打开日志文件失败")
	}
	defer file.Close()

	stat, _ := file.Stat()
	// 文件大小
	fileSize := stat.Size()

	// 游标
	var cursor int64

	// 从文件末尾开始读取，找到任务开始标志
	mission := findMission(file, fileSize, &cursor)
	if mission == "" {
		log.Error().Msg("未找到任务开始标志")
		return
	}

	// 判断任务所在星图
	starChart := getStarChart(mission)
	if starChart == "" {
		log.Error().Msg("未找到任务所在星图，请设置yaml配置文件")
		return
	}

	homes := findHome(starChart, file, fileSize, &cursor)
	if len(homes) <= 0 {
		log.Error().Msg("未找到目标地形")
		return
	}
	log.Info().Strs("地形", homes).Msg("找到目标地形")
}
