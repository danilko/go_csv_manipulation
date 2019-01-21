package main

import (
  "encoding/csv"
  "strings"
  "log"
  "os"
  "io"
  "fmt"
)

func main() {
  // setup reader
  csvIn, err := os.Open("input.csv")
  if err != nil {
    log.Fatal(err)
  }
  r := csv.NewReader(csvIn)

  // setup writer
  csvOut, err := os.Create("output.csv")
  if err != nil {
    log.Fatal("Unable to open output")
  }
  w := csv.NewWriter(csvOut)
  defer csvOut.Close()

  // handle header
  rec, err := r.Read()
  if err != nil {
    log.Fatal(err)
  }
  rec = append(rec, "score")
  if err = w.Write(rec); err != nil {
    log.Fatal(err)
  }

  for {
    rec, err = r.Read()
    if err != nil {
      if err == io.EOF {
        break
      }
      log.Fatal(err)
    }

    splitmfgPNos := strings.Split(rec[15], "~~")
    splitmfgCompanies := strings.Split(rec[16], "~~")

    for key, _ := range splitmfgPNos {
    fmt.Printf("additional column \n")

        if len(splitmfgPNos) > key {
        rec[15] = splitmfgPNos[key]
        }

        if len(splitmfgCompanies) > key {
        rec[16] = splitmfgCompanies[key]
        } else {
          rec[16] = "ERROR"
        }

         if err = w.Write(rec); err != nil {
         log.Fatal(err)
         }

         w.Flush()
    }


  }
}