package logfile

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"ffCommon/log/log"
	"ffCommon/util"
)

// 输出日志到文本的 logger
type loggerFileNormal struct {
	// 输出相关
	out          io.Writer        // 输出目标
	chLogRequest chan *logRequest // 等待写入到 out 的请求

	// 停止输出相关
	run     bool
	chClose chan struct{}

	file *os.File // 文件对象

	filePath    string // 日志文件存储的绝对目录
	filePrefix  string // 日志文件名的前缀
	outLen      int    // 累计输出的日志长度
	outLenLimit int    // 累计输出的日志长度达到此值后，将自动换输出目标，目前仅用于日志文件输出

	fileLenLimit int // 文件长度限制
}

// 写 log 对象内容到日志文件
func (l *loggerFileNormal) closeFile() {
	// 关闭文件
	l.out = nil
	if l.file != nil {
		l.file.Close()
		l.file = nil
	}
	l.outLen = 0
}

// 切换输出目标(本方法，不会被外界直接调用到，模块内部确保锁的正确性)
func (l *loggerFileNormal) switchOut(forceSwitch bool) (err error) {
	// 关闭
	l.closeFile()

	// 实际应该使用的名称, 以及写入长度限制
	latestName, outLenLimit := latestName(l.filePath, l.filePrefix, l.fileLenLimit, forceSwitch)
	l.outLenLimit = outLenLimit

	// 创建文件
	if l.file, err = util.CreateFile(path.Join(l.filePath, latestName)); err != nil {
		return err
	}

	l.out = l.file

	return err
}

// 写 log 对象内容到日志文件
func (l *loggerFileNormal) write(one *logRequest) (err error) {
	defer lrq.back(one)

	// 正常写入文件
	nWriteLen, err := l.out.Write(one.buf)
	if err == nil {
		// 输出到标准输出
		if one.stdout {
			os.Stdout.Write(one.buf)
		}

		l.outLen += nWriteLen
		if l.outLenLimit > 0 && l.outLen >= l.outLenLimit {
			l.switchOut(true)
		}
		return err
	}

	// 尝试切换文件
	if err = l.switchOut(true); err != nil {
		return err
	}

	// 再次尝试写入文件
	nWriteLen, err = l.out.Write(one.buf)
	if err == nil {
		// 输出到标准输出
		if one.stdout {
			os.Stdout.Write(one.buf)
		}

		l.outLen += nWriteLen
		if l.outLenLimit > 0 && l.outLen >= l.outLenLimit {
			l.switchOut(true)
		}
		return err
	}

	return err
}

// goroutine 持续写
func (l *loggerFileNormal) goWrite(extras ...interface{}) {
	defer func() {
		l.closeFile()
	}()

	var err error
	var one *logRequest
	isOutOk := true

deadloop:
	for {
		select {
		case one = <-l.chLogRequest:
			if isOutOk {
				// 尝试正常写入
				if err = l.write(one); err != nil {
					isOutOk = false

					// 输出错误
					fmt.Println(err)

					// 向标准输出写入
					if one.stdout {
						fmt.Println(string(one.buf))
					}
				}
			} else if one.stdout {
				// 向标准输出写入
				fmt.Println(string(one.buf))
			}
		case <-l.chClose:
			break deadloop
		}
	}
}

// 追加到输出管道末尾
func (l *loggerFileNormal) logAsync(s string, stdout bool) bool {
	now := time.Now() // get this early.

	one := lrq.apply()

	// 缓冲区重置
	one.buf = one.buf[:logPrefixSize]
	one.stdout = stdout

	// 时间前缀
	hour, min, sec := now.Clock()
	one.buf[0] = byte('0' + hour/10)
	one.buf[1] = byte('0' + hour%10)
	one.buf[2] = byte(':')

	one.buf[3] = byte('0' + min/10)
	one.buf[4] = byte('0' + min%10)
	one.buf[5] = byte(':')

	one.buf[6] = byte('0' + sec/10)
	one.buf[7] = byte('0' + sec%10)
	one.buf[8] = byte(':')

	// 微秒, 6位
	microsecond := now.Nanosecond() / 1e3 // [0, 999999]
	base := 100000
	for index := 0; index < 6; index++ {
		v := microsecond / base
		microsecond -= v * base
		base /= 10
		one.buf[9+index] = byte('0' + v)
	}
	one.buf[15] = byte(' ')

	// 具体的日志内容
	one.buf = append(one.buf, s...)

	// 是否需要追加换行
	if len(s) == 0 || s[len(s)-1] != '\n' {
		one.buf = append(one.buf, '\n')
	}

	// 添加到待写入 out 的管道内
	l.chLogRequest <- one

	return true
}

// Printf calls l.logAsync to print to the loggerFileNormal.
// Arguments are handled in the manner of fmt.Printf.
func (l *loggerFileNormal) Printf(format string, v ...interface{}) {
	if l.run {
		l.logAsync(fmt.Sprintf(format, v...), false)
	}
}

// Print calls l.logAsync to print to the loggerFileNormal.
// Arguments are handled in the manner of fmt.Print.
func (l *loggerFileNormal) Print(v ...interface{}) {
	if l.run {
		l.logAsync(fmt.Sprint(v...), false)
	}
}

// Println calls l.logAsync to print to the loggerFileNormal.
// Arguments are handled in the manner of fmt.Println.
func (l *loggerFileNormal) Println(v ...interface{}) {
	if l.run {
		l.logAsync(fmt.Sprintln(v...), false)
	}
}

// Stop stop or recover output. fatal not support.
func (l *loggerFileNormal) Stop(stop bool) {
	// 关闭
	l.run = !stop
}

// Close close output. fatal not support.
func (l *loggerFileNormal) Close() {
	// 不再向待输出缓冲区写入
	// 停止输出协程(将待输出缓冲区内的内容, 输出到目标)
	// 待输出缓冲区内有残留, 直到实例对象被回收
	l.run = false
	close(l.chClose)
}

// newFileLoggerNormal creates a new file Logger.
func newFileLoggerNormal(filePath string, filePrefix string, fileLenLimit int) (l log.Logger, err error) {
	f := &loggerFileNormal{
		chLogRequest: make(chan *logRequest, logsCacheCount),

		run:     true,
		chClose: make(chan struct{}),

		filePath:    filePath,
		filePrefix:  filePrefix,
		outLenLimit: fileLenLimit,

		fileLenLimit: fileLenLimit,
	}

	if err = f.switchOut(false); err == nil {
		go util.SafeGo(f.goWrite, nil)
	}

	return f, err
}
