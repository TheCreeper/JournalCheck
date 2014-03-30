package main

import (

    //    "code.google.com/p/go.crypto/ssh"
    //    "github.com/pkg/sftp"
    "log"
    "strings"
    "time"
    "fmt"
)

func watchJournal(triggerwords []string) {

    if (journal_open() != 0) {

        log.Fatal("Failed to open the journal")
    }
    previousLine := journal_get_cursor()

    log.Print("Now watching Journal")
    for {

        time.Sleep(1000 * time.Millisecond)

        event := journal_read_line()
        currentLine := journal_test_cursor(previousLine)
        if (currentLine > 0) {

            // previous and currentline are the same
            continue;
        }
        for _, v := range triggerwords {

            if strings.Contains(event, v) {

                    log.Print(fmt.Sprintf("Login at %s", GetHostName()), " ", event)
                }
            }
        }
        previousLine = journal_get_cursor()
    }
    if (journal_close() != 0) {

        log.Fatal("Failed to close the journal!")
    }
}

func main() {

    triggerwords := []string{"session opened for user"}
    watchJournal(triggerwords)
}
