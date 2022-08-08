package brightdata

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	cache2 "github.com/patrickmn/go-cache"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

func NewClient(clientOptions BrightDataClientOptions) (*BrightDataClient, error) {

	if clientOptions.LuminatiUrl == "" {
		return nil, errors.New("LuminatiUrl empty in clientoptions")
	}

	parsedUrl, err := url.Parse(clientOptions.LuminatiUrl)
	if err != nil {
		return nil, err
	}

	var cache *BDRedisClient
	if clientOptions.UseRedisCache {
		cache = NewBrightDataRedis(&clientOptions.RedisOpts)
	}

	localCache := cache2.New(30*time.Minute, 30*time.Minute)

	return &BrightDataClient{
		options:    &clientOptions,
		proxy:      parsedUrl,
		cache:      cache,
		localCache: localCache,
	}, nil
}

func (bdc *BrightDataClient) DelCache(search SearchOptions) error {

	err := bdc.cache.RdsDel(bdc.getCacheKey(search))
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (bdc *BrightDataClient) CloseClient() {
	if bdc.options.UseRedisCache {
		_ = bdc.cache.CloseConn()
	}
}

func (bdc *BrightDataClient) Search(search SearchOptions) (Serps, error) {

	var rawSerps brightDataResponse
	var serpsResponse Serps

	search.CheckDomain = strings.TrimSpace(search.CheckDomain)
	search.Keyword = strings.TrimSpace(search.Keyword)
	search.Country = strings.TrimSpace(search.Country)

	data, localfound := bdc.localCache.Get(bdc.getCacheKey(search))
	if localfound {
		fmt.Println("from local cache")
		rawSerps = data.(brightDataResponse)
	} else {

		reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
		if err != nil {
			return serpsResponse, err
		}
		processedKeyword := strings.Join(strings.Fields(reg.ReplaceAllString(search.Keyword, "")), "+")

		if bdc.options.UseRedisCache {
			err = bdc.cache.RdsGet(bdc.getCacheKey(search), &rawSerps)
			if err != nil {
				fmt.Println("from live")
				rawSerps, err = bdc.getRawSerps(search, processedKeyword)
				if err != nil {
					return serpsResponse, err
				}
			} else {
				bdc.localCache.Set(bdc.getCacheKey(search), rawSerps, 30*time.Minute)
				fmt.Println("from redis")
			}

		} else {
			fmt.Println("from live")
			rawSerps, err = bdc.getRawSerps(search, processedKeyword)
			if err != nil {
				return serpsResponse, err
			}
		}
	}

	features, featuresForUrl, err := bdc.featureMap(&rawSerps, search.CheckDomain)
	if err != nil {
		return serpsResponse, err
	}

	serpsResponse.CheckedUrl, err = bdc.getRanksForDomain(&rawSerps, search.CheckDomain)
	if err != nil {
		return serpsResponse, err
	}

	serpsResponse.Organic = rawSerps.Organic
	serpsResponse.Features = features
	serpsResponse.CheckedUrl.Query.Features = featuresForUrl

	return serpsResponse, nil
}

func (bdc *BrightDataClient) getRawSerps(search SearchOptions, processedKeyword string) (brightDataResponse, error) {

	var rawSerps brightDataResponse

	timeout := 1 * time.Minute

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		DialContext:     (&net.Dialer{Timeout: timeout}).DialContext,
		IdleConnTimeout: 20 * time.Second,
		MaxIdleConns:    5,
		Proxy:           http.ProxyURL(bdc.proxy),
	}

	client := &http.Client{
		Timeout:   timeout,
		Transport: tr,
	}

	device := ""
	if !search.Desktop {
		device = "&lum_mobile=1"
	}

	googleUrl := "https://www.google.com/search?q=" + processedKeyword + "&lum_json=1&num=100&gl=" + search.Country + device

	req, err := http.NewRequest("GET", googleUrl, nil)
	if err != nil {
		return rawSerps, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return rawSerps, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Println("Error closing body : ", err)
		}
	}(resp.Body)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return rawSerps, err
	}

	err = json.Unmarshal(bodyBytes, &rawSerps)
	if err != nil {
		log.Println("--------- JSON FAILED : RAW BODY ---------")
		log.Println("BODY:", string(bodyBytes))
		log.Println("QUERY:", googleUrl)
		log.Println("CHECK_DOMAIN:", search.CheckDomain)
		log.Println("KEYWORD:", search.Keyword)
		log.Println("PROCESSED_KEYWORD:", processedKeyword)
		log.Println("COUNTRY:", search.Country)
		log.Println("------------------------------------------")
		return rawSerps, err
	}

	if len(rawSerps.Organic) < 1 {
		return rawSerps, errors.New("empty data returned from bright data : " + googleUrl)
	}

	if bdc.options.UseRedisCache {
		err = bdc.cache.RdsSet(bdc.getCacheKey(search), rawSerps)
		if err != nil {
			log.Println("Error setting cache in redis (will continue anyway) : ", err)
		}
	}

	bdc.localCache.Set(bdc.getCacheKey(search), rawSerps, 30*time.Minute)

	return rawSerps, nil
}

func (bdc *BrightDataClient) getCacheKey(search SearchOptions) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		log.Println(err)
	}
	processedKeyword := strings.Join(strings.Fields(reg.ReplaceAllString(search.Keyword, "")), "+")

	return processedKeyword + "+" + search.Country
}
