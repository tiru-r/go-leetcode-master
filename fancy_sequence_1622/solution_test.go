package fancy_sequence_1622

import "testing"

func TestFancy(t *testing.T) {
	tests := []struct {
		name string
		ops  func(*Fancy)
		idx  int
		want int
	}{
		{
			name: "empty sequence",
			ops:  func(f *Fancy) {},
			idx:  0,
			want: -1,
		},
		{
			name: "append only",
			ops: func(f *Fancy) {
				f.Append(2)
				f.Append(3)
			},
			idx:  1,
			want: 3,
		},
		{
			name: "add then mult",
			ops: func(f *Fancy) {
				f.Append(1)
				f.Append(2)
				f.AddAll(3)  // [4,5]
				f.MultAll(2) // [8,10]
			},
			idx:  0,
			want: 8,
		},
		{
			name: "mult then add",
			ops: func(f *Fancy) {
				f.Append(2)
				f.MultAll(5) // [10]
				f.AddAll(7)  // [17]
			},
			idx:  0,
			want: 17,
		},
		{
			name: "negative modulo wrap",
			ops: func(f *Fancy) {
				f.Append(0)
				f.AddAll(-1) // under 0 mod 1e9+7
			},
			idx:  0,
			want: 1_000_000_006,
		},
		{
			name: "out of bounds",
			ops: func(f *Fancy) {
				f.Append(10)
			},
			idx:  5,
			want: -1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			f := NewFancy()
			tc.ops(f)
			got := f.GetIndex(tc.idx)
			if got != tc.want {
				t.Errorf("GetIndex(%d) = %d; want %d", tc.idx, got, tc.want)
			}
		})
	}
}
