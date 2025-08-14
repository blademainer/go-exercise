// Package main 实现了针对数字订单号的TLV压缩算法
//
// 算法特点：
// 1. 针对主要由数字组成的订单号进行优化压缩
// 2. 使用base62编码将数字转换为更短的字符串
// 3. 使用TLV格式处理混合字符串（数字+字母）
// 4. 针对纯数字订单号使用特殊优化（N前缀格式）
// 5. 针对常见字符串使用映射表（m前缀格式，无长度字段）
//
// 压缩效果：
// - 纯数字订单号：23位 -> 14-15位（压缩率约60-65%）
// - 混合订单号：24位 -> 20位左右（压缩率约80-85%）
//
// 支持的字符集：[a-zA-Z0-9]
// 目标长度：< 18位
//
// 使用示例：
//
//	packer := NewTLVPacker(1)
//	compressed := packer.Compress("20250811074413372133086")
//	decompressed, _ := packer.Decompress(compressed)
package main

import (
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

// TLVPacker TLV压缩打包器
type TLVPacker struct {
	lengthDigits        int               // 长度字段占用的位数，默认为2
	mappingTable        map[string]string // 字符串到映射码的映射表
	reverseMappingTable map[string]string // 映射码到字符串的反向映射表
}

// NewTLVPacker 创建新的TLV打包器
func NewTLVPacker(lengthDigits int) *TLVPacker {
	if lengthDigits <= 0 {
		lengthDigits = 2 // 默认长度字段占2位
	}

	// 初始化常见字符串的映射表
	mappingTable := map[string]string{
		"SP":  "1", // SP -> 1
		"SB":  "2", // SB -> 2
		"PAY": "3", // PAY -> 3
		"ORD": "4", // ORD -> 4
		"TXN": "5", // TXN -> 5
		"REF": "6", // REF -> 6
	}

	// 创建反向映射表
	reverseMappingTable := make(map[string]string)
	for k, v := range mappingTable {
		reverseMappingTable[v] = k
	}

	return &TLVPacker{
		lengthDigits:        lengthDigits,
		mappingTable:        mappingTable,
		reverseMappingTable: reverseMappingTable,
	}
}

// Segment 表示一个数据段
type Segment struct {
	Type  string // 'n' for number, 's' for string, 'm' for mapped
	Value string // 原始值
}

// parseSegments 解析输入字符串，识别连续的数字和字符串
func (p *TLVPacker) parseSegments(input string) []Segment {
	var segments []Segment

	// 如果整个字符串都是数字，直接作为一个数字段处理
	if isNumeric(input) {
		segments = append(segments, Segment{Type: "n", Value: input})
		return segments
	}

	// 使用正则表达式识别连续的数字和非数字
	re := regexp.MustCompile(`(\d+|[^\d]+)`)
	matches := re.FindAllString(input, -1)

	for _, match := range matches {
		if isNumeric(match) {
			segments = append(segments, Segment{Type: "n", Value: match})
		} else {
			// 检查是否在映射表中
			if mappedValue, exists := p.mappingTable[match]; exists {
				segments = append(segments, Segment{Type: "m", Value: mappedValue})
			} else {
				segments = append(segments, Segment{Type: "s", Value: match})
			}
		}
	}

	return segments
}

// isNumeric 检查字符串是否为纯数字
func isNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// encodeNumber 将数字字符串转换为base62
func (p *TLVPacker) encodeNumber(numStr string) string {
	bigInt := &big.Int{}
	bigInt.SetString(numStr, 10)
	return bigInt.Text(62)
}

// decodeNumber 将base62字符串转换回数字
func (p *TLVPacker) decodeNumber(base62Str string) (string, error) {
	bigInt := &big.Int{}
	_, ok := bigInt.SetString(base62Str, 62)
	if !ok {
		return "", fmt.Errorf("invalid base62 string: %s", base62Str)
	}
	return bigInt.String(), nil
}

// formatLength 格式化长度字段，使用紧凑的编码来减少长度开销
func (p *TLVPacker) formatLength(length int) string {
	if length < 10 {
		// 对于小于10的长度，直接用数字表示
		return strconv.Itoa(length)
	} else if length < 36 {
		// a=10, b=11, ..., z=35
		return string(rune('a' + length - 10))
	} else if length < 62 {
		// A=36, B=37, ..., Z=61
		return string(rune('A' + length - 36))
	} else {
		// 如果长度超过61，这种情况在我们的场景中很少见
		// 使用特殊标记，但通常base62编码不会这么长
		panic(fmt.Sprintf("length %d is too large for single character encoding", length))
	}
}

// Compress 压缩输入字符串
func (p *TLVPacker) Compress(input string) string {
	if input == "" {
		return ""
	}

	segments := p.parseSegments(input)

	// 如果只有一个数字段，直接使用无TLV格式
	if len(segments) == 1 && segments[0].Type == "n" {
		value := p.encodeNumber(segments[0].Value)
		return "N" + value // 使用N前缀表示单一数字，无需长度字段
	}

	var result strings.Builder

	for _, segment := range segments {
		var value string

		if segment.Type == "n" {
			// 数字类型，进行base62压缩
			value = p.encodeNumber(segment.Value)
			// 构建TLV格式：Type + Length + Value
			length := p.formatLength(len(value))
			tlv := segment.Type + length + value
			result.WriteString(tlv)
		} else if segment.Type == "m" {
			// 映射类型，直接使用 m + 映射值，无需长度字段
			tlv := segment.Type + segment.Value
			result.WriteString(tlv)
		} else {
			// 字符串类型，不压缩
			value = segment.Value
			// 构建TLV格式：Type + Length + Value
			length := p.formatLength(len(value))
			tlv := segment.Type + length + value
			result.WriteString(tlv)
		}
	}

	return result.String()
}

// parseLength 解析长度字段
func (p *TLVPacker) parseLength(compressed string, pos int) (int, int, error) {
	if pos >= len(compressed) {
		return 0, 0, fmt.Errorf("insufficient data for length at position %d", pos)
	}

	firstChar := compressed[pos]
	if firstChar >= '0' && firstChar <= '9' {
		// 数字表示的长度（0-9）
		length := int(firstChar - '0')
		return length, 1, nil
	} else if firstChar >= 'a' && firstChar <= 'z' {
		// 小写字母表示的长度（10-35）
		length := int(firstChar - 'a' + 10)
		return length, 1, nil
	} else if firstChar >= 'A' && firstChar <= 'Z' {
		// 大写字母表示的长度（36-61）
		length := int(firstChar - 'A' + 36)
		return length, 1, nil
	}

	return 0, 0, fmt.Errorf("invalid length character: %c", firstChar)
}

// Decompress 解压缩字符串
func (p *TLVPacker) Decompress(compressed string) (string, error) {
	if compressed == "" {
		return "", nil
	}

	// 检查是否为单一数字格式（N前缀）
	if len(compressed) > 0 && compressed[0] == 'N' {
		value := compressed[1:] // 去掉N前缀
		decodedValue, err := p.decodeNumber(value)
		if err != nil {
			return "", fmt.Errorf("failed to decode single number: %v", err)
		}
		return decodedValue, nil
	}

	var result strings.Builder
	i := 0

	for i < len(compressed) {
		// 检查剩余长度是否足够读取Type
		if i >= len(compressed) {
			return "", fmt.Errorf("invalid TLV format: insufficient data for type at position %d", i)
		}

		// 读取Type (1位)
		segmentType := string(compressed[i])
		i++

		if segmentType == "m" {
			// 映射类型特殊处理：m + 单字符映射值，无长度字段
			if i >= len(compressed) {
				return "", fmt.Errorf("invalid mapping format: insufficient data for mapped value at position %d", i)
			}

			mappedValue := string(compressed[i])
			i++

			// 通过反向映射表还原
			if originalValue, exists := p.reverseMappingTable[mappedValue]; exists {
				result.WriteString(originalValue)
			} else {
				return "", fmt.Errorf("unknown mapped value: %s", mappedValue)
			}
		} else {
			// 其他类型需要读取长度字段
			// 读取Length
			length, lengthBytes, err := p.parseLength(compressed, i)
			if err != nil {
				return "", fmt.Errorf("failed to parse length: %v", err)
			}
			i += lengthBytes

			// 检查剩余长度是否足够读取Value
			if i+length > len(compressed) {
				return "", fmt.Errorf(
					"invalid TLV format: insufficient data for value at position %d, need %d bytes but only %d remaining",
					i, length, len(compressed)-i,
				)
			}

			// 读取Value
			value := compressed[i : i+length]
			i += length

			// 根据类型处理Value
			if segmentType == "n" {
				// 数字类型，进行base62解压
				decodedValue, err := p.decodeNumber(value)
				if err != nil {
					return "", fmt.Errorf("failed to decode number: %v", err)
				}
				result.WriteString(decodedValue)
			} else if segmentType == "s" {
				// 字符串类型，直接添加
				result.WriteString(value)
			} else {
				return "", fmt.Errorf("unknown segment type: %s", segmentType)
			}
		}
	}

	return result.String(), nil
}

// CalculateCompressionRatio 计算压缩率
func (p *TLVPacker) CalculateCompressionRatio(original, compressed string) float64 {
	if len(original) == 0 {
		return 0
	}
	return float64(len(compressed)) / float64(len(original))
}

// Demo 演示函数
func main() {
	// 创建TLV打包器（使用1位长度字段来减少开销）
	packer := NewTLVPacker(1)

	// 测试用例
	testCases := []string{
		"20250813133226111106311",  // 纯数字订单号（生产环境）
		"20250813132554SB10942143", // 混合订单号（开发环境）
		"202501011234567890123456", // 长数字订单号
		"20250101123ABC456DEF789",  // 多段混合订单号
		"PAY20250101123456789012",  // 前缀+数字
		"ORD123456789",             // 短订单号
	}

	fmt.Println("TLV Compression Demo")
	fmt.Println("====================")

	for i, testCase := range testCases {
		fmt.Printf("\nTest Case %d: %s\n", i+1, testCase)
		fmt.Printf("Original Length: %d\n", len(testCase))

		// 压缩
		compressed := packer.Compress(testCase)
		fmt.Printf("Compressed: %s\n", compressed)
		fmt.Printf("Compressed Length: %d\n", len(compressed))

		// 计算压缩率
		ratio := packer.CalculateCompressionRatio(testCase, compressed)
		fmt.Printf("Compression Ratio: %.2f\n", ratio)

		// 解压缩验证
		decompressed, err := packer.Decompress(compressed)
		if err != nil {
			fmt.Printf("Decompression Error: %v\n", err)
			continue
		}

		fmt.Printf("Decompressed: %s\n", decompressed)

		// 验证正确性
		if decompressed == testCase {
			fmt.Printf("✓ Compression/Decompression successful!\n")
		} else {
			fmt.Printf("✗ Compression/Decompression failed!\n")
		}

		// 检查是否满足18位限制
		if len(compressed) < 18 {
			fmt.Printf("✓ Length requirement satisfied (< 18 chars)\n")
		} else {
			fmt.Printf("✗ Length requirement not satisfied (>= 18 chars)\n")
		}
	}

	// 详细分析一个例子
	fmt.Println("\n\nDetailed Analysis:")
	fmt.Println("==================")
	example := "20250813113048SP11000002"
	fmt.Printf("Input: %s\n", example)

	segments := packer.parseSegments(example)
	fmt.Println("Parsed Segments:")
	for i, seg := range segments {
		fmt.Printf("  Segment %d: Type=%s, Value=%s\n", i+1, seg.Type, seg.Value)
		if seg.Type == "n" {
			encoded := packer.encodeNumber(seg.Value)
			fmt.Printf("    Base62 Encoded: %s (length: %d)\n", encoded, len(encoded))
		} else if seg.Type == "m" {
			fmt.Printf("    Mapped from original string\n")
		}
	}

	compressed := packer.Compress(example)
	fmt.Printf("Final Compressed: %s (length: %d)\n", compressed, len(compressed))

	// 算法总结
	fmt.Println("\n\nAlgorithm Summary:")
	fmt.Println("===================")
	fmt.Println("✓ 纯数字订单号压缩效果优秀，可将23位压缩到14-15位")
	fmt.Println("✓ 混合订单号通过映射表优化，显著减少TLV开销")
	fmt.Println("✓ 算法支持完整的压缩和解压缩流程")
	fmt.Println("✓ 输出字符集限制在[a-zA-Z0-9]范围内")
	fmt.Println("✓ 映射表支持常见字符串模式，使用m+映射值格式（无长度字段）")
	fmt.Println("\nMapping Table:")
	for original, mapped := range packer.mappingTable {
		oldFormat := fmt.Sprintf("s%d%s", len(original), original)
		newFormat := fmt.Sprintf("m%s", mapped)
		saved := len(oldFormat) - len(newFormat)
		fmt.Printf("  %s -> %s (format: %s -> %s, saves %d chars)\n", original, mapped, oldFormat, newFormat, saved)
	}
	fmt.Println("\nOptimization Suggestions:")
	fmt.Println("- 对于主要为数字的订单号，建议使用此算法")
	fmt.Println("- 映射表可根据实际业务场景中的常见模式进行扩展")
	fmt.Println("- 可根据实际订单号分布调整算法参数")
}
