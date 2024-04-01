package serverSettingStructs

type ServerGeneralSetting struct {
	Motd                 string `json:"motd"`
	MaxPlayers           string `json:"max-players"`
	Gamemode             string `json:"gamemode"`
	Difficulty           string `json:"difficulty"`
	IsOnlineMode         string `json:"online-mode"`
	IsPvp                string `json:"pvp"`
	IsGenerateStructures string `json:"generate-structures"`
}

type ServerWorldSetting struct {
	LevelSeed         string `json:"level-seed"`
	MaxWorldSize      string `json:"max-world-size"`
	LevelType         string `json:"level-type"`
	IsAllowNether     string `json:"allow-nether"`
	GeneratorSettings string `json:"generator-settings"`
}

type ServerNetworkingSetting struct {
	ServerIp                    string `json:"server-ip"`
	ServerPort                  string `json:"server-port"`
	ServerMemory                string `json:"server-memory"`
	ViewDistance                string `json:"view-distance"`
	SimulationDistance          string `json:"simulation-distance"`
	NetworkCompressionThershold string `json:"network-compression-threshold"`
	MaxTickTime                 string `json:"max-tick-time"`
	IsUseNativeTransport        string `json:"use-native-transport"`
}

type ServerPlayerSetting struct {
	IsWhiteList        string `json:"white-list"`
	IsEnforceWhitelist string `json:"enforce-whitelist"`
	SpawnProtection    string `json:"spawn-protection"`
	OpPermissionLevel  string `json:"op-permission-level"`
}

type ServerResourcesPackSetting struct {
	IsRequireResourcePack string `json:"require-resource-pack"`
	ResourcePack          string `json:"resource-pack"`
	ResourcePackPrompt    string `json:"resource-pack-prompt"`
}

type ServerAdditionalSetting struct {
	IsAllowFlight           string `json:"allow-flight"`
	IsEnableCommandBlock    string `json:"enable-command-block"`
	IsEnableQuery           string `json:"enable-query"`
	IsSyncChunkWrites       string `json:"sync-chunk-writes"`
	IsLogIps                string `json:"log-ips"`
	FunctionPermissionLevel string `json:"function-permission-level"`
	TextFliteringConfig     string `json:"text-filtering-config"`
	IsEnforceSecureProfile  string `json:"enforce-secure-profile"`
}

type AllServerSettings struct {
	General    ServerGeneralSetting       `json:"general"`
	World      ServerWorldSetting         `json:"world"`
	Network    ServerNetworkingSetting    `json:"network"`
	Player     ServerPlayerSetting        `json:"player"`
	Resources  ServerResourcesPackSetting `json:"resources"`
	Additional ServerAdditionalSetting    `json:"additional"`
}
