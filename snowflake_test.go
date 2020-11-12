package snowflake_test

import (
	"testing"

	"github.com/Yu-33/snowflake"
)

func TestGenerateId(t *testing.T) {
	var err error
	var id int64

	idWorker, err := snowflake.New(0)
	if err != nil {
		t.Errorf("New snowflake fail: %v", err)
		return
	}

	number := 2 << 16
	ids := make([]int64, number)
	for i := 0; i < number; i++ {
		id, err = idWorker.NextID()
		if err != nil {
			t.Errorf("Generate id fail: %v", err)
			return
		}
		if id == 0 {
			t.Errorf("invalid id")
			return
		}
		ids[i] = id
	}

	for i := 0; i < number-1; i++ {
		if ids[i] >= ids[i+1] {
			t.Errorf("invalid id")
			return
		}
	}
}

func BenchmarkSnowFlake_NextId(b *testing.B) {
	idWorker, err := snowflake.New(0)
	if err != nil {
		b.Errorf("New snowflake fail: %v", err)
		return
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			id, err := idWorker.NextID()
			_ = id
			_ = err
		}
	})
}
