package main

import (
	"fmt"
	"math"
)

// LinearRegression 用于计算多元线性回归模型中的系数
type LinearRegression struct {
	Coefficients []float64 // 回归系数，包括截距
}

// Fit 函数用于拟合数据，计算出多元线性回归的系数
func (lr *LinearRegression) Fit(X [][]float64, Y []float64) {
	if len(X) == 0 || len(X) != len(Y) {
		panic("输入的 X 和 Y 长度必须相同且不能为空")
	}

	n := len(Y)
	m := len(X[0]) + 1 // 增加一列用于表示截距

	// 构造增广矩阵 A 和向量 b
	A := make([][]float64, m)
	b := make([]float64, m)
	for i := 0; i < m; i++ {
		A[i] = make([]float64, m)
	}

	// 计算 A 和 b
	for i := 0; i < n; i++ {
		// 增加常数项 x0 = 1
		xi := append([]float64{1.0}, X[i]...)
		for j := 0; j < m; j++ {
			b[j] += Y[i] * xi[j]
			for k := 0; k < m; k++ {
				A[j][k] += xi[j] * xi[k]
			}
		}
	}

	// 使用高斯消元法解方程组 A * Coefficients = b
	lr.Coefficients = gaussElimination(A, b)
}

// Predict 函数用于给定 X 值进行预测
func (lr *LinearRegression) Predict(x []float64) float64 {
	// 增加常数项 x0 = 1
	x = append([]float64{1.0}, x...)
	if len(x) != len(lr.Coefficients) {
		panic("输入的特征数量与模型不匹配")
	}
	sum := 0.0
	for i := 0; i < len(lr.Coefficients); i++ {
		sum += lr.Coefficients[i] * x[i]
	}
	return sum
}

// gaussElimination 用于解决线性方程组
func gaussElimination(A [][]float64, b []float64) []float64 {
	m := len(b)
	for i := 0; i < m; i++ {
		// 找到主元
		maxRow := i
		for k := i + 1; k < m; k++ {
			if math.Abs(A[k][i]) > math.Abs(A[maxRow][i]) {
				maxRow = k
			}
		}
		// 交换行
		A[i], A[maxRow] = A[maxRow], A[i]
		b[i], b[maxRow] = b[maxRow], b[i]

		// 消元
		for k := i + 1; k < m; k++ {
			c := A[k][i] / A[i][i]
			b[k] -= c * b[i]
			for j := i; j < m; j++ {
				A[k][j] -= c * A[i][j]
			}
		}
	}

	// 回代求解
	x := make([]float64, m)
	for i := m - 1; i >= 0; i-- {
		sum := b[i]
		for j := i + 1; j < m; j++ {
			sum -= A[i][j] * x[j]
		}
		x[i] = sum / A[i][i]
	}
	return x
}

func main() {
	// 示例数据 (多元线性回归)
	X := [][]float64{
		{1, 2},
		{2, 4},
		{3, 5},
		{4, 4},
		{5, 5},
	}
	Y := []float64{2, 4, 5, 4, 5}

	// 创建线性回归模型并进行拟合
	lr := LinearRegression{}
	lr.Fit(X, Y)

	// 输出拟合结果
	fmt.Printf("回归系数: ")
	for i, coef := range lr.Coefficients {
		if i == 0 {
			fmt.Printf("%.2f", coef)
		} else {
			fmt.Printf(" + %.2f*x%d", coef, i)
		}
	}
	fmt.Println()

	// 使用模型进行预测
	predictX := []float64{6, 7}
	predictedY := lr.Predict(predictX)
	fmt.Printf("预测结果: 当 x = %v 时, y = %.2f\n", predictX, predictedY)
}
