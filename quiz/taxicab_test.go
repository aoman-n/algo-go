package quiz

import (
	"fmt"
	"testing"
)

func Test_taxiCabNumber(t *testing.T) {
	type args struct {
		maxAnswerNum   int
		matchAnswerNum int
		startAnswer    int
	}
	tests := []struct {
		name string
		args args
	}{
		// {
		// 	name: "Case1",
		// 	args: args{
		// 		maxAnswerNum:   1,
		// 		matchAnswerNum: 2,
		// 		startAnswer:   0,
		// 	},
		// },
		// {
		// 	name: "Case2",
		// 	args: args{
		// 		maxAnswerNum:   2,
		// 		matchAnswerNum: 2,
		// 		startAnswer:   0,
		// 	},
		// },
		{
			name: "Case3",
			args: args{
				maxAnswerNum:   1,
				matchAnswerNum: 3,
				startAnswer:    87539319,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := taxiCabNumber(tt.args.maxAnswerNum, tt.args.matchAnswerNum, tt.args.startAnswer)

			fmt.Printf("got length: %v\n", len(got))
			for i, data := range got {
				fmt.Printf("[%v] : %v\n", i, data)
			}
		})
	}
}
