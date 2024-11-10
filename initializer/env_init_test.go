package initializer_test

import (
	"os"
	"sqli/initializer"
	"testing"
	"testing/fstest"
)

func Test(t *testing.T) {
	fs := fstest.MapFS{
		".env": {Data: []byte(`DBUSER=sqli
DBPASS=sqlishouldbedeadalready
DBADDR=192.168.92.43:3306
DBNAME=sqlidb`)},
	}

	initializer.LoadEnv(fs)
	got := os.Getenv("DBNAME")
	want := "sqlidb"
	if got != want {
		t.Errorf("got %v, want %v",got, want )
	}
}
