package file

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

// CopyLog 复制日志文件到指定路径
//
// 参数:
//
//	backupPath: 备份路径
//
// 返回值:
//
//	无
//
// 示例:
//
//	CopyLog("H:/backup/logs")
//
// 复制游戏日志文件到指定路径。
func CopyLog(backupPath string) {
	// 游戏日志文件路径
	logPath := filepath.Join(os.Getenv("LOCALAPPDATA"), "Warframe", "EE.log")

	// 复制日志文件
	cmd := exec.Command("cmd", "/C", "copy", "/Y", logPath, backupPath)
	err := cmd.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("复制日志文件失败")
	}
	log.Info().Str("备份日志文件名", backupPath).Msg("复制日志文件成功")
}

func DeleteLog(backupPath string) {
	// 删除日志文件
	cmd := exec.Command("cmd", "/C", "del", "/F", "/Q", backupPath)
	err := cmd.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("删除日志文件失败")
	}
	log.Info().Str("删除日志文件名", backupPath).Msg("删除日志文件成功")
}
