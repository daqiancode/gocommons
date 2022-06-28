package jsons

import "github.com/daqiancode/jsoniter"

var JSON = jsoniter.Config{Decapitalize: true}.Froze()
