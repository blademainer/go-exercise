package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/moov-io/iso8583"
	"github.com/moov-io/iso8583/specs"
)

func main() {
	// 定义ISO 8583的报文规范
	// spec := &iso8583.MessageSpec{
	// 	Fields: map[int]field.Field{
	// 		0: field.NewString(
	// 			&field.Spec{
	// 				Length:      4,
	// 				Description: "Message Type Indicator",
	// 				Enc:         encoding.ASCII,
	// 				Pref:        prefix.ASCII.Fixed,
	// 			},
	// 		),
	// 		1: field.NewBitmap(
	// 			&field.Spec{
	// 				Description: "Bitmap",
	// 				Enc:         encoding.BytesToASCIIHex,
	// 				Pref:        prefix.Hex.Fixed,
	// 			},
	// 		),
	// 		2: field.NewString(
	// 			&field.Spec{
	// 				Length:      19,
	// 				Description: "Primary Account Number (PAN)",
	// 				Enc:         encoding.ASCII,
	// 				Pref:        prefix.ASCII.LL,
	// 			},
	// 		),
	// 		3: field.NewComposite(
	// 			&field.Spec{
	// 				Length:      6,
	// 				Description: "Processing Code",
	// 				Pref:        prefix.ASCII.Fixed,
	// 				Tag: &field.TagSpec{
	// 					Sort: sort.StringsByInt,
	// 				},
	// 				Subfields: map[string]field.Field{
	// 					"01": field.NewString(
	// 						&field.Spec{
	// 							Length:      2,
	// 							Description: "Transaction Type",
	// 							Enc:         encoding.ASCII,
	// 							Pref:        prefix.ASCII.Fixed,
	// 						},
	// 					),
	// 					"02": field.NewString(
	// 						&field.Spec{
	// 							Length:      2,
	// 							Description: "From Account",
	// 							Enc:         encoding.ASCII,
	// 							Pref:        prefix.ASCII.Fixed,
	// 						},
	// 					),
	// 					"03": field.NewString(
	// 						&field.Spec{
	// 							Length:      2,
	// 							Description: "To Account",
	// 							Enc:         encoding.ASCII,
	// 							Pref:        prefix.ASCII.Fixed,
	// 						},
	// 					),
	// 				},
	// 			},
	// 		),
	// 		4: field.NewNumeric(
	// 			&field.Spec{
	// 				Length:      12,
	// 				Description: "Transaction Amount",
	// 				Enc:         encoding.ASCII,
	// 				Pref:        prefix.ASCII.Fixed,
	// 				Pad:         padding.Left('0'),
	// 			},
	// 		),
	// 		7: field.NewString(
	// 			&field.Spec{
	// 				Length:      10,
	// 				Description: "Transmission Date & Time",
	// 				Enc:         encoding.ASCII,
	// 				Pref:        prefix.ASCII.Fixed,
	// 				Subfields: map[string]field.Field{
	// 					"01": field.NewString(
	// 						&field.Spec{
	// 							Length:      4,
	// 							Description: "Transmission Date, This subfield must contain a valid date in MMDD format.",
	// 							Enc:         encoding.ASCII,
	// 							Pref:        prefix.ASCII.Fixed,
	// 						},
	// 					),
	// 					"02": field.NewString(
	// 						&field.Spec{
	// 							Length:      6,
	// 							Description: "Transmission Time, Time must contain a valid time in hhmmss format.",
	// 							Enc:         encoding.ASCII,
	// 							Pref:        prefix.ASCII.Fixed,
	// 						},
	// 					),
	// 				},
	// 			},
	// 		),
	// 		11: field.NewNumeric(
	// 			&field.Spec{
	// 				Length:      6,
	// 				Description: "System Trace Audit Number (STAN)",
	// 				Enc:         encoding.ASCII,
	// 				Pref:        prefix.ASCII.Fixed,
	// 				Pad:         padding.Left('0'),
	// 			},
	// 		),
	// 		18: field.NewNumeric(
	// 			&field.Spec{
	// 				Length:      4,
	// 				Description: "Merchant Type",
	// 				Enc:         encoding.ASCII,
	// 				Pref:        prefix.ASCII.Fixed,
	// 			},
	// 		),
	// 		41: field.NewString(
	// 			&field.Spec{
	// 				Length:      8,
	// 				Description: "Card Acceptor Terminal ID",
	// 				Enc:         encoding.ASCII,
	// 				Pref:        prefix.ASCII.Fixed,
	// 			},
	// 		),
	// 		49: field.NewNumeric(
	// 			&field.Spec{
	// 				Length:      3,
	// 				Description: "Transaction Currency Code",
	// 				Enc:         encoding.ASCII,
	// 				Pref:        prefix.ASCII.Fixed,
	// 			},
	// 		),
	// 	},
	// }

	// 构建ISO 8583 报文
	message := iso8583.NewMessage(specs.Spec87ASCII)

	// 设置字段值
	message.MTI("0100") // MTI 表示授权请求
	// message.Field(1, "F23E8000000000000000000000000000")
	err := message.Field(2, "5432123456789012")
	if err != nil {
		panic(err)
	} // PAN
	message.Field(3, "886688")     // Processing Code
	message.Field(4, "100000")     // 交易金额
	message.Field(7, "0912123456") // 传输时间
	message.Field(11, "123456")    // 系统跟踪号
	message.Field(14, "1234")      // 卡有效期
	message.Field(18, "5812")      // 商户类型
	message.Field(22, "012")       // POS输入方式码
	message.Field(23, "733")       // 卡序列号
	message.Field(32, "123456")    // 受理方标识码
	message.Field(48, "123456")
	message.Field(49, "156")    // 交易货币代码
	message.Field(61, "123456") // Point of Service Data Code
	// message.Field(41, "TERMID01")        // 终端ID
	// message.Field(49, "840")             // 货币代码

	// 打包ISO 8583报文为二进制格式
	packedMessage, err := message.Pack()
	if err != nil {
		log.Fatalf("Failed to pack ISO 8583 message: %v", err)
	}

	fmt.Println("Packed ISO 8583 message: ", string(packedMessage))
	// 打印打包后的报文
	fmt.Printf("Packed ISO 8583 message(hexed): %x\n", packedMessage)

	// 解包ISO 8583报文
	unpackedMessage := iso8583.NewMessage(specs.Spec87ASCII)
	err = unpackedMessage.Unpack(packedMessage)
	if err != nil {
		panic(err.Error())
		log.Fatalf("Failed to unpack ISO 8583 message: %v", err)
	}

	out := bytes.NewBuffer([]byte{})
	err = iso8583.Describe(unpackedMessage, out)
	if err != nil {
		panic(err.Error())
		return
	}
	fmt.Println(out.String())

	// 打印解包后的字段
	// mti, err := unpackedMessage.GetMTI()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Unpacked ISO 8583 message MTI: %s\n", mti)
	// fmt.Printf("Unpacked Field 2 (PAN): %s\n", unpackedMessage.GetField(2))
	// fmt.Printf("Unpacked Field 3 (Processing Code): %s\n", unpackedMessage.GetField(3))
	// fmt.Printf("Unpacked Field 4 (Transaction Amount): %s\n", unpackedMessage.GetField(4))
	// fmt.Printf("Unpacked Field 7 (Transmission Date & Time): %s\n", unpackedMessage.GetField(7))
	// fmt.Printf("Unpacked Field 11 (STAN): %s\n", unpackedMessage.GetField(11))
	// fmt.Printf("Unpacked Field 41 (Card Acceptor Terminal ID): %s\n", unpackedMessage.GetField(41))
	// fmt.Printf("Unpacked Field 49 (Transaction Currency Code): %s\n", unpackedMessage.GetField(49))
}
