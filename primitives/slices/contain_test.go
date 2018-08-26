package slices

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestContain(t *testing.T) {

	Convey("TestContainsString", t, func() {
		tests := []struct {
			Description     string
			Container       []string
			Item            string
			ExpectContained bool
		}{
			{
				Description:     "Empty Slice + Empty String",
				Container:       []string{},
				Item:            "",
				ExpectContained: false,
			},
			{
				Description:     "Nil Slice + Empty String",
				Container:       nil,
				Item:            "",
				ExpectContained: false,
			},
			{
				Description:     "Empty Slice + NonEmpty String",
				Container:       []string{},
				Item:            "nonempty",
				ExpectContained: false,
			},
			{
				Description:     "NonEmpty Slice + Not Contained String",
				Container:       []string{"a", "b", "c"},
				Item:            "somestring",
				ExpectContained: false,
			},
			{
				Description:     "NonEmpty Slice + Contained String",
				Container:       []string{"a", "b", "c", "somestring"},
				Item:            "somestring",
				ExpectContained: true,
			},
		}

		for _, t := range tests {
			Convey(t.Description, func() {
				So(ContainsString(t.Container, t.Item), ShouldEqual, t.ExpectContained)
			})
		}

	})

	Convey("TestAddStringIfNotContained", t, func() {
		tests := []struct {
			Description string
			Container   []string
			Items       []string
		}{
			{
				Description: "Empty Slice + Empty String",
				Container:   []string{},
				Items:       []string{""},
			},
			{
				Description: "Nil Slice + Empty String",
				Container:   nil,
				Items:       []string{""},
			},
			{
				Description: "Empty Slice + NonEmpty String",
				Container:   []string{},
				Items:       []string{"nonempty"},
			},
			{
				Description: "NonEmpty Slice + Not Contained String",
				Container:   []string{"a", "b", "c"},
				Items:       []string{"something"},
			},
			{
				Description: "NonEmpty Slice + Contained String",
				Container:   []string{"a", "b", "c", "somestring"},
				Items:       []string{"somestring"},
			},
			{
				Description: "Miscellaneous 1",
				Container:   []string{"a", "b", "c", "d", "e", "f"},
				Items:       []string{"b", "d", "f"},
			},
		}

		for _, t := range tests {
			Convey(t.Description, func() {
				itemsToAdd := 0
				for _, item := range t.Items {
					if !ContainsString(t.Container, item) {
						itemsToAdd++
					}
				}
				after := AddStringIfNotContained(t.Container, t.Items...)
				So(len(after), ShouldEqual, len(t.Container)+itemsToAdd)
			})
		}
	})

}
