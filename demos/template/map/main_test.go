package main

import "testing"

func BenchmarkExecute(b *testing.B) {
	Init()
	for i := 0; i < b.N; i++ {
		Execute()
	}
}

func TestExecuteData(t *testing.T) {
	Init()
	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "t1",
			args: args{
				data: map[string]interface{}{
					"person": Person{"Zhangsan", 11},
					"m": map[string]interface{}{
						"name": "zhangsan",
						"age":  11,
					},
				},
			},
			want: `My dear master!
your age: 11 multiply twice age: 22
map age:11
Say: Hello, my name is Zhangsan`,
		},
		{
			name: "t2",
			args: args{
				data: map[string]interface{}{
					"person": Person{"Wangwu", 11},
					"m": map[string]interface{}{
						"name": "wangwu",
					},
				},
			},
			want: `Hello Wangwu
your age: 11 multiply twice age: 22
map age:0
Say: Hello, my name is Wangwu`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExecuteData(tt.args.data); got != tt.want {
				t.Errorf("ExecuteData() = %v, want %v", got, tt.want)
			}
		})
	}
}
