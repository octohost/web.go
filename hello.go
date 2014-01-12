package main

import (
    "github.com/hoisie/web"
)

func hello(val string) string { 
    return "<h1 id='title'>hello " + val + " from octohost</h1>"
} 

func main() {
    web.Get("/(.*)", hello)
    web.Run("0.0.0.0:9999")
}
