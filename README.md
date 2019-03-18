# version

a small package for version compare

usage:

```go
	version.Set("1.2.8")

	t.Logf("1.2.8 == 1.2.8: %t",  version.CheckVersionOk("1.2.8")) // 1.2.8 == 1.2.8 is true
	t.Logf("1.2.8 < 1.2.9: %t",  version.CheckVersionOk("<1.2.9")) // 1.2.8 < 1.2.9 is true
	t.Logf("1.2.8 <= 1.2.8: %t",  version.CheckVersionOk("<= 1.2.8")) // 1.2.8 <= 1.2.8 is true
	t.Logf("1.2.8 < 1.2.8: %t",  version.CheckVersionOk("< 1.2.8")) // 1.2.8 < 1.2.8 is false
	t.Logf("1.2.8 > 1.2.7: %t",  version.CheckVersionOk(">1.2.7")) // 1.2.8 < 1.2.8 is true
	t.Logf("1.2.8 is in 1.2.7 ~ 1.2.9: %t", version.CheckVersionOk("1.2.7 ~ 1.2.9")) // 1.2.8 is in  1.2.7 ~ 1.2.9 is true
	t.Logf("1.2.8 is in 1.2.7 ~ 1.2.8: %t", version.CheckVersionOk("1.2.7 ~ 1.2.8")) // 1.2.8 is in  1.2.7 ~ 1.2.8 is true
	t.Logf("1.2.8 is in 1.2.6 ~ 1.2.7: %t", version.CheckVersionOk("1.2.6 ~ 1.2.7")) // 1.2.8 is in  1.2.6 ~ 1.2.7 is false
```


