package fortniteapi

import "testing"

func TestNoFlags(t *testing.T) {
	result := CombineFlags()
	if result != 0 {
		t.Errorf("Expected 0, got %v", result)
	}
}

func TestCombineSingleFlag(t *testing.T) {
	result := CombineFlags(FlagIncludePaths)
	if result != FlagIncludePaths {
		t.Errorf("Expected %v, got %v", FlagIncludePaths, result)
	}
}

func TestCombineMultipleFlags(t *testing.T) {
	result := CombineFlags(FlagIncludePaths, FlagIncludeGameplayTags)
	expected := FlagIncludePaths | FlagIncludeGameplayTags

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCombineDuplicateFlags(t *testing.T) {
	result := CombineFlags(FlagIncludePaths, FlagIncludePaths)
	if result != FlagIncludePaths {
		t.Errorf("Expected %v, got %v", FlagIncludePaths, result)
	}
}

func TestFlagAllConstant(t *testing.T) {
	expected := FlagIncludePaths | FlagIncludeGameplayTags | FlagIncludeShopHistory
	if FlagAll != expected {
		t.Errorf("Expected %v, got %v", expected, FlagAll)
	}
}
