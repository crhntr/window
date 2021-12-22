package selector

import (
	"testing"

	"github.com/crhntr/please"
)

func Test_runeLen(t *testing.T) {
	t.Logf("%08x", 0)
	t.Logf("%08x", '$')
	t.Logf("%08x", '©')
	t.Logf("%08x", '⬟')
	t.Logf("%08x", '🍊')

	please.ExpectEqual(t, runeLen(0), 0)
	please.ExpectEqual(t, runeLen('$'), 1)
	please.ExpectEqual(t, runeLen('©'), 2)
	please.ExpectEqual(t, runeLen('⬟'), 3)
	please.ExpectEqual(t, runeLen('🍊'), 4)
}
