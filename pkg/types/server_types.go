/*
Copyright 2020 The Laputa Cloud Co.

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
*/

package types

type ServerTypes string

const (
	Airplane      ServerTypes = "AIRPLANE"
	Bukkit        ServerTypes = "BUKKIT"
	Catserver     ServerTypes = "CATSERVER"
	CurseForge    ServerTypes = "CURSEFORGE"
	FTB           ServerTypes = "FTBA"
	Fabric        ServerTypes = "FABRIC"
	Forge         ServerTypes = "FORGE"
	Magma         ServerTypes = "MAGMA"
	Mohist        ServerTypes = "MOHIST"
	Paper         ServerTypes = "PAPER"
	Purpur        ServerTypes = "PURPUR"
	Spiget        ServerTypes = "SPIGET"
	Spigot        ServerTypes = "SPIGOT"
	SpongeVanilla ServerTypes = "SPONGEVANILLA"
	Tuinity       ServerTypes = "TUINITY"
	Yatopia       ServerTypes = "YATOPIA"
)

type ServerOptions struct {
	AutopausePeriod                int         `json:"autopausePeriod,omitempty" env:"AUTPAUSE_PERIOD,omitempty"`
	AutopauseTimeoutEst            int         `json:"autopauseTimeoutEst,omitempty" env:"AUTPAUSE_TIMEOUT_EST,omitempty"`
	AutopauseTimeoutInit           int         `json:"autopauseTimeoutInit,omitempty" env:"AUTPAUSE_TIMEOUT_INIT,omitempty"`
	AutopauseTimeoutKn             int         `json:"autopauseTimeoutKn,omitempty" env:"AUTPAUSE_TIMEOUT_KN,omitempty"`
	AutpauseKnockInterface         string      `json:"autpauseKnockInterface,omitempty" env:"AUTOPAUSE_KNOCK_INTERFACE,omitempty"`
	BroadcastConsoleToOps          bool        `json:"broadcastConsoleToOps,omitempty" env:"BROADCAST_CONSOLE_TO_OPS,omitempty"`
	BroadcastRconToOps             bool        `json:"broadcastRconToOps,omitempty" env:"BROADCAST_RCON_TO_OPS,omitempty"`
	BuildFromSource                bool        `json:"buildFromSource,omitempty" env:"BUILD_FROM_SOURCE,omitempty"`
	BukkitDownloadURL              string      `json:"bukkitDownloadURL,omitempty" env:"BUKKIT_DOWNLOAD_URL,omitempty"`
	Console                        bool        `json:"console,omitempty" env:"CONSOLE,omitempty"` // default True
	CopyConfigDest                 string      `json:"copyConfigDest,omitempty" env:"COPY_CONFIG_DEST,omitempty"`
	CopyModsDest                   string      `json:"copyModsDest,omitempty" env:"COPY_MODS_DEST,omitempty"`
	CurseForgeServerMod            string      `json:"curseForgeServerMod,omitempty" env:"CF_SERVER_MOD,omitempty"`
	CustomServer                   string      `json:"customServer,omitempty" env:"CUSTOM_SERVER,omitempty"`
	DisableHealthcheck             bool        `json:"disableHealthcheck,omitempty" env:"DISABLE_HEALTHCHECK,omitempty"`
	EULA                           bool        `json:"EULA,omitempty" env:"EULA,omitempty"`
	EnableAutopause                bool        `json:"enableAutopause,omitempty" env:"ENABLE_AUTOPAUSE,omitempty"`
	EnableJmxMonitoring            bool        `json:"enableJmxMonitoring,omitempty" env:"ENABLE_JMX_MONITORING,omitempty"`
	EnableQuery                    bool        `json:"enableQuery,omitempty" env:"ENABLE_QUERY,omitempty"`
	EnableRcon                     bool        `json:"enableRcon,omitempty" env:"ENABLE_RCON,omitempty"`
	EnableRollingLogs              bool        `json:"enableRollingLogs,omitempty" env:"ENABLE_ROLLING_LOGS,omitempty"`
	EnableStatus                   bool        `json:"enableStatus,omitempty" env:"ENABLE_STATUS,omitempty"`
	EnforceWhitelist               bool        `json:"enforceWhitelist,omitempty" env:"ENFORCE_WHITELIST,omitempty"`
	EntityBroadcastRangePercentage int         `json:"entityBroadcastRangePercentage,omitempty" env:"ENTITY_BROADCAST_RANGE_PERCENTAGE,omitempty"`
	ExecDirectly                   bool        `json:"execDirectly,omitempty" env:"EXEC_DIRECTLY,omitempty"`
	FTBLegacyJavaFixer             bool        `json:"FTBLegacyJavaFixer,omitempty" env:"FTB_LEGACYJAVAFIXER,omitempty"`
	FTBModpackID                   int         `json:"FTBModpackID,omitempty" env:"FTB_MODPACK_ID,omitempty"`
	FTBModpackVersionID            int         `json:"FTBModpackVersionID,omitempty" env:"FTB_MODPACK_VERSION_ID,omitempty"`
	FabricInstalerURL              string      `json:"fabricInstalerURL,omitempty" env:"FABRIC_INSTALLER_URL,omitempty"`
	FabricInstaller                string      `json:"fabricInstaller,omitempty" env:"FABRIC_INSTALLER,omitempty"`
	ForceRedownload                bool        `json:"forceRedownload,omitempty" env:"FORCE_REDOWNLOAD,omitempty"`
	ForceWorldCopy                 bool        `json:"forceWorldCopy,omitempty" env:"FORCE_WORLD_COPY,omitempty"`
	ForgeInstaller                 string      `json:"forgeInstaller,omitempty" env:"FORGE_INSTALLER,omitempty"`
	ForgeInstallerURL              string      `json:"forgeInstallerURL,omitempty" env:"FORGE_INSTALLER_URL,omitempty"`
	ForgeVersion                   string      `json:"forgeVersion,omitempty" env:"FORGE_VERSION,omitempty"`
	FunctionPermissionLevel        int         `json:"functionPermissionLevel,omitempty" env:"FUNCTION_PERMISSION_LEVEL,omitempty"`
	GID                            int         `json:"GID,omitempty" env:"GID,omitempty"`
	GUI                            bool        `json:"GUI,omitempty" env:"GUI,omitempty"`
	GenerateStructures             bool        `json:"generateStructures,omitempty" env:"GENERATE_STRUCTURES,omitempty"`
	GeneratorSettings              string      `json:"generatorSettings,omitempty" env:"GENERATOR_SETTINGS,omitempty"`
	Icon                           string      `json:"icon,omitempty" env:"ICON,omitempty"`
	InitMemory                     string      `json:"initMemory,omitempty" env:"INIT_MEMORY,omitempty"`
	JVMDDOpts                      string      `json:"JVMDDOpts,omitempty" env:"JVM_DD_OPTS,omitempty"`
	JVMOpts                        string      `json:"JVMOpts,omitempty" env:"JVM_OPTS,omitempty"`
	JVMXXOpts                      string      `json:"JVMXXOpts,omitempty" env:"JVM_XX_OPTS,omitempty"`
	MaxMemory                      string      `json:"maxMemory,omitempty" env:"MAX_MEMORY,omitempty"`
	MaxPlayers                     int         `json:"maxPlayers,omitempty" env:"MAX_PLAYERS,omitempty"`
	MaxTickTime                    int         `json:"maxTickTime,omitempty" env:"MAX_TICK_TIME,omitempty"`
	MaxWorldSize                   int         `json:"maxWorldSize,omitempty" env:"MAX_WORLD_SIZE,omitempty"`
	Memory                         string      `json:"memory,omitempty" env:"MEMORY,omitempty"`
	MohistBuild                    string      `json:"mohistBuild,omitempty" env:"MOHIST_BUILD,omitempty"`
	MOTD                           string      `json:"MOTD,omitempty" env:"MOTD,omitempty"`
	NetworkCompressionThreshold    int         `json:"networkCompressionThreshold,omitempty" env:"NETWORK_COMPRESSION_THRESHOLD,omitempty"`
	OpPermissionLevel              int         `json:"opPermissionLevel,omitempty" env:"OP_PERMISSION_LEVEL,omitempty"`
	OverrideIcon                   string      `json:"overrideIcon,omitempty" env:"OVERRIDE_ICON,omitempty"`
	OverrideServerProperties       bool        `json:"overrideServerProperties,omitempty" env:"OVERRIDE_SERVER_PROPERTIES,omitempty"`
	PaperDownloadURL               string      `json:"paperDownloadURL,omitempty" env:"PAPER_DOWNLOAD_URL,omitempty"`
	PlayerIdleTimeout              int         `json:"playerIdleTimeout,omitempty" env:"PLAYER_IDLE_TIMEOUT,omitempty"`
	PreventProxyConnections        bool        `json:"preventProxyConnections,omitempty" env:"PREVENT_PROXY_CONNECTIONS,omitempty"`
	Proxy                          string      `json:"proxy,omitempty" env:"PROXY,omitempty"`
	PurpurBuild                    string      `json:"purpurBuild,omitempty" env:"PURPUR_BUILD,omitempty"`
	Pvp                            bool        `json:"pvp,omitempty" env:"PVP,omitempty"`
	QueryPort                      int         `json:"queryPort,omitempty" env:"QUERY_PORT,omitempty"`
	RconPassword                   string      `json:"rconPassword,omitempty" env:"RCON_PASSWORD,omitempty"`
	RconPort                       int         `json:"rconPort,omitempty" env:"RCON_PORT,omitempty"`
	Release                        string      `json:"release,omitempty" env:"RELEASE,omitempty"`
	RemoveOldMods                  bool        `json:"removeOldMods,omitempty" env:"REMOVE_OLD_MODS,omitempty"`
	RemoveOldModsDepth             int         `json:"removeOldModsDepth,omitempty" env:"REMOVE_OLD_MODS_DEPTH,omitempty"`
	RemoveOldModsExclude           string      `json:"removeOldModsExclude,omitempty" env:"REMOVE_OLD_MODS_EXCLUDE,omitempty"`
	RemoveOldModsInclude           string      `json:"removeOldModsInclude,omitempty" env:"REMOVE_OLD_MODS_INCLUDE,omitempty"`
	ServerIp                       string      `json:"serverIp,omitempty" env:"SERVER_IP,omitempty"`
	ServerName                     string      `json:"serverName,omitempty" env:"SERVER_NAME,omitempty"`
	ServerPort                     int         `json:"serverPort,omitempty" env:"SERVER_PORT,omitempty"`
	SnooperEnabled                 bool        `json:"snooperEnabled,omitempty" env:"SNOOPER_ENABLED,omitempty"`
	SpigetResources                []int       `json:"spigetResources,omitempty" env:"SPIGET_RESOURCES,omitempty"`
	SpigotDownloadURL              string      `json:"spigotDownloadURL,omitempty" env:"SPIGOT_DOWNLOAD_URL,omitempty"`
	SpongeBranch                   string      `json:"spongeBranch,omitempty" env:"SPONGEBRANCH,omitempty"`
	StopDuration                   int         `json:"stopDuration,omitempty" env:"STOP_DURATION,omitempty"`
	SyncChunkWrites                bool        `json:"syncChunkWrites,omitempty" env:"SYNC_CHUNK_WRITES,omitempty"`
	Timezone                       string      `json:"timezone,omitempty" env:"TZ,omitempty"`
	Type                           ServerTypes `json:"type,omitempty" env:"TYPE,omitempty"`
	UID                            int         `json:"UID,omitempty" env:"UID,omitempty"`
	UseAikarFlags                  bool        `json:"useAikarFlags,omitempty" env:"USE_AIKAR_FLAGS,omitempty"`
	UseFlareFlags                  bool        `json:"useFlareFlags,omitempty" env:"USE_FLARE_FLAGS,omitempty"`
	UseLargePages                  bool        `json:"useLargePages,omitempty" env:"USE_LARGE_PAGES,omitempty"`
	UseNativeTransport             bool        `json:"useNativeTransport,omitempty" env:"USE_NATIVE_TRANSPORT,omitempty"`
	Version                        string      `json:"version,omitempty" env:"VERSION,omitempty"`
	World                          string      `json:"world,omitempty" env:"WORLD,omitempty"`
	WorldIndex                     int         `json:"worldIndex,omitempty" env:"WORLD_INDEX,omitempty"`
}
