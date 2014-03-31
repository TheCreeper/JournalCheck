#include <stdio.h>
#include <string.h>
#include <systemd/sd-journal.h>

sd_journal *j;

int journal_open() {

    return sd_journal_open(&j, SD_JOURNAL_LOCAL_ONLY);
}

int journal_close() {

    if (j) { // sd_journal_close does not return a value

        sd_journal_close(j);
        j = NULL;
        return 0;
    }
    return 1;
}

int journal_add_match(char *m) {

    return sd_journal_add_match(j, (const void **)m, 0);
}

int journal_seek_tail() {

    return sd_journal_seek_tail(j);
}

int journal_next() {

    return sd_journal_next(j);
}

int journal_previous() {

    return sd_journal_previous(j);
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

    return sd_journal_test_cursor(j, c);
}