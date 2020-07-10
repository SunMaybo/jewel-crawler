package common

import (
	"github.com/SunMaybo/jewel-crawler/common/sign"
	"sync"
	"testing"
)

var SafeMap sync.Map = sync.Map{}

func BenchmarkGenerateRandomID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				u := GenerateRandomID()
				_, load := SafeMap.LoadOrStore(u, "")
				if load {
					b.Fatal(u)
				}
				b.Log(u)
			}
		})

	}
}

type Student struct {
	Name     string    `sign:"name"`
	Age      int       `sign:"age"`
	Teachers []Teacher `sign:"teachers"`
}
type Teacher struct {
	Name string `sign:"name"`
	Age  int    `sign:"age"`
}

func TestSignature(t *testing.T) {
	stu := Student{
		Name:     "张三",
		Age:      34,
		Teachers: []Teacher{{Name: "wang", Age: 45}, {Name: "tian", Age: 65}},
	}
	t.Log(sign.Encode(stu))
	t.Log(Signature(stu))
	t.Log(SignatureMap(map[string]string{"name": "test", "age": "45"}))
}
