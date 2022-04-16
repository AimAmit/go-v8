package server

import (
	"fmt"
	"github.com/kraken-hpc/go-fork"
	"io/ioutil"
	"log"
	"os"
	v8 "rogchap.com/v8go"
	"time"
)

var clientIso = New(0, 100, func(k, v any) {
	v.(*v8.Isolate).Dispose()
})
var clientCtx = New(0, 100, func(k, v any) {
	v.(*v8.Context).Close()
})

var clientFunc = New(0, 100, func(k, v any) {})

func RegisterGet() {

	t := time.Now()

	log.Println("args", os.Args)
	log.Println("pid", os.Getpid())

	key := os.Args[0]

	var ok bool
	if iso, ok = clientIso.Get(key).(*v8.Isolate); !ok {
		iso = v8.NewIsolate()
		iso.GetHeapStatistics()
		iso.TerminateExecution()
		clientIso.Put(key, iso)
	}

	if ctx, ok = clientCtx.Get(key).(*v8.Context); !ok {
		ctx = v8.NewContext(iso)
		clientCtx.Put(key, ctx)
	}

	var function *v8.Function

	if function, ok = clientFunc.Get(key).(*v8.Function); !ok {

		file, err := ioutil.ReadFile(fmt.Sprintf("js/%s.js", key))
		runScript, err := ctx.RunScript(string(file), "index.js")
		log.Println("runScript", runScript, err)
		if err != nil {
			//return number, err
		}

		global, err := ctx.Global().Get("global")
		log.Println("global", global, err)
		if err != nil {
			//return number, err
		}

		function, err = global.AsFunction()
		if err != nil {
			log.Println("function", function, err)
			//return number, err
		}

		clientFunc.Put(key, function)
	}

	fmt.Println("function fetched: ", time.Since(t).Microseconds(), " us \n\n")
	obj := v8.NewObjectTemplate(iso)
	array := v8.NewObjectTemplate(iso)

	array.Set("a", "a")

	log.Println(obj.Set("array", array))
	log.Println(obj.Set("b", int32(29)))
	log.Println(obj.Set("c", "amit"))

	valObj, _ := obj.NewInstance(ctx)

	call, err := function.Call(valObj, valObj)
	if err != nil {
		log.Println("call", call, err)
		return
	}

	fmt.Println(call.Number(), err)

	fmt.Println("res: ", time.Since(t).Microseconds(), " us \n\n")
	//return call.Number(), err
}

func HandleGet(key string) (number float64, err error) {

	//cmd := reexec.Command("/bin/bash")

	//cmd := exec.CommandContext()
	//log.Println("path ", cmd.Path)
	//cmd.Args = []string{key}
	//cmd.Stdin = os.Stdin
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	//
	//log.Println(cmd.Start())
	//
	//log.Println(cmd.Wait())
	//
	//os.Process

	fmt.Printf("main() pid: %d\n", os.Getpid())
	if err := fork.Fork("client"); err != nil {
		log.Fatalf("failed to fork: %v", err)
	}

	return
}
