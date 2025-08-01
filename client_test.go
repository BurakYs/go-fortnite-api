package fortniteapi

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

const (
	testCosmeticName = "Peely"
	testCosmeticID1  = "CID_349_Athena_Commando_M_Banana"
	testCosmeticID2  = "CID_049_Athena_Commando_M_HolidayGingerbread"

	testStatsName = "BurakYhs"
	testStatsID   = "05006cb489c347beaad83551a1b9544e"

	testPlaylistID  = "Playlist_DefaultSolo"
	testCreatorCode = "Ninja"
)

var (
	testClient *Client
	testCtx    = context.Background()
)

func requireAPIKey(t *testing.T) {
	if err := testClient.checkAPIKey(); err != nil {
		t.Skip("API_KEY is not set")
	}
}

func TestMain(m *testing.M) {
	godotenv.Load()

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Println("API_KEY is not set in .env file, skipping tests that require it")
	}

	testClient = NewClient(LanguageEnglish, apiKey)

	os.Exit(m.Run())
}

func Test_GetAESKey(t *testing.T) {
	_, err := testClient.GetAESKey(testCtx, nil)
	assert.NoError(t, err)
}

func Test_GetBanners(t *testing.T) {
	_, err := testClient.GetBanners(testCtx, nil)
	assert.NoError(t, err)
}

func Test_GetBannerColors(t *testing.T) {
	_, err := testClient.GetBannerColors(testCtx)
	assert.NoError(t, err)
}

func Test_GetAllCosmetics(t *testing.T) {
	_, err := testClient.GetAllCosmetics(testCtx, nil)
	assert.NoError(t, err)
}

func Test_GetNewCosmetics(t *testing.T) {
	_, err := testClient.GetNewCosmetics(testCtx, nil)
	assert.NoError(t, err)
}

func Test_GetBRCosmeticsList(t *testing.T) {
	_, err := testClient.GetBRCosmeticsList(testCtx, nil)
	assert.NoError(t, err)
}

func Test_GetTrackCosmeticsList(t *testing.T) {
	_, err := testClient.GetTrackCosmeticsList(testCtx, nil)
	assert.NoError(t, err)
}

func Test_GetInstrumentCosmeticsList(t *testing.T) {
	_, err := testClient.GetInstrumentCosmeticsList(testCtx, nil)
	assert.NoError(t, err)
}

func Test_GetCarCosmeticsList(t *testing.T) {
	_, err := testClient.GetCarCosmeticsList(testCtx, nil)
	assert.NoError(t, err)
}

func Test_GetLegoCosmeticsList(t *testing.T) {
	_, err := testClient.GetLegoCosmeticsList(testCtx, nil)
	assert.NoError(t, err)
}

func Test_GetLegoKitCosmeticsList(t *testing.T) {
	_, err := testClient.GetLegoKitCosmeticsList(testCtx, nil)
	assert.NoError(t, err)
}

func Test_GetBeanCosmeticsList(t *testing.T) {
	_, err := testClient.GetBeanCosmeticsList(testCtx, nil)
	assert.NoError(t, err)
}

func Test_GetBRCosmeticByID(t *testing.T) {
	resp, err := testClient.GetBRCosmeticByID(testCtx, testCosmeticID1, nil)
	assert.NoError(t, err)
	assert.Equal(t, testCosmeticID1, resp.ID)
}

func Test_GetBRCosmeticByID_WithEmptyID(t *testing.T) {
	_, err := testClient.GetBRCosmeticByID(testCtx, "", nil)
	assert.ErrorIs(t, err, ErrEmptyParameter)
}

func Test_SearchBRCosmetic(t *testing.T) {
	resp, err := testClient.SearchBRCosmetic(testCtx, &SearchBRCosmeticParams{Name: testCosmeticName})
	assert.NoError(t, err)
	assert.Equal(t, testCosmeticName, resp.Name)
}

func Test_SearchBRCosmetics(t *testing.T) {
	_, err := testClient.SearchBRCosmetics(testCtx, &SearchBRCosmeticsParams{Name: testCosmeticName})
	assert.NoError(t, err)
}

func Test_SearchBRCosmeticsByIDs(t *testing.T) {
	ids := []string{testCosmeticID1, testCosmeticID2}
	resp, err := testClient.SearchBRCosmeticsByIDs(testCtx, ids, nil)

	assert.NoError(t, err)
	assert.Len(t, resp, 2)

	for _, cosmetic := range resp {
		assert.Contains(t, ids, cosmetic.ID)
	}
}

func Test_SearchBRCosmeticsByIDs_WithEmptyIDs(t *testing.T) {
	_, err := testClient.SearchBRCosmeticsByIDs(testCtx, []string{}, nil)
	assert.ErrorIs(t, err, ErrEmptyParameter)
}

func Test_GetCreatorCode(t *testing.T) {
	_, err := testClient.GetCreatorCode(testCtx, testCreatorCode, nil)
	assert.NoError(t, err)
}

func Test_GetCreatorCode_WithEmptyCode(t *testing.T) {
	_, err := testClient.GetCreatorCode(testCtx, "", nil)
	assert.ErrorIs(t, err, ErrEmptyParameter)
}

func Test_GetBRMap(t *testing.T) {
	_, err := testClient.GetBRMap(testCtx, nil)
	assert.NoError(t, err)
}

func Test_GetNews(t *testing.T) {
	_, err := testClient.GetNews(testCtx, nil)
	assert.NoError(t, err)
}

func Test_GetBRNews(t *testing.T) {
	_, err := testClient.GetBRNews(testCtx, nil)
	assert.NoError(t, err)
}

func Test_GetSTWNews(t *testing.T) {
	_, err := testClient.GetSTWNews(testCtx, nil)
	assert.NoError(t, err)
}

func Test_GetCreativeNews(t *testing.T) {
	t.Skip("Creative news are not available anymore")

	_, err := testClient.GetCreativeNews(testCtx, nil)
	assert.NoError(t, err)
}

func Test_GetPlaylists(t *testing.T) {
	_, err := testClient.GetPlaylists(testCtx, nil)
	assert.NoError(t, err)
}

func Test_GetPlaylistByID(t *testing.T) {
	resp, err := testClient.GetPlaylistByID(testCtx, testPlaylistID, nil)
	assert.NoError(t, err)
	assert.Equal(t, testPlaylistID, resp.ID)
}

func Test_GetPlaylistByID_WithEmptyID(t *testing.T) {
	_, err := testClient.GetPlaylistByID(testCtx, "", nil)
	assert.ErrorIs(t, err, ErrEmptyParameter)
}

func Test_GetShop(t *testing.T) {
	_, err := testClient.GetShop(testCtx, nil)
	assert.NoError(t, err)
}

func Test_GetBRStatsByName(t *testing.T) {
	requireAPIKey(t)

	resp, err := testClient.GetBRStatsByName(testCtx, testStatsName, nil)
	assert.NoError(t, err)
	assert.Equal(t, testStatsName, resp.Account.Name)
}

func Test_GetBRStatsByName_WithEmptyName(t *testing.T) {
	requireAPIKey(t)

	_, err := testClient.GetBRStatsByName(testCtx, "", nil)
	assert.ErrorIs(t, err, ErrEmptyParameter)
}

func Test_GetBRStatsByAccountID(t *testing.T) {
	requireAPIKey(t)

	resp, err := testClient.GetBRStatsByID(testCtx, testStatsID, nil)
	assert.NoError(t, err)
	assert.Equal(t, testStatsID, resp.Account.ID)
}

func Test_GetBRStatsByAccountID_WithEmptyID(t *testing.T) {
	requireAPIKey(t)

	_, err := testClient.GetBRStatsByID(testCtx, "", nil)
	assert.ErrorIs(t, err, ErrEmptyParameter)
}
