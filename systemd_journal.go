package main

/*
    #include <stdlib.h>

    #cgo CFLAGS:
    #cgo LDFLAGS: -lsystemd

    extern int journal_open();
    extern int journal_close();

    extern int journal_seek_tail();

    extern int journal_next();
    extern int journal_previous();

    extern int journal_get_data(char **);

    extern int journal_get_cursor(char**);
    extern int journal_test_cursor(const char*);

    extern int journal_add_match(char*);
    extern int journal_add_disjunction();
    extern int journal_add_conjunction();
    extern int journal_flush_matches();
*/
import "C"
import "unsafe"

/*
    Open the system journal for reading
*/
func journal_open() int {

    return int(C.journal_open())
}

func journal_close() int {

    return int(C.journal_close())
}

/*
    Seek to a position in the journal
*/
func journal_seek_tail() int {

    return int(C.journal_seek_tail())
}

/*
    Advance or set back the read pointer in the journal
*/
func journal_next() int {

    return int(C.journal_next())
}

func journal_previous() int {

    return int(C.journal_previous())
}

/*
    Read data fields from the current journal entry
*/
func journal_get_data(data *string) int {

    var r int
    var n *C.char

    defer C.free(unsafe.Pointer(n))
    r = int(C.journal_get_data(&n))
    *data = C.GoString(n)

    return r
}

/*
    Get cursor string for or test cursor string against the current journal entry
*/
func journal_get_cursor(cursor *string) int {

    var r int
    var n *C.char

    defer C.free(unsafe.Pointer(n))
    r = int(C.journal_get_cursor(&n))
    *cursor = C.GoString(n)

    return r
}

func journal_test_cursor(cursor string) int {

    return int(C.journal_test_cursor(C.CString(cursor)))
}

/*
    Add or remove entry matches
*/
func journal_add_match(m string) int {

    return int(C.journal_add_match(C.CString(m)))
}

func journal_add_disjunction() int {

    return int(C.journal_add_disjunction())
}

func journal_add_conjunction() int {

    return int(C.journal_add_disjunction())
}

func journal_flush_matches() int {

    return int(C.journal_flush_matches())
}