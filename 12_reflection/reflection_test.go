package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with a non int field",
			struct {
				Name string
				Age  int
			}{"Chris", 42},
			[]string{"Chris"},
		},
		{
			"nested fields",
			Person{
				Name: "Chris",
				Profile: Profile{
					Age:  42,
					City: "London",
				},
			},
			[]string{"Chris", "London"},
		},
		{
			"pointers to things",
			&Person{
				Name: "Chris",
				Profile: Profile{
					Age:  33,
					City: "London",
				},
			},
			[]string{"Chris", "London"},
		},
		{
			"slices",
			[]Profile{
				{33, "London"},
				{42, "San Francisco"},
			},
			[]string{"London", "San Francisco"},
		},
		{
			"slices",
			[2]Profile{
				{33, "London"},
				{42, "San Francisco"},
			},
			[]string{"London", "San Francisco"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v but want %v", got, test.ExpectedCalls)
			}
		})
	}

}
