package main

import (
  "net/http"
  ghttp "github.com/gorilla/http"
  "math/rand"
  "fmt"
  "os"
)

func randomCage() (string) {
  i := []string{"HU6xsjD.gif","7OG2fIf.gif", "SHonBbB.gif", "PU905S9.gif", "fyNnZp4.gif", "BS0ZEVI.gif", "3RrRn3h.gif", "f5d1Xfe.gif", "yn4ZfQE.gif", "SQQyvc7.gif", "dEvEDSu.gif", "4a7bMrb.gif", "jB9Dhj1.gif", "KczdvZO.gif", "L0cFhcN.gif", "nBnuTKA.gif", "sVQOGnJ.gif", "2LYEtiF.gif", "DXysaDu.gif", "VWsTPVf.gif", "r4ADVzs.gif", "h0F4gSq.gif","Tv7OQlE.gif", "qOFISNE.gif", "yCRnTM7.gif", "SZr4tYC.gif", "F9Inc03.gif", "nKK9ypc.gif", "9GozEe4.gif", "Xsa20sk.gif", "uWXehlO.gif", "6w6iHuE.gif", "7WQ8QdH.gif", "61VFxMk.gif", "QvSZsm1.gif", "LdQ0RYc.gif", "kZqTgje.gif", "Vvww1bs.gif", "5j4tYw4.gif", "fBefOV1.gif", "QM454Ad.gif", "VAlBXwz.gif", "ydwMYzU.gif", "UMasOvJ.gif", "vvoDLfa.gif", "AcDBbIc.gif", "KYrG5B6.gif", "34bqWoT.gif", "SyE8y1W.gif", "mW9kYU4.gif", "yLC5nwk.gif", "asCYIkp.gif", "hRDHEmN.gif", "bocb1fy.gif", "Teope2r.gif", "m2vKPhQ.gif", "0j1ACjv.gif", "v9cJWKp.gif", "fxQ6eJM.gif", "unD3l0N.gif", "qh9e6lH.gif", "9DZQt3j.gif", "2ANvuVr.gif", "Op6M1P3.gif", "kXAhrR3.gif", "cFu89aE.gif", "iaITdUA.gif", "rK3pZbh.gif", "Lvpja57.gif", "O87V731.gif", "aDcP9gK.gif", "hxQ5Mdh.gif", "rB59gzS.gif", "F92XLqV.gif", "RFhFxrS.gif", "8G4cOzG.gif", "LybCtZr.gif", "J2vBcXr.gif", "zlY0P6U.gif", "h7lNZs2.gif", "Ps3bGwU.gif", "3OGh68I.gif", "l5aaeBI.gif", "BmATpoS.gif", "xi0wMLS.gif", "qUsn0zs.gif", "Zw0UqBr.gif", "7CN476X.gif", "cyCUBr9.gif", "l6NtJxz.gif", "tdfRS8A.gif", "uem9663.gif", "8LfbDci.gif", "8yIsCWI.gif", "eV6u9DH.gif", "O1KK3Bd.gif", "Oz2EYrc.gif", "ZJa1yv7.gif", "5yHWz5g.gif", "ZSJY9uO.gif", "Ha8fjnA.gif", "xQqVyU8.gif", "l8dVk1H.gif", "aQIMXiu.gif", "ed2Bt7g.gif", "RpiqN25.gif", "4SErI2k.gif", "N7ltm9n.gif", "j9gNGln.gif", "DJhlwKh.gif", "eF0weNf.gif", "sPz1wdk.gif", "tHj7ZAM.gif", "AIw9oVp.gif", "yG2SEDa.gif", "zIqL5VF.gif", "bGp4Imu.gif", "3oFN19f.gif", "gGl2WE6.gif", "Qs5sW7g.gif", "7fifBuo.gif", "6PyqlfX.gif", "ySAakbj.gif", "fyXfLYT.gif", "mMHOcKI.gif", "gCbLTeU.gif", "CoJlDiZ.gif", "6XAjOye.gif", "w4PIbMT.gif", "BtbHAYz.gif", "G57qDzH.gif", "S0NTLhA.gif", "XwlrGn6.gif", "N3YbZQt.gif", "okr87mU.gif", "fPqAfi2.gif", "UT2WnkS.gif", "fldLpe5.gif", "kxOBsnX.gif", "edki8Wj.gif", "RJoEgaj.gif", "G81FeHm.gif", "fdYIAAk.gif", "NXfw1ZM.gif", "kyuoCXu.gif", "i3roGc0.gif", "IKRbPAd.gif", "WzpgffI.gif", "dHrsiBs.gif", "yBrYZhT.gif", "8028O7D.gif", "4tXNJo8.gif", "3oU7xZJ.gif", "70SfSQ1.gif", "iXar9.gif", "0kAX0Mh.gif", "tcgtCr8.gif", "0Vg8G8M.gif", "SVEh952.gif", "aiwzq6K.gif", "5uCLN2l.gif", "pq0Y3ww.gif", "4MSlO3j.gif", "R8vFib3.gif", "WPkeiRv.gif", "3kcr7dt.gif", "pbNvjy7.gif", "BJd9OPv.gif", "PsCyRtX.gif", "fTQCfNV.gif", "TITHCQK.gif", "73kPOYV.gif", "INU4O7C.gif", "vhvmiSX.gif", "LUMscJS.gif", "NCgwvyh.gif", "Px5FHPR.gif", "SWAtpkJ.gif", "aC8QP14.gif", "wrpkqze.gif", "voLnaiY.gif", "9tdFbhz.gif", "vOMMe2D.gif", "JTGWsKh.gif", "Q0zAFhC.gif", "s5BMqkS.gif", "Kjv0L7x.gif", "G2pyTCS.gif", "j2wVO68.gif", "KSHPBDf.gif", "RGUxiCn.gif", "PLoHlx3.gif", "n6J9igY.gif", "Whzkwup.gif", "nqyjZsK.gif", "AAIhwqb.gif", "us1Rd9F.gif", "eaaI3GJ.gif", "C8OSxDi.gif", "Srmf8NB.gif", "b89kUsS.gif", "vDz7TlC.gif", "fVTydpZ.gif", "aSx1z8h.gif", "yAJLqVU.gif", "WS3ABt3.gif", "otuabfM.gif", "XesLBlE.gif", "cyngNXi.gif", "jqNIoaO.gif", "1iP0Ui8.gif", "1lxwfRA.gif", "UZvvZ3C.gif", "tTt1wQs.gif", "XihBank.gif"}
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
