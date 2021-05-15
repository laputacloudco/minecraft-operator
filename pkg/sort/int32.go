/*
Copyright 2020 The Laputa Cloud Co.

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
*/

package sort

import (
	stdsort "sort"
)

// Int32Slice attaches the methods of Interface to []int32,
// sorting in increasing order.
type Int32Slice []int32

func (p Int32Slice) Len() int           { return len(p) }
func (p Int32Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int32Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Int32Slice) Sort()              { stdsort.Sort(p) }

// FirstUnusedInRange finds and returns the first unused value in an
// ascendingly sorted range, or -1 if it is fully populated
func FirstUnusedInRange(used []int32, lower, upper int32) int32 {
	for i := int32(0); i <= (upper - lower); i++ {
		try := lower + i
		if int(i) >= len(used) {
			return try
		}
		if try < used[i] {
			return try
		}
	}
	return -1
}
