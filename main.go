/*
    Allow running as cron job
*/

package main

import (

    //    "code.google.com/p/go.crypto/ssh"
    //    "github.com/pkg/sftp"
    "log"
    "strings"
    "time"
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "runtime"
)

func watchJournal(sysd Journald, ntf Notifiers) {

    /*
        Implement a que
    */

    if (journal_open() < 0) {

        log.Fatal("Failed to open the journal!")
    }
    if (sysd.Match != "") {

        if (journal_add_match(sysd.Match) < 0) {

            log.Fatal("Failed to add match in journal!")
        }
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

        time.Sleep(time.Duration(sysd.Sleep) * time.Second)

        next := journal_next()
        if (next == 0) {

            // at end of journal
            continue;
        } else if (next < 0) {

            // failed to iterate to next entry
            log.Print("Failed to iterate to next entry in journal!")
            break;
        }
        for (next > 0) {

            event := journal_get_data()

            for _, v := range sysd.TriggerWords {

                if strings.Contains(event, v) {

                    notice := fmt.Sprintf("System Event Occurred on %s %s %s", GetHostName(), "at", time.Now())
                    log.Print(notice)

                    err := SendEmail(
                    ntf.Email.Host,
                    ntf.Email.Port,
                    ntf.Email.Username,
                    ntf.Email.Password,
                    ntf.Email.To,
                    notice,
                    strings.Split(event, "MESSAGE=")[1])
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

func init() {

    runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

    cfg := GetCFG()

    go watchJournal(cfg.SystemdJournal, cfg.Notifications)

    sigc := make(chan os.Signal, 1)
    signal.Notify(sigc,
        syscall.SIGHUP)

    for {

        s := <-sigc

        switch s {

            case syscall.SIGHUP:

                log.Print("Reloading Configuration File")
                cfg = GetCFG()
        }
    }
}
