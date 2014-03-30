package main

/*
   #cgo CFLAGS:
   #cgo LDFLAGS: -lsystemd

   #include <stdlib.h>

   extern int journal_open();
   extern int journal_close();
   extern int journal_next();
   extern char* journal_get_data();
   extern char* journal_get_cursor();
   extern int journal_test_cursor(const char*);
*/
import "C"

func journal_open() int {

    return int(C.journal_open())
}

func journal_close() int {

    return int(C.journal_close())
}

func journal_next() int {

  return int(C.journal_next())
}

func journal_get_data() string {

    return C.GoString(C.journal_get_data())
}

func journal_get_cursor() string {

    return C.GoString(C.journal_get_cursor())
}

func journal_test_cursor(cursor string) int {

    return int(C.journal_test_cursor(C.CString(cursor)))
}