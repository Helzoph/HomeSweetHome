package file

import (
	"regexp"
	"strings"
	"wfhome/conf"

	"github.com/rs/zerolog/log"
)

// isMissionStar 判断字符串中是否包含特定的文本。
// 如果字符串包含 "ThemedSquadOverlay.lua: Lobby::Host_StartMatch: launching level for"，则返回 true，否则返回 false。
func isMissionStar(str string) bool {
	return strings.Contains(str, "launching level for")
}

// getMissionName 根据输入的字符串获取任务名称。
// 如果字符串不符合任务星级的格式，则返回空字符串。
// 如果找到匹配的任务名称，则返回任务名称。
// 如果未找到任务名称，则记录错误日志并返回空字符串。
func getMissionName(str string) string {
	if !isMissionStar(str) {
		return ""
	}

	re := regexp.MustCompile(`launching level for (.+?) \((.*?)\)`)
	matches := re.FindStringSubmatch(str)

	if len(matches) > 0 {
		return matches[1]
	} else {
		log.Error().Msg("未找到任务名称")
		return ""
	}
}

// hasHome函数用于检查字符串是否包含特定信息，并统计出现次数以及包含的房间名称。
// 参数str为待检查的字符串，参数count为出现次数的计数器指针，参数room为包含房间名称的切片指针。
// 函数内部通过正则表达式匹配字符串，如果匹配成功则增加出现次数计数器。
// 如果出现次数达到2次及以上，则函数直接返回。
// 函数还会遍历指定的房间列表，如果字符串中包含房间路径，则将房间名称添加到room切片中。
func hasHome(str string, count *int, room *[]string) {
	re := regexp.MustCompile(`^\d+.\d+ Sys \[Info\]: $`)
	if re.MatchString(str) {
		*count++
	}

	if *count >= 2 {
		return
	}

	// 是否包含指定的房间
	for _, r := range conf.Void.Rooms {
		if strings.Contains(str, r.Path) {
			*room = append(*room, r.Name)
		}
	}
}
