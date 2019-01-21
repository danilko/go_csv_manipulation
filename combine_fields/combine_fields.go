package main

import (
  "encoding/csv"
  "log"
  "os"
  "io"
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

  combineRec := []string{}
  previousLevel := "-1";

  for {
    rec, err = r.Read()
    if err != nil {
      if err == io.EOF {
        if len(combineRec) > 0 {
           if err = w.Write(combineRec); err != nil {
           log.Fatal(err)
           }

           w.Flush()
        }

        break
      }
      log.Fatal(err)
    }

    currentLevel := rec[1]

    if currentLevel != previousLevel {
       if len(combineRec) > 0 {
          if err = w.Write(combineRec); err != nil {
           log.Fatal(err)
          }

           w.Flush()
        }

        combineRec = rec

        previousLevel = currentLevel
    } else  {
      combineRec[15] = combineRec[15] + ";" + rec[15]
      combineRec[16] = combineRec[16] + ";" + rec[16]
    }

  }

  w.Flush()
}