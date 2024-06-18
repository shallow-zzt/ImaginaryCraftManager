package serverConfig

import (
	"ImaginaryCraftManager/generic/serverCmd"
	"ImaginaryCraftManager/jsonStructs/requestStructs/serverSettingStructs"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadServerConfig(filePath string) (map[string]string, error) {
	properties := make(map[string]string)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line: %s", line)
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		properties[key] = value
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return properties, nil
}

func WriteProperty(filename, key, value string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")

	found := false
	for i, line := range lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 && strings.TrimSpace(parts[0]) == key {
			lines[i] = fmt.Sprintf("%s=%s", key, value)
			found = true
			break
		}
	}

	if !found {
		lines = append(lines, fmt.Sprintf("%s=%s", key, value))
	}

	updatedContent := strings.Join(lines, "\n")
	err = os.WriteFile(filename, []byte(updatedContent), 0644)
	if err != nil {
		return err
	}

	return nil
}

func WriteServerConfig2Json(serverConfig map[string]string, serverPath string) *serverSettingStructs.AllServerSettings {
	var General serverSettingStructs.ServerGeneralSetting
	var World serverSettingStructs.ServerWorldSetting
	var Network serverSettingStructs.ServerNetworkingSetting
	var Player serverSettingStructs.ServerPlayerSetting
	var Resources serverSettingStructs.ServerResourcesPackSetting
	var Additional serverSettingStructs.ServerAdditionalSetting
	var ServerJson serverSettingStructs.AllServerSettings

	General.Motd = serverConfig["motd"]
	General.MaxPlayers = serverConfig["max-players"]
	General.Gamemode = serverConfig["gamemode"]
	General.Difficulty = serverConfig["difficulty"]
	General.IsOnlineMode = serverConfig["online-mode"]
	General.IsPvp = serverConfig["pvp"]
	General.IsGenerateStructures = serverConfig["generate-structures"]

	World.LevelSeed = serverConfig["level-seed"]
	World.MaxWorldSize = serverConfig["max-world-size"]
	World.LevelType = serverConfig["level-type"]
	World.IsAllowNether = serverConfig["allow-nether"]
	World.GeneratorSettings = serverConfig["generator-settings"]

	Network.ServerIp = serverConfig["server-ip"]
	Network.ServerPort = serverConfig["server-port"]
	Network.ServerMemory = serverCmd.GetCmdParameter(serverPath)
	Network.ViewDistance = serverConfig["view-distance"]
	Network.SimulationDistance = serverConfig["simulation-distance"]
	Network.NetworkCompressionThershold = serverConfig["network-compression-threshold"]
	Network.IsUseNativeTransport = serverConfig["use-native-transport"]
	Network.MaxTickTime = serverConfig["max-tick-time"]

	Player.IsWhiteList = serverConfig["white-list"]
	Player.IsEnforceWhitelist = serverConfig["enforce-whitelist"]
	Player.OpPermissionLevel = serverConfig["op-permission-level"]
	Player.SpawnProtection = serverConfig["spawn-protection"]

	Resources.IsRequireResourcePack = serverConfig["require-resource-pack"]
	Resources.ResourcePack = serverConfig["resource-pack"]
	Resources.ResourcePackPrompt = serverConfig["resource-pack-prompt"]

	Additional.IsAllowFlight = serverConfig["allow-flight"]
	Additional.IsEnableCommandBlock = serverConfig["enable-command-block"]
	Additional.IsEnableQuery = serverConfig["enable-query"]
	Additional.IsSyncChunkWrites = serverConfig["sync-chunk-writes"]
	Additional.IsLogIps = serverConfig["log-ips"]
	Additional.FunctionPermissionLevel = serverConfig["function-permission-level"]
	Additional.TextFliteringConfig = serverConfig["text-filtering-config"]
	Additional.IsEnforceSecureProfile = serverConfig["enforce-secure-profile"]

	ServerJson.General = General
	ServerJson.World = World
	ServerJson.Network = Network
	ServerJson.Player = Player
	ServerJson.Resources = Resources
	ServerJson.Additional = Additional

	return &ServerJson
}
