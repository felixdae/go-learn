package main

import "github.com/garyburd/redigo/redis"
import "time"
import "fmt"
import "log"
import "os"
import "bufio"
import "strings"
import "strconv"
import "encoding/json"

func pushRedis(gmps map[string]string){
    conn, err := redis.DialTimeout("tcp", "192.168.1.18:6382", 0, 1*time.Second, 1*time.Second)
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    // size, err := conn.Do("DBSIZE")
    // fmt.Printf("size is %d \n",size)
    // vs := "";
    // for k, v := range gmps{
    //     vs += " zhizi.gmp." + k + " " + v
    // }

    // fmt.Println(vs)
    // rr, err := conn.Do("hmset", "zhizi_gmp", vs)
    _, err = conn.Do("HMSET", redis.Args{}.Add("zhizi_gmp").AddFlat(gmps)...)
    if err != nil {
        log.Fatal(err)
    }
    // h, err := conn.Do("get", "hhhhhhhjjjjjjjj")
    // fmt.Println(rr)
    // fmt.Printf("hmset %d items\n", len(gmps))
}

func string2Double(s string) float64{
    f, err := strconv.ParseFloat(s, 64)
    if err != nil {
        panic(err)
    }
    return f
}

func makeJson(fields []string) string{
    // {"id":"14174435","click":210.975,"pv":10973.63,"gmp":0.019225634543902065,"ts":"201604132000"}
    // type T struct {
    //         FieldA int    `json:"field_a"`
    //             FieldB string `json:"field_b,omitempty"`
    //         }

    type rgt struct {
        Id      string      `json:"id"`
        Click   float64     `json:"click"`
        Pv      float64     `json:"pv"`
        Gmp     float64     `json:"gmp"`
        Ts      string      `json:"ts"`
    }
    rg := rgt {
        Id:      fields[0],
        Click:   string2Double(fields[1]),
        Pv:      string2Double(fields[2]),
        Gmp:     string2Double(fields[3]),
        Ts:      fields[4],
    }
    // fmt.Println(rg)
    b, err := json.Marshal(rg)
    if err != nil {
        // fmt.Println("fdafafdsf")
        log.Fatal(err)
    }
    return string(b)
}

func main(){
    file, err := os.Open("gmp.log")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    cc := 0
    mm := make(map[string]string)
    for scanner.Scan() {
        // fmt.Println(scanner.Text())
        line := scanner.Text()
        fields := strings.Fields(line)
        mm["zhizi.gmp."+fields[0]] = makeJson(fields)
        cc += 1
        if cc % 100 == 0 {
            pushRedis(mm)
            // fmt.Println(fields[0])
            fmt.Println(cc);
            mm = make(map[string]string)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
