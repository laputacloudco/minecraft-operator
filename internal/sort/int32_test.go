/*
Copyright 2020 The Laputa Cloud Co.

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
*/

package sort

import "testing"

func TestFirstUnusedInRange(t *testing.T) {
	tests := []struct {
		name  string
		used  []int32
		lower int32
		upper int32
		want  int32
	}{
		{
			name:  "in range",
			used:  []int32{0, 1, 3},
			lower: 0,
			upper: 3,
			want:  2,
		},
		{
			name:  "first",
			used:  []int32{1, 2, 3},
			lower: 0,
			upper: 3,
			want:  0,
		},
		{
			name:  "last",
			used:  []int32{0, 1, 2},
			lower: 0,
			upper: 3,
			want:  3,
		},
		{
			name:  "multi",
			used:  []int32{0, 1, 3, 5},
			lower: 0,
			upper: 5,
			want:  2,
		},
	}
	for _, tt := range tests {
		if got := FirstUnusedInRange(tt.used, tt.lower, tt.upper); got != tt.want {
			t.Errorf("%s : FirstUnusedInRange() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
