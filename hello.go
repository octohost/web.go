package main

import (
    "github.com/hoisie/web"
)

func hello(val string) string { 
    return "<html><head><title>web.go</title></head><body><h1 id='title'>hello " + val + " from octohost</h1></body></html>"
} 

func main() {
    web.Get("/(.*)", hello)
    web.Run("0.0.0.0:9999")
}
