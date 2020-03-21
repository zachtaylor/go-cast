# Benchmarks

This document compares version benchmarks, including the latest benchmarks

## v0.0.11
```
goos: windows
goarch: amd64
pkg: ztaylor.me/cast
BenchmarkBuiltinStringBytes-4                   222637688                5.22 ns/op            0 B/op          0 allocs/op
BenchmarkCastStringBytes-4                      1000000000               0.285 ns/op           0 B/op          0 allocs/op
BenchmarkBuiltinBytesS-4                        196402045                6.03 ns/op            0 B/op          0 allocs/op
BenchmarkCastBytesS-4                           1000000000               0.274 ns/op           0 B/op          0 allocs/op

BenchmarkSTDFmtSprintBuiltinSlice-4               922111              1101 ns/op             352 B/op         10 allocs/op
BenchmarkSTDJsonMarshalBuiltinSlice-4            1511382               787 ns/op             240 B/op          5 allocs/op
BenchmarkCastStringBuiltinSlice-4                1301565               914 ns/op             216 B/op          6 allocs/op
BenchmarkCastStringCastArray-4                   1305806               920 ns/op             216 B/op          6 allocs/op
BenchmarkCastStringNBuiltinSlice-4               1683080               703 ns/op             216 B/op          5 allocs/op

BenchmarkSTDJsonMarshalBuiltinMap1-4              155829              7161 ns/op            2368 B/op         41 allocs/op
BenchmarkCastDictEncodeJSON1-4                    315804              3328 ns/op             822 B/op         28 allocs/op
BenchmarkCastJSONString1-4                        444295              2514 ns/op             769 B/op         13 allocs/op

BenchmarkSTDJsonMarshalBuiltinMap2-4                7495            156899 ns/op           50221 B/op       1005 allocs/op
BenchmarkCastDictEncodeJSON2-4                      9223            117311 ns/op           32390 B/op        393 allocs/op
BenchmarkCastJSONString2-4                         10000            116796 ns/op           32353 B/op        389 allocs/op

BenchmarkSTDJsonMarshalBuiltinMap3-4                4800            237703 ns/op           77127 B/op       1537 allocs/op
BenchmarkCastDictEncodeJSON3-4                      6315            177983 ns/op           50682 B/op        590 allocs/op
BenchmarkCastJSONString3-4                          6312            176645 ns/op           50650 B/op        586 allocs/op

BenchmarkSTDFmtSprintfBuiltinMap1-4               134821              8515 ns/op            1992 B/op         41 allocs/op
BenchmarkCastStringBuiltinMap1-4                  444306              2541 ns/op             767 B/op         13 allocs/op
BenchmarkCastDictString1-4                        631245              1788 ns/op             496 B/op          5 allocs/op

BenchmarkSTDFmtSprintfBuiltinMap2-4                 6662            165561 ns/op           44605 B/op        992 allocs/op
BenchmarkCastStringBuiltinMap2-4                    9224            117298 ns/op           32353 B/op        389 allocs/op
BenchmarkCastDictString2-4                          9234            116629 ns/op           32353 B/op        389 allocs/op

BenchmarkSTDFmtSprintfBuiltinMap3-4                 4442            253031 ns/op           66167 B/op       1507 allocs/op
BenchmarkCastStringBuiltinMap3-4                    6315            178141 ns/op           50650 B/op        586 allocs/op
BenchmarkCastDictString3-4                          6312            177910 ns/op           50650 B/op        586 allocs/op
```

## v0.0.10
```
goos: windows
goarch: amd64
pkg: ztaylor.me/cast
BenchmarkBuiltinStringBytes-4                   223883353                5.31 ns/op            0 B/op          0 allocs/op
BenchmarkCastStringBytes-4                      1000000000               0.271 ns/op           0 B/op          0 allocs/op
BenchmarkBuiltinBytesS-4                        195759649                6.08 ns/op            0 B/op          0 allocs/op
BenchmarkCastBytesS-4                           1000000000               0.271 ns/op           0 B/op          0 allocs/op

BenchmarkSTDFmtSprintBuiltinSlice-4               923154              1128 ns/op             352 B/op         10 allocs/op
BenchmarkSTDJsonMarshalBuiltinSlice-4            1479709               810 ns/op             240 B/op          5 allocs/op
BenchmarkCastStringBuiltinSlice-4                1240986               954 ns/op             216 B/op          6 allocs/op
BenchmarkCastStringCastArray-4                   1256560               933 ns/op             216 B/op          6 allocs/op
BenchmarkCastStringNBuiltinSlice-4               1641657               715 ns/op             216 B/op          5 allocs/op

BenchmarkSTDJsonMarshalBuiltinMap1-4              151873              7499 ns/op            2368 B/op         41 allocs/op
BenchmarkCastDictEncodeJSON1-4                    333346              3423 ns/op             828 B/op         28 allocs/op
BenchmarkCastJSONString1-4                        428590              2539 ns/op             768 B/op         13 allocs/op

BenchmarkSTDJsonMarshalBuiltinMap2-4                6662            161058 ns/op           50224 B/op       1005 allocs/op
BenchmarkCastDictEncodeJSON2-4                      9237            122223 ns/op           32391 B/op        393 allocs/op
BenchmarkCastJSONString2-4                          9241            119681 ns/op           32353 B/op        389 allocs/op

BenchmarkSTDJsonMarshalBuiltinMap3-4                4444            249768 ns/op           77133 B/op       1537 allocs/op
BenchmarkCastDictEncodeJSON3-4                      5454            187016 ns/op           50682 B/op        590 allocs/op
BenchmarkCastJSONString3-4                          6001            186297 ns/op           50652 B/op        586 allocs/op

BenchmarkSTDFmtSprintfBuiltinMap1-4               129036              8842 ns/op            1992 B/op         41 allocs/op
BenchmarkCastStringBuiltinMap1-4                  428596              2580 ns/op             768 B/op         13 allocs/op
BenchmarkCastDictString1-4                        631694              1844 ns/op             496 B/op          5 allocs/op

BenchmarkSTDFmtSprintfBuiltinMap2-4                 6315            170858 ns/op           44608 B/op        992 allocs/op
BenchmarkCastStringBuiltinMap2-4                    9223            119154 ns/op           32353 B/op        389 allocs/op
BenchmarkCastDictString2-4                          9949            120512 ns/op           32353 B/op        389 allocs/op

BenchmarkSTDFmtSprintfBuiltinMap3-4                 4285            259036 ns/op           66175 B/op       1507 allocs/op
BenchmarkCastStringBuiltinMap3-4                    6000            182661 ns/op           50651 B/op        586 allocs/op
BenchmarkCastDictString3-4                          5996            184786 ns/op           50651 B/op        586 allocs/op
```

## v0.0.9
```
goos: windows
goarch: amd64
pkg: ztaylor.me/cast

BenchmarkBuiltinStringBytes-4                   225989678                5.26 ns/op            0 B/op          0 allocs/op
BenchmarkCastStringBytes-4                      1000000000               0.273 ns/op           0 B/op          0 allocs/op
BenchmarkBuiltinBytesS-4                        201333802                5.92 ns/op            0 B/op          0 allocs/op
BenchmarkCastBytesS-4                           1000000000               0.273 ns/op           0 B/op          0 allocs/op

BenchmarkSTDJsonMarshalBuiltinMap1-4              196040              6187 ns/op            2144 B/op         37 allocs/op
BenchmarkSTDFmtSprintfBuiltinMap1-4               164388              6923 ns/op            1736 B/op         36 allocs/op
BenchmarkCastDictEncodeJSON1-4                    545451              2064 ns/op             724 B/op         22 allocs/op
BenchmarkCastStringBuiltinMap1-4                  797490              1422 ns/op             654 B/op          9 allocs/op
BenchmarkCastJSONString1-4                        749572              1467 ns/op             655 B/op          9 allocs/op
BenchmarkCastDictString1-4                       1410130               866 ns/op             432 B/op          3 allocs/op

BenchmarkSTDJsonMarshalBuiltinMap2-4                6666            161259 ns/op           50225 B/op       1005 allocs/op
BenchmarkSTDFmtSprintfBuiltinMap2-4                 6662            169916 ns/op           44608 B/op        992 allocs/op
BenchmarkCastDictEncodeJSON2-4                     19543             60634 ns/op           47862 B/op        423 allocs/op
BenchmarkCastStringBuiltinMap2-4                   19448             61445 ns/op           49088 B/op        419 allocs/op
BenchmarkCastJSONString2-4                         18265             60884 ns/op           49074 B/op        419 allocs/op
BenchmarkCastDictString2-4                         19671             60443 ns/op           49080 B/op        419 allocs/op

BenchmarkSTDFmtSprintBuiltinSlice-4              1000000              1101 ns/op             352 B/op         10 allocs/op
BenchmarkSTDJsonMarshalBuiltinSlice-4            1487028               805 ns/op             240 B/op          5 allocs/op
BenchmarkCastStringBuiltinSlice-4                1673650               717 ns/op             216 B/op          6 allocs/op
BenchmarkCastStringCastArray-4                   1724162               690 ns/op             216 B/op          6 allocs/op
BenchmarkCastStringNBuiltinSlice-4               1898758               625 ns/op             184 B/op          5 allocs/op
```

## v0.0.7

```
goos: windows
goarch: amd64
pkg: ztaylor.me/cast
BenchmarkBuiltinStringBytes-4                   220974615                5.27 ns/op            0 B/op          0 allocs/op
BenchmarkCastStringBytes-4                      1000000000               0.271 ns/op           0 B/op          0 allocs/op
BenchmarkBuiltinBytesS-4                        201013476                5.95 ns/op            0 B/op          0 allocs/op
BenchmarkCastBytesS-4                           1000000000               0.270 ns/op           0 B/op          0 allocs/op

BenchmarkSTDJsonMarshalBuiltinMap1-4              169264              6237 ns/op            2144 B/op         37 allocs/op
BenchmarkSTDFmtSprintfBuiltinMap1-4               159984              7026 ns/op            1736 B/op         36 allocs/op
BenchmarkCastDictEncodeJSON1-4                    545223              2169 ns/op             728 B/op         22 allocs/op
BenchmarkCastStringBuiltinMap1-4                  856671              1422 ns/op             654 B/op          9 allocs/op
BenchmarkCastJSONString1-4                        799519              1412 ns/op             654 B/op          9 allocs/op
BenchmarkCastDictString1-4                       1444072               816 ns/op             432 B/op          3 allocs/op

BenchmarkSTDJsonMarshalBuiltinMap2-4                6319            162995 ns/op           50222 B/op       1005 allocs/op
BenchmarkSTDFmtSprintfBuiltinMap2-4                 6315            170857 ns/op           44608 B/op        992 allocs/op
BenchmarkCastDictEncodeJSON2-4                     19448             60621 ns/op           47901 B/op        423 allocs/op
BenchmarkCastStringBuiltinMap2-4                   19292             62045 ns/op           49093 B/op        419 allocs/op
BenchmarkCastJSONString2-4                         19231             61777 ns/op           49106 B/op        419 allocs/op
BenchmarkCastDictString2-4                         19448             60673 ns/op           49074 B/op        419 allocs/op

BenchmarkSTDFmtSprintBuiltinSlice-4              1000000              1109 ns/op             352 B/op         10 allocs/op
BenchmarkSTDJsonMarshalBuiltinSlice-4            1487037               800 ns/op             240 B/op          5 allocs/op
BenchmarkCastStringBuiltinSlice-4                1680716               716 ns/op             216 B/op          6 allocs/op
BenchmarkCastStringCastArray-4                   1697365               688 ns/op             216 B/op          6 allocs/op
BenchmarkCastStringNBuiltinSlice-4               1935546               615 ns/op             184 B/op          5 allocs/op
```

## v0.0.6

```
goos: windows
goarch: amd64
pkg: ztaylor.me/cast
BenchmarkBuiltinStringBytes-4                   218590940                5.42 ns/op            0 B/op          0 allocs/op
BenchmarkCastStringBytes-4                      1000000000               0.276 ns/op           0 B/op          0 allocs/op
BenchmarkBuiltinBytesS-4                        196084774                6.09 ns/op            0 B/op          0 allocs/op
BenchmarkCastBytesS-4                           1000000000               0.274 ns/op           0 B/op          0 allocs/op

BenchmarkSTDJsonMarshalBuiltinMap1-4              181796              5957 ns/op            2144 B/op         37 allocs/op
BenchmarkSTDFmtSprintfBuiltinMap1-4               164355              7015 ns/op            1736 B/op         36 allocs/op
BenchmarkCastDictEncodeJSON1-4                    500341              2222 ns/op             872 B/op         24 allocs/op
BenchmarkCastStringBuiltinMap1-4                  705914              1605 ns/op             798 B/op         11 allocs/op
BenchmarkCastJSONString1-4                        666684              1642 ns/op             799 B/op         11 allocs/op
BenchmarkCastDictString1-4                       1505674               798 ns/op             432 B/op          3 allocs/op

BenchmarkSTDJsonMarshalBuiltinMap2-4                6666            157961 ns/op           50224 B/op       1005 allocs/op
BenchmarkSTDFmtSprintfBuiltinMap2-4                 7058            172566 ns/op           44604 B/op        992 allocs/op
BenchmarkCastDictEncodeJSON2-4                      6002            188931 ns/op           99948 B/op       1654 allocs/op
BenchmarkCastStringBuiltinMap2-4                    6000            182161 ns/op          101131 B/op       1649 allocs/op
BenchmarkCastJSONString2-4                          6002            179769 ns/op          101158 B/op       1649 allocs/op
BenchmarkCastDictString2-4                          8572            134153 ns/op           80732 B/op       1074 allocs/op

BenchmarkSTDFmtSprintBuiltinSlice-4               923119              1136 ns/op             352 B/op         10 allocs/op
BenchmarkSTDJsonMarshalBuiltinSlice-4            1465266               821 ns/op             240 B/op          5 allocs/op
BenchmarkCastStringBuiltinSlice-4                1666755               715 ns/op             216 B/op          6 allocs/op
BenchmarkCastStringCastArray-4                   1621671               719 ns/op             216 B/op          6 allocs/op
BenchmarkCastStringNBuiltinSlice-4               1793746               666 ns/op             184 B/op          5 allocs/op
```

## v0.0.5

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

## v0.0.4

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

## v0.0.3

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

## v0.0.2

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
