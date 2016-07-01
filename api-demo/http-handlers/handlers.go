package handlers

import (
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"

    "github.com/gorilla/mux"
)

func DefaultHandler(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "Hello from home!\n")
}

func ParseJSONHandlerGET(w http.ResponseWriter, req *http.Request) {
    msg := "Hello from parse-json!"
    vars := mux.Vars(req)

    parser, ok := vars["parser"]
    if !ok {
        panic("OH GOD SOMETHING BAD HAPPENED")
    }

    msg = fmt.Sprintf("%s\nYou chose to parse %s!\n", msg, parser)

    // write message to byte buffer
    buf := make([]byte, len(msg))

    for i, c := range msg {
        buf[i] = byte(c)
    }

    w.Write(buf[0:])
}

func ParseJSONHandlerPOST(w http.ResponseWriter, req *http.Request) {
    type resultFmt struct {
        Foo string `json:"foo"`
    }

    buf, _ := ioutil.ReadAll(req.Body)
    str := string(buf)
    fmt.Printf("%v\n", str)

    result := resultFmt{}
    json.Unmarshal(buf, &result)

    io.WriteString(w, fmt.Sprintf("unmarshalled: %v\n", result))
    io.WriteString(w, "done!\n")
}
