# ztaylor.me/cast

A swiss-pocket samurai-bayonet type-casting multi-tool for go

## What

A centralized API for type casting in go, replacing the standard library where I have written faster code

Faster encoding than the standard library

Types useful for construction in larger programs

Aliases into the standard library where I haven't written anything faster

## Changelog

- v0.0.10 @2020-01-25
  - import `reflect.Map`
  - change `Reflect` to be an alias of `reflect.ValueOf`
- v0.0.9 @2020-01-23
  - import `time.Tick`
  - import `sort.Slice`
  - import `strings.Split`
  - import `reflect.Type`, `reflect.TypeOf`
  - add `cast.KindOf`
- v0.0.8 @2020-01-11
  - import `time.Timer`, `time.NewTimer`
- v0.0.7 @2019-12-27
  - fix auto-initialize `Dict`
  - import `fmt.Fprintf`
  - import `io.Copy`, `io.EOF`, `io.Reader`, `io.Writer`, `io.Closer`, `io.ReadWriter`, `io.ReadCloser`, `io.ReadWriteCloser`, `io.WriteCloser`
  - import `time.Ticker`, `time.NewTicker`
  - add `Write`, `WriteString`, `WriteN`
  - add support for errors
  - add `json.GetKeys()`
  - import `sort.Strings`, `sort.Ints`, `sort.StringSlice`, `sort.IntSlice`
  - added benchmark data taken from `elemen7s.com`@`v0.0.1`
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

### JSON encoding

Data suggests that `cast/JSON.String()` can be up to 200% faster than (runtime 33% of) `encoding/json.Marshal()` with a `map[string]interface{}` of negligable size

This scales to about 30% faster than (runtime 77% of) `encoding/json.Marshal()`  for payload size about 50Kb

This is achieved by using about 40% of the `alloc`s used by the standard library, and significant memory savings as well

### Map encoding

`cast.Dict` encoding is similarly more performant than `map[interface{}]interface{}`

### Details

The numbers are kept in the file [benchmarks.md](benchmarks.md)
