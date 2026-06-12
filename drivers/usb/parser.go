package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "2.3" )

func LN9WZqLJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oKY2vfujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func URehj1whWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AxWqbdBoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8MqUcRl2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FXtPkMXFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pekpxahTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qXjWadTUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z7kFWDdjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HGj6Cu82Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lJcFMCdQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mGe7xT9CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CvDJeZXAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sfcCeB31Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wxRY9IjNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WXIWRsHrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 04iZ0cWbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e365ktZFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Af3zzj4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rw8TyDsfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jAUnkXUQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vftjpl4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O5BvmIYHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func igMRnhsGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Evvab0rkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VSSgoOw3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sZEo74g5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tO4IZXhCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LwbGbTLlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hWAJfOwJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hhPdWhAeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 47YTy9p6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OXD4RlXXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fyyvzwMbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M8SIYh7nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tnJdRL21Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func npEPKxkfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pexRWCdhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Eb7Lig91Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 49C2WOBsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VCeEJL4dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hPas9IsdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0UXO8oLSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CFBVka7HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 03TQVuCWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dhXjCqH0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func darjDGgUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sXY312QvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CbKN8IekWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kKYPugpeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YoUdGKGbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L5yYhwk6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func necUo5SUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lB8PULxtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G5cYHCc3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j8VKRBcHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7jnuPX0XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NiTE7PPvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5egm8lX1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DRfZJ7bWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IQnCcfZXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YCpJmha9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A0DbwVszWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WGbk8L70Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pWZul6PSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iM9L92eXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p5qTQZnZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jcVX24EQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YqoVajmFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rGK0zYyyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x70BWL8BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M2VFwbGMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w4xh5K0vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lKerceUJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XM9XoMeZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5Eq5dv1SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TOx6LaqUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4bXKIz7aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tAkJ6219Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pfnwWXAXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o89OGydUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UlvFkHrmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func flSB6mkHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6HjowstJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7nFFw61zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DWuJcJymWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6SujBefqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7pBmZT0hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rvwt0eZ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vu7FOd7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Um5wb8XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FAllbPL8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iyDWDzoqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gCVGfXPvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O26xXaoeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qtHCcJw7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 66zqrdhLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q59jjd8PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m1DueWH7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b6iOnRgDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TGVXmDi3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WAgWm1MoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UfXxgvmEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a3kNvwzBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BJK8NfusWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7arKfMwJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ujtNlKPwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WQ4GIjizWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CpjCotlbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iXExrXdvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tfvN20mvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WMAU48d2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AQu8zNhcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LVa5LiErWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2kmnJch6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cih4fm95Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pTtm2dTZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WjfzM4tRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aESg1ayWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3YrtzNUEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K4NorJSuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wUOTn8SyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qb8AIXJyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4OzNRynKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mSZbSpLHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YV166ZeZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ud5StjgcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CpAVl54LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wOQLLzzQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jBB7VrQTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aISItJ3FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4nnE879zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6J8q9WoaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vjLo2gP8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x2CeJK7gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x7ixKP3SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b7GmKr98Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xLGdvJ1DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FA2AIbkZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yth8CyXuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sn7Fu56SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f2lYIUUBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a8d0AwlZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K9mF6lkDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W5dRZ4txWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c6OldOj0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xcM6IvZ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZdiZXJAnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CTgzeLFAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qjAFpPOgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c4SiFoPbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zFK2S191Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bjYhSTEaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aCnMIMuyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dCg5QHnmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8YEI3YwXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g1WGU2ogWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JUvDbz3FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k5SDWlzdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A8EGqYQcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Kb95WVmVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t6feA8z6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RvQRqCvjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gDEqG2LLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0kntmlRfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kzUVqj1aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SQAfDwepWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uNqHv1JTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tmQxcqNWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WSt7U1rsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rbPpbE6IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func muWCK5CkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2LGjQZPdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XcmnfX8cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YGyqVUejWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wNSfCZf6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3eZEo6vAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rKH7ZCFLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ip6hLIy9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func go1zOEGrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KLuKhNcAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8AEISkUuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WIhQxVycWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WZypO6zcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NYah719tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kFvnDEfcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6pvS0Vv9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rASWWGMNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BUBNFRRwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fQLkFmWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lzbRfBicWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7HwdLnlCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ySnBBmtIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gy9TmMB7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wph4iEZ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5uxg4ShnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A2skaeh4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i2kCs2ZYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hlw9Di6yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l7aUG9bXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sX81cM3EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rRws3xKJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t6sYgpaCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LtpuANpMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iFboC0t1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3EeM8MtBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W2Agmix9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NtQBjYq9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WFQXD7VKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n1jZKxPEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FGrOahI8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qDOdM52VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AL40pGOkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func neq75MqEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vZwnQjI6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wD9Y48PgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tuvSF5sdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p3Od5hbLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zjqwGp8wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CgYI85hrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tibWVMqvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fesuk6D8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g1hObPExWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CX4ayJ5wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5RSjAvPIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t9KOwNw1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0hSAIyUMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pnNro60HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xo5oYcnbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jxA368l2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wC2IsPpoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GNkrylZlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bfMjbue6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mxq1NY5YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G28j1fV1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func exkpVrshWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xOMCUVvjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cvUNehDCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WnxDCBDbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func joP3FnR4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func idr0p3bqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FlflJdwyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XFZOMDApWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4okBr648Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TzxTRiqvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PzD1RK1VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uefkVQCuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func REbUgbg5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FBr8hBtYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J4ZjKpIqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QQBURtDEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4VktJNDJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wc3ohzw7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func REvCMdnYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qAk7HE9uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5PLCFayMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3649lnxGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CimXP5j0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pEVDcCj4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YMwStk7zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5S9VMYpjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ykfD6pZ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xo3e73KnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 65OC8oEbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y59QsPvmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KtbOih3PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xjcmKSNvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zpmWPUUwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mKptIciNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aUfznbNZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yu3suOCYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vm6PI7mkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uyXyprfIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1HlPLGyeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bUukUcG5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VcbOIBqrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q83rjH3DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SnAqssFsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gcy4qXKYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1aN84gAeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QTFsr3PiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I91cA02XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6SLsvPCPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JnBrqbNyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 23VqVRTUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UQZqY0XXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kWEYBNF3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6qhRZV4dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JRVnqylmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8oNLpS1kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CzEfuP07Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o9u05JYtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qaJGWMh3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9vGvUUeqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MSlL1YpAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8xg92b0WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3FB2VZ2UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nkeC2L7uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EjjLETuvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fbrf3DeYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sn3Q9K9xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b6w4y446Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zHYwa8oaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JV3yyKZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cHXTMh4fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aL0iDnXLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NlPKf57hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rBBuR7a0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QeXqq3VFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OMjUTshdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nCr16j3KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aqz2MoB6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GZVi7W9CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4mQwgLLmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oUE0saYYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oHfUrYD1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func llJnqCvwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O9wULPRkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O039FTCEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ioUHtBTmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 42OTWmr2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TvOe3WhlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xu0694CCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sx7ekrgeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EfUmfhvLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t2wR0XUcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZKYhpQEkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func frhRK5tyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OOp8qecNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RTXaJKj1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jwdj6EPNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QfalKvo8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dIXeA8BxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UJvx18wPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PnygWh50Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wAqKmfwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w1kFuBCQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bPVmpDArWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0JXgILbxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xIOq14oZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RhEBpQzEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hcdgtRJxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YfvzdkvYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kPWvVzmpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MojbshaxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ti0InV78Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CylqjgGzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MzxFQAZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S1upPj8kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WdF4VVKXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func amKtW5GlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dMvhmKnHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zJgtbwnjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7djbTgqPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qOumkcSgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hShJAHkLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6jokX8jkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8hfAqqn2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fjF5spz9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qdGKIi3ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a2z5iQMgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QEYRiPEHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8oj9vl3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UTbp8L7zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HWlu4NCEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LoTvtvWxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MAevXalNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DyK9OAtUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bZtWJcWuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NlqqWx1RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ppm4MJiuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GnZ00tPVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KaEXgrGtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Apz3CstzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OR9MqJmVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0Wid0wIjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func izfPNKbSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nUAnyARMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3QAJevKyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gsdnEUwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a1j49nDQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rK6INct6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yh6JbxB7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qOUTDGfHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func amgnbQwsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HUMDbDDIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qF2kDzs8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cokLre6UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1PUiBbp3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DZtgGN5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 76UNbT8sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DvcTrK4aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O6Njw9IEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QNmIKudfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N54m3VWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eCNniissWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FvEuHOXMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i1sm3ppLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BDa48zzWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VZzlnFsLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mCHHDJ7mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aoV3RoTVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9pLzfepmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AozMIsqwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WMkSgHCsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func acsbEF27Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qlpyxG4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wwEYBZMNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xVlPtkZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0JwIcU4JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HnrpNPQoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TXWPN28YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func owoPcr21Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D0Bsi8eaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gjYphJAfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a3GGxe3hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CYRwvxqFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NRFijpwiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SIOkqc68Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MmtE9aRvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XbtEcAI7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4eJ7N1DTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zlqQZJEGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sJu005oKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func elTGqWcgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KZgddoimWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4JvDDEBIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WoF1BEgAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4HH8oysPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xtvwhCNtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GLq7nIRZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gq195XQtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xtNI7gHWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RV3m04wfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pDw0sn2VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ALbP4wf2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HP761AonWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WE0JIHf5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2FiGradzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bnxZELzhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yH1ANDMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3iBH2p7JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oyzZF0LPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H3nkROu5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jQdO2ghaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yskriOlxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 33KH9bBHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Uk7IU8WaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3I6pTwqvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CCYH39OVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9txGijsJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Ivs58edWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GqC92J0DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3DP3iJcPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VSfyIf0uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pQOh0YPyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2bV3yAbZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YJgK6tEoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SgSURhI1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ThSRtzWYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xfKxqg3zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P1TU6UeHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2jfNBwofWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xh2heGxqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QXWveQR9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func avpHZPBOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eo1fu3LqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wenbQERlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3wPrNeLiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7cllyRd3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GlGMGwVIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vb2EOnOAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1ZFHNJE4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DDx0vYRBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UPpxoEfzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BR7eEdGEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 23SBUbSeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tRNFmS4TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dl3qoM9BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rUGo2j1xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func URANFzj7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LlZeKO0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bHV3UjyvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gX4rwXqWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bljOL71nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s7uKpNQQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LMV1j5T4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zr3287G8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XsQV0gJ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ajcwv59wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dmoRHLisWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZE9PYPD1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ioS8sZ0EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cglOEAbFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tLciOlGcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gcMxUOTYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JF9htXt0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UJvqaUxDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iHevB6QUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2tEgO3kcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ERGtymx8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U7gTzWeOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B8LW3CW1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wIXqksFjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nz1gvJryWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 85BnisWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zLXp9MP5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M4GVJpsnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BtKunQPlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NvB1Q17qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jpZzS9A5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xgADJCLZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oj4WdH8wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r090rmViWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vah7w0ThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5U2ot66kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6uVNxardWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6uKoTnA9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YfEeMpYSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jhp4GABvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jAO9B6pYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z9jLk8q3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uWwbAiUbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xs66VVEXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YfqBIVo7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lChGF1cRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EgIAaaVGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rZcFqJDSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u391GpvbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7TqNqoL4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XzmSNEeVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y0Kn3RbaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dMveME4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AhBKpnUOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DiGBL4stWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MI2DasbcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KhozllMZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k7OnkhAwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hBIvZYeQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I3yDLRM9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XwG2YAafWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hHpZLLTpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k7muLYjTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hMjHzEOgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y4ShQVZsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AUbWMWNbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZIGZcgDPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VMMz1hReWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f3iJuztQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H4n10BWbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 61AjayiuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tWxtWlAWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pN4lsW9EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RS4sv1uKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func epNxrxiNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FHmn5VetWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SzDHAF38Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rOHrwxaTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BqLcMqkpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h0nIPdGDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ChgjqdkJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bb3Edtf4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xz6ApgueWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gi7J3M4FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 66ZBJh4XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G1MTnhOJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w4CRpD9UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QN5R692zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ivqO9VcKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GjkgVmcyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4cVWTVIgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func to0RSnzhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rwAAS5vAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3OajBi93Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E1maJlM6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LMPEnb07Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HVkjSmPbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xecupb0tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9oW5Ix2KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Hcbk4lGQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0yf1BComWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cKZSekiAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KMJrHOGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A3C4d3CSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uQZX8KDcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kULyET0bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pBEN4akyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YRMn6KVCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5rdUDvV4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ECQms3YUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nXYQfukrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pEPQdVMrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YNRgWs0EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kHGXkoWjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iPfyMgL1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7NnoCNrlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SnfOX6KrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BYDIgy9nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GktP6amlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FZEfu9wkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tLGnPppXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V8tecTKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xVxFvlfcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VcZgzAbvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1YqeUrgeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DNwXf1DqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oqMQBEElWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TQrtRoG9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FYiM6R7iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R2JhB5pyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ldiWolZ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ECQN5v38Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8OkuOw2AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rLv6B6BkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RUwoSOlcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OWcrkshQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lDt1ptscWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X3887BffWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uaYS8bSpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nrOQmFF0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func radKUrghWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HXz3XJA6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZBXV6TcsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lYlWDeFOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xqmhuQzJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MfqpukOGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oPDVoF4VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OHMFYiz4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jg5jALPvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eVOg7SXQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s8eceJTDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8jXGQWPLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ova0rXJCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1LOtdTpJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uP1oly5jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ik6FUpMlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AMrhF0OTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HGEbCddxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q7pATKUJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NAf9hTVTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TYX3o4n9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yk3uUpJ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rw8ijSdsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hHA7oW7OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y6tWIQrLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func suxjhFefWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HNonn27gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BDxRr2KaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q2GCUi7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZwreEkrUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q8lSN8BUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kHgj7TLlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ve29GjjQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y3EV3cydWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 44zaodAyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nInzdSESWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CGoGMOVGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3Kl8opKZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DKxYEMZyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YvU8e7i3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oGbLd0PvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 80h71l2CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DEXZPanLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dJCZLj2IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OSbSqRBjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1eu0AbIRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func opj8l7qgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dvjyEe9aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FAWCu2OKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VLb1pucJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O0r37tlaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rc0OrU4iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6ex4KTdZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D7YjAKCJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CA1EieOiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oRcVq2Z4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7zpV1aLpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lINhXpMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fkjV2tHVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func asDT2QbZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cORncXPeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mLSEmYjHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bInaIi3pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lxIFlPMmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8EkopOBFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cVdU3STXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GZproaZPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2tMKo01SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mTuIPDubWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3BaZ1JHhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3lT6n2XmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GN6n2oxvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OTAYkor9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9YBHEi01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uxemdiKgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ka31O25fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q2uY10BCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FrJzhDl0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JnuE42CeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1JpjTbCvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1lrKCTCbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pY52mGb2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GG1y00rzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WiCeGQhYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nBthWv6VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lT6CXjarWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8Pkfm1tfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O2QBrxf9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BMRftcSZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u8kxFdKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3yqMtKj4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func krDYG4aIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YPar0w6VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YzaYSQqRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 27938e8nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Aa2xgsOGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eEwTz2sxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UJu3Soe8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PZpOLNdtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r4uK7STEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Ur1ulAfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uKKSufN8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AY0JX5a0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IGCbLSftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IJEw55i7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kNynFpqhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mparNW9EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lJIhEYIEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9VJ5JQ5iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v3PCDNaRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NTSMjJKGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FM5lZIWqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N3twzcJaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jOwLkigNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k7Q12gp6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ANO97qEkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ND5hyP5RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4HgRblzVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5GhKf9eHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jPeanAYMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WBgekI9bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SShajk01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uLYGnjPTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ptAA4uO8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W6gnRHT3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CBFFg4RYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M9SV2oIuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tb1JxiEyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TB8jRLdpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3x7AZlCVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Uf2OvpXnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4vZCznrxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fisPZIRzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ccVMU5ijWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lI0EcGSpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5Msh2PDpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8VLYzWtuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H6CM6lq1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7fC07IVBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aaLpCs7eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YmljLvjGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I8BvrPV9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 03M9Hdp4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2BKa8N77Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ad3zi9BFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XrojYus1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uhCzeaVBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gzlFC91rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pN31xsWeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LbD2SkyAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9Pl3Gp9aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HBP1cO20Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wa0Ik9HtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XsZtgObKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qVshkVocWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fjltHq1NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8p3aGvhyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TzZVMQHzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NBOOe1JWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3HX5PmQuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K7Vwe9oKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4YBNAOLlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FHraQ74LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b85ewc2dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SME7FMiaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iZ93d6zNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func splb5aV6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tryqI9etWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TikEAy2HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R4SKvBtyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ujbJsIawWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bvU3kmq7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vYZF8BsAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NT9tF9SqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1BdZ2iXeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q9HjnzXFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NCvHjThAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PuCGTdXOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vZYDyGmzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HzuWPBQiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X9LuJW8hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nYXMhSmjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Oj7QTEIrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Shy2oVNaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KGFzbEZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZP8Ieo0rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HsMqUo64Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DOWtEX97Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func StPY54CbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9QxsnHl0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K5ke3qudWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s0SdFV5wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sw4bDL6JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5NDDFNGaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pZwdTWLUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func drF53jE0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NYyblCLFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xMpsJXJeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wE4y5SmNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zWF6k3CeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 06dThqasWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0TDkZW2hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9hIffTkQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MuOkInFiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VxQJt4AyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a6kIXCv0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g3sDhbfUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qe9yujaFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uyhVYyvBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tKBd6Z4VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yTm7dGHtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func scq5CTimWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GAQbnPnGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DOnNz1sKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3TXYaEguWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cyqb02ETWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0oQ9w8mrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tkr6geSEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g3hBbQxMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8C3zDOuYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hQgwucyeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fYzqdqDRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tn7Mgd1zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dUJLmqnCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZRsGrmfbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vSmq6KacWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func INXwIUkIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5oJGUkZDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m7q6qON2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jYYEWC2EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 26xsqrq8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NAWUtIHCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YmHzEyo5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zIWcrOxlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jFMCn8qMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QZcvMQJWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v3ptVwNcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func olGwHk8jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qKTT6a6sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lUnRQd3yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func euLkX0oRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iLJa2xAeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TWHOwPf3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JpdE53HFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uTU5hY4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mkAAFUvOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e1xifJmfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xNqHBtdjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y80tVw2RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tuJW9cq6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vaiCHf9lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yDw0EJQxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ixbKdvqZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PUL1jsXUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sN5Z9mqBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NFQIpWQiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MuVx1rcMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OYzTZZQsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GjQOhGAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0TwPA6aAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9fyc6Mj4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eHXtwqFBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WeMTo5eNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dB2MFYPqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func STPwxd75Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GSoKUSYaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NVQQrb2uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IQKgWt8CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bnaJb5elWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Voz821hrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eDyMeR7MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J2RbNomfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5BbHZ2cmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p8ESSaZMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l7otWgfnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a0vMGIpcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Iw3cx6cdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dBJEWhUbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JiQkHdqwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c40P4AZpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iSrNAxbeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3Kwy4YNVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Psd2fJHnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cIIHGoiwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SxsGdN6WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z9UArhNtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LoVXC5uhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MclVQ7AiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bjXaDqyvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bxx60djZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func stsjeXeNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tc2VkOOlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6R4O5bG2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tLQnwhzqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JmfNcnbgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rRtgSFmWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ngPbUjfAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uazOAK8qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5oHfT72DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QDoM9hnYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uefososWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ldvGudu9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ylIPLHE3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XyCUgljDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kklz8HzlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gEu4ETOxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZUz29QorWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z22kYdcwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mvNZd6WyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ebU4FnaKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Bt8ai5PqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VV7xm9itWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eZNfMp5sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9XEgFjZZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pO9F2M8AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6MKyUlO7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BXgCLt1SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cCO0qSYVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YRTS5hRGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UzMJanGoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 15QylG6wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GOrYrfWkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func szlYEWWoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Le5nY6M4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ty8cIa7RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VbkkBayqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oOk3gLNoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NpSrSd9WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fzRHQyBCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xpp7YkIJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4UsmuMzKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qu7qW9YUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qrc9lhlAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ds9IY46eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yZ7dcNJ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5yvHpT1xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BfqtVh9ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GhP1a2lOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I9Vl6sBbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NwfpCOo7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xk8J22SZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oGRQN3lhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KbtV5XjcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sfXi9M4yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EN9ixrN7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pvuda0BZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TufCQys5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E1mY6YOEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MfiDaCpNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xDnp7es2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dV2eODTZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7H2F1XCXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iL1qxwM7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GS5vwnOuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iiMLymrEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5kgNJEr9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wildM8bHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PebzmfHDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ryqWj74kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vcLjB3iuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iVVvJPmSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PKBUzqr0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mxjDEBaMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xJnLII9JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BCjkBirTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qnW2Q9frWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 42PGN0hdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gp1T7z42Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uxJMJWdPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qOgTVU2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iv52vwYWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aPPDKWpMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mw1PCGsHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rHUDx8UlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1xTVQw0HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sGSHoZgcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1WXWV2zuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Oh7l9Do5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Bk5dGg7fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eJVufFHIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Me2jn8cfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gZsnNkS9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wFcPTE6UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6v38YGa5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4GvTSw5zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Chcxc6TnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3xE1OWOYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UUQIZPcxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 40ZsrAYaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vx5tGgnsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XcWfWDI0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YBP95MwUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f7yf8gziWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nOX7rKxqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LOQxKLQDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g4ZQ5wnpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jQ26bUGbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FUiOqIHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8NfpfWAjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 83CbuDseWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VJSl6KUjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0j2IFOFqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gcnBN6ilWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mrhWokMIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W6raa1NbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ms1E0BS4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jmXUoXpNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S9CYBvwlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d5CbyYnlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1gt553J0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vf76p4avWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sA7r99SxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZhOTDiqaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HNODzRcTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ohXLfhPrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rK79wGj7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jEBHOFNkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ndxTEchEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F7uGiK0lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZDE0BkaUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w0f9FzxBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3lfq5EHhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nJQyd1GuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pNV5xULSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uO30AMCDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func un1NdlMcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B7VC07fMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AiL9jINzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 89WtNgheWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CaxQZCzKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xWVLCI7DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sMdmMVbeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KH0JsDjHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3OTiAwGfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0PQLzQIyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t9TuuhlYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ecww5AXLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3DpxOcdvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zLljMnDEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WNtr0z5BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bq1OG6nTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mdanRr2zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qF5sxUH7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NtVmCWLlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mDzBaJvHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uS54TP5eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L9Jsur1YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4754wh5qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UnPHNWNHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UoHGyRWvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rkzakeFvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QI9CA0BPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dKyY1QtcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9ABLzGOyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0QknlkzxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I7HAcLxuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eypjUPakWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nyhaCKoPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yhtruxfTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5Ex5rqtLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gmOTElHCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func emrSvIlwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qn9GQZPLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4ly9sBkJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qv263X3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hQ9jfH5zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sw0AjP7JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DA8RhsSgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tl1WtQAsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ccNngC7yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3dMCv5StWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dkZeHj6QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6N4gCTuLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oikQTJ7yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DI5wWM1iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kckVmddJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K2R8Twp5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rHrR93QoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MwaNPImRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qpWxI86jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M6MQiMYIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xW3uOGkmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0PQqtWmoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BCCv7K7MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sXYbTlyOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Yd4LXThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fy0Rvd2mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ej84zpmPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S9Yf7jIkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qbL9jgWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vqegJUoMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9Rj75UPTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t0ShUdRMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dmvkwmk2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OwNpWQu9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CwlVxwjmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yGH9GBC3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BsQMaoqrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sTQ4raUSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Avnq6HNUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8WmPgLGgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IE55oTUjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GYbJMb2DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jLqJwovVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 85lmmXxMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qS8tDAUmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2fVI7xvwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YxxLf3epWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5n8G5lJnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lvgpRArMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YWci0XGgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sq1LscanWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q6oOTnbbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B6oNGCQHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aEAqRd0wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6SOoct0RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sf0Dm6RPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OilNWznrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func abMVQYKmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J2mO9dTKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4wveihWGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KKEkLRdtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tHlnGXefWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UrxB56teWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HRDsWI2dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FlkLNWWgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RqYSEO1WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gVSW770QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TwJJX5KyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 80Am7CDNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oCFtLAMWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func He82tEHXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I7SBZIt6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UaLa6V4rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jCTcwOjnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yeUCvtd0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pXAZyY8xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QMNq8yHaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PpmCLphYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dDkMQjrMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xsSaJ14kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1MjpHXluWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pq1shMuiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P5xWWaMOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TWSwCJ8jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FDM1pqXJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vYF8Y9ecWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qX0B739kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XvJwZRbOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eppSPjnGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1IPb7ZtSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mfEHeEpNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DQ0wzgeKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZOugfPe1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NbocKcB7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jRgpsstOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GSarGnV8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4O5vih7fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6E9vRRUEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yC5hH7j4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IuOvjh7kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iGtmQ5eeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XgoXijBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iBR0P3XPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F9tmwoOgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TqkwxQMeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YIuEX2RSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6yNoDUx0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WtUVM7YKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TiNJETN9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j0eooRfqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KJBaOAlgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NOqOy6cxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uFNSheRCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xdQSUiU9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zuAge2epWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DWo44zTbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AzT0gEpEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BBOp6KpbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rAp5s91NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bhHSlp8mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gNO2Cr2RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 51t6bbKNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pvualHNbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TZ1yOJFMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gf56YjUIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N3gzGkBmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2Ed0lF1mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yhoPXet9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r79vCgYRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rlLqORWwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fk4SK0sCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gLTkWdwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NzzTGPHPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5XIjAEM3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y00dXnNoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BV8QEUR0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qQh6FbhSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8LwovdgCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PmXgpJfZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xgVNj22bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HlcyfmMrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oO4CE00aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R5gcwIfHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wmc73gMEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xOrKqmpUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5MrxxyUQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jiDq6qcTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wyR0sITTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9tFh4DJXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iil9FT2gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XMtZvPKoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5GEEtdCUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kWDU8jNPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pRXpCl9SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F7SScMYoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DIl7jdHYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ntShbiNVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UFiewuCcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DgN1v0lZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TrywBBDwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0hv46YWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ydc1elBrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ltAGEA2WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1Q8Avx6kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lIL1VnhFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GnajduX9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PRPW0MROWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MBy5cKgSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kbk518xwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HFRnQTYQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func luLx6UEZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EaXpyZOeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q7XRsFbgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yxkGEb5iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yjSX30qyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kNZ1VXBbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lJ8OCF1OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V89kH41kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m3PVfLIzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 07JaHDKmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qoTZ6GBgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WVoYoHhHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func utDTUb3LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hpKQw9h6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ll2Ug7C1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0deReugEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oxptW0YAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func reqNpSPcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uxJM8BTXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OttKTgk8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PLx660eRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func avQV4FIkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iu8dh9ozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oh6KnTfyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fvkCGNS7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uvb8TM8NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9NoN6x8GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UVPxU7oWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AfjogGHJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7s9vEKuYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func olCbblWEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Reo0DpmoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EmlD98XJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P1J935MyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sRTSIvg8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U8AOPpOnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hrknpW6bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e0JitDXEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0SWMMWQzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pXyLrV4pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3IDAJBaRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func an87xVGyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jUfoGHCZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vS6K6iCJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K9QNV4IHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qDknrYOcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZAd7zEiXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vHakKMraWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g2Nz2BSzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dugpTWC8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tn1hInAKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eX6u4r2cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9LtVQkuxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qd40j6cLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MiS2m2GzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nuPiMskhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0mWKOmt1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hp1UiKv6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9vvx2bT7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0CAFXAyHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZUAlQtozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tbk3vb5GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 51SKZWliWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uXoOe3A5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func imGogs6oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xYFHvfqaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LLOXDarPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qFxDCEcoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kOdgVc4gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QLZcQdi8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jsqw4lxMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FnywC6lCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n4cBg8IxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q4lWt2D9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LoHhD0RKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y56DfCu7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z8VMIxEaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 77kMqQaSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dx72ExNfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2LTl7PXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eWLe53kGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HUs3dUXGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fnFZ4JnbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func re5kPKgpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GTcHYTsNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TMJGbcfxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fHEH1uw7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EMAd4wA8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PaFdvlZwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v1ZWUxlvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7QwAUDUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wyq9kApYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8ICYWJZSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qzaOVvRkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5g2He46sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SsMTozxKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9RgHqnBkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VqOe7MqmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func brnvmjSqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EpYGdFs6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NUn4o1TBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z1MC8fdzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ajzvzAwLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WgXcheKTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ewg93NQrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DPdHFdaZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gu7rjxBZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aBEFIEcbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kBHpMWVzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2ZV5KiubWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PgEYhioRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0n3WNxuRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xRkkZqN5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BxWCpwz7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BEnK8T8eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p2BBcuV2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s7iiNCSPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JUdZ8dkSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yzmc4DuFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1413GWtkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JMuHcsEIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XDLwsxmgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AbhSaq7XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BGBcAv26Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Si0ILuJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WGEwAk3SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bttnn9boWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VtsdIvlNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DY8zkutjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tPC11tDNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GH1aV3tfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J0PylEbeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func scmK3ceJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OoRIpfH2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vPBvLxzaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WGzW3t1jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZaqVMOCgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Jm5WM3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I55wITDuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WPH1JLvgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hJ4WKIIMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qvk6wxEFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F4KloBHVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AMt12JjAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SE8VRam3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zECQ2fYcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4pELQZffWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h7SrtJLMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N9t0YTwSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bx5qWfZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yq9P05uAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZwSZnXfbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MPxAb5e6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PQxGwgkkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fMGbHIkpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BwH1E65DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t2HEGnQkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LS7hfiA5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 71eYdB1iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BT0Jvy5xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GYzpDScFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GfZWs5IcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XSz7MIdTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TfVjVpqOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yNR16ot3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qSzVO9OJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O1hKadS7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ILLhEM9kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lb6yPQ7sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gQTAc8BFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lsXssiDZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hQSLd36WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 92PUqfDhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MODiVpeaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IM1TgBkJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zVK3bvpIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J0XGbHy6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K4UrQqGLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SIoT55L2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 435IDvwbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gdrdkAAyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AirdicHZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ylq4Y1ZZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TmLfI4EBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aTT8C8RNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WcfPU8hXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hgirxGsXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2Q7SBzmPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K9JZCgAfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pgA0YmZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G35lniLtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YecjH2zPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sGEXl7rxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rMjiH6s6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lo04j0HoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mb0650IUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func abELcrj8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XpPv21wzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uYCYBxZ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6sfbXfx5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Hl2u8O3wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gMIhZvvJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wJ2nbpeGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m5xt3WjEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fJWTgbjNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TaOu8kheWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UGQ3HjcQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hVm6VBmFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i8BZu8lTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pRS6LzpaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eMcdWxKYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aF8LDpICWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K5lIN6qpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jQBH4aj3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fE8cecdyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fthfWwPxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xAKRdXQ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jsgmMj13Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EscC6TkOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oQPLOU37Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k35BkJEYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VaybDhbHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oXmcqJw8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OfrRWWLvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j7iaqXVfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XquSEQduWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rWhXBYblWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6OBaN029Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7nyrWO8MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UW9EsF2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZLhF8YzJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BGmfBuv0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ut20KXVlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1A4AD8STWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vKmKVW4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j66Jx3orWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9Z4eJgkuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fvvlxzwjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9ZSqW2I1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G3CajFwHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NzAwBoedWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rC8lxoGmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tEin32xZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 92wk9AAXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func csQmvPj6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T0yotw64Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func juYFAIT9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gxK6TA9GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vssz6c5UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jTiuUHK7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sF4CUknVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rUO2cWFMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JcoIvCCfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s1syH11PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zjhhtThhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KzJrw1lhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q7vrgzWdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A4rFdqcRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zButkGiCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vRzArzUDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b4o2ZRlfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fiMlqHTAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F9RtqCYbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WThnxuy5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TqOgFMSdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OqZkR2VpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kviO8NebWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2KJmcwmVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GmXm5IMNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pIfsyAnmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func puk4UNUSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SaCe74lWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 43EpoxhfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zd8igY4KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qWKk3rOZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func swoGAtg1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yUvgoSaXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J4KzyOFwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kDA1fQ61Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SUBaZ6BDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EeJFpIB3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QIf7d7AtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hJANIqL6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ol63wK6PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nPVLEoxCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WfzdUxNJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ltcJySIpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mRiH83ONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J7k2B8XfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CeqMFyJaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FsqrPFdJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6TXt6kzdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ozyc7LFlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YFeyNOpHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z6gGbvjYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GSz6nc4MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lIA7oD58Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6LNjcS7bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mhZMUZDTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1tuhhwRWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DFoK35SnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kUdul3aQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W45Ur0odWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oKr2jNN3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8uQCc6qdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NGh8ddW5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lIWlIuRqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Om7G3JCjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lSDajIXxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AVxcthC2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MjIN0sjeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p14goKrkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TPs7XYf4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jLjV8eTKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gR3akblSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ktnNt6MCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qPgCPkB1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hrxuGT2CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T10W1jFFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mJkvV6X3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uEFQSuu2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BIDZCvJDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C1KeC4leWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EHi970sEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IGncDcaUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qvGf1knHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FQMOpYyqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YPFXQhUiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J8bMpv2JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ucHfxgklWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PPI6FViYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ihrqjB8zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ns7dsGncWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KHUJCwehWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0GlduMMQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NHgTmjtNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1bHHwKBMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cgMUogZbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func giDUw0a4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JMgeHVEcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qChNuMUFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func emfw5hb5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BcF0s5ZMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cGQE59YdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8q4e1UtZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BuD9qzmGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DkdLuKX5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sUBSzsCHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pXCkQDjlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BAzhL0AWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JLxprdkUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wH30ZSEJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qDOU3t4xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R5QFaeE5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5WIS2OBaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dse5jXirWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v7OfHRZFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S56CyaIFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func afieVSquWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yXJltG2wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sOmguWc0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 08EHjpPkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0B5pcrQ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nTQghwWbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7tk3djtjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4UwxXLZiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uRNfopndWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MWUWoBAjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MjNHxGJhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g0FWJLfQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OFsgXubRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tpqsvz2LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Pq8DPzeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rRmk7V9hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DuNbnEIfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vv8k0JRQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xJyBr1TUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F5mdp2PiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wMnuIMuFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XykCRCDxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Enwu2CmnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DLrshLLZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dAMvI1H8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gv2Jm3HtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UjlQ0Ib1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IuUweqAzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wYZGFyoCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xNOlo29IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gz0a4wdvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XHRadVAMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dM2vxvZ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5QAiF4XQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9nd5DaNHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9VCBPNd5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4c7m2INxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2JVkRlAgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 74yfhkZiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TeZyAnFvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AESp51EJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QR2gKbNjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E4Fo1UgOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MJ2HUBvgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AzQTqB3JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1ZhxMyrBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mQ7kN7zyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mDPFHjwDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yoLrhgGhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jDxcisd8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wMKonPhAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9SQPY111Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Db62SYjZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JEXgJIb9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KRH6cIVOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9AXuoCVOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9VzzN1CfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VLTuZPPUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6LdN0ZIFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XVEGO6T8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1wKTDUnpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P2f3XJDIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IiEd8mYDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b4MDe44sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zvMbc3eHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XsgojH3yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iNxHAS7vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J5lj1LxGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FxPaHcq9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9RnPV6UeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p3C8fizXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fYqKtU9GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fC0lnZfyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TzhZ1s6PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2xhvbxGfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n6bTP9urWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xMngHUC9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hISU2q9LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rKhFmhOkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R3FGRLUyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dgi63nq0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XardgnYIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sPFBWMF0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fEwImcxSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5bzfWOIYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0ZANWIVnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6UEZa2uBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ygWZ6akrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pD6AldnDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1Pl8KD8AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pUrc2IjzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SgwaFQetWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZwTZMFlSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3ObeRlTkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9aeF1QsdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UCLNWPPVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Grdew9sGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gdm2u1kTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gsWjIDB0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vBvBej9DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wg48KgpHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lCEvTF8lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OBv0FPMGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vh7GCb54Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tupIWllZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g5tWSaJOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3xBBPRTuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sA0XEsjXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KXtvggZXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LY7enVkzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KlCtz6p2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EcvGteqKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FU5qtyD3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l35sc8ntWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UHqPKtnhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 24pzyggwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H8kJ2BZZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6jNrIfRfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LjjFmsxUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func egRWKsXsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CRV5ty6aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TM3f9h9cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LyJP9GmKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lXlKklSnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zHfJr8XaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Tzky9vlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0tXPbqR7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SVwMjx3UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sNkH7wMUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SRey0ppUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0wGMuASeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cxfCvdyKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uFeLJhGaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ed5eIDn6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4GnF8qp7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YTUQnZhgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JP4Sykf5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PmAzg4udWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EHEivgT6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xEQzJPbIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PeGNwJrHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UD0XvzU9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ts5LstU6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vhkkD9fMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b67jiFLpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KRnY385LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9eLYkfw3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qYftX4CVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4EtWvVj6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GrwgH9R7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MMiZnSIBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0kNVrsqIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XgaO8PNJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qz6odWr3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A75Jn0UaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6c1weD61Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eCiqHf5SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zDre6J1VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mb98T9EsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oesG4WK2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gLMTb5ztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 74LPVne0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9L0WoSpWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aAEy4OppWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vhkaYGMMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func forrIdpJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dxKmJJwsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f2aV7SpdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5Cjr2pk4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pxxQ6yauWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HQdKHMwhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sWsO6IFBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AHV3D4phWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lmj7UOwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A1Tm1WauWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ISOsvMQoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xc3kn4eDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B7T4A8InWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T1Q9ejU3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9uPNOwnDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yVZuEFToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oh5Ig5rcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RcVuYuR5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QxdfXF7xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dnRVIWutWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xyOEqoU4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HDrCj5QAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zuNECL9jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FNObQoJRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nF5wNTK9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zMJy7YsZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gqtspIDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QIqPDoaWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HQmoxgQ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nyc6bdOwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func waRRYjo7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tlKDgcMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wCvSuv3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R4P5fq80Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RiszRPwyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vsYFhB9gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jRciML5tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x7tXwzQwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0mo6e7dKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VJ4hpuT5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tZB5NqS3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OPgKQU73Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oI3ka2x9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FrKK69zpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UUp6zpEsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1pLGWhSIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q4slah6AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SiW6kzAqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z9N9rp1kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EYqkT9DAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 282DEnpYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8AuIkmARWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FFFd5NHqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SOg2uMReWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sKl2o5OJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TKBDggItWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3P6N0UsmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OjZ6GydGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UNkn3xhkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6Q8540HYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hRsgqJocWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xG8TMEeUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aY6hU2YPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SI5KPoGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wpXZ3eqjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Srme3eCKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0Wf4RmY5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hxgFydfZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HjPzP9o7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 396fvxaCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lItbF16AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sm2gZWwvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 18WsrdpdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QXlSmOE3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ntwFPcojWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xMeQQX2IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iqjyZwKcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T2PbLVoNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FuzMAVZvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ON2Ju2f5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OuK1HlThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nlGYuz17Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3usO0EqbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yWBYPHB1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EajWvHXKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UY9s0tOXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KlbkHxE7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HDnHWB10Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mAb2cObSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rvHeCJ5hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3zuTJRySWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DM3Tg1hKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N2ynDfrVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E055VDSdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nlIpK6TrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z5yRx9JRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GBNzzB4eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z3TvQti8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h1l3sGqVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zp19iT5oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lKm4WvzCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qIr9BBMvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1bKBFdOYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lf7ScdOPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NVaMaOj9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZsiN3TYhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TKzkuJxEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zj0vjN0jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KonFEyUVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DFcFbBGuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ckHpoAF9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2z1BaopBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z5jaYeGJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gHvF37kTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nFamG6gtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HgwOkU3bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fzs1uxLQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XmpLs4VzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b6KGj6C2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nNKUZXjAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iM05JqcQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GuGirtf1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZJYTOrbpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bXXZgp5tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nRWcI4QXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sYIVdaiBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DCY7XH0iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WZh5oDHjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SV6MhiN1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QA2k9T2qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func axbOyBhZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eM8CLNPLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5cT8DmZBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1qnmjcnmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zFuFMsfjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vo3gairSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eVwr2z2UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PEJv2ke9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FDQ3mZIzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rnyKegKsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jZ2ZFXcBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mQjixSaLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZQEYmFYFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OC6r9u5CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1oEQ1P69Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R8VuACrzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 03g4MUrUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rBXsF7BUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XLW1WyDiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KscvAQd8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XeuptdtrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EZVmixg6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KxTKiIFGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hLhGceOOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UFxlHpTzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xdG8h6grWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func onxDhx9AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Kl1kUjk4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pCVIpev4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VZ8sTfr9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5RpUBsZOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RKoVgFLZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IFYIreE9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XwQHM9HgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kvTxZowSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mcuTdO8qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TTWYkWlTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dwgxz7YqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 81Xe4pqvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wcObz295Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wVDr9iTwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZY2GTAfdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4wHM5VArWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ov3DkWGpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C6iDouTwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tnu5FacvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IlrXdBkpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 37jgEGuCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8FhqDf9rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hRuxJpn2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HdlFyvk7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w3KxRmQWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vc4ohRZtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hIN7EpuQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IuzGcObyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KNucpbuzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DtyXmGp5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LnjiyEHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GDlaIkVUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zsxzm1M5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AvNm7TwDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N52lLV9PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B5YNt7fhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lO4tDqkoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func suWBfeKyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yf4RVdSkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cU3F3B9nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func viYUTrBKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oargi8WCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EzHfpDA2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uFx0LFe1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CCvYiArAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7IOg5BX6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b4dpbWEQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1xBvk7WzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rAFGFE6jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MXNb7vnDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kd67X9A9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fRByYt3eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aRa8Pci2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lSwCVoXEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YNXKp5zMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func coBKsljJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 83ZAhSpkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tEPaJL6TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S5UkcQSvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z6fkvZLjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nuu89yD8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yxGw3mzwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hqC1l4nDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Aam2xUuQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ffrXYmliWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q9DlX0sFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O6AokS7bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8VYubgGjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UZo6Bo0fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6ln1BdUyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MrlfXTFOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tziz3ArQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aU4iaC2fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lhf4jqLJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jy6GWmjWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gtX3U6wQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K1fw16CJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LBGh9jyRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hPitHxkVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u4hBYv1lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1yCZ3NLzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tovYM5HgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V1nXS5XLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3VlOpBjUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RX7jygHGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cs7sHU7SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oQPQ6n0PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uDhr3zD0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iuz9PEc8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8YJ63POaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TDmPuLZtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9oaWq8ogWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 13wF2yXaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qc8FqKbjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nSD0uBYPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SE14EPLkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PddcWyCkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JSi0IkNoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aGaNPTnKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n8yY2H9OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aXR0HiFaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AQjf0X4FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b2G8ISGuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kH0zo38pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8jBT7287Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vguAoMLJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dyUft7iHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M1bP5vMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xszeqLDKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OPEVb0J2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yId2Zm0eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uPJJgHHFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uT9xdP9iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HlzEnZAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iEY9gdcaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QpqB5Bl7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SoRm20UlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DLEZ8dNpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3sLRqvGgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4Kw244dHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wSBvucNDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Iv3E7JkSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5wmxHSi1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FtuAelNmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EmiZuxDCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g035fLv0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2NmSpU1WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eyIFJwogWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XXGv2LDoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 735uvnVVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 44pRFNNBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bJgNs8uEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 38531jFbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HNKN4BJeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func amSuCrIkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xkUed7voWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u7ONzgFeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func svk8hhb3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7nRvAQOrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mwiJI51vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hpA3JpoGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SFq7FJMBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QnN0kZxdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7kq9mSBSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zr7HGBaDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func od3y62B9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xp3UmFYeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0nrWMqwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9rE66CwaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tsijB4TfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xjesgrNvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N00anbXRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func svKzgBhiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vKYajzHKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4hhVSIQoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0SD16OiZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jk1dfBauWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LGTqbH7bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 31rnVukWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cPPGny4uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9MwnSGR5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TgHrLzCnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PRDwXeEEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eUpHKx4nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xD7uZd6AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eXfCSbuZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T9jOQQdZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wCorYH3tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SYyk3tEzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WDOw4s9kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lLMzF9zSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9pAO5xnJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SUHYuQ6JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uGZKghyVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XCw4FNmsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KkiFJZpcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IS9OD9zWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ga4EYar4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0pdZjF5mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PA7FjwbRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vnFNcxrRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lrdcOD6NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gNOP46a5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 46m7x6HLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Txmlv1czWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zST6GEkaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mkdbc7pMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jtcPo3jlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wXIDCQr4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q9t0DLXpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZYpt2ZGyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pYJIofPHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FdiTHER2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ryTxeRwSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DVAJR3TSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9BAYOpp2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fqd6Zsn3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UCCPBArRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uC1GHg0IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j4HOJhcoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j9mxdghyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HxKi3em5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HJBakTOlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pcDDCFCEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PnsnK7oTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UDQysuWzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fSIJRPynWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gpjOQUIaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z6C3tnRqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9S10P7fEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3rKB511LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3XqMPux2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tUlMWuSNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LvJeKFpPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9To3orePWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uWEXYCqMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VMjC30BmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yYZBnBAFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 07DOoENdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ARWySgliWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XBd7DISJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func glYZB71eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func shtGcrlRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lGrVUEZ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yFNgycNoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZREGHs09Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BVrVa3FuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gT4NDWDDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S5Pp1wigWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lSfVG5pyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YiO4bpZ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TDrcxKYDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qPLJfGdTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZyueE9drWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dp3cAT4eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GcCvKYy7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M1F1v3cOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 34Taz6TQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PE7MqJEnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YHLZx91GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AHDJgYWrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VMCg7bopWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0oH6Q8uOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lwes5rC0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d99ZInW8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ziQQlo0YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mSS2iQMJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rPYIbGzzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M88bZFXtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YskmSkInWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xFwkfzaJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pELWwOsnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xOqOxRtcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CUMr15wEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m28I6T2wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7muIdrAQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gtIjBtxeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jq4Kd57PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WiUOh0lkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rYDubDheWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3VSZh7FDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cPMktaiKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D7QETB94Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HnoteWhFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Eojn1xDYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 86IipvisWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YgPABFuBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OHgXvOX9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vzTP4orTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9T0ciKVjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iPFsR6CGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hwAVDR0oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mrrOq4jZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cRdak1m7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eXvpEeB9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vFsIWf1kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tH40Wfv0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WmvOWkcMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gxaRzh4eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LTnmuEOVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lIkxUlruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UE8VCYqQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j4ZGmx4fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0fMTWp9pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sB2asODVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nX3rfY5nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bryZtmToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vuomm6lQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jrWkTSj5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X1BM8p5dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CwNorQkpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YuYrpCrwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gK0N4kvuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jbuoB8bjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9k2OASqTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EuAwoXk0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IglrCp1MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5On2DjDRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gDiAZwZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZfbWKxxTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q95kcQTGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gS4VVsE7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MZD2gKBiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VYFgdo6nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D1FWAgNGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P42iQ8PdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KmxLYKVpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DxanmYvOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cj50MGYNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QUlNI9n9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a5qbGgYAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4Qo542MKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GcmdqIdlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dtNATMsGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WqAr9K2nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5zeFfskNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 68dy7SjwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KgHCfKKyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ENw4PIDkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SRZ2no1zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OjBCbrFzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iPabUHBfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AstkHcsmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3P02XjurWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nUeKyd6MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lPI1LJfnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ncI6KpdBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rf78xUPTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U6hS6VhcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7QncT19wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KLegb8umWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KzjFUAL5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PtGlcCZgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DSWZBBttWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0NBXW06KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cAXLU7LrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 04iG9KJ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W36ulXHAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8YmN3A9uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vtoLxVt1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W4xYjf9RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z5clcTidWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IeI4j4kXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QQkFRJsHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4GxZOa18Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ScobSzFMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ogabmO5LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sh8F8WNJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dejFNsKiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gy0TrLN4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6dt7SCuxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AdrcXMkJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mIGgx9UHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 024YH8JvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7775saYWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FRUAfNsIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func raL10MxWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u2VoingfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B4eTiPTMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 87Zh42dyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9giVqMnaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YKjT9ewlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CaFGRmVFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lyt7XZinWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hN1SNFb1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zses1Sp2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7bshw8naWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zgnrP1jzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nwTxQ6hKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xv1pdvXTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ppLUHPaCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IVUnkKPVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B9Poj4urWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4tu4e12AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6gDNlNgHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hKSqs6F2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oVjyiMaNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CPsCqBcdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 95qYi9Q0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jNFfrzhfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HJiimc1KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 60ElRhevWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hgN1Wv0kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AyTXasq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dQUWHTFdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V8wulDGwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eHt89kWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pG6BSfYoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YwtiIKhrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5Xx0oftWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jDVQu2TCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SjzkqvPpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hCbSrAR2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mXaYlVZKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0SP2U8Q2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eAZAOoDhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5ZxIaRnvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BizKI0WsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sNmUn7pQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 29mMKdVUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XRz0oV14Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c1hBmUb2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qR5BMXsgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mHLqlgWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qS8ecqZjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bWDWxYWYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WByhKujiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H3OCvmwCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CoFDhWXOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7LmEQMBbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tgkppcq0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6vjH2Q2vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7TRHceHjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4VcWrFlWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hyWK5cgmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LjuMjX1LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ooz84kdIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xOLKUa2QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func feNMkHuGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OTntXJFNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a446GU85Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4akplbdFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dMNWZGQ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9WDZBBfTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vs00iI5HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yOhgXOb0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mlJVncBgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CkcDt6fiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ic4Gtku5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DIc1A6CUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WDlyPDFnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v83jaNcPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BfGjenqJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HONQp75xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i0rj3OTzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qRfmDYAGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IiOB0nPrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CckJpxQoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KnI34eaHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func toDyWrMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yDmZkjfVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tgxvYWlHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bFLvC5WqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4pyYNXc0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XMI0XGF4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func opW62lKpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3UFCrxzmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gMa240ZXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ycODGC18Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cX0jn9JVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qeudjar5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 48DLz29FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rbPcjoWCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DcMGDekfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Oq9lbdJiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QOX7eV73Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WlrMaZPHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HqOfmZh5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L0s7eld7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9S1nrGA6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JzVZR9R6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3TtT8WaNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QLvSGwMFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qZaeho53Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l5tFXogNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NadOVFmCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iPmL5vjXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5LSa32seWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3bocaSxbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zQhCAzFDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3zaYyg93Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sveM8Qe6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cgNUSw9OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZuoJWp90Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m5U8byjbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func drTmpwJAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s3gCW9V9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DAg3lfibWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sfDV36JQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VoguKLilWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XvnVAF2pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VnlUMKFkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RvTeBLjGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wWQXknF5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nuRGUfweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EQNMQt5WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s4iLhI3TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QTGYMtVMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3eLvbduXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mUs2iNnBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wb7Lk1uXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NhnICJobWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EU1cFEhiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Aw3vrQ1vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SXnTDK9HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qet5lOCBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cOLcrokQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8GzQUgDMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ji8IGFM5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ipsqL7fPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MQpCZDJ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sgUe4SDQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WqaS82NPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tH9RzmabWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bZdEnGvcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gvGljJB9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GYkXo76wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FhEuDjk5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K7bJmgUuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UQvvtHfCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xMSGXpafWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BeN1Dm0TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U51ydWcAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KMumM5rqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pIN3obbyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jks2LfbwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oy0p3S2ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7dP8NlbyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dNnoJyiWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 65CZ70ERWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jJwakjsXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aAkETm6eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G81Ubz0cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1Qyh9yf3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zPfuNJe5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LEIsxzNJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xerltig1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tagkkgw8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2Mj74b9oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GUo83BipWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kLk6CPGvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ikAWD2imWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w3F77NUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FJ8ATaF3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rUhGlxBfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xtosaJiJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TQ7COH1PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pOiinSe5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8pPJ1sDbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4djAYxJYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eq81i6fhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ufOf6c3jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mkvT4C3RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r2ZZ6SvFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ituyx1nlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Hfn2Zh5ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k18QpztUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eJiyra4JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GsE6dXsyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RvQiyOtuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func re2X3LqJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d4FzrXS5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I1irHY1VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sa2Dn7p4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pW1wb5ctWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vN0PzqGXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dBipHqeRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lsLFlZEZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gtykzl9iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L9BJxii6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0hcS2E39Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func co5Oqhv9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aNENlVhLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bdavRv3iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yf8CWl3DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EPuaLCiBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 99wGWNdxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DvBnbBHoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N8ZfU5wgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vy9b7KuZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fOaXywfNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N1nV4lpuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vmb3OoC3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KSgjaX0fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X5OgDfUWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rQ4ldeKbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bkAicAYNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R0tD4IuvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lVH2H6aYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 74y5DSrbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d5Qwq9oUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func etjnahOmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZNXlGmENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cXTorgDHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func elPm9joOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lKC24G7oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dww3Kd06Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 378WtAwNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YoTuXTVaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cUWeBbOsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jdG7sIBoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ch3NFnnpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VIILM9EyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FqA1NWiXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UREI7HWtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MFIMhDdRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xsGC9bw2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FBho0auXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 78ugpMSxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cTlVTgFTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OGSnyNvhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dwXA7nsuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nWIo8nfeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bDQLt3K0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hy9TL6QdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TopRFmQTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Os7c5tNsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xKMRgDfiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QsK6PygEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HOmbWOYlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wDOcuo2dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aUhnigxzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wZ2vxKJVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YLAciEF0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rcVHaVBaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T9myoQMUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OtvVBTR6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q2BTzuyHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P499PqIJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func seVrYPlCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n14Kwmh2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MKzxAGCSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DrA4lygHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lpUMNUWYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c0rbtLhIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jYlLKFMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ly2y95ysWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lTTIHktCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y3uH726FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BQcvSkuNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7oEJYDtmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uTH5kh3sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5EH4XtB5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mbi0u7nBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SMgkwxY0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wxhEESCjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OBn2kBndWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7p4EVamUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3IZByCveWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I8IvJf8SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yw0M6225Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nh97Y36MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z1qSr1jJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0R7DygLRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wa6SpMmiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func omSlJ2DRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func malGaOCPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gKLEaBj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f0P30Us1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y57cbrAfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2CLtMlNaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VpphE7OOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bi9URc1rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SoRS2erKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZfUL9DRlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kWx3GbBlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UX4rbgFhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gXi1BwUHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ooxFwntbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fd1sLJspWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qP4YLmGbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pmKknR0ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zseWqzNgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r0tq7dX0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X4ZgTGjrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4rhtxMwXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AM3ipkIqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Om9PJeeyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func az0bLRpIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6SFsUm5JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OTo3ATyvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func io4YP9C0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1nwDOGDQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yEHQEu7yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vnIQ1LvGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qGmMteAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Aaf622hBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sGMbXCsGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j18fLoKRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PWxzGDl0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QY1k9IsjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TC4wnCL0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6z7keNWmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zWrKrZugWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KQjkj6qLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NBPBF3cGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ou7uPV1mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pvjB43RVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eDubrBKtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ud5MuamAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 31btdT4IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F5nsim46Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iz7hBDP9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R8LayfFcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rqy8C8STWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Oq0xTBSPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oOFmAuV4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WWnFZmP6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6Rp76K1xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ug8LEbQ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WD9a37T2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2WAjSNIaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AAUDvTusWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W25gqXTiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0rKAOgG5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t5hwIPAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AX97BCiTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SY82qu0MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RTYNIYONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KoaEYrkaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oUvYKIsAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3BFprNN0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hGZUwjDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eCAD2msSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yto1KWE5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6iPYK3PxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XcisOlRpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VMe0wdZDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SaLakvkDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uAqKyNrAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VUm0ic8mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ToGTysBoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u4nduyvWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7ZmnlJIoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sl9n1uroWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x80VMvMLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q2XZBxf1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u8HCKpWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FmYREmuTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oPI4Zvi9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KubeEXFpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wfTQxEz4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4CIRWimiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dueoQ8TcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b6nUf76gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ARterruRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lhCUfpJYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yWNgBnWBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oFicSUkUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func edPeNVbCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WUFDyspWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DMYFxsGaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func whZtx41nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GlG2SCCgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O4IfrE2cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mVDq2qbGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ylJ24xTqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5E2snRx5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func anuG2rUFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mYIBSv08Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eBttInHKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oOUxBhxbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y933hXDiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6R0avLzdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rh7za5ERWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c6OnT4x9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YDQ5wmTyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JjVtWvPFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vfYf9HcDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UxhXPvv9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UU3uCGGTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tEYnw5ufWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jLb6vsPeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wbjxv3XiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wzUp6r01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yn9LvXDlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OA5hi32RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pZf2hB1hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZtCTPUL8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p3mkBD3zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mkfPQLENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iJ2kFplXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EDEbILylWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a8f1Q2uMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func voFNdiPaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OH7OKXAtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wBKwEJ4OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func reftB3nVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n0DoLCs6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TyDGuAt2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sFE3O0awWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uo8DbQ2gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PupA5hx6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func itIA9CRIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E8AxLUbiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OtJdJ9xeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6Ud0GtMVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4hkWvcKBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LjtAtbppWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JE5wclRFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JiNwY736Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EjqafBWJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4rTc0L00Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U1GrGs3eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b2Ego6KAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gBw4WqIsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PIIWn3jLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G0PpqkkMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 71t8VhCAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9BV8ftRzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sa115uXyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p26VqEJuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hfzokyuuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5QIqgrq0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DfiHf1OCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dJjIVQ6KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v65p6PEwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DX4FuAsYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func obY8Vm33Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2B6KX6PXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ULYTNvsxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xHU2xEOXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WMxyFpzMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EzeuABtPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r1ZVngEGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7UxTXf1iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p1lEmBIoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 55PcMsylWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func psXJ6jIwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X5ne1yhQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JKdFXUElWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zUVX2xsEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ItSrOKWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 41QYeGzdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a0QtZHizWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pd7GhJVTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pMIesaRzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9Ms2rYGbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SJTuXydQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3UImZbA2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DNhlxqhVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PZOERsNsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x11Jk08IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1kLJKN14Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W4Z9wvl3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ofx40f9QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func teeGaVPjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qPuyFEwSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I0TXioPNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CnTBBlTeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DCMYbpM1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VcfjPvz3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yCmU0cRKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OMqZReqxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZzWDbHpGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LNuvpgwNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hPR5aT0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5xy66NayWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vz2qAKkYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hS6wWjusWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9WFPAaDxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8utrdaUDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SvMPSEFgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XAobUY49Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HLvm5yXyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cTFmt9aOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1Se6bvQQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kVOJGhhTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IOl9gHfHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K8mzoj1lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 84La3Pn2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hfdK7QzYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cdRfPab3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 46WdsZr5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZIR7HhyDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o1Coh2YzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uqHeGjEDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jOmiKsYQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S3muCABsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p81gaMrJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ich9U2tkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aR7r73QCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kTTIMXbWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sK5SjnqGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WzQZqsDeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RQYzNi8fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wllEZkwfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yweYBHSzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kdiRd4XpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tp49pTUHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eONEslc3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cX9HSxKRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0I3NAKeTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bZjZmMNSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rz17sGdhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RSRgzcrPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

