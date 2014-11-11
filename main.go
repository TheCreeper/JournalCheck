/*
	Allow running as cron job
*/

package main

import (

	//    "code.google.com/p/go.crypto/ssh"
	//    "github.com/pkg/sftp"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"
)

func WatchJournal(cfg ClientConfig) {

	if journal_open() < 0 {

		log.Fatal("Failed to open the journal!")
	}
	if journal_flush_matches() < 0 {

		log.Fatal("Failed to flush the journal filter!")
	}
	if len(cfg.Journald.Match) < 1 {

		for _, v := range cfg.Journald.Match {

			if journal_add_match(v) < 0 {

				log.Fatal("Failed to add match in journal!")
			}
		}
	}
	if journal_seek_tail() < 0 {

		log.Fatal("Failed to skip to end of journal!")
	}
	/*
		systemd feature/ bug: without a sd_journal_previous,
		sd_journal_seek_tail has no effect
	*/
	if journal_previous() < 0 {

		log.Fatal("Failed to go back in journal!")
	}

	log.Print("Now watching Journal")
	for {

		time.Sleep(time.Duration(cfg.Journald.Sleep) * time.Second)

		next := journal_next()
		if next == 0 {

			// at end of journal
			continue
		} else if next < 0 {

			// failed to iterate to next entry
			log.Print("Failed to iterate to next entry in journal!")
			break
		}
		for next > 0 {

			var event string
			if journal_get_data(&event) < 0 {

				log.Print("Failed to get journal data!")
			}

			event = strings.Split(event, "MESSAGE=")[1]

			for _, v := range cfg.Journald.TriggerWords {

				if strings.Contains(event, v) {

					notice := fmt.Sprintf("System Event Occurred on %s %s %s", GetHostName(), "at", time.Now())
					log.Print(notice)
					log.Print(event)

					err := SendEmail(
						cfg.Notifications.Email.Host,
						cfg.Notifications.Email.Port,
						cfg.Notifications.Email.Username,
						cfg.Notifications.Email.Password,
						cfg.Notifications.Email.To,
						notice,
						event)
					if err != nil {

						log.Print(err)
					}
				}
			}
			next = journal_next()
		}
	}
	if journal_close() < 0 {

		log.Fatal("Failed to close the journal!")
	}
}

func init() {

	flag.StringVar(&ConfigFile, "f", "./config.json", "The configuration file")
	flag.Parse()
}

func main() {

	cfg, err := GetCFG(ConfigFile)
	if err != nil {

		log.Fatalf("Could not parse config settings. You may have to remove %s", ConfigFile)
	}
	WatchJournal(cfg)
}
