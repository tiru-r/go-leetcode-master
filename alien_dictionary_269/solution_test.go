package alien_dictionary_269

import "testing"

func TestAlienOrder(t *testing.T) {
	type args struct{ words []string }
	tests := []struct {
		name string
		args args
		want string
	}{
		// Original cases
		{"wrt…rftt", args{[]string{"wrt", "wrf", "er", "ett", "rftt"}}, "wertf"},
		{"z x", args{[]string{"z", "x"}}, "zx"},
		{"z x z", args{[]string{"z", "x", "z"}}, ""},
		{"za…cb", args{[]string{"za", "zb", "ca", "cb"}}, "azbc"},

		// Extra cases that frequently break implementations
		{"empty list", args{[]string{}}, ""},
		{"single word", args{[]string{"abc"}}, "abc"},
		{"same word twice", args{[]string{"abc", "abc"}}, "abc"},
		{"prefix contradiction", args{[]string{"abc", "ab"}}, ""},
		{"cycle via transitivity", args{[]string{"a", "b", "a"}}, ""},
		{"multi-letter cycle", args{[]string{"ab", "bc", "ca"}}, ""},
		{"all same char", args{[]string{"aaa", "aaa"}}, "a"},
		{"longer valid chain", args{[]string{"abc", "abx", "axx", "axy"}}, "abcxy"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AlienOrder(tt.args.words)
			if got != tt.want {
				t.Errorf("AlienOrder(%v) = %q, want %q", tt.args.words, got, tt.want)
			}
		})
	}
}
