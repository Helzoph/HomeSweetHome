package file

import "github.com/rs/zerolog/log"

// revers 函数将字节切片中的元素反转。
// 参数 b 是要反转的字节切片。
func revers(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}

// processBuffer 是一个处理缓冲区的函数。
// 它接收一个指向游标的指针、一个字节切片 buff 和一个字节切片 char 作为参数。
// 函数返回一个更新后的缓冲区 buff、一个任务名称 mission 和一个布尔值，表示是否找到任务开始标志。
func processBuffer(cursor *int64, buff []byte, char []byte) ([]byte, string, bool) {
	var mission string
	if char[0] == '\n' {
		if len(buff) > 0 {
			// 反转 buff
			buff = revers(buff)

			// 获取任务名称
			mission = getMission(string(buff))
			if mission != "" {
				log.Info().Str("任务", mission).Msg("找到任务开始标志")
				log.Debug().Int64("游标", *cursor).Msg("游标位置")
				return buff, mission, true
			}
		}
		// 清空 buff
		buff = buff[:0]
	} else {
		buff = append(buff, char[0])
	}
	return buff, mission, false
}
