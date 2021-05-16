/*
Copyright 2020 The Laputa Cloud Co.

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
*/

package component

import (
	gamev1alpha2 "github.com/laputacloudco/minecraft-operator/api/v1alpha2"
)

func standardLabels(mc gamev1alpha2.Minecraft) map[string]string {
	return map[string]string{
		"app":      "minecraft",
		"instance": mc.Name,
		"managed":  "true",
		"owner":    mc.Name,
	}
}
