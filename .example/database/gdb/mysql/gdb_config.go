package main

import (
	"fmt"

	"github.com/jin502437344/gf/frame/g"
)

func main() {
	if r, err := g.DB().Table("user").Where("uid=?", 1).One(); err == nil {
		fmt.Println(r["uid"].Int())
		fmt.Println(r["name"].String())
	} else {
		fmt.Println(err)
	}
}
