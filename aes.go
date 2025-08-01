package fortniteapi

type AESKeyParams struct {
	// Enum: "aes", "hex"
	//
	// Default: "hex"
	KeyFormat     string       `url:"keyFormat"`
	ResponseFlags ResponseFlag `url:"responseFlags,omitempty"`
}

type AESDynamicKey struct {
	PakFilename string `json:"pakFilename"`
	PakGUID     string `json:"pakGuid"`
	Key         string `json:"key"`
}

type AESKeyResponse struct {
	Build       string          `json:"build"`
	MainKey     string          `json:"mainKey"`
	DynamicKeys []AESDynamicKey `json:"dynamicKeys"`
	Updated     string          `json:"updated"`
}
