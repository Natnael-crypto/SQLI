package initializer_test

import (
	"os"
	"sqli/initializer"
	"testing"
	"testing/fstest"
)

func TestEnv(t *testing.T) {
	fs := fstest.MapFS{
		".env": {Data: []byte(`DBUSER=something
DBPASS=somethingortheother
DBADDR=133.133.133.133:3306
DBNAME=sqlidb`)},
	}

	initializer.LoadEnv(fs)
	got := os.Getenv("DBNAME")
	want := "sqlidb"
	if got != want {
		t.Errorf("got %v, want %v",got, want )
	}
}
