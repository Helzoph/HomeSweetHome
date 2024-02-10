package file

import (
	"homesweethome/conf"
	"strings"

	"github.com/rs/zerolog/log"
)

// getStarChart函数根据给定的任务名称查找任务所在的星图。
// 如果找到匹配的星图，则返回该星图的名称；否则返回空字符串。
// 参数：
//
//	mission：任务名称
//
// 返回值：
//
//	starChart：星图名称，如果找到匹配的星图；否则为空字符串。
func getStarChart(mission string) (starChart string) {
	log.Debug().Str("任务", mission).Msg("开始查找任务所在星图")

	// 遍历星图列表，查找任务所在星图
	for _, m := range conf.Missions {
		if strings.Contains(mission, m.InternalName) {
			starChart = m.StarChart
			log.Info().Str("任务", mission).Str("星图", starChart).Msg("找到任务所在星图")
			return
		}
	}

	return ""
}
