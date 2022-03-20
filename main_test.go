package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type test struct {
	name string
	run  func() interface{}
	res  interface{}
	err  error
}

func TestCases(t *testing.T) {
	t.Parallel()

	tests := []test{
		{
			name: "testValidity",
			run: func() interface{} {
				return testValidity("23-ab-48-caba-56-haha")
			},
			res: true,
			err: nil,
		},
		{
			name: "averageNumber",
			run: func() interface{} {
				return averageNumber("23-ab-48-caba-56-haha")
			},
			res: 42.333333333333336,
			err: nil,
		},
		{
			name: "wholeStory",
			run: func() interface{} {
				return wholeStory("23-ab-48-caba-56-haha")
			},
			res: "ab caba haha",
			err: nil,
		},
		{
			name: "storyStats",
			run: func() interface{} {
				s1, s2, _, _ := storyStats("23-ab-48-caba-56-haha")
				return s1 + s2
			},
			res: "abcaba",
			err: nil,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := tc.run()

			require.EqualValues(t, res, tc.res)
		})
	}
}
