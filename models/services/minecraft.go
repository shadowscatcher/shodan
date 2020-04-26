package services

type Minecraft struct {
	Version     MinecraftServerVersion `json:"version"`
	Players     MinecraftPlayersInfo   `json:"players"`
	ForgeData   MinecraftForgeInfo     `json:"forgeData"`
	Description string                 `json:"description"`
}

type MinecraftServerVersion struct {
	Protocol int    `json:"protocol"`
	Name     string `json:"name"`
}

type MinecraftPlayersInfo struct {
	Max    int `json:"max"`
	Online int `json:"online"`
}

type MinecraftForgeInfo struct {
	Channels          []MinecraftForgeChannel `json:"channels"`
	Mods              []MinecraftMod          `json:"mods"`
	FmlNetworkVersion int                     `json:"fmlNetworkVersion"`
}

type MinecraftForgeChannel struct {
	Res      string `json:"res"`
	Version  string `json:"version"`
	Required bool   `json:"required"`
}

type MinecraftMod struct {
	ModMarker string `json:"modmarker"`
	ModId     string `json:"modId"`
}
