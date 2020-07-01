# strmod

`strmod` allows to modify string in very readable way.

It is helpful when you need to do multiple operations on single string, like trimming, replacing etc.

Currently it supports most string operations from stdlib `strings` package and few custom made.

```Go
in := "Some text/Other text/Something else "

out := strmod.Modify(in,
	strmod.SplitAndReturnLastPart("/"),
	strmod.TrimSpace(),
	strmod.ReplaceAll(" ", "_"),
	strmod.ToLower(),
)
fmt.Println(out)
// something_else
```

