package date

import (
	"testing"
	"fmt"
)

func TestFormat(t *testing.T) {
	fmt.Println(DATE_FORMAT.Format())
	fmt.Println(DATE_TIME_FORMAT.Format())
}
