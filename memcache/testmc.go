package main

import "github.com/bradfitz/gomemcache/memcache"
import "fmt"
import "time"
import "strconv"
import "runtime"
import "os"
// import "reflect"

func stressMc(idx, n int) {
    // mc := memcache.New("192.168.1.227:11213", "192.168.1.226:11213")
    mc := memcache.New("192.168.1.226:11212")
    for {
        nsf := 0;
        ts1 := time.Now().Unix()
        for i := 0; i < n; i += 1 {
            // fmt.Println("set", i, "goroutine", idx)
            key := "foo-" + strconv.Itoa(idx) + "-" + strconv.Itoa(i)
            err := mc.Set(&memcache.Item{Key: key, Value: []byte("my valuejjjjjjjjjjjjjjjjjjjjjjjjjjjjjllllllllllllllllllllllllllllllllllllaaaaaaaaaaaaaaaaaaaaaaaaf" + strconv.Itoa(i))})
            if err != nil {
                // fmt.Println("err happend setting: ", key);
                nsf += 1
            }
        }
        ts2 := time.Now().Unix()
        fmt.Println("index", idx, "set " + strconv.Itoa(n) + " keys, time:", ts2-ts1, "fail:", nsf)

        ngf := 0
        for i := 0; i < n; i += 1 {
            // fmt.Println("get", i, "goroutine", idx)
            key := "foo-" + strconv.Itoa(idx) + "-" + strconv.Itoa(i)
            _, err := mc.Get(key);
            if (err != nil) {
                // fmt.Println("err happend getting: ", key);
                ngf += 1
            }
        }
        ts3 := time.Now().Unix()
        fmt.Println("index", idx, "get " + strconv.Itoa(n) + " keys, time:", ts3-ts2, "fail:", ngf)
        // fmt.Println("index", idx, "set fail: ", nsf,  "get fail: ", ngf)
        time.Sleep(1 * time.Second)
    }
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    ngr, _ := strconv.Atoi(os.Args[1])
    // fmt.Println(os.Args[1], reflect.TypeOf(os.Args[1]))
    for j := 0; j < ngr; j += 1{
        go stressMc(j, 10000)
    }
    c := make(chan int)
    <- c
}

