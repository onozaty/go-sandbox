`go test ./...` で複数パッケージをテストすると、各パッケージが並列に実行される。

```
go test ./... -v -count=1
=== RUN   Test1
test20220522/one.Test1 start: 2022-05-22 12:02:39.5633225 +0900 JST m=+0.002258101
test20220522/one.Test1 end  : 2022-05-22 12:02:40.5970111 +0900 JST m=+1.035946701
--- PASS: Test1 (1.03s)
=== RUN   Test2
test20220522/one.Test2 start: 2022-05-22 12:02:40.5970111 +0900 JST m=+1.035946701
test20220522/one.Test2 end  : 2022-05-22 12:02:41.6221348 +0900 JST m=+2.061070401
--- PASS: Test2 (1.03s)
PASS
ok      test20220522/one        2.273s
=== RUN   Test1
test20220522/two.Test1 start: 2022-05-22 12:02:39.5909363 +0900 JST m=+0.002622301
test20220522/two.Test1 end  : 2022-05-22 12:02:40.632877 +0900 JST m=+1.044563001
--- PASS: Test1 (1.04s)
=== RUN   Test2
test20220522/two.Test2 start: 2022-05-22 12:02:40.632877 +0900 JST m=+1.044563001
test20220522/two.Test2 end  : 2022-05-22 12:02:41.6432918 +0900 JST m=+2.054977801
--- PASS: Test2 (1.01s)
PASS
ok      test20220522/two        2.254s
```

`go help testflag` で`parallel`について下記のように記載がある。

```
        -parallel n
            Allow parallel execution of test functions that call t.Parallel.
            The value of this flag is the maximum number of tests to run
            simultaneously; by default, it is set to the value of GOMAXPROCS.
            Note that -parallel only applies within a single test binary.
            The 'go test' command may run tests for different packages
            in parallel as well, according to the setting of the -p flag
            (see 'go help build').
```

`-parallel=1`としても並列だった。

```
go test ./... -v -count=1 -parallel=1
=== RUN   Test1
test20220522/one.Test1 start: 2022-05-22 12:22:01.7306837 +0900 JST m=+0.002627601
test20220522/one.Test1 end  : 2022-05-22 12:22:02.7675848 +0900 JST m=+1.039528701
--- PASS: Test1 (1.04s)
=== RUN   Test2
test20220522/one.Test2 start: 2022-05-22 12:22:02.7675848 +0900 JST m=+1.039528701
test20220522/one.Test2 end  : 2022-05-22 12:22:03.7747177 +0900 JST m=+2.046661601
--- PASS: Test2 (1.01s)
PASS
ok      test20220522/one        2.264s
=== RUN   Test1
test20220522/two.Test1 start: 2022-05-22 12:22:01.7457338 +0900 JST m=+0.002596601
test20220522/two.Test1 end  : 2022-05-22 12:22:02.7828459 +0900 JST m=+1.039708701
--- PASS: Test1 (1.04s)
=== RUN   Test2
test20220522/two.Test2 start: 2022-05-22 12:22:02.7829719 +0900 JST m=+1.039834701
test20220522/two.Test2 end  : 2022-05-22 12:22:03.790036 +0900 JST m=+2.046898801
--- PASS: Test2 (1.01s)
PASS
ok      test20220522/two        2.260s
```

`go help build` を読むと、`-p`で指定することが書いてあった。

```
        -p n
                the number of programs, such as build commands or
                test binaries, that can be run in parallel.
                The default is GOMAXPROCS, normally the number of CPUs available.
```

`-p=1`とすることで回避できた。


```
go test ./... -v -count=1 -p=1
=== RUN   Test1
test20220522/one.Test1 start: 2022-05-22 12:26:35.3316419 +0900 JST m=+0.003329101
test20220522/one.Test1 end  : 2022-05-22 12:26:36.3682487 +0900 JST m=+1.039935901
--- PASS: Test1 (1.04s)
=== RUN   Test2
test20220522/one.Test2 start: 2022-05-22 12:26:36.3682487 +0900 JST m=+1.039935901
test20220522/one.Test2 end  : 2022-05-22 12:26:37.3818506 +0900 JST m=+2.053537801
--- PASS: Test2 (1.01s)
PASS
ok      test20220522/one        2.257s
=== RUN   Test1
test20220522/two.Test1 start: 2022-05-22 12:26:37.9276788 +0900 JST m=+0.002951901
test20220522/two.Test1 end  : 2022-05-22 12:26:38.9719354 +0900 JST m=+1.047208501
--- PASS: Test1 (1.04s)
=== RUN   Test2
test20220522/two.Test2 start: 2022-05-22 12:26:38.9726733 +0900 JST m=+1.047946401
test20220522/two.Test2 end  : 2022-05-22 12:26:39.9831677 +0900 JST m=+2.058440801
--- PASS: Test2 (1.01s)
PASS
ok      test20220522/two        2.259s
```
