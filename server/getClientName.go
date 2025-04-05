package main

import (
    "bufio"
    "fmt"
    "net"
    "strings"
)

func getClientName(conn net.Conn) string {
    logo := 
`
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    .        | . \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
      -'        --'
`
    fmt.Fprintln(conn, logo)

    reader := bufio.NewReader(conn)
    var name string

    for {
        fmt.Fprint(conn, "[ENTER YOUR NAME]: ")
        rawName, err := reader.ReadString('\n')
        if err != nil {
            fmt.Fprintln(conn, "\nConnection closed while entering name.")
            return ""
        }

        name = strings.TrimSpace(rawName)

        if len(name) == 0 {
            fmt.Fprintln(conn, "Name cannot be empty.")
            continue
        }

        if len(name) > 20 {
            fmt.Fprintln(conn, "Name is too long. Max 20 characters.")
            continue
        }

        break
    }

    return name
}
