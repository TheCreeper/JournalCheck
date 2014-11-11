#include <systemd/sd-journal.h>

sd_journal *j;

/*
	Open the system journal for reading
*/
int journal_open() {

	return sd_journal_open(&j, SD_JOURNAL_LOCAL_ONLY);
}

int journal_close() {

	if (j) { // sd_journal_close does not return a value

		sd_journal_close(j);
		j = NULL;
		return 0;
	}
	return -1;
}

/*
	Seek to a position in the journal
*/
int journal_seek_tail() {

	return sd_journal_seek_tail(j);
}

/*
	Advance or set back the read pointer in the journal
*/
int journal_next() {

	return sd_journal_next(j);
}

int journal_previous() {

	return sd_journal_previous(j);
}

/*
	Read data fields from the current journal entry
*/
int journal_get_data(char **d) {

	size_t l;

	return sd_journal_get_data(j, "MESSAGE", (const void **)d, &l);
}

/*
	Get cursor string for or test cursor string against the current journal entry
*/
int journal_get_cursor(char **c) {

	return sd_journal_get_cursor(j, c);
}

int journal_test_cursor(const char *c) {

	return sd_journal_test_cursor(j, c);
}

/*
	Add or remove entry matches
*/
int journal_add_match(char *m) {

	return sd_journal_add_match(j, (const void **)m, 0);
}

int journal_add_disjunction() {

	return sd_journal_add_disjunction(j);
}

int journal_add_conjunction() {

	return sd_journal_add_conjunction(j);
}

int journal_flush_matches() {

	if (j) { // sd_journal_flush_matches returns no value

		sd_journal_flush_matches(j);
		return 0;
	}
	return -1;
}