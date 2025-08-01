package fortniteapi

type ResponseFlag uint32

type ResponseFlagsParams struct {
	ResponseFlags ResponseFlag `url:"responseFlags,omitempty"`
}

const (
	FlagIncludePaths ResponseFlag = 1 << iota
	FlagIncludeGameplayTags
	FlagIncludeShopHistory

	FlagAll = FlagIncludePaths | FlagIncludeGameplayTags | FlagIncludeShopHistory
)

func CombineFlags(flags ...ResponseFlag) ResponseFlag {
	var result ResponseFlag
	for _, flag := range flags {
		result |= flag
	}

	return result
}
