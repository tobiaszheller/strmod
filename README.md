# Strmod

## Warning: Alpha!


```Go
in := "Some text/Other text/Something else "

out := strmod.Modify(in,
	strmod.SplitAndReturnLast("/"),
	strmod.TrimSpace(),
	strmod.ReplaceAll(" ", "_"),
	strmod.ToLower(),
)
fmt.Println(out)
// something_else
```

