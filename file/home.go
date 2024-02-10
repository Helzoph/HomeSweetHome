package file

import (
	"bufio"
	"homesweethome/conf"
	"os"
	"regexp"
	"strings"

	"github.com/rs/zerolog/log"
)

// getHome函数根据给定的字符串和星图路径，将符合条件的房间名称添加到家的列表中。
// 参数str是要检查的字符串，参数starChart是星图路径，参数home是家的列表指针。
// 函数会遍历星图中的房间，如果字符串中包含星图的路径，则将房间名称添加到家的列表中。
func getHome(str string, starChart string, home *[]string) {
	for _, r := range conf.StarChartMap[starChart].Rooms {
		if strings.Contains(str, r.Path) {
			*home = append(*home, r.Name)
		}
	}
}

// findHome函数从给定的日志文件中查找家庭信息。
// 它接受星图字符串、文件指针、文件大小和游标指针作为参数，并返回一个字符串切片。
// 函数会正序读取日志文件，逐行查找包含特定信息的行，并将符合条件的行添加到家庭信息切片中。
// 函数内部使用了带缓冲的读取器来提高读取效率，并根据游标位置设置读取器的起始位置。
// 在读取文件内容的过程中，函数会根据特定的正则表达式匹配行，计算符合条件的行的数量。
// 当数量达到1时，函数会调用getHome函数处理该行，并将结果添加到家庭信息切片中。
// 当数量达到2或更多时，函数会终止读取过程。
// 如果读取器发生错误，函数会抛出一个致命错误。
func findHome(starChart string, file *os.File, fileSize int64, cursor *int64) (home []string) {
	log.Debug().Msg("开始正序读取日志文件")

	defer log.Debug().Msg("正序读取日志文件完毕")

	// 创建一个带缓冲的读取器
	reader := bufio.NewScanner(file)

	// 计算起始游标位置
	offset := fileSize + *cursor
	log.Debug().Int64("游标位置", offset).Msg("移动游标")

	// 设置读取器的起始位置
	_, err := file.Seek(offset, 0)
	if err != nil {
		log.Fatal().Err(err).Msg("移动游标失败")
	}

	var count int

	// 逐行读取文件内容
	for reader.Scan() {
		line := reader.Text()
		// log.Debug().Msg(line)

		re := regexp.MustCompile(`^\d+.\d+ Sys \[Info\]: $`)
		if re.MatchString(line) {
			count++
		}

		if count == 1 {
			// TODO: 多线程处理是否包含目标地形
			getHome(line, starChart, &home)
		} else if count >= 2 {
			break
		}
	}

	// 检查读取器是否发生错误
	if err := reader.Err(); err != nil {
		log.Fatal().Err(err).Msg("读取日志文件失败")
	}

	return
}
