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
)

func watchJournal(sysd Journald, ntf Notifiers) {

    if (journal_open() < 0) {

        log.Fatal("Failed to open the journal!")
    }
    if (journal_flush_matches() < 0) {

        log.Fatal("Failed to flush the journal filter!")
    }
    if (sysd.Match[0] != "") {

        for _, v := range sysd.Match {

            if (journal_add_match(v) < 0) {

                log.Fatal("Failed to add match in journal!")
            }
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

            var event string
            if (journal_get_data(&event) < 0) {

                log.Print("Failed to get journal data!")
            }

            event = strings.Split(event, "MESSAGE=")[1]

            for _, v := range sysd.TriggerWords {

                if strings.Contains(event, v) {

                    notice := fmt.Sprintf("System Event Occurred on %s %s %s", GetHostName(), "at", time.Now())
                    log.Print(notice)
                    log.Print(event)

                    err := SendEmail(
                    ntf.Email.Host,
                    ntf.Email.Port,
                    ntf.Email.Username,
                    ntf.Email.Password,
                    ntf.Email.To,
                    notice,
                    event)
                    if (err != nil) {

                        log.Print(err)
                    }
                }
            }
            next = journal_next()
        }
    }
    if (journal_close() < 0) {

        log.Fatal("Failed to close the journal!")
    }
}

func main() {

    cfg, err := GetCFG()
    if (err != nil) {

        log.Fatalf("Could not parse config settings. You may have to remove %s", cfgfile)
    }

    go watchJournal(cfg.SystemdJournal, cfg.Notifications)

    sigc := make(chan os.Signal, 1)
    /*
        Only interrupt and kill is guaranteed to work across
        multiple platforms.
    */
    signal.Notify(sigc,
        os.Interrupt)

    for {

        // Main loop

        s := <-sigc

        if (s == os.Interrupt) {

            // Gracefully close log and exit here
            os.Exit(0)
        }
    }
}
