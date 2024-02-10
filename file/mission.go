package file

import (
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/rs/zerolog/log"
)

// isMissionStar 判断字符串中是否包含"launching level for"子串。
func isMissionStar(str string) bool {
	return strings.Contains(str, "launching level for")
}

// getMission函数从给定的字符串中提取任务名称。
// 如果字符串不符合任务名称的格式，则返回空字符串。
// 任务名称的格式为"launching level for {任务名称} ({任务难度})"。
// 如果找到任务名称，则返回去掉任务难度后的任务名称。
// 如果未找到任务名称，则记录错误日志并返回空字符串。
func getMission(str string) (mission string) {
	if !isMissionStar(str) {
		return ""
	}

	// 利用正则表达式提取任务名称
	re := regexp.MustCompile(`launching level for (.+?) \((.*?)\)`)
	matches := re.FindStringSubmatch(str)

	if len(matches) > 0 {
		mission = matches[1]
	} else {
		log.Error().Msg("未找到任务名称")
		return ""
	}

	// 去掉任务难度
	mission = strings.Replace(mission, "_Hard", "", -1)

	return
}

// findMission函数从给定的文件中逆序读取日志，查找并返回任务字符串。
// 参数file是要读取的文件指针，fileSize是文件的大小，cursor是当前读取位置的游标。
// 函数返回找到的任务字符串。
func findMission(file *os.File, fileSize int64, cursor *int64) (mission string) {
	log.Debug().Msg("开始逆序读取日志文件")

	defer log.Debug().Msg("逆序读取日志文件完毕")

	// 存储每行日志
	buff := make([]byte, 0, 1024)
	// 存储每个字符
	char := make([]byte, 1)

	var flag bool

	// 从文件末尾开始读取
	for {
		*cursor -= 1
		// 移动游标到文件末尾
		_, _ = file.Seek(*cursor, io.SeekEnd)

		// 读取一个字符
		_, err := file.Read(char)
		if err != nil {
			log.Fatal().Err(err).Msg("读取日志文件失败")
		}

		buff, mission, flag = processBuffer(cursor, buff, char)
		if flag {
			return
		}

		// 到达文件头部
		if *cursor == -fileSize {
			break
		}
	}

	return
}
