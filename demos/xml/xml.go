package main

import (
	"encoding/xml"
	"fmt"
)

type XMLProduct struct {
	XMLName     xml.Name `xml:"optin-request"`
	ProductId   string   `xml:"product_id"`
	ProductName Cdata    `xml:"product_name"`
	// ProductName      string  `xml:",cdata"`
	OriginalPrice    string  `xml:"original_price"`
	BargainPrice     string  `xml:"bargain_price"`
	TotalReviewCount int     `xml:"total_review_count"`
	AverageScore     float64 `xml:"average_score"`
	Timeout          struct {
		After string `xml:"after,attr"`
	} `xml:"timeout"`
}

type Cdata struct {
	Value string `xml:",cdata"`
}

func main() {
	prod := XMLProduct{
		ProductId:        "ProductId",
		ProductName:      Cdata{Value: "ProductName"},
		OriginalPrice:    "OriginalPrice",
		BargainPrice:     "BargainPrice",
		TotalReviewCount: 20,
		AverageScore:     2.1,
		Timeout: struct {
			After string `xml:"after,attr"`
		}{
			After: "10",
		},
	}

	out, err := xml.MarshalIndent(prod, " ", "  ")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(string(out))

	// out = bytes.Replace(out, []byte("<![CDATA[>"), []byte("<![CDATA["), -1)
	// out = bytes.Replace(out, []byte("</![CDATA[>"), []byte("]]>"), -1)
	// fmt.Println(string(out))
}
