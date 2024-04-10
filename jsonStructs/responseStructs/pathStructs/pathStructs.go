package pathStructs

type ModConfigPath struct {
	Configs    []string `json:"configs"`
	ConfigNums int      `json:"confignums"`
}

type ModPath struct {
	Mods    []string `json:"mods"`
	ModNums int      `json:"modnums"`
}
