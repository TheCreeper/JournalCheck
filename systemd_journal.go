package main

/*
   #cgo CFLAGS:
   #cgo LDFLAGS: -lsystemd

   #include <stdlib.h>

   extern int journal_open();
   extern char* journal_read_line();
   extern char* journal_get_cursor();
   extern int journal_test_cursor(const char*);
   extern int journal_close();
*/
import "C"

func journal_open() int {

    return int(C.journal_open())
}

func journal_read_line() string {

    return C.GoString(C.journal_read_line())
}

func journal_get_cursor() string {

    return C.GoString(C.journal_get_cursor())
}

func journal_test_cursor(cursor string) int {

    return int(C.journal_test_cursor(C.CString(cursor)))
}

func journal_close() int {

    return int(C.journal_close())
}