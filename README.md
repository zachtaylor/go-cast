# ztaylor.me/cast

Provides some super-easy functions for casting values in go

Now I will never write these functions again

## Changelog

- v0.0.6 @2019-12-21
  - import `strings.Contains`, `strings.LastIndex`, `strings.Trim`
  - add `InCharset`, `OutCharset`, package `charset`
  - separate filename by imports `fmt`, `strconv`
  - import `reflect.Kind`, `reflect.Value`, `reflect.Slice`
  - import `time.Duration`, `time.Time`, `time.After`, `time.Now`, `time.Unix`, `time.Millisecond`, `time.Second`, `time.Minute`, `time.Hour`
  - import `bytes.Buffer`, add `NewBuffer` as `bytes.NewBufferString`
  - added futher benchmark data and performance tuning
- v0.0.5 @2019-11-05
  - add `Pool = sync.Pool`
  - add `EncodeJSON(any)`
  - change JSON string encoding, now string escape `\` and `"` with an extra `\`
  - added test data for heavier bench
- v0.0.4 @2019-11-01
  - add type `cast.Array`
  - add func `cast.Arr`, `cast.NewArray`, `cast.SpreadArray`
  - remove json string presize `Grow` guess logic
- v0.0.3 @ 2019-10-03
  - repack `cast.Fields` as `cast.JSON`
    - change getter function names
  - fix `cast.String(map)`
  - add `cast.StringBuilder`
  - add tests and benchmarks
- fields v0.0.2 @ 2019-03-18
  - add methods to `cast.Fields`
- break Stringer v0.0.1 @ 2019-03-12
  - fix `cast.String([]byte)`, `cast.String(error)`, `cast.Bool(string)`, `cast.Int(bool)`
    - add tests
  - break `cast.Stringer`
    - remove type alias `fmt.Stringer`
    - now returns `fmt.Stringer`
  - add `[]byte` casts: `cast.BytesS`
    - add benchmark
  - add `int` casts: `cast.IntS`
  - add `string` casts: `cast.StringBytes`, `cast.StringI`, `cast.StringUI`, `cast.StringI64`, `cast.StringF`, `cast.IntS`
    - add benchmark
  - add `fmt.Stringer` casts: `cast.StringerB`, `cast.StringerE`, `cast.StringerFunc`, `cast.StringerI`, `cast.StringerS`
- init v0.0.0 @ 2018-09-13
  - add `cast.String`, `cast.Int`, `cast.Bool`

## Benchmarks

### v0.0.6

```
goos: windows
goarch: amd64
pkg: ztaylor.me/cast
BenchmarkBuiltinStringBytes-4                   218590940                5.42 ns/op            0 B/op          0 allocs/op
BenchmarkCastStringBytes-4                      1000000000               0.276 ns/op           0 B/op          0 allocs/op
BenchmarkBuiltinBytesS-4                        196084774                6.09 ns/op            0 B/op          0 allocs/op
BenchmarkCastBytesS-4                           1000000000               0.274 ns/op           0 B/op          0 allocs/op

BenchmarkSTDFmtSprintfBuiltinMap-4                160014              6993 ns/op            1736 B/op         36 allocs/op
BenchmarkSTDJsonMarshalBuiltinMap-4               184621              6169 ns/op            2144 B/op         37 allocs/op
BenchmarkCastDictEncodeJSON-4                     521524              2314 ns/op             868 B/op         24 allocs/op
BenchmarkCastStringBuiltinMap-4                   666696              1693 ns/op             798 B/op         11 allocs/op
BenchmarkCastJSONString-4                         631615              1689 ns/op             798 B/op         11 allocs/op
BenchmarkCastDictString-4                        1440621               835 ns/op             432 B/op          3 allocs/op

BenchmarkSTDFmtSprintBuiltinSlice-4               923119              1136 ns/op             352 B/op         10 allocs/op
BenchmarkSTDJsonMarshalBuiltinSlice-4            1465266               821 ns/op             240 B/op          5 allocs/op
BenchmarkCastStringBuiltinSlice-4                1666755               715 ns/op             216 B/op          6 allocs/op
BenchmarkCastStringCastArray-4                   1621671               719 ns/op             216 B/op          6 allocs/op
BenchmarkCastStringNBuiltinSlice-4               1793746               666 ns/op             184 B/op          5 allocs/op
```

### v0.0.5

```
goos: windows
goarch: amd64
pkg: ztaylor.me/cast

BenchmarkBuiltinStringBytes-4           225151906                5.37 ns/op            0 B/op          0 allocs/op
BenchmarkCastStringBytes-4              1000000000               0.266 ns/op           0 B/op          0 allocs/op

BenchmarkBuiltinBytesS-4                198681680                6.02 ns/op            0 B/op          0 allocs/op
BenchmarkCastBytesS-4                   1000000000               0.263 ns/op           0 B/op          0 allocs/op

BenchmarkLoadedMapSprintf-4               173976              6570 ns/op            1768 B/op         37 allocs/op
BenchmarkLoadedMapJsonMarshal-4           203464              5922 ns/op            2144 B/op         37 allocs/op
BenchmarkLoadedJSONString-4               600020              2023 ns/op            1406 B/op         19 allocs/op
BenchmarkLoadedDictString-4               922408              1325 ns/op            1024 B/op         13 allocs/op

BenchmarkStringSprint-4                 10909456               110 ns/op              64 B/op          2 allocs/op
BenchmarkStringString-4                 40006533                30.3 ns/op            16 B/op          1 allocs/op
```

```
goos: linux
goarch: amd64
pkg: ztaylor.me/cast

BenchmarkBuiltinStringBytes   	39096788	        25.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkCastStringBytes      	308717432	         3.87 ns/op	       0 B/op	       0 allocs/op

BenchmarkBuiltinBytesS        	29764130	        36.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkCastBytesS           	303753740	         3.90 ns/op	       0 B/op	       0 allocs/op

BenchmarkLoadedMapSprintf     	   40836	     28837 ns/op	    1768 B/op	      37 allocs/op
BenchmarkLoadedMapJsonMarshal 	   44342	     25758 ns/op	    2144 B/op	      37 allocs/op
BenchmarkLoadedJSONString     	  130828	      8370 ns/op	    1426 B/op	      19 allocs/op
BenchmarkLoadedDictString     	  185958	      5523 ns/op	    1019 B/op	      12 allocs/op

BenchmarkStringSprint         	 2340438	       497 ns/op	      64 B/op	       2 allocs/op
BenchmarkStringString         	 7812093	       148 ns/op	      16 B/op	       1 allocs/op
```

### v0.0.4

```
goos: windows
goarch: amd64
pkg: ztaylor.me/cast

BenchmarkBuiltinStringBytes-4           225261331                5.32 ns/op            0 B/op          0 allocs/op
BenchmarkCastStringBytes-4              1000000000               0.267 ns/op           0 B/op          0 allocs/op
BenchmarkBuiltinBytesS-4                191671183                6.22 ns/op            0 B/op          0 allocs/op
BenchmarkCastBytesS-4                   1000000000               0.267 ns/op           0 B/op          0 allocs/op
BenchmarkLoadedMapFmtSprintf-4            175710              6600 ns/op            1736 B/op         36 allocs/op
BenchmarkLoadedCastJSONString-4           946446              1227 ns/op            1142 B/op         11 allocs/op
BenchmarkStringFmtSprint-4              11029989               113 ns/op              64 B/op          2 allocs/op
BenchmarkStringCastString-4             39680570                31.0 ns/op            16 B/op          1 allocs/op
```

```
goos: linux
goarch: amd64
pkg: ztaylor.me/cast

BenchmarkBuiltinStringBytes   	46444358	        25.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkCastStringBytes      	296433331	         4.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltinBytesS        	36157374	        32.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkCastBytesS           	304336578	         3.75 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoadedMapFmtSprintf  	   40173	     28539 ns/op	    1736 B/op	      36 allocs/op
BenchmarkLoadedCastJSONString 	  240616	      4965 ns/op	    1141 B/op	      11 allocs/op
BenchmarkStringFmtSprint      	 2410371	       491 ns/op	      64 B/op	       2 allocs/op
BenchmarkStringCastString     	 7351100	       147 ns/op	      16 B/op	       1 allocs/op
```

### v0.0.3

```
goos: windows
goarch: amd64
pkg: ztaylor.me/cast

BenchmarkDefaultStringBytes-4           222851612                5.26 ns/op            0 B/op          0 allocs/op
BenchmarkCastStringBytes-4              1000000000               0.284 ns/op           0 B/op          0 allocs/op
BenchmarkDefaultBytesS-4                194495706                6.27 ns/op            0 B/op          0 allocs/op
BenchmarkCastBytesS-4                   1000000000               0.277 ns/op           0 B/op          0 allocs/op
BenchmarkLoadedMapString-4                222242              4859 ns/op            1240 B/op         30 allocs/op
BenchmarkLoadedJSONString-4              1252633               970 ns/op             336 B/op          6 allocs/op
BenchmarkNilJSONFmtSprint-4              6091582               196 ns/op              16 B/op          2 allocs/op
BenchmarkNilJSONCastString-4            24996874                45.5 ns/op             8 B/op          1 allocs/op
BenchmarkStringFmtSprint-4              11112447               105 ns/op              24 B/op          2 allocs/op
BenchmarkStringCastSprint-4             37501171                30.4 ns/op            16 B/op          1 allocs/op
BenchmarkNilMapFmtSprint-4              18750409                60.9 ns/op             5 B/op          1 allocs/op
BenchmarkNilMapCastString-4             17914858                64.7 ns/op             5 B/op          1 allocs/op
```

```
goos: linux
goarch: amd64
pkg: ztaylor.me/cast

BenchmarkDefaultStringBytes 	48460549	        25.0 ns/op
BenchmarkCastStringBytes    	298395208	         3.88 ns/op
BenchmarkDefaultBytesS      	37349222	        32.2 ns/op
BenchmarkCastBytesS         	296883522	         3.83 ns/op
BenchmarkLoadedMapString    	   50401	     21670 ns/op
BenchmarkLoadedJSONString   	  277009	      4331 ns/op
BenchmarkNilJSONFmtSprint   	 1425506	       825 ns/op
BenchmarkNilJSONCastString  	 4715850	       232 ns/op
BenchmarkStringFmtSprint    	 2506722	       475 ns/op
BenchmarkStringCastSprint   	 7341931	       149 ns/op
BenchmarkNilMapFmtSprint    	 4287756	       263 ns/op
BenchmarkNilMapCastString   	 4106491	       278 ns/op
```

### v0.0.2

```
goos: windows
goarch: amd64
pkg: ztaylor.me/cast

BenchmarkStringBytes-4                  1000000000               0.265 ns/op
BenchmarkDefaultStringBytes-4           224837287                5.32 ns/op
BenchmarkBytesS-4                       1000000000               0.264 ns/op
BenchmarkDefaultBytesS-4                196981875                5.97 ns/op
```

```
goos: linux
goarch: amd64
pkg: ztaylor.me/cast

BenchmarkStringBytes        	305869677	         3.85 ns/op
BenchmarkDefaultStringBytes 	46919974	        25.0 ns/op
BenchmarkBytesS             	298804626	         3.82 ns/op
BenchmarkDefaultBytesS      	36722206	        32.1 ns/op
```
