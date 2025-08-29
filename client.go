package fortniteapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	querypkg "github.com/google/go-querystring/query"
)

const (
	version = "v1.0.0"
	baseURL = "https://fortnite-api.com"
)

var (
	ErrNoAPIKey       = errors.New("an API key is required for this request")
	ErrEmptyParameter = errors.New("parameter cannot be empty")
)

type APIResponse[T any] struct {
	Status int `json:"status"`
	Data   T   `json:"data"`
}

type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"error"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("api error: %d - %s", e.Status, e.Message)
}

type Client struct {
	language   Language
	httpClient *http.Client
	apiKey     string
}

func NewClient(language Language, apiKey string) *Client {
	return &Client{
		language:   language,
		httpClient: &http.Client{},
		apiKey:     apiKey,
	}
}

func (c *Client) Fetch(ctx context.Context, method, path string, query, body, out any) error {
	u, err := url.Parse(baseURL + path)
	if err != nil {
		return fmt.Errorf("failed to parse URL: %w", err)
	}

	params := url.Values{}

	if query != nil {
		if values, ok := query.(url.Values); ok {
			params = values
		} else {
			params, err = querypkg.Values(query)
			if err != nil {
				return fmt.Errorf("failed to encode query: %w", err)
			}
		}
	}

	if c.language != "" && !params.Has("language") {
		params.Set("language", string(c.language))
	}

	u.RawQuery = params.Encode()

	var bodyReader io.Reader
	if body != nil {
		jsonBytes, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}

		bodyReader = bytes.NewReader(jsonBytes)
	}

	request, err := http.NewRequestWithContext(ctx, method, u.String(), bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create new request: %w", err)
	}

	request.Header.Set("User-Agent", "go-fortnite-api/"+version)

	if c.apiKey != "" {
		request.Header.Set("Authorization", c.apiKey)
	}

	if body != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	defer response.Body.Close() //nolint:errcheck

	decoder := json.NewDecoder(response.Body)

	if response.StatusCode != http.StatusOK {
		var apiError APIError
		if err := decoder.Decode(&apiError); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}

		return &apiError
	}

	var apiResponse APIResponse[json.RawMessage]
	if err := decoder.Decode(&apiResponse); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if out != nil {
		if err := json.Unmarshal(apiResponse.Data, out); err != nil {
			return fmt.Errorf("failed to unmarshal data from response: %w", err)
		}
	}

	return nil
}

func (c *Client) Get(ctx context.Context, path string, params, result any) error {
	return c.Fetch(ctx, http.MethodGet, path, params, nil, result)
}

func (c *Client) GetAESKey(ctx context.Context, params *AESKeyParams) (*AESKeyResponse, error) {
	result := new(AESKeyResponse)
	err := c.Get(ctx, "/v2/aes", params, result)
	return result, err
}

func (c *Client) GetBanners(ctx context.Context, params *BannersParams) (*BannersResponse, error) {
	result := new(BannersResponse)
	err := c.Get(ctx, "/v1/banners", params, result)
	return result, err
}

func (c *Client) GetBannerColors(ctx context.Context) (*BannerColorsResponse, error) {
	result := new(BannerColorsResponse)
	err := c.Get(ctx, "/v1/banners/colors", nil, result)
	return result, err
}

func (c *Client) GetAllCosmetics(ctx context.Context, params *AllCosmeticsParams) (*AllCosmeticsResponse, error) {
	result := new(AllCosmeticsResponse)
	err := c.Get(ctx, "/v2/cosmetics", params, result)
	return result, err
}

func (c *Client) GetNewCosmetics(ctx context.Context, params *NewCosmeticsParams) (*NewCosmeticsResponse, error) {
	result := new(NewCosmeticsResponse)
	err := c.Get(ctx, "/v2/cosmetics/new", params, result)
	return result, err
}

func (c *Client) GetBRCosmeticsList(ctx context.Context, params *BRCosmeticsListParams) (*BRCosmeticsListResponse, error) {
	result := new(BRCosmeticsListResponse)
	err := c.Get(ctx, "/v2/cosmetics/br", params, result)
	return result, err
}

func (c *Client) GetTrackCosmeticsList(ctx context.Context, params *TrackCosmeticsListParams) (*TrackCosmeticsListResponse, error) {
	result := new(TrackCosmeticsListResponse)
	err := c.Get(ctx, "/v2/cosmetics/tracks", params, result)
	return result, err
}

func (c *Client) GetInstrumentCosmeticsList(ctx context.Context, params *InstrumentCosmeticsListParams) (*InstrumentCosmeticsListResponse, error) {
	result := new(InstrumentCosmeticsListResponse)
	err := c.Get(ctx, "/v2/cosmetics/instruments", params, result)
	return result, err
}

func (c *Client) GetCarCosmeticsList(ctx context.Context, params *CarCosmeticsListParams) (*CarCosmeticsListResponse, error) {
	result := new(CarCosmeticsListResponse)
	err := c.Get(ctx, "/v2/cosmetics/cars", params, result)
	return result, err
}

func (c *Client) GetLegoCosmeticsList(ctx context.Context, params *LegoCosmeticsListParams) (*LegoCosmeticsListResponse, error) {
	result := new(LegoCosmeticsListResponse)
	err := c.Get(ctx, "/v2/cosmetics/lego", params, result)
	return result, err
}

func (c *Client) GetLegoKitCosmeticsList(ctx context.Context, params *LegoKitCosmeticsListParams) (*LegoKitCosmeticsListResponse, error) {
	result := new(LegoKitCosmeticsListResponse)
	err := c.Get(ctx, "/v2/cosmetics/lego/kits", params, result)
	return result, err
}

func (c *Client) GetBeanCosmeticsList(ctx context.Context, params *BeanCosmeticsListParams) (*BeanCosmeticsListResponse, error) {
	result := new(BeanCosmeticsListResponse)
	err := c.Get(ctx, "/v2/cosmetics/beans", params, result)
	return result, err
}

func (c *Client) GetBRCosmeticByID(ctx context.Context, id string, params *BRCosmeticByIDParams) (*BRCosmeticByIDResponse, error) {
	if id == "" {
		return nil, emptyParamErr("id")
	}

	result := new(BRCosmeticByIDResponse)
	err := c.Get(ctx, fmt.Sprintf("/v2/cosmetics/br/%s", id), params, result)
	return result, err
}

func (c *Client) SearchBRCosmetic(ctx context.Context, params *SearchBRCosmeticParams) (*SearchBRCosmeticResponse, error) {
	if params == nil {
		params = &SearchBRCosmeticParams{}
	}

	if params.SearchLanguage == "" {
		params.SearchLanguage = c.language
	}

	result := new(SearchBRCosmeticResponse)
	err := c.Get(ctx, "/v2/cosmetics/br/search", params, result)
	return result, err
}

func (c *Client) SearchBRCosmetics(ctx context.Context, params *SearchBRCosmeticsParams) (*SearchBRCosmeticsResponse, error) {
	if params == nil {
		params = &SearchBRCosmeticsParams{}
	}

	if params.SearchLanguage == "" {
		params.SearchLanguage = c.language
	}

	result := new(SearchBRCosmeticsResponse)
	err := c.Get(ctx, "/v2/cosmetics/br/search/all", params, result)
	return result, err
}

func (c *Client) SearchBRCosmeticsByIDs(ctx context.Context, ids []string, params *BRCosmeticsByIDsParams) (*BRCosmeticsByIDsResponse, error) {
	if len(ids) == 0 {
		return nil, emptyParamErr("ids")
	}

	result := new(BRCosmeticsByIDsResponse)
	err := c.Fetch(ctx, "POST", "/v2/cosmetics/br/search/ids", params, ids, result)
	return result, err
}

func (c *Client) GetCreatorCode(ctx context.Context, name string, params *CreatorCodeParams) (*CreatorCodeResponse, error) {
	if name == "" {
		return nil, emptyParamErr("name")
	}

	if params == nil {
		params = &CreatorCodeParams{}
	}

	params.Name = name

	result := new(CreatorCodeResponse)
	err := c.Get(ctx, "/v2/creatorcode", params, result)
	return result, err
}

func (c *Client) GetBRMap(ctx context.Context, params *BRMapParams) (*BRMapResponse, error) {
	result := new(BRMapResponse)
	err := c.Get(ctx, "/v1/map", params, result)
	return result, err
}

func (c *Client) GetNews(ctx context.Context, params *NewsParams) (*NewsResponse, error) {
	result := new(NewsResponse)
	err := c.Get(ctx, "/v2/news", params, result)
	return result, err
}

func (c *Client) GetBRNews(ctx context.Context, params *BRNewsParams) (*BRNewsResponse, error) {
	result := new(BRNewsResponse)
	err := c.Get(ctx, "/v2/news/br", params, result)
	return result, err
}

func (c *Client) GetSTWNews(ctx context.Context, params *STWNewsParams) (*STWNewsResponse, error) {
	result := new(STWNewsResponse)
	err := c.Get(ctx, "/v2/news/stw", params, result)
	return result, err
}

func (c *Client) GetCreativeNews(ctx context.Context, params *CreativeNewsParams) (*CreativeNewsResponse, error) {
	result := new(CreativeNewsResponse)
	err := c.Get(ctx, "/v2/news/creative", params, result)
	return result, err
}

func (c *Client) GetPlaylists(ctx context.Context, params *PlaylistsParams) (*PlaylistsResponse, error) {
	result := new(PlaylistsResponse)
	err := c.Get(ctx, "/v1/playlists", params, result)
	return result, err
}

func (c *Client) GetPlaylistByID(ctx context.Context, id string, params *PlaylistByIDParams) (*PlaylistByIDResponse, error) {
	if id == "" {
		return nil, emptyParamErr("id")
	}

	result := new(PlaylistByIDResponse)
	err := c.Get(ctx, fmt.Sprintf("/v1/playlists/%s", id), params, result)
	return result, err
}

func (c *Client) GetShop(ctx context.Context, params *ShopParams) (*ShopResponse, error) {
	result := new(ShopResponse)
	err := c.Get(ctx, "/v2/shop", params, result)
	return result, err
}

func (c *Client) GetBRStatsByName(ctx context.Context, name string, params *BRStatsByNameParams) (*BRStatsResponse, error) {
	if err := c.checkAPIKey(); err != nil {
		return nil, err
	}

	if name == "" {
		return nil, emptyParamErr("name")
	}

	if params == nil {
		params = &BRStatsByNameParams{}
	}

	params.Name = name
	result := new(BRStatsResponse)
	err := c.Get(ctx, "/v2/stats/br/v2", params, result)
	return result, err
}

func (c *Client) GetBRStatsByID(ctx context.Context, id string, params *BRStatsByIDParams) (*BRStatsResponse, error) {
	if err := c.checkAPIKey(); err != nil {
		return nil, err
	}

	if id == "" {
		return nil, emptyParamErr("id")
	}

	result := new(BRStatsResponse)
	err := c.Get(ctx, fmt.Sprintf("/v2/stats/br/v2/%s", id), params, result)
	return result, err
}

func (c *Client) checkAPIKey() error {
	if c.apiKey == "" {
		return ErrNoAPIKey
	}

	return nil
}

func emptyParamErr(name string) error {
	return fmt.Errorf("%w: %s", ErrEmptyParameter, name)
}
