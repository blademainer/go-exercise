package main

import (
	"fmt"
)

// LinearRegression 用于计算简单线性回归模型中的截距和系数
type LinearRegression struct {
	A float64 // 截距
	B float64 // 回归系数
}

// Fit 函数用于拟合数据，计算出线性回归的系数和截距
func (lr *LinearRegression) Fit(X []float64, Y []float64) {
	if len(X) != len(Y) || len(X) == 0 {
		panic("输入的 X 和 Y 长度必须相同且不能为空")
	}

	sumX, sumY, sumXY, sumX2 := 0.0, 0.0, 0.0, 0.0
	n := float64(len(X))

	// 计算必要的统计值
	for i := 0; i < len(X); i++ {
		sumX += X[i]
		sumY += Y[i]
		sumXY += X[i] * Y[i]
		sumX2 += X[i] * X[i]
	}

	// 计算回归系数 B 和截距 A
	lr.B = (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	lr.A = (sumY - lr.B*sumX) / n
}

// Predict 函数用于给定 X 值进行预测
func (lr *LinearRegression) Predict(x float64) float64 {
	return lr.A + lr.B*x
}

func main() {
	// 示例数据
	X := []float64{1, 2, 3, 4, 5}
	Y := []float64{2, 4, 5, 4, 5}

	// 创建线性回归模型并进行拟合
	lr := LinearRegression{}
	lr.Fit(X, Y)

	// 输出拟合结果
	fmt.Printf("回归方程: y = %.2f + %.2f*x\n", lr.A, lr.B)

	// 使用模型进行预测
	predictX := 6.0
	predictedY := lr.Predict(predictX)
	fmt.Printf("预测结果: 当 x = %.2f 时, y = %.2f\n", predictX, predictedY)
}
