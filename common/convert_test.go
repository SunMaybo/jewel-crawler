package common

import (
	"log"
	"testing"
)

func TestConvertAssign(t *testing.T) {
	var dst1 float64
	var dst2 int

	src := "123"

	if err := ConvertAssign(src,&dst1); err != nil {
		log.Fatalf("convert failed, %v", err)
	}else {
		log.Printf("convert ok: %f", dst1)
	}

	if err := ConvertAssign(&dst2, src); err != nil {
		log.Fatalf("convert failed, %v", err)
	}else {
		log.Printf("convert ok: %d", dst2)
	}
}
