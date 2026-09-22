package metricsql

import "testing"

func TestParseLeadingZeroNumbers(t *testing.T) {
	f := func(expr string) {
		t.Helper()
		if _, err := Parse(expr); err != nil {
			t.Fatalf("cannot parse %q: %s", expr, err)
		}
	}
	// Numbers with leading zeros must parse as decimal, like in Prometheus/Mimir.
	// See https://github.com/VictoriaMetrics/VictoriaMetrics/issues/11621
	f("hour() > 09")
	f("hour() >= 08")
	f("09 + 007")
	f("09.5")
	f("foo{bar='baz'} > 0099")

	// Valid octal, hex, binary and octal-prefix forms must keep working.
	if _, err := Parse("0123"); err != nil {
		t.Fatalf("cannot parse octal number: %s", err)
	}
	if _, err := Parse("0x1f"); err != nil {
		t.Fatalf("cannot parse hex number: %s", err)
	}
	if _, err := Parse("0b1010"); err != nil {
		t.Fatalf("cannot parse binary number: %s", err)
	}
	if _, err := Parse("0o17"); err != nil {
		t.Fatalf("cannot parse 0o number: %s", err)
	}
}
