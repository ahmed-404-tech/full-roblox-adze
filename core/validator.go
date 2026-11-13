package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "5.5" )

func S73tqdBgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W7h1TL2VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wQyP9x8TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kLGhuSwRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QWlq8mNOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qf4zCDlvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wPD6UvEWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xoUStVPRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i9LN5setWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hx4VVtfBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hqY2eFboWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zGo1k7jNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PMcNEABvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K6OSSFoVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yHwDtIIMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MRP9cbcyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WuoGlmAqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AZP7mZDVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MelueCadWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fOQfwhloWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hFwO2JAKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8j4I7vzdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zjyWuOVAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yb44zuSSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0efwdlvRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kNnB2d0dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BoiHbDF9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CWOYwjrqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hj1iO9jaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vl7hO7dYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VvFkTHblWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eKGClrJtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UAC846SYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w6KE8nwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bNIYxedZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RCA0DxhiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n98dxUoMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iET31NzGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func McOSjQfdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func baU4VFMpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K8NFwrGzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1jeTAkE2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fpbeNxXnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GohWgjUuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c9zWyM5OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YfIoJB5rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bh2cmtiyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 45bji0r5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uQDaAx0OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6iFwCuNJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tpgXfRbKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pjqvnltYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ih78OnS7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VtJ9cRUMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LnwHmTARWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9wfQ43zKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q1EHhXq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UVYYVYCKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vtvQwwKPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fMrKflOrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NSuSwAavWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YR2A4b1FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BWIpOYfzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gxQ3Io2VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hdMONwZjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B4nwrwd3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0cekIenAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mIBljUKcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L4rY40iCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CaeTgBc8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iBihSowHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9XrqaJHyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gQBOwjsyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HpZkJCCpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uQFpFEX2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nwFo57n7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hZJoXHKWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ku0LjswgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JwyohbaLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 65USkI8rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ipVtwApLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WmF1DkGmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 69mHbbX5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EoP2UQAxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pq8M6Q2WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 93tMJ8g0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rfq26Xr6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hIao5t6DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GlZ2OE4yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2RZtWHhUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func viz15HudWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mTS9CtG2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TCmC0UsrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wLNsEd4KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bS2UHi1pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fMlG3HzcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6uCtW5xVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mVgWQ6aoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V5yjFr39Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hrf3KcPZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hkRiZdFrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lAw5htobWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ksnYDgsnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zWAYsOrKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ngWpwre4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zAqRyAoTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lTsfulWsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KFqROe7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SyiqfAVWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HPPSO8vXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rytWqLx5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func twrMWVhmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3TbeEPIGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yIPEhnoYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bvY2qSDTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LLXI6rvjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fFeuNDjmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hsKRucVnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ekOpQnmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G7anxxp2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hGgo7fzHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vApzozHvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nG7am8FXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jXcew0ZeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IcQQuRdYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h6qDkxrFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func INPkbD8wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TetPBlrDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NKpSxjtQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QXvbHz6uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aLUbq231Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AKP7Mh3kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y4U3ZaOOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cpQPzL1uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mcmJfvczWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fBiR7ZTJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bqyB3SweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dk5LlpMpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gc1PuELLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FOkYlniBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NeAzmCpaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3BAtHuO6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A4IuWr6iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4RXguRntWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IlTf4mr7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9kL3T3OEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M2qEbl3vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sqmgO3hQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7s7MLFyQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2id4oRXmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ChySOJVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YmtGqbxJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vel3bxDNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oQhXq4yTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TCLB4hiJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WpR2CRsSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eL0p6V4AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mjELmI7iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zfX9ydOXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rtq25IOLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M59GS68jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yzIbgb0lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dkxbV472Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ghhhqcwxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QiUGqW8dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VR8UdtTGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CDpXBLszWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bZNj25y9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7f7hbV3HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mJb9gF0YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SfWYvGdaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i3vJXIsAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AcWjZuiqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h7YQDiG4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Inf78i7xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r5TcPCaiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rWwHCxQEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EGkwNphLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D8KMPSRaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ELjmTztBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UtWCUj9bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d7dQEICjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AWzL1mqoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EzEsh4kLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u2bgHufiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7jDNa8gRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h94NfQemWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GuJ9AwDHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HdFR2d4xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1nVzeaNtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func soO3b9hBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FHIdbNHvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qeumzzWLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fCxZWTriWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5LnpB8LXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ECpb1mB4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8AN3oFfLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7GdJHCBAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IoiKvjoiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CZO8epKZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M5MSabMnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M0ceW3icWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gOAzo44xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uht4bMuRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JbBHgLrfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XR7ODQOvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2CIPICzSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PgAAREh5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o7Sqn1s9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HpnclUPvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0hn9FPyTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AfP5n4LJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lBqs58AeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KMhUtw7CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ENh6YQqqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gUJdHEj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FGca1OB4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0sQDersCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fjg4FcZAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tqAMR0EiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5S6rrYOFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yOZhhRhyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H9YawT53Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lVR8Hw8VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ucMETmJ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IJq2Chj6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CZElF8hyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wfJpQhjMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CbBQ7mVuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oboSWd8MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BkCPqhFiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YPa6cmSfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xK3dWRzmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RDqPCaqvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f2kRighSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qs6EQhk0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LtRZK3FcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XnlfIAKeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TQ4NZuziWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EpSYta4gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ElXNTRZgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fq76g9e8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func scmbYLq6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VzT9A9dxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K6LbJSJ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WYqXs8PhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qXArj9lTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FqnLSfOuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wNEuegIPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hvmmA7I2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RSOznNx3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jbGafG2HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9yfBC2SfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YCL4VFRwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uzMc2G5tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XKh8Fl47Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ccd5m8TrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DLtTVyQrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cGyoQHdAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0aTOGswSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func imDuLh5HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ub3axJtFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F0pdEYUnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E5bMpzPLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1l58xTPoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

