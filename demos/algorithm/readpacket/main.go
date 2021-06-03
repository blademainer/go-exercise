package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * 计算随机值
 * min  最小金额(默认为1, 0.01元)
 * max 最大金额(默认为20000， 200元)
 * total 剩余总金额
 * num 剩余总人数
 * 本次随机金额
 */
func calcRandomValue(min int64, max int64, total int64, num int64) int64 {
	if num == 1 {
		return total
	}

	// 确定本次随机范围
	low := min
	if (total - (num-1)*max) >= min {
		low = total - (num-1)*max
	}

	high := max
	if (total - (num-1)*min) <= max {
		high = total - (num-1)*min
	}

	ave := total / num
	if ave <= 1 {
		ave = 1
	}
	// 调整上限
	if high <= 2*ave {
		high = 2 * ave
	}

	// 生成随机值
	ram := rand.Int63n(high)

	// 防止溢出
	if ram < low {
		ram = low
	}

	if ram > high {
		ram = high
	}

	return ram
}

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))

	total := int64(1000)
	num := int64(10)
	for i := 0; i < 10; i++ {
		value := calcRandomValue(1, 200, total, num)
		fmt.Println(value)
		total -= value
		num--
	}

}
