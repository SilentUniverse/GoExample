package main

import (
    "fmt"
    "time"
)

func main() {
    p := fmt.Println

    now := time.Now()
    p(now)

    then := time.Date(
        2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)

    p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())

    p(then.Weekday())

    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))

    diff := now.Sub(then)
    p(diff)

    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())

    p(then.Add(diff))
    p(then.Add(-diff))


	fmt.Println(now.Unix())
    fmt.Println(now.UnixMilli())
    fmt.Println(now.UnixNano())

    fmt.Println(time.Unix(now.Unix(), 0))
    fmt.Println(time.Unix(0, now.UnixNano()))
}

/*2025-01-12 00:39:49.85957422 +0800 CST m=+0.000034004
2009-11-17 20:34:58.651387237 +0000 UTC
2009
November
17
20
34
58
651387237
UTC
Tuesday
true
false
false
132812h4m51.208186983s
132812.08089116306
7.968724853469783e+06
4.78123491208187e+08
478123491208186983
2025-01-11 16:39:49.85957422 +0000 UTC
1994-09-24 00:30:07.443200254 +0000 UTC

获取unix时间
1736613662
1736613662410
1736613662410378518
2025-01-12 00:41:02 +0800 CST
2025-01-12 00:41:02.410378518 +0800 CST
*/