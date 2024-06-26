## go-array

<p align="center">
<a href="https://pkg.go.dev/github.com/deatil/go-array" ><img src="https://pkg.go.dev/badge/deatil/go-array.svg" alt="Go Reference"></a>
<a href="https://codecov.io/gh/deatil/go-array" ><img src="https://codecov.io/gh/deatil/go-array/graph/badge.svg?token=SS2Z1IY0XL"/></a>
<img src="https://goreportcard.com/badge/github.com/deatil/go-array" />
<a href="https://github.com/avelino/awesome-go"><img src="https://awesome.re/mentioned-badge.svg" alt="Mentioned in Awesome Go"></a>
</p>

<p align="center">
A Go package that read or set data from map, slice or json
</p>

[中文](README_CN.md) | English


### Download

~~~go
go get -u github.com/deatil/go-array
~~~


### Get Starting

~~~go
import "github.com/deatil/go-array/array"

arrData := map[string]any{
    "a": 123,
    "b": map[string]any{
        "c": "ccc",
        "d": map[string]any{
            "e": "eee",
            "f": map[string]any{
                "g": "ggg",
            },
        },
        "dd": []any{
            "ccccc",
            "ddddd",
            "fffff",
        },
        "ff": map[any]any{
            111: "fccccc",
            222: "fddddd",
            333: "dfffff",
        },
        "hh": map[int]any{
            1115: "hccccc",
            2225: "hddddd",
            3335: map[any]string{
                "qq1": "qq1ccccc",
                "qq2": "qq2ddddd",
                "qq3": "qq3fffff",
            },
        },
        "kJh21ay": map[string]any{
            "Hjk2": "fccDcc",
            "23rt": "^hgcF5c",
        },
    },
}

data := array.Get(arrData, "b.d.e")
// output: eee

data := array.Get(arrData, "b.dd.1")
// output: ddddd

data := array.Get(arrData, "b.hh.3335.qq2")
// output: qq2ddddd

data := array.Get(arrData, "b.kJh21ay.Hjk2", "defValString")
// output: fccDcc

data := array.Get(arrData, "b.kJh21ay.Hjk23333", "defValString")
// output: defValString
~~~


### Examples

* Exists data
~~~go
var res bool = array.New(arrData).Exists("b.kJh21ay.Hjk2")
// output: true

var res bool = array.New(arrData).Exists("b.kJh21ay.Hjk12")
// output: false
~~~

* Get data
~~~go
var res any = array.New(arrData).Get("b.kJh21ay.Hjk2")
// output: fccDcc

var res any = array.New(arrData).Get("b.kJh21ay.Hjk12", "defVal")
// output: defVal
~~~

* Find data
~~~go
var res any = array.New(arrData).Find("b.kJh21ay.Hjk2")
// output: fccDcc

var res any = array.New(arrData).Find("b.kJh21ay.Hjk12")
// output: nil
~~~

* Use Sub to Find data
~~~go
var res any = array.New(arrData).Sub("b.kJh21ay.Hjk2").Value()
// output: fccDcc

var res any = array.New(arrData).Sub("b.kJh21ay.Hjk12").Value()
// output: nil
~~~

* Use Search to Find data
~~~go
var res any = array.New(arrData).Search("b", "kJh21ay", "Hjk2").Value()
// output: fccDcc

var res any = array.New(arrData).Search("b", "kJh21ay", "Hjk12").Value()
// output: nil
~~~

* Use Index to Find data
~~~go
var res any = array.New(arrData).Sub("b.dd").Index(1).Value()
// output: ddddd

var res any = array.New(arrData).Sub("b.dd").Index(6).Value()
// output: nil
~~~

* Use Set to set data
~~~go
arr, err := array.New(arrData).Set("qqqyyy", "b", "ff", 222)
// arr.Get("b.ff.222") output: qqqyyy
~~~

* Use SetIndex to set data
~~~go
arr, err := array.New(arrData).Sub("b.dd").SetIndex("qqqyyySetIndex", 1)
// arr.Get("b.dd.1") output: qqqyyySetIndex
~~~

* Use Delete to delete data
~~~go
arr, err := array.New(arrData).Delete("b", "hh", 2225)
// arr.Get("b.hh.2225") output: nil
~~~

* Use DeleteKey to delete data
~~~go
arr, err := array.New(arrData).DeleteKey("b.d.e")
// arr.Get("b.d.e") output: nil
~~~


### LICENSE

*  The library LICENSE is `Apache2`, using the library need keep the LICENSE.


### Copyright

*  Copyright deatil(https://github.com/deatil).
