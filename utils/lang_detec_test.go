package utils

import (
	"fmt"
	"testing"
)

func TestLangDetect(t *testing.T) {
	fmt.Printf("LangDetect(\"中英互译\"): %v\n", LangDetect("中英互译"))
	fmt.Printf("LangDetect(\"中英互译\"): %v\n", LangDetect("中英互译"))
	fmt.Printf("LangDetect(\"中英互译\"): %v\n", LangDetect("中英互译"))
	fmt.Printf("LangDetect(\"中英互译\"): %v\n", LangDetect("中英互译"))
}
