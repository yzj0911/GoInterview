package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"sync"
)

func main() {

	var pemData = []byte(`
-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC9akD8Sloc+r1H
daU7xQoc5a91z4zLj8OIQuIVS2u/KiWKuirJpg0vEj5rM+mXYRU2D9zHuZgVdcJF
QZN9TVfF1opLDWsKQc3abSWKH1UolFG6J3wg2AfkqsWlYNY4xe7lyIiEvVD37wCf
mVleUDixN2DWA09ZMSejBesyrPlVh3/WcpEZJSc5g3tq/cgKeNjo7D1JkUy+Y3dk
hkRVpUfyoc4Ja/1MNBpnilgBpz1ZqDMoiGSN9iNxNY0YMN5gxMCAv1mdxpYYR6z4
AWgXM9XfTVJh8W07w7TnC6a5ya9htj2NpANFKibooCHaoOiDrNZT0xLCYSq6fevh
2nNUIJRtAgMBAAECggEAARzbrt+88cijTgUp6FT6/zp2Mmn6uMNqcaV68IcV5fSY
bd00HSUllK41wal3aNVAK6YiNOYpN48OihudgzHuHxJok5JLm67RR8Q2YT5X4Lom
VMKgnwzF1xkNui/8ci7vfVgVPTpjfGFiDo33EW/FNX1Oy1MPp0V7pyCExjJv8Imt
pjk3oW87Wf/H7OQ5egqfSN2ou2KVh1CPBKqlXktbFMu832Wd1BTfSRJjvtTlLON0
l+HgKazQyIcAFpHGC6DXo/ZjSimm97/LYo6cPugXXU+TdqpJmJ1MO0fc07OSneJy
vBSUAcdVhcxyJ/lnUvqHwQ1631J35ZP1M91+2SodQQKBgQDowtnQaLs8uwrRiAge
WjwUVqtIpRfH294DI2Bp7sMCTUZaSde3XCWZz5YVNXbKjZ5yfJHaGfgRoKNZer2d
AoFplD45O9EYkjPHD66vNbPJI2bk4jJcdveoBQlerP7h70/R5TnZTPdv0jTL8ekB
1jkw/CT4yAYy9xifGlOrMrnGsQKBgQDQU4Tm0mq5Ip9Hsaee32HJkPYV3B56+mal
T1W0ik8SWFhZD7YI9uexK8kpYgngp7ZQvPCGwyJWEzVdysH5SbmE7yUyon8LjYzT
uHNvZO7p49X0UggI2s96NmSeyAnCgPxIutMqZa+Fy01xjpR2DKJeeMVsj707jkwt
pSaOgdeQfQKBgQDSrzSrQXFhqkhDmucGWlURb9XAfrdEz45otsfZeyYG2l148mgQ
75aVX+IQtoEdHQ0zwe/fRCxYAFh7cO9axF7Raz7bXXqJzCST5W0P6QMgaCwFt30w
Vvsamdx+VwarCYvtiJhRSiqai+IATKrFX9wKq+DnU17RGPqvYQwk5VhlMQKBgGs2
kG86WzJsXwzGoT1iOTFDKWKWphkkRS9OZQ1FIOyQCufK7iQu7Y6AukZR7kNwDKQA
mMjCJCmoOQ7MCogBKTkA2mP0vO11K8TKaJ4rk8lLOBFFJl1oPt7mn2IYEO3I9A16
GLL5Ihv5RSHr/vvCBM4Z2YDFeN3tncbf97ffmtEBAoGAZbBZRx3xSfx4Tbr4l9jy
rOYgwNJLSpa2VnT7L0N+xNPZZQ7wnma+ZXlalRdKYIDVJJiNJI1W8CZzxzfJmhgU
1Psu/79mhSomd+s6s1K1axxr0Cmx9VVrFslN6zAPqOlymu7c4Q83faIihRQfUeUA
qgOXGr/oGOdUfvz0VhAXdek=
-----END PRIVATE KEY-----`)

	block, rest := pem.Decode(pemData)
	if block == nil || block.Type != "PRIVATE KEY" {
		fmt.Println("failed to decode PEM block containing public key")
	}

	pri, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Got a %T, with remaining data: %q", pri, rest)

}

var lock sync.Mutex
var wg sync.WaitGroup
var s = 1000

func add(count int) {
	lock.Lock()
	fmt.Printf("加锁----第%d个携程\n", count)
	defer func() {
		fmt.Printf("解锁----第%d个携程\n", count)
		lock.Unlock()
	}()
	for i := 0; i < 4; i++ {
		s++
		fmt.Printf("j %d gorount %d \n", s, count)
	}
	wg.Done()
}
