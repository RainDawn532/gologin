package biz

import (
	"time"
)

var (
	machineID     int64 // 机器 id 占10位, 十进制范围是 [ 0, 1023 ]
	serialNumber  int64 // 序列号占 12 位,十进制范围是 [ 0, 4095 ]
	lastTimeStamp int64 // 上次的时间戳(毫秒级), 1秒=1000毫秒, 1毫秒=1000微秒,1微秒=1000纳秒
)

func init() {
	// 程序初始化，初始化更新时间
	lastTimeStamp = time.Now().UnixMilli()
}

func SetMachineId(mid int64) {
	// 把机器 id 左移 12 位,让出 12 位空间给序列号使用
	machineID = mid << 12
}

func GetSnowflakeId() int64 {
	// 读当前时间
	curTimeStamp := time.Now().UnixMilli()
	// 如果是同一毫秒（与上一个请求时间对比）
	if curTimeStamp == lastTimeStamp {
		serialNumber++
		// 序列号占 12 位,十进制范围是 [ 0, 4095 ]
		if serialNumber > 4095 {
			// 休眠一毫秒
			time.Sleep(time.Millisecond)
			// 再次读当前时间
			curTimeStamp = time.Now().UnixMilli()
			// 以当前时间作为下一个请求的对比时间
			lastTimeStamp = curTimeStamp
			serialNumber = 0
		}

		// 取 41 位的二进制数 0x1FFFFFFFFFF 1111111111 1111111111 1111111111 1111111111 1( 这里共 41 个 1 )和时间戳进行并操作
		// 并结果( 右数 )第 42 位必然是 0,  低 41 位也就是时间戳的低 41 位
		rightBinValue := curTimeStamp & 0x1FFFFFFFFFF
		// 机器 id 占用10位空间,序列号占用12位空间,所以左移 22 位; 经过上面的并操作,左移后的第 1 位,必然是 0
		rightBinValue <<= 22
		// 生成你构造雪花ID
		id := rightBinValue | machineID | serialNumber
		return id
	}
	if curTimeStamp > lastTimeStamp {
		serialNumber = 0
		lastTimeStamp = curTimeStamp
		// 取 41 位的二进制数 0x1FFFFFFFFFF 1111111111 1111111111 1111111111 1111111111 1( 这里共 41 个 1 )和时间戳进行并操作
		// 并结果( 右数 )第 42 位必然是 0,  低 41 位也就是时间戳的低 41 位
		rightBinValue := curTimeStamp & 0x1FFFFFFFFFF
		// 机器 id 占用10位空间,序列号占用12位空间,所以左移 22 位; 经过上面的并操作,左移后的第 1 位,必然是 0
		rightBinValue <<= 22
		// 生成你构造雪花ID
		id := rightBinValue | machineID | serialNumber
		return id
	}
	if curTimeStamp < lastTimeStamp {
		return 0
	}
	return 0
}

func GenID() int64 {
	SetMachineId(1)
	return GetSnowflakeId()
}
