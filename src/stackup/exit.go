package main

import "os"

type OsExit struct { Code int }

func ExitHandler() {
  if r := recover(); r != nil {
    switch ex := r.(type) {
    case OsExit:
      os.Exit(ex.Code)
    default:
      panic(r)
    }
  }
}

func Exit(code int) {
  panic(OsExit{code})
}
