package file

import (
	"bufio"
	"io"
	"os"

	"github.com/rs/zerolog/log"
)

func ReadLog(logFile string) {
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

	// 从文件末尾开始逆向读取日志，直到找到任务开始标志或到达文件头部
	mission := reversRead(file, fileSize, &cursor)
	if mission == "" {
		log.Error().Msg("未找到任务开始标志")
		return
	}

	// 从任务开始标志开始正向读取日志
	homes := read(file, fileSize, &cursor)
	if len(homes) <= 0 {
		log.Error().Msg("未找到目标房间")
		return
	}
	log.Info().Strs("房间", homes).Msg("找到目标房间")
}

// reversRead 从文件末尾开始逆向读取日志，直到找到任务开始标志或到达文件头部。
// 参数：
//   - file: 要读取的文件指针
//   - fileSize: 文件大小
//   - cursor: 游标位置，用于记录读取的位置
//
// 返回值：无
// 功能：
//   - 从文件末尾开始逆向读取每行日志
//   - 检查每行日志是否包含任务开始标志
//   - 如果找到任务开始标志，则输出游标位置和任务开始标志，并结束读取
//   - 如果到达文件头部仍未找到任务开始标志，则输出错误信息并结束读取
func reversRead(file *os.File, fileSize int64, cursor *int64) (mission string) {
	log.Debug().Msg("开始逆向读取日志文件")

	defer log.Debug().Msg("逆向读取日志文件完毕")

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

// read函数从给定的文件中读取内容，并返回包含目标房间的字符串切片。
// 参数file是要读取的文件指针，fileSize是文件的大小，cursor是游标位置。
// 返回值room是包含目标房间的字符串切片。
func read(file *os.File, fileSize int64, cursor *int64) (room []string) {
	log.Debug().Msg("开始正向读取日志文件")

	defer log.Debug().Msg("正向读取日志文件完毕")

	// 创建一个带缓冲的读取器
	reader := bufio.NewScanner(file)

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

		// TODO: 多线程处理是否包含目标房间
		hasHome(line, &count, &room)
		if count >= 2 {
			break
		}
	}

	// 检查读取器是否发生错误
	if err := reader.Err(); err != nil {
		log.Fatal().Err(err).Msg("读取日志文件失败")
	}

	return
}
