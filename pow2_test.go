package pow2

import "testing"

var testCase = map[Int]bool{
	0:                   false,
	1:                   true,
	2:                   true,
	3:                   false,
	4:                   true,
	5:                   false,
	7:                   false,
	8:                   true,
	9:                   false,
	15:                  false,
	16:                  true,
	17:                  false,
	31:                  false,
	32:                  true,
	33:                  false,
	63:                  false,
	64:                  true,
	65:                  false,
	127:                 false,
	128:                 true,
	129:                 false,
	255:                 false,
	256:                 true,
	257:                 false,
	511:                 false,
	512:                 true,
	513:                 false,
	1023:                false,
	1024:                true,
	1025:                false,
	2047:                false,
	2048:                true,
	2049:                false,
	4095:                false,
	4096:                true,
	4097:                false,
	8191:                false,
	8192:                true,
	8193:                false,
	16383:               false,
	16384:               true,
	16385:               false,
	32767:               false,
	32768:               true,
	32769:               false,
	65535:               false,
	65536:               true,
	65537:               false,
	131071:              false,
	131072:              true,
	131073:              false,
	262143:              false,
	262144:              true,
	262145:              false,
	524287:              false,
	524288:              true,
	524289:              false,
	1048575:             false,
	1048576:             true,
	1048577:             false,
	2097151:             false,
	2097152:             true,
	2097153:             false,
	4194303:             false,
	4194304:             true,
	4194305:             false,
	8388607:             false,
	8388608:             true,
	8388609:             false,
	16777215:            false,
	16777216:            true,
	16777217:            false,
	33554431:            false,
	33554432:            true,
	33554433:            false,
	67108863:            false,
	67108864:            true,
	67108865:            false,
	134217727:           false,
	134217728:           true,
	134217729:           false,
	268435455:           false,
	268435456:           true,
	268435457:           false,
	536870911:           false,
	536870912:           true,
	536870913:           false,
	1073741823:          false,
	1073741824:          true,
	1073741825:          false,
	2147483647:          false,
	2147483648:          true,
	2147483649:          false,
	4294967295:          false,
	4294967296:          true,
	4294967297:          false,
	8589934591:          false,
	8589934592:          true,
	8589934593:          false,
	17179869183:         false,
	17179869184:         true,
	17179869185:         false,
	34359738367:         false,
	34359738368:         true,
	34359738369:         false,
	68719476735:         false,
	68719476736:         true,
	68719476737:         false,
	137438953471:        false,
	137438953472:        true,
	137438953473:        false,
	274877906943:        false,
	274877906944:        true,
	274877906945:        false,
	549755813887:        false,
	549755813888:        true,
	549755813889:        false,
	1099511627775:       false,
	1099511627776:       true,
	1099511627777:       false,
	2199023255551:       false,
	2199023255552:       true,
	2199023255553:       false,
	4398046511103:       false,
	4398046511104:       true,
	4398046511105:       false,
	8796093022207:       false,
	8796093022208:       true,
	8796093022209:       false,
	17592186044415:      false,
	17592186044416:      true,
	17592186044417:      false,
	35184372088831:      false,
	35184372088832:      true,
	35184372088833:      false,
	70368744177663:      false,
	70368744177664:      true,
	70368744177665:      false,
	140737488355327:     false,
	140737488355328:     true,
	140737488355329:     false,
	281474976710655:     false,
	281474976710656:     true,
	281474976710657:     false,
	562949953421311:     false,
	562949953421312:     true,
	562949953421313:     false,
	1125899906842623:    false,
	1125899906842624:    true,
	1125899906842625:    false,
	2251799813685247:    false,
	2251799813685248:    true,
	2251799813685249:    false,
	4503599627370495:    false,
	4503599627370496:    true,
	4503599627370497:    false,
	9007199254740991:    false,
	9007199254740992:    true,
	9007199254740993:    false,
	18014398509481983:   false,
	18014398509481984:   true,
	18014398509481985:   false,
	36028797018963967:   false,
	36028797018963968:   true,
	36028797018963969:   false,
	72057594037927935:   false,
	72057594037927936:   true,
	72057594037927937:   false,
	144115188075855871:  false,
	144115188075855872:  true,
	144115188075855873:  false,
	288230376151711743:  false,
	288230376151711744:  true,
	288230376151711745:  false,
	576460752303423487:  false,
	576460752303423488:  true,
	576460752303423489:  false,
	1152921504606846975: false,
	1152921504606846976: true,
	1152921504606846977: false,
	2305843009213693951: false,
	2305843009213693952: true,
	2305843009213693953: false,
	4611686018427387903: false,
	4611686018427387904: true,
	4611686018427387905: false,
	9223372036854775807: false,
	9223372036854775808: true,
	9223372036854775809: false,
}

func TestAndMinus1(t *testing.T) {
	for k, v := range testCase {
		if e := AndMinus1(k); e != v {
			t.Errorf("AndMinus1(%d) = %v, expected = %v", k, e, v)
		}
	}
}

func TestOnesCount(t *testing.T) {
	for k, v := range testCase {
		if e := OnesCount(k); e != v {
			t.Errorf("OnesCount(%d) = %v, expected = %v", k, e, v)
		}
	}
}

func TestForCountAll(t *testing.T) {
	for k, v := range testCase {
		if e := ForCountAll(k); e != v {
			t.Errorf("ForCountAll(%d) = %v, expected = %v", k, e, v)
		}
	}
}

func TestForCountOnce(t *testing.T) {
	for k, v := range testCase {
		if e := ForCountOnce(k); e != v {
			t.Errorf("ForCountOnce(%d) = %v, expected = %v", k, e, v)
		}
	}
}

func TestInSlice(t *testing.T) {
	for k, v := range testCase {
		if e := InSlice(k); e != v {
			t.Errorf("InSlice(%d) = %v, expected = %v", k, e, v)
		}
	}
}

func TestInMap(t *testing.T) {
	for k, v := range testCase {
		if e := InMap(k); e != v {
			t.Errorf("InMap(%d) = %v, expected = %v", k, e, v)
		}
	}
}

// --

func BenchmarkAndMinus1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k := range testCase {
			_ = AndMinus1(k)
		}
	}
}

func BenchmarkOnesCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k := range testCase {
			_ = OnesCount(k)
		}
	}
}

func BenchmarkForCountAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k := range testCase {
			_ = ForCountAll(k)
		}
	}
}

func BenchmarkForCountOnce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k := range testCase {
			_ = ForCountOnce(k)
		}
	}
}

func BenchmarkInSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k := range testCase {
			_ = InSlice(k)
		}
	}
}

func BenchmarkInMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k := range testCase {
			_ = InMap(k)
		}
	}
}
