package chinaid

import (
	"fmt"
	"testing"
	"time"
)

func TestIDCardDetail_GetAge(t1 *testing.T) {
	type fields struct {
		Sex      int
		Birthday time.Time
		Addr     Addr
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "",
			fields: fields{},
			want:   41,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t,err := IDCard("41040319810315551x").Decode()
			if err != nil {
				fmt.Println(err)
				return
			}
			if got := t.GetAge(); got != tt.want {
				t1.Errorf("GetAge() = %v, want %v", got, tt.want)
			}
		})
	}
}