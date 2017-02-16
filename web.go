package main

import (
  "net/http"
  ghttp "github.com/gorilla/http"
  "math/rand"
  "fmt"
  "os"
)

func randomCage() (string) {
  i := []string{"Lh2VJ4O.gif", "0fQk5qs.gif", "GKi89OO.gif", "PJLRxyQ.gif", "Z0CfLVA.gif", "rVs2s44.gif", "blCmYiu.gif", "QtFbQBu.gif", "s50GnoI.gif", "dlqj03p.gif"}
  return i[rand.Intn(len(i))]
}

func proxy(w http.ResponseWriter, r *http.Request) {
  ghttp.Get(w, "http://i.imgur.com/" + randomCage())
}

func main() {
  rand.Seed(823)
  http.HandleFunc("/", proxy)
  http.HandleFunc("/random.gif", proxy)
  bind := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
  http.ListenAndServe(bind, nil)
}
