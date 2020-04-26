package services

type Minecraft struct {
	Version     MinecraftServerVersion `json:"version"`
	Players     MinecraftPlayersInfo   `json:"players"`
	ForgeData   MinecraftForgeInfo     `json:"forgeData"`
	ModInfo     MinecraftModInfo       `json:"modinfo,omitempty"`
	Description string                 `json:"description"`
	Favicon     string                 `json:"favicon,omitempty"`
	Whitelisted bool                   `json:"whitelisted,omitempty"`
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

type MinecraftModInfo struct {
	Type    string        `json:"type"`
	ModList []ModInfoItem `json:"modList"`
}

type ModInfoItem struct {
	Version string `json:"version"`
	ModId   string `json:"modId"`
}

type MinecraftPlayer struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
