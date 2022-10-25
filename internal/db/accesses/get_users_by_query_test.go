package accesses

import (
	"reflect"
	"testing"

	"github.com/bots-empire/ams-service/internal/entity"
)

func TestStorage_getMatchIDs(t *testing.T) {
	type testCase struct {
		name   string
		admins []*entity.Access
		check  []string
		result []int64
	}
	testCases := []*testCase{
		{
			name: "no need to update limit",
			admins: []*entity.Access{
				{
					UserID: 1,
					Additional: []string{
						"1",
						"2",
						"3",
					},
				},
				{
					UserID: 2,
					Additional: []string{
						"1",
						"2",
						"4",
					},
				},
				{
					UserID: 3,
					Additional: []string{
						"1",
						"3",
						"5",
					},
				},
			},
			check: []string{
				"1",
				"3",
			},
			result: []int64{
				1,
				3,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resIDs := getMatchIDs(tc.admins, tc.check)

			if !reflect.DeepEqual(resIDs, tc.result) {
				t.Fail()
			}
		})
	}
}
