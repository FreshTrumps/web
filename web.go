package main

import (
  "net/http"
  ghttp "github.com/gorilla/http"
  "math/rand"
  "fmt"
  "os"
)

func randomCage() (string) {
  i := []string{"Lh2VJ4O.gif", "0fQk5qs.gif", "GKi89OO.gif", "PJLRxyQ.gif", "Z0CfLVA.gif", "rVs2s44.gif", "blCmYiu.gif", "QtFbQBu.gif", "s50GnoI.gif", "dlqj03p.gif", "qsxHrHF.gif", "FVndoUj.gif", "ivF8q8L.gif", "ynrXDiD.gif", "sZA29Pe.gif", "hm6BVL3.gif", "Pnj9hiF.gif", "4MbTxN9.gif", "Zx4RYfV.gif", "hvRAEGF.gif", "YCOUmLU.gif", "IvKeQ1A.gif", "ZmvXumi.gif", "6ZiL4DQ.gif", "Z0XVacD.gif", "Yb8oxxC.gif", "rZT6Ohe.gif", "F1i8osw.gif", "gpwLxSf.gif", "Th9iFHU.gif", "tJ0BQfL.gif", "7CwEdpy.gif", "BlKuIY2.gif", "drUcoRJ.gif", "CXPBe4A.gif", "Znl2pPb.gif", "GOo9GIR.gif", "EzKoTuy.gif", "e18ykmu.gif", "hlZSAiy.gif", "AC8sDLj.gif", "LWZzTi5.gif", "zG6H7r6.gif", "BVsUZAp.gif", "XECKfqm.gif", "Bemwu1u.gif", "AbSSHHM.gif", "hZnqF55.gif", "FCq0Az5.gif", "Oplh1lG.gif"}
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
