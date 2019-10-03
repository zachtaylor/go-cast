# ztaylor.me/cast

Provides some super-easy functions for casting values in go

Now I will never write these functions again

## Changelog

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

```
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
