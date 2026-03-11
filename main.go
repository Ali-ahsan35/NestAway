package main

import (
	_ "smartours/routers"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	// Add custom template functions
    beego.AddFuncMap("mul", func(a, b float64) float64 {
        return a * b
    })
    beego.AddFuncMap("contains", func(s, substr string) bool {
        return strings.Contains(s, substr)
    })
	beego.AddFuncMap("dict", func(values ...interface{}) map[string]interface{} {
		dict := map[string]interface{}{}
		for i := 0; i < len(values); i += 2 {
			key, _ := values[i].(string)
			dict[key] = values[i+1]
		}
		return dict
	})
	beego.AddFuncMap("replace", func(s, old, new string) string {
    return strings.ReplaceAll(s, old, new)
	})
	beego.Run()
}

