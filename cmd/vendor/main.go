package main

import (
	"fmt"

	"github.com/cheekybits/genny/generic"
	_ "github.com/containerd/containerd" // security vulnerability force pull
	gennySarulabsdi "github.com/fluffy-bunny/sarulabsdi/genny"
	_ "github.com/valyala/fasthttp"
)

/*
This project is needed to force
   go mod vendor
to bring down files we need for such things as CI/CD genny code generation
*/
var gennyType generic.Type

func main() {
	// forces sarulabsdi genny generics to come down
	fmt.Printf("%v", gennyType)
	// forces sarulabsdi genny generics to come down
	fmt.Printf("%v", gennySarulabsdi.ReflectTypeInterfaceType)

}
