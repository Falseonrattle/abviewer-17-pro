package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "3.4" )

func e3m5McZFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Bq0l8d6iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JxZlJkI2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EKVgBifYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KynxGWEQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func byNFsMY5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3crgHwROWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lCFyZKvpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TTZOl4QaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rN5R7qdGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func axGWzKd9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kr2LnvKFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RcV6UpOsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UFfjFUyGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PRD0xQXTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E0XkRnKVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BVcFiXPHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uOUPaW0nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ONQ7uaVuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nbrA9rNzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dK0nZngTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zCkNfU1qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TtltSZySWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G8OGeehWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gViSUTGNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dmaeiSQXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 858np0aEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xdrsI9khWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4Czk4H8sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EMv7NVeFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yJttj519Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DlUpIcUMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nWNd9rPBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FStw38HsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7j9C218MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func odU5xeyfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func na4NCJMiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3rwOCC5lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UCtC3TONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PBHFYL8AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lwgvDTNEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mIRAUn7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KJ7tiMWuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B5fcCWw7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func POlsJ14nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AnS9062eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZOt2JZq9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d7jN7odaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zcGQsOSKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I4f93xBeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VpRWHxYrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fjVBfrLBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SAbDZ652Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kthKu6jPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RnoyTH3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KCtiP3kTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VbYx4GM7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sOVSlxSGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fvKiC0lMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ThpDVt8IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hpx7lrBDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0l96y1gVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EwXfPn8ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mliv4FUIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6BWUcoreWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p07ZzXs6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z5nCcNvpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Izy5CjPvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9M0o71I9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BdPJL29hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LnuxbROEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kgsG2wCaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yg9Liv7rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 704KPKdOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func egWD8TjpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NBnJGMRWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NBWntkHrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ckGvnLbjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DfzfTCOrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4XBv8R7bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n75GmmiYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RDkWXbFtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JlurKwdaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gW3Gy0QtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p1IAuwzjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EB13UJ9oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9pGkPQZBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6WeSa1EvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vcfZfrKsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JysCrujvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gl6x95ALWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JRBIM2FHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iRHYfF5cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E5i476hmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4wj0sBJaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L6jXRfzJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H75IO4Y1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M6knHAXgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C8myUr3MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zp9F8wXpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WSUEeC4KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eJhyLH2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GfnR6iz8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7abkQNxhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ziFv9TUIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d089vIpVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8ggVvdDYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6zb8FB8iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XI7rVCY5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7jkohnXoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rQ1mqAlqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R5TfJXOxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PDuzNh6QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iSUDv69ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func drGvFZE0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YK99OZT6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func riRu0JU5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lHvEgZGgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jZ1GwkCGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func awOJFLjYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hyU5Z6mAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HCJCdjuJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x5eBqtnmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zlr7gAyZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oR3t72YRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l3J901RTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UrdjwymHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bIbINJLwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WZuDLGURWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3dHA6docWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wnREP1erWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9N6P7jXyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uyiUIg10Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vRbNrikxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gajR4qX0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FCsqwAOjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m4gWHzSAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I8Ucr3lPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iVOWU5GaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A82wVgovWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ehO19rdwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wZ4wr3EKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0OvS1kz7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UJWgJRFSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4j4He2OUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qwldIywQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fz39aoLnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H0Nv0EeNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C9Ib79xUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qxW33LGUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func omKtH9R7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hCJkEn16Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iOhfkG23Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M5DvXmBxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rhQdDJ3HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sbh0HOk5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ogvSnksQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oFEsiiCDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 47UeXGDSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func niP4RfObWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b7pMj1kqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P58NunJaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vS9ZZRe1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D4S3F4XvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9tlvcEvnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vVyXdrzrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yToZpgSZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R1YmGpBxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NZ0bh87DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XgMLLP6vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KbBK9WT3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sZ7mSWpqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m0A0z8TrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mGwnF0wdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Hl4y5JjAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sSkEGK7RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UtGCvUfnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 886yUq63Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nsnkz2cMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vxXr0czFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MkomztA8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6SsuYnH2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wfvuiX4xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IqeKuPpGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WTE0hSiOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K4F7V9CdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zIwkVgl6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vaaNSWeeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5kSUdpMkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RQfmJ0WdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iy5w9YFqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7psFxGgwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cN1EdpVUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3QGdUdFFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3w4OwIAQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cn0kN655Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qkBDenZ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XfYasyzfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8QKfBiSOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H8lXtuY0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oqd47h8sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c5R13T0tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9RUYTfegWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5ZCEJ6LbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yTamGpP2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C5JjJsjbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AoZkaaM7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2OimEta9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B6A6Tn4xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6Bf8IltEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i7fhLzHUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gjp2lqXsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uOvuZGNBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jT7qsUqZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oCbHnjvXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u1aRI6EoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MhrZ5fQ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S5deq544Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

