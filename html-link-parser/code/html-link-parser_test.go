package htmllinkparser

import (
	"reflect"
	"strings"
	"testing"
)

type TestSuite struct {
	label  string
	input  string
	output []Link
}

func TestHtmlLinkParser(t *testing.T) {
	testSuites := []TestSuite{
		{
			label: "Single link",
			input: `<a href="/dog">dog</a>`,
			output: []Link{
				{
					href: "/dog",
					text: "dog",
				},
			},
		},
		{
			label: "Multiple links",
			input: `<a href="/dog">dog</a><a href="/cat">cat</a>`,
			output: []Link{
				{
					href: "/dog",
					text: "dog",
				},
				{
					href: "/cat",
					text: "cat",
				},
			},
		},
		{
			label: "Nested links",
			input: `<a href="/dog"><span>dog</span></a>`,
			output: []Link{
				{
					href: "/dog",
					text: "dog",
				},
			},
		},
		{
			label:  "No links",
			input:  `<p>no links</p>`,
			output: []Link{},
		},
		{
			label: "Complex HTML",
			input: `
				<a href="/dog"><span>dog</span></a>
				<a href="/cat">cat</a>
				<a href="/rabbit">rabbit</a>
				<a href="/fish">fish</a>
			`,
			output: []Link{
				{
					href: "/dog",
					text: "dog",
				},
				{
					href: "/cat",
					text: "cat",
				},
				{
					href: "/rabbit",
					text: "rabbit",
				},
				{
					href: "/fish",
					text: "fish",
				},
			},
		},
	}
	for _, testSuite := range testSuites {
		t.Run(testSuite.label, func(t *testing.T) {
			got := ParseHTML(strings.NewReader(testSuite.input))
			if reflect.DeepEqual(got, testSuite.output) {
				t.Errorf("Got %v, expected %v", got, testSuite.output)
			}
		})
	}
}
