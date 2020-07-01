package strmod_test

import (
	"strings"
	"testing"

	"github.com/tobiaszheller/strmod"
)

func TestModify(t *testing.T) {
	modifiers := func(mods ...strmod.Modifier) []strmod.Modifier { return mods }
	testCases := []struct {
		desc      string
		in        string
		out       string
		modifiers []strmod.Modifier
	}{
		{
			desc: "No modifiers",
			in:   "a",
			out:  "a",
		},
		{
			desc: "TrimPrefix & custom suffix",
			in:   "aaaHelloWorld1",
			out:  "HelloWorld",
			modifiers: modifiers(
				// builtin
				strmod.TrimPrefix("aaa"),
				// custom defined
				func(in string) string {
					return strings.TrimSuffix(in, "1")
				},
			),
		},
		{
			desc: "SplitAndReturnLastPart - return last",
			in:   "first/middle/last",
			out:  "last",
			modifiers: modifiers(
				strmod.SplitAndReturnLastPart("/"),
			),
		},
		{
			desc: "SplitAndReturnLastPart - unchanged for no matches",
			in:   "first/middle/last",
			out:  "first/middle/last",
			modifiers: modifiers(
				strmod.SplitAndReturnLastPart("?"),
			),
		},
		{
			desc: "SplitAndReturnLastPart - last char for empty sep",
			in:   "first/middle/last",
			out:  "t",
			modifiers: modifiers(
				strmod.SplitAndReturnLastPart(""),
			),
		},
		{
			desc: "SplitAndReturnLastPart - empty output for in & sep empty",
			in:   "",
			out:  "",
			modifiers: modifiers(
				strmod.SplitAndReturnLastPart(""),
			),
		},
		{
			desc: "SplitAndReturnFirstPart - return first",
			in:   "first/middle/last",
			out:  "first",
			modifiers: modifiers(
				strmod.SplitAndReturnFirstPart("/"),
			),
		},
		{
			desc: "SplitAndReturnFirstPart - unchanged for no matches",
			in:   "first/middle/last",
			out:  "first/middle/last",
			modifiers: modifiers(
				strmod.SplitAndReturnFirstPart("?"),
			),
		},
		{
			desc: "SplitAndReturnFirstPart - last char for empty sep",
			in:   "first/middle/last",
			out:  "f",
			modifiers: modifiers(
				strmod.SplitAndReturnFirstPart(""),
			),
		},
		{
			desc: "SplitAndReturnFirstPart - empty output for in & sep empty",
			in:   "",
			out:  "",
			modifiers: modifiers(
				strmod.SplitAndReturnFirstPart(""),
			),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if out := strmod.Modify(tC.in, tC.modifiers...); out != tC.out {
				t.Errorf("Output mismatch, exp: %q, got: %q", tC.out, out)
			}
		})
	}
}
