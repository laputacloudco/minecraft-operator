/*
Copyright 2020 The Laputa Cloud Co.

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
*/

package config

// Controller is the operator config struct
type Controller struct {
	Defaults map[string]string `json:"defaults,omitempty"`
}
