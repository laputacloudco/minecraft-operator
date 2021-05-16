/*
Copyright 2020 The Laputa Cloud Co.

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
*/

package types

type LevelType string

const (
	Amplified     LevelType = "AMPLIFIED"
	BiomesOP      LevelType = "BIOMESOP"
	BiomesOPlenty LevelType = "BIOMESOPLENTY"
	Buffet        LevelType = "BUFFET"
	Customized    LevelType = "CUSTOMIZED"
	Default       LevelType = "DEFAULT"
	Flat          LevelType = "FLAT"
	LargeBiomes   LevelType = "LARGEBIOMES"
)

type GameMode string

const (
	Adventure GameMode = "adventure"
	Creative  GameMode = "creative"
	Spectator GameMode = "spectator"
	Survival  GameMode = "survival"
)

type GameOptions struct {
	AllowFlight                bool      `json:"allowFlight,omitempty" env:"ALLOW_FLIGHT,omitempty"`
	AllowNether                bool      `json:"allowNether,omitempty" env:"ALLOW_NETHER,omitempty"`
	AnnouncePlayerAchievements bool      `json:"announcePlayerAchievements,omitempty" env:"ANNOUNCE_PLAYER_ACHIEVEMENTS,omitempty"`
	Difficulty                 int       `json:"difficulty,omitempty" env:"DIFFICULTY,omitempty"`
	EnableCommandBlock         bool      `json:"enableCommandBlock,omitempty" env:"ENABLE_COMMAND_BLOCK,omitempty"`
	ForceGamemode              bool      `json:"forceGamemode,omitempty" env:"FORCE_GAMEMODE,omitempty"`
	Gamemode                   GameMode  `json:"gamemode,omitempty" env:"MODE,omitempty"`
	GeneratorSettings          []string  `json:"generatorSettings,omitempty" env:"GENERATOR_SETTINGS,omitempty"`
	Hardcore                   bool      `json:"hardcore,omitempty" env:"HARDCORE,omitempty"`
	LevelName                  string    `json:"levelName,omitempty" env:"LEVEL,omitempty"`
	LevelType                  LevelType `json:"levelType,omitempty" env:"LEVEL_TYPE,omitempty"`
	MaxBuildHeight             int       `json:"maxBuildHeight,omitempty" env:"MAX_BUILD_HEIGHT,omitempty"`
	Modpack                    string    `json:"modpack,omitempty" env:"MODPACK,omitempty"`
	Mods                       []string  `json:"mods,omitempty" env:"MODS,omitempty"`
	OnlineMode                 bool      `json:"onlineMode,omitempty" env:"ONLINE_MODE,omitempty"`
	Ops                        []string  `json:"ops,omitempty" env:"OPS,omitempty"`
	ResourcePack               string    `json:"resourcePack,omitempty" env:"RESOURCE_PACK,omitempty"`
	ResourcePackSha1           string    `json:"resourcePackSha1,omitempty" env:"RESOURCE_PACK_SHA1,omitempty"`
	Seed                       string    `json:"seed,omitempty" env:"SEED,omitempty"`
	SpawnAnimals               bool      `json:"spawnAnimals,omitempty" env:"SPAWN_ANIMALS,omitempty"`
	SpawnMonsters              bool      `json:"spawnMonsters,omitempty" env:"SPAWN_MONSTERS,omitempty"`
	SpawnNpcs                  bool      `json:"spawnNpcs,omitempty" env:"SPAWN_NPCS,omitempty"`
	SpawnProtection            int       `json:"spawnProtection,omitempty" env:"SPAWN_PROTECTION,omitempty"`
	TexturePack                string    `json:"texturePack,omitempty" env:"TEXTURE_PACK,omitempty"`
	ViewDistance               int       `json:"viewDistance,omitempty" env:"VIEW_DISTANCE,omitempty"`
	WhiteList                  bool      `json:"whiteList,omitempty" env:"WHITE_LIST,omitempty"`
}
