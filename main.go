package main

import (

    //    "code.google.com/p/go.crypto/ssh"
    //    "github.com/pkg/sftp"
    "log"
    "strings"
    "time"
    "fmt"
    "os"
)

func closeJournal() {

    n := journal_close()
    if (n != 0) {

        log.Fatal("Failed to close the journal!")
    }
}

func watchJournal(triggerwords []string, ntf Notifiers) {

    if (journal_open() != 0) {

        log.Fatal("Failed to open the journal!")
    }

    log.Print("Now watching Journal")
    for {

        time.Sleep(1000 * time.Millisecond)

        next := journal_next()
        if (next == 0) {

            // at end of journal
            continue;
        } else if (next < 0) {

            // failed to iterate to next entry
            log.Print("Failed to iterate to next entry in journal!")
            closeJournal()
            os.Exit(1)
        }
        for (next > 0) {

            if (next < 0) {

                log.Print("Failed to iterate to next entry in journal!")
                closeJournal()
                os.Exit(1)
            }

            event := journal_get_data()

            for _, v := range triggerwords {

                if strings.Contains(event, v) {

                    notice := fmt.Sprintf("Login at %s, %s", GetHostName(), time.Now())
                    log.Print(notice)
                }
            }
            next = journal_next()
        }
    }
    closeJournal()
}

func main() {

    triggerwords := []string{"session opened for user"}
    cfg := GetCFG()
    watchJournal(triggerwords, cfg.Notifications)
}
