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
	"time"

	querypkg "github.com/google/go-querystring/query"
)

const (
	Version = "v1.0.0"
	BaseURL = "https://fortnite-api.com"
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
	HTTPClient *http.Client
	Language
	APIKey string
}

func NewClient(language Language, apiKey string) *Client {
	return &Client{
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
		Language:   language,
		APIKey:     apiKey,
	}
}

func (c *Client) Fetch(ctx context.Context, method, path string, query, body, out any) error {
	u, err := url.Parse(BaseURL + path)
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

	if c.Language != "" && !params.Has("language") {
		params.Set("language", string(c.Language))
	}

	u.RawQuery = params.Encode()

	var bodyReader io.Reader
	if body != nil {
		jsonBytes, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}

		bodyReader = bytes.NewReader(jsonBytes)
	} else {
		bodyReader = nil
	}

	request, err := http.NewRequestWithContext(ctx, method, u.String(), bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create new request: %w", err)
	}

	request.Header.Set("User-Agent", "go-fortnite-api/"+Version)

	if c.APIKey != "" {
		request.Header.Set("Authorization", c.APIKey)
	}

	if body != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	defer func() {
		io.Copy(io.Discard, response.Body)
		response.Body.Close()
	}()

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
	return c.Fetch(ctx, "GET", path, params, nil, result)
}

func (c *Client) checkAPIKey() error {
	if c.APIKey == "" {
		return ErrNoAPIKey
	}

	return nil
}

func (c *Client) GetAESKey(ctx context.Context, params *AESKeyParams) (AESKeyResponse, error) {
	var result AESKeyResponse
	err := c.Get(ctx, "/v2/aes", params, &result)
	return result, err
}

func (c *Client) GetBanners(ctx context.Context, params *BannersParams) (BannersResponse, error) {
	var result BannersResponse
	err := c.Get(ctx, "/v1/banners", params, &result)
	return result, err
}

func (c *Client) GetBannerColors(ctx context.Context) (BannerColorsResponse, error) {
	var result BannerColorsResponse
	err := c.Get(ctx, "/v1/banners/colors", nil, &result)
	return result, err
}

func (c *Client) GetAllCosmetics(ctx context.Context, params *AllCosmeticsParams) (AllCosmeticsResponse, error) {
	var result AllCosmeticsResponse
	err := c.Get(ctx, "/v2/cosmetics", params, &result)
	return result, err
}

func (c *Client) GetNewCosmetics(ctx context.Context, params *NewCosmeticsParams) (NewCosmeticsResponse, error) {
	var result NewCosmeticsResponse
	err := c.Get(ctx, "/v2/cosmetics/new", params, &result)
	return result, err
}

func (c *Client) GetBRCosmeticsList(ctx context.Context, params *BRCosmeticsListParams) (BRCosmeticsListResponse, error) {
	var result BRCosmeticsListResponse
	err := c.Get(ctx, "/v2/cosmetics/br", params, &result)
	return result, err
}

func (c *Client) GetTrackCosmeticsList(ctx context.Context, params *TrackCosmeticsListParams) (TrackCosmeticsListResponse, error) {
	var result TrackCosmeticsListResponse
	err := c.Get(ctx, "/v2/cosmetics/tracks", params, &result)
	return result, err
}

func (c *Client) GetInstrumentCosmeticsList(ctx context.Context, params *InstrumentCosmeticsListParams) (InstrumentCosmeticsListResponse, error) {
	var result InstrumentCosmeticsListResponse
	err := c.Get(ctx, "/v2/cosmetics/instruments", params, &result)
	return result, err
}

func (c *Client) GetCarCosmeticsList(ctx context.Context, params *CarCosmeticsListParams) (CarCosmeticsListResponse, error) {
	var result CarCosmeticsListResponse
	err := c.Get(ctx, "/v2/cosmetics/cars", params, &result)
	return result, err
}

func (c *Client) GetLegoCosmeticsList(ctx context.Context, params *LegoCosmeticsListParams) (LegoCosmeticsListResponse, error) {
	var result LegoCosmeticsListResponse
	err := c.Get(ctx, "/v2/cosmetics/lego", params, &result)
	return result, err
}

func (c *Client) GetLegoKitCosmeticsList(ctx context.Context, params *LegoKitCosmeticsListParams) (LegoKitCosmeticsListResponse, error) {
	var result LegoKitCosmeticsListResponse
	err := c.Get(ctx, "/v2/cosmetics/lego/kits", params, &result)
	return result, err
}

func (c *Client) GetBeanCosmeticsList(ctx context.Context, params *BeanCosmeticsListParams) (BeanCosmeticsListResponse, error) {
	var result BeanCosmeticsListResponse
	err := c.Get(ctx, "/v2/cosmetics/beans", params, &result)
	return result, err
}

func (c *Client) GetBRCosmeticByID(ctx context.Context, id string, params *BRCosmeticByIDParams) (BRCosmeticByIDResponse, error) {
	var result BRCosmeticByIDResponse

	if id == "" {
		return result, emptyParamErr("id")
	}

	err := c.Get(ctx, fmt.Sprintf("/v2/cosmetics/br/%s", id), params, &result)
	return result, err
}

func (c *Client) SearchBRCosmetic(ctx context.Context, params *SearchBRCosmeticParams) (SearchBRCosmeticResponse, error) {
	var result SearchBRCosmeticResponse

	if params == nil {
		params = &SearchBRCosmeticParams{}
	}

	if params.SearchLanguage == "" {
		params.SearchLanguage = c.Language
	}

	err := c.Get(ctx, "/v2/cosmetics/br/search", params, &result)
	return result, err
}

func (c *Client) SearchBRCosmetics(ctx context.Context, params *SearchBRCosmeticsParams) (SearchBRCosmeticsResponse, error) {
	var result SearchBRCosmeticsResponse

	if params == nil {
		params = &SearchBRCosmeticsParams{}
	}

	if params.SearchLanguage == "" {
		params.SearchLanguage = c.Language
	}

	err := c.Get(ctx, "/v2/cosmetics/br/search/all", params, &result)
	return result, err
}

func (c *Client) SearchBRCosmeticsByIDs(ctx context.Context, ids []string, params *BRCosmeticsByIDsParams) (BRCosmeticsByIDsResponse, error) {
	var result BRCosmeticsByIDsResponse

	if len(ids) == 0 {
		return result, emptyParamErr("ids")
	}

	err := c.Fetch(ctx, "POST", "/v2/cosmetics/br/search/ids", params, ids, &result)
	return result, err
}

func (c *Client) GetCreatorCode(ctx context.Context, name string, params *CreatorCodeParams) (CreatorCodeResponse, error) {
	var result CreatorCodeResponse

	if name == "" {
		return result, emptyParamErr("name")
	}

	if params == nil {
		params = &CreatorCodeParams{}
	}

	params.Name = name

	err := c.Get(ctx, "/v2/creatorcode", params, &result)
	return result, err
}

func (c *Client) GetBRMap(ctx context.Context, params *BRMapParams) (BRMapResponse, error) {
	var result BRMapResponse
	err := c.Get(ctx, "/v1/map", params, &result)
	return result, err
}

func (c *Client) GetNews(ctx context.Context, params *NewsParams) (NewsResponse, error) {
	var result NewsResponse
	err := c.Get(ctx, "/v2/news", params, &result)
	return result, err
}

func (c *Client) GetBRNews(ctx context.Context, params *BRNewsParams) (BRNewsResponse, error) {
	var result BRNewsResponse
	err := c.Get(ctx, "/v2/news/br", params, &result)
	return result, err
}

func (c *Client) GetSTWNews(ctx context.Context, params *STWNewsParams) (STWNewsResponse, error) {
	var result STWNewsResponse
	err := c.Get(ctx, "/v2/news/stw", params, &result)
	return result, err
}

func (c *Client) GetCreativeNews(ctx context.Context, params *CreativeNewsParams) (CreativeNewsResponse, error) {
	var result CreativeNewsResponse
	err := c.Get(ctx, "/v2/news/creative", params, &result)
	return result, err
}

func (c *Client) GetPlaylists(ctx context.Context, params *PlaylistsParams) (PlaylistsResponse, error) {
	var result PlaylistsResponse
	err := c.Get(ctx, "/v1/playlists", params, &result)
	return result, err
}

func (c *Client) GetPlaylistByID(ctx context.Context, id string, params *PlaylistByIDParams) (PlaylistByIDResponse, error) {
	var result PlaylistByIDResponse

	if id == "" {
		return result, emptyParamErr("id")
	}

	err := c.Get(ctx, fmt.Sprintf("/v1/playlists/%s", id), params, &result)
	return result, err
}

func (c *Client) GetShop(ctx context.Context, params *ShopParams) (ShopResponse, error) {
	var result ShopResponse
	err := c.Get(ctx, "/v2/shop", params, &result)
	return result, err
}

func (c *Client) GetBRStatsByName(ctx context.Context, name string, params *BRStatsByNameParams) (BRStatsByNameResponse, error) {
	var result BRStatsByNameResponse

	if err := c.checkAPIKey(); err != nil {
		return result, err
	}

	if name == "" {
		return result, emptyParamErr("name")
	}

	if params == nil {
		params = &BRStatsByNameParams{}
	}

	params.Name = name

	err := c.Get(ctx, "/v2/stats/br/v2", params, &result)
	return result, err
}

func (c *Client) GetBRStatsByID(ctx context.Context, id string, params *BRStatsByIDParams) (BRStatsByIDResponse, error) {
	var result BRStatsByIDResponse

	if err := c.checkAPIKey(); err != nil {
		return result, err
	}

	if id == "" {
		return result, emptyParamErr("id")
	}

	err := c.Get(ctx, fmt.Sprintf("/v2/stats/br/v2/%s", id), params, &result)
	return result, err
}

func emptyParamErr(name string) error {
	return fmt.Errorf("%s %w", name, ErrEmptyParameter)
}
