package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

type XMLProduct struct {
	XMLName          xml.Name `xml:"row"`
	ProductId        string   `xml:"product_id"`
	ProductName      string   `xml:"![CDATA["`
	OriginalPrice    string   `xml:"original_price"`
	BargainPrice     string   `xml:"bargain_price"`
	TotalReviewCount int      `xml:"total_review_count"`
	AverageScore     float64  `xml:"average_score"`
}

func main() {
	prod := XMLProduct{
		ProductId:        "ProductId",
		ProductName:      "ProductName",
		OriginalPrice:    "OriginalPrice",
		BargainPrice:     "BargainPrice",
		TotalReviewCount: 20,
		AverageScore:     2.1,
	}

	out, err := xml.MarshalIndent(prod, " ", "  ")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(string(out))

	out = bytes.Replace(out, []byte("<![CDATA[>"), []byte("<![CDATA["), -1)
	out = bytes.Replace(out, []byte("</![CDATA[>"), []byte("]]>"), -1)
	fmt.Println(string(out))
}
