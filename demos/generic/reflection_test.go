package generic

import (
	"testing"
)

func TestReflectType(t *testing.T) {
	ReflectType(&Server[int]{})
}
