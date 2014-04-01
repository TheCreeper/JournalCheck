package main

import (
    "encoding/json"
    "io/ioutil"
    "log"
//    "os"
)

var cfgfile string = "./config.json"

type GConfig struct {

    Notifications Notifiers
    SystemdJournal Journald
}

type Journald struct {

    Sleep int64
    Match string
    TriggerWords []string
}

type Notifiers struct {

    Email NEmail
    PushOver NPushOver
    PushAlot NPushAlot
}

/* Notifiers */
type NEmail struct {

    Host string
    Port int
    Username string
    Password string
    To []string
}

type NPushOver struct {

    UserToken string
    AppToken  string
}

type NPushAlot struct {

    Token string
}
/*
func GetDefaultConfig() GConfig {

    var tfg GConfig

    var nfr Notifiers
    var mail NEmail
    var pusho NPushOver
    var pusha NPushAlot

    mail.Host = "127.0.0.1"
    mail.Port = 25
    mail.Username = "notify@example.com"
    mail.Password = ""
    mail.To = []string {

        "session opened for user"
    }
    nfr.Email = mail

    pusho.UserToken = "token"
    pusho.AppToken = "token"
    nfr.PushOver = pusho

    pusha.Token = "token"
    nfr.PushAlot = pusha

    tfg.Notifications = nfr
    return tfg
}

func CheckIfResetConfig(args []string) {

    if len(args) == 2 {

        if args[1] == "reset" {

            e := os.Remove(cfgfile)
            if e != nil {

                log.Fatal("Could not remove current config file. Permissions issue?")
            }
            Default := GetDefaultConfig()
            out, e := json.Marshal(Default)
            e = ioutil.WriteFile(cfgfile, out, 600)
            if e != nil {

                log.Fatal("cannot open settings file :(")
            }
            log.Fatal("Built config file. please fill it in.")
        }
    }
}
*/
func GetCFG() GConfig {

    b, e := ioutil.ReadFile(cfgfile)
    /*
    tfg := GetDefaultConfig()
    if e != nil {

        out, e := json.Marshal(tfg)
        e = ioutil.WriteFile(cfgfile, out, 600)
        if e != nil {

            log.Fatal("cannot open settings file :(")
        }
        log.Fatal("Built config file. please fill it in.")
    }*/
    var tfg GConfig

    e = json.Unmarshal(b, &tfg)
    if e != nil {

        log.Fatalf("Could not parse config settings. You may have to remove %s", cfgfile)
    }
    return tfg
}