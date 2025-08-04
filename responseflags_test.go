package fortniteapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CombineFlags_Empty(t *testing.T) {
	result := CombineFlags()
	assert.Equal(t, ResponseFlag(0), result)
}

func Test_CombineFlags_Single(t *testing.T) {
	result := CombineFlags(FlagIncludePaths)
	assert.Equal(t, FlagIncludePaths, result)
}

func Test_CombineFlags_Multiple(t *testing.T) {
	result := CombineFlags(FlagIncludePaths, FlagIncludeGameplayTags)
	expected := FlagIncludePaths | FlagIncludeGameplayTags

	assert.Equal(t, expected, result)
}

func Test_CombineFlags_Duplicate(t *testing.T) {
	result := CombineFlags(FlagIncludePaths, FlagIncludePaths)
	assert.Equal(t, FlagIncludePaths, result)
}

func Test_CombineFlags_AllConstant(t *testing.T) {
	expected := FlagIncludePaths | FlagIncludeGameplayTags | FlagIncludeShopHistory
	assert.Equal(t, expected, FlagAll)
}
