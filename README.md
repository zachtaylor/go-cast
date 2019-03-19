# ztaylor.me/cast

Provides some super-easy functions for casting values in go

Now I will never write these functions again

## Changelog

- fields v0.0.2 @ 2019-03-18
  - add methods to `Fields`
- break Stringer v0.0.1 @ 2019-03-12
  - fix `cast.String([]byte)`, `cast.String(error)`, `cast.Bool(string)`, `cast.Int(bool)`
    - add tests
  - break `cast.Stringer`
    - remove type alias `fmt.Stringer`
    - new functionality returns `fmt.Stringer`
  - add `[]byte` casts: `cast.BytesS`
    - add benchmark
  - add `int` casts: `cast.IntS`
  - add `string` casts: `cast.StringBytes`, `cast.StringI`, `cast.StringUI`, `cast.StringI64`, `cast.StringF`, `cast.IntS`
    - add benchmark
  - add `fmt.Stringer` casts: `cast.StringerB`, `cast.StringerE`, `cast.StringerFunc`, `cast.StringerI`, `cast.StringerS`
- init v0.0.0 @ 2018-09-13
  - add `cast.String`, `cast.Int`, `cast.Bool`
