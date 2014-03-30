#include <stdio.h>
#include <string.h>
#include <systemd/sd-journal.h>

#define MATCH "_SYSTEMD_UNIT=sshd.service"

sd_journal *j;

int journal_open() {

    int r;

    r = sd_journal_open(&j, SD_JOURNAL_LOCAL_ONLY);
    if (r < 0) {

        return 1; // Failed to open journal
    }
    r = sd_journal_add_match(j, MATCH, 0);
    if (r < 0) {

        return 1;
    }
    r = sd_journal_seek_tail(j);
    if (r < 0) {

        return 1;
    }
    /*
        systemd feature/ bug: without a sd_journal_previous,
        sd_journal_seek_tail has no effect
    */
    r = sd_journal_previous(j);
    if (r < 0) {

        return 1;
    }
    return 0;
}

int journal_close() {

    if (j) {

        sd_journal_close(j);
        j = NULL;
        return 0;
    }
    return 1;
}

int journal_next() {

    int r;

    r = sd_journal_next(j);
    return r;
}

char* journal_get_data() {

    int r;
    char *d;
    size_t l;

    r = sd_journal_get_data(j, "MESSAGE", (const void **)&d, &l);
    if (r < 0) {

        return "";
    }
    return d;
}

char* journal_get_cursor() {

    int r;
    char *c;

    r = sd_journal_get_cursor(j, &c);
    if (r < 0) {

        return "";
    }
    return c;
}

int journal_test_cursor(const char *c) {

    int r;

    r = sd_journal_test_cursor(j, c);
    return r;
}