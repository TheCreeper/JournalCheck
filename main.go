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

var (

    exitCode = 0
)

func watchJournal(triggerwords []string, ntf Notifiers, match string) {

    if (journal_open() < 0) {

        log.Fatal("Failed to open the journal!")
    }
    if (journal_add_match(match) < 0) {

        log.Fatal("Failed to add match in journal!")
    }
    if (journal_seek_tail() < 0) {

        log.Fatal("Failed to skip to end of journal!")
    }
    /*
        systemd feature/ bug: without a sd_journal_previous,
        sd_journal_seek_tail has no effect
    */
    if (journal_previous() < 0) {

        log.Fatal("Failed to go back in journal!")
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
            exitCode = 1
            break;
        }
        for (next > 0) {

            event := journal_get_data()

            for _, v := range triggerwords {

                if strings.Contains(event, v) {

                    notice := fmt.Sprintf("Login at %s, %s", GetHostName(), time.Now())
                    log.Print(notice)

                    err := SendEmail(
                    ntf.Email.Host,
                    ntf.Email.Port,
                    ntf.Email.Username,
                    ntf.Email.Password,
                    ntf.Email.To,
                    fmt.Sprintf("System Event Occurred on %s, %s, %s", GetHostName(), "at", time.Now()),
                    notice)
                    if (err != nil) {

                        log.Print(err)
                    }
                }
            }
            next = journal_next()
        }
    }
    if (journal_close() != 0) {

        log.Fatal("Failed to close the journal!")
    }
}

func main() {

    triggerwords := []string{"session opened for user"}
    cfg := GetCFG()
    m := "_SYSTEMD_UNIT=sshd.service"
    watchJournal(triggerwords, cfg.Notifications, m)
    os.Exit(exitCode)
}
