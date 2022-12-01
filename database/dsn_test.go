package database

import (
	"testing"
)

func TestDSNFromURI(t *testing.T) {

	tests := map[string]string{
		"example://mem":                        "file::memory:?mode=memory&cache=shared",
		"example:///usr/local/test.db":         "file:/usr/local/test.db?cache=shared&mode=rwc",
		"example:///usr/local/test.db?foo=bar": "file:/usr/local/test.db?cache=shared&mode=rwc&foo=bar",
	}

	for uri, expected := range tests {

		dsn, err := DSNFromURI(uri)

		if err != nil {
			t.Fatalf("Failed to derive DSN from '%s', %v", uri, err)
		}

		if dsn != expected {
			t.Fatalf("Unexpected DSN for '%s', expected '%s' but got '%s'", uri, expected, dsn)
		}
	}

}
