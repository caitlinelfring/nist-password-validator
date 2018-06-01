package password

import "testing"

func TestPasswordIsASCII(t *testing.T) {
	var testStrings = []struct {
		value   string
		isASCII bool
	}{
		{"", true},
		{"abcd", true},
		{"qwertyui", true},
		{"qwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnmqwertyuiopas", true},
		{"asdlfdja;sfjajfpojp8jp8efpw3jpowefjohefpowahepfo", true},
		{"asdlfdja;sfjajfpojp8jp8efpw3jpowsa efjohefpowahepfo", true},
		{"asdlfdja;sfjajfpojp8jp8efpw3jpowsa efjohefpowahepfo", true},
		{"asdlfdja;lsdijfa;sdfpoaisdufpoaisuf9oarpfauhrfiuehfpiudhfioufgoiudfhgoiudfgpdupodsifpuosiUFPAOSIDUFPAOSDIUFP", true},
		{"☺☻☹", false},
		{"日a本b語ç日ð本Ê語þ日¥本¼語i日©", false},
		{"日a本b語ç日ð本Ê語þ日¥本¼語i日©日a本b語ç日ð本Ê語þ", false},
		{"\x80\x80\x80\x80asdasdaslkjdalskjf", false},
		{"ಠ_ಠasdsadಠ_ಠ", false},
		{"zzzzfitt", true},
		{"vjht008", true},
		{"123456789", true},
	}

	for _, test := range testStrings {
		pass := Password(test.value)
		is := pass.isASCII()
		if is != test.isASCII {
			t.Errorf("%s -- expected: %t, got: %t", test.value, test.isASCII, is)
		}
	}
}
