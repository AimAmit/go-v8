package main

import (
	_ "embed"
	"github.com/kraken-hpc/go-fork"
	"go-v8/server"
	"go.kuoruan.net/v8go-polyfills/console"
	"io/ioutil"
	"log"
	"os"
	v8 "rogchap.com/v8go"
	"strconv"
	"sync"
	"time"
)

func init() {
	//
	//	log.Println("init")
	//	reexec.Register("client1", func() {
	//		log.Println("child command", os.Getpid())
	//	})

	fork.RegisterFunc("client", func() {
		key := "demo"
		log.Println(key, ": child cmd ", os.Getpid())
		os.Stdout.Write([]byte(key + ": child cmd " + strconv.Itoa(os.Getpid())))
	})
	//fork.Init()
}

//func main2() {
//	foo := 4
//	bar := 10
//	id, a, b := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
//	fmt.Println(a, b)
//	if id == 0 {
//		foo++
//		fmt.Println("In child:", id, foo, bar, syscall.Getpid())
//	} else {
//		bar++
//		fmt.Println("In parent:", id, foo, bar, syscall.Getpid())
//	}
//
//	//syscall.ForkExec()
//
//	time.Sleep(time.Second * 10)
//	fmt.Println("back")
//	syscall.Kill(syscall.Getpid(), syscall.SIGINT)
//}

func main2() {
	wg := sync.WaitGroup{}

	wg.Add(3)
	for i := 0; i < 3; i++ {
		//i := i
		go func(i int) {
			defer wg.Done()
			//sys.Cloneflags
			//log.Println(i, os.Getpid(), syscall.CLONE)
			time.Sleep(time.Second * 5)
		}(i)
	}

	wg.Wait()
}

func main() {

	//main2()
	//return
	//get, err := server.HandleGet("key2")
	//log.Println(get, err)
	//return
	//server.Bundle()
	//return
	server.Server()

	iso := v8.NewIsolate()
	ctx := v8.NewContext(iso)

	if err := console.InjectTo(ctx); err != nil {
		panic(err)
	}

	//log.Println(jsCode)

	file, err := ioutil.ReadFile("js/out.js")
	runScript, err := ctx.RunScript(string(file), "index.js")
	log.Println("runScript", runScript, err)
	if err != nil {
		return
	}

	global, err := ctx.Global().Get("global")
	log.Println("global", global, err)
	if err != nil {
		return
	}

	//object, err := global.AsObject()
	//if err != nil {
	//	return
	//}
	//
	//main, err := object.Get("main")
	//if err != nil {
	//	return
	//}
	//main, err := ctx.RunScript("main", "index.js")
	//log.Println("runScript", runScript, err)
	//if err != nil {
	//	return
	//}
	//
	//log.Println(main)

	function, err := global.AsFunction()
	if err != nil {
		log.Println("function", function, err)
		return
	}

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

}
