package brightdata

import "net/url"

func (bdc *BrightDataClient) featureMap(br *brightDataResponse, checkUrl string) ([]string, []string, error) {

	var features []string
	var featuresForCheckurl []string

	parsedCheckUrl, err := url.Parse(checkUrl)
	if err != nil {
		return nil, nil, err
	}

	//check for videos
	if len(br.Videos) > 0 {
		features = append(features, "videos")
	}

	//check for images
	if len(br.Images) > 0 {
		features = append(features, "images")
	}

	//check for tweets
	if len(br.Twitter.Items) > 0 {
		features = append(features, "twitter")
	}

	//check for featured snippets
	if len(br.FeaturedSnippets) > 0 {
		features = append(features, "featured_snippets")

		for _, fs := range br.FeaturedSnippets {
			tUrl, err := url.Parse(fs.Link)
			if err != nil {
				continue
			}

			if tUrl.Host == parsedCheckUrl.Host {
				featuresForCheckurl = append(featuresForCheckurl, "featured_snippets")
				break
			}
		}
	}

	//check for popular products
	if len(br.PopularProducts.Items) > 0 {
		features = append(features, "popular_products")

		for _, fs := range br.PopularProducts.Items {
			tUrl, err := url.Parse(fs.Link)
			if err != nil {
				continue
			}

			if tUrl.Host == parsedCheckUrl.Host {
				featuresForCheckurl = append(featuresForCheckurl, "popular_products")
				break
			}
		}
	}

	//check for recipes
	if len(br.Recipes.Items) > 0 {
		features = append(features, "recipes")

		for _, fs := range br.Recipes.Items {
			tUrl, err := url.Parse(fs.Link)
			if err != nil {
				continue
			}

			if tUrl.Host == parsedCheckUrl.Host {
				featuresForCheckurl = append(featuresForCheckurl, "recipes")
				break
			}
		}
	}

	//check for top stories
	if len(br.TopStories.Items) > 0 {
		features = append(features, "top_stories")

		for _, fs := range br.TopStories.Items {
			tUrl, err := url.Parse(fs.Link)
			if err != nil {
				continue
			}

			if tUrl.Host == parsedCheckUrl.Host {
				featuresForCheckurl = append(featuresForCheckurl, "top_stories")
				break
			}
		}
	}

	//check for jobs
	if len(br.Jobs.Items) > 0 {
		features = append(features, "jobs")
	}

	//check for apps
	if len(br.Apps.Items) > 0 {
		features = append(features, "recipes")
	}

	//check for shopping results
	if len(br.Shopping) > 0 {
		features = append(features, "shopping")

		for _, fs := range br.Shopping {
			tUrl, err := url.Parse(fs.Link)
			if err != nil {
				continue
			}

			if tUrl.Host == parsedCheckUrl.Host {
				featuresForCheckurl = append(featuresForCheckurl, "shopping")
				break
			}
		}
	}

	//check for product offers
	if len(br.ProductOffers) > 0 {
		features = append(features, "product_offers")

		for _, fs := range br.ProductOffers {
			tUrl, err := url.Parse(fs.Link)
			if err != nil {
				continue
			}

			if tUrl.Host == parsedCheckUrl.Host {
				featuresForCheckurl = append(featuresForCheckurl, "product_offers")
				break
			}
		}
	}

	//check for top knowledge
	if len(br.TopKnowledgeCarousel.List) > 0 {
		features = append(features, "top_knowledge_carousel")

		for _, fs := range br.TopKnowledgeCarousel.List {
			tUrl, err := url.Parse(fs.Link)
			if err != nil {
				continue
			}

			if tUrl.Host == parsedCheckUrl.Host {
				featuresForCheckurl = append(featuresForCheckurl, "top_knowledge_carousel")
				break
			}
		}
	}

	//check for knowledge
	if len(br.Knowledge.DescriptionLink) > 0 {
		features = append(features, "knowledge")

		tUrl, err := url.Parse(br.Knowledge.DescriptionLink)
		if err == nil {
			if tUrl.Host == parsedCheckUrl.Host {
				featuresForCheckurl = append(featuresForCheckurl, "knowledge")
			}
		}
	}

	//check for hotels
	if len(br.HotelsSelection.Link) > 0 {
		features = append(features, "hotels")

		tUrl, err := url.Parse(br.HotelsSelection.Link)
		if err == nil {
			if tUrl.Host == parsedCheckUrl.Host {
				featuresForCheckurl = append(featuresForCheckurl, "hotels")
			}
		}
	}

	//check for snackpack map
	if len(br.SnackPackMap.Link) > 0 {
		features = append(features, "snack_pack_map")

		tUrl, err := url.Parse(br.SnackPackMap.Link)
		if err == nil {
			if tUrl.Host == parsedCheckUrl.Host {
				featuresForCheckurl = append(featuresForCheckurl, "snack_pack_map")
			}
		}
	}

	//check for snackpack
	if len(br.SnackPack) > 0 {
		features = append(features, "snack_pack")

		for _, fs := range br.SnackPack {
			tUrl, err := url.Parse(fs.Link)
			if err != nil {
				continue
			}

			if tUrl.Host == parsedCheckUrl.Host {
				featuresForCheckurl = append(featuresForCheckurl, "snack_pack")
				break
			}
		}
	}

	//check for top ads
	if len(br.TopAds) > 0 {
		features = append(features, "top_ads")

		for _, fs := range br.TopAds {
			tUrl, err := url.Parse(fs.Link)
			if err != nil {
				continue
			}

			if tUrl.Host == parsedCheckUrl.Host {
				featuresForCheckurl = append(featuresForCheckurl, "top_ads")
				break
			}
		}
	}

	//check for bottom ads
	if len(br.BottomAds) > 0 {
		features = append(features, "bottom_ads")

		for _, fs := range br.BottomAds {
			tUrl, err := url.Parse(fs.Link)
			if err != nil {
				continue
			}

			if tUrl.Host == parsedCheckUrl.Host {
				featuresForCheckurl = append(featuresForCheckurl, "bottom_ads")
				break
			}
		}
	}

	//check for top pla
	if len(br.TopPla) > 0 {
		features = append(features, "top_pla")

		for _, fs := range br.TopPla {
			tUrl, err := url.Parse(fs.Link)
			if err != nil {
				continue
			}

			if tUrl.Host == parsedCheckUrl.Host {
				featuresForCheckurl = append(featuresForCheckurl, "top_pla")
				break
			}
		}
	}

	//check for bottom pla
	if len(br.BottomPla) > 0 {
		features = append(features, "bottom_pla")

		for _, fs := range br.BottomPla {
			tUrl, err := url.Parse(fs.Link)
			if err != nil {
				continue
			}

			if tUrl.Host == parsedCheckUrl.Host {
				featuresForCheckurl = append(featuresForCheckurl, "bottom_pla")
				break
			}
		}
	}

	//check for jackpot pla
	if len(br.JackpotPla) > 0 {
		features = append(features, "jackpot_pla")

		for _, fs := range br.JackpotPla {
			tUrl, err := url.Parse(fs.Link)
			if err != nil {
				continue
			}

			if tUrl.Host == parsedCheckUrl.Host {
				featuresForCheckurl = append(featuresForCheckurl, "jackpot_pla")
				break
			}
		}
	}

	//check for people also ask
	if len(br.PeopleAlsoAsk) > 0 {
		features = append(features, "people_also_ask")

		for _, fs := range br.PeopleAlsoAsk {
			tUrl, err := url.Parse(fs.AnswerLink)
			if err != nil {
				continue
			}

			if tUrl.Host == parsedCheckUrl.Host {
				featuresForCheckurl = append(featuresForCheckurl, "people_also_ask")
				break
			}

			for _, fsl := range fs.Answers {
				tUrl, err = url.Parse(fsl.Link)
				if err != nil {
					continue
				}

				if tUrl.Host == parsedCheckUrl.Host {
					featuresForCheckurl = append(featuresForCheckurl, "people_also_ask")
					break
				}
			}
		}
	}

	//check for news
	if len(br.News) > 0 {
		features = append(features, "news")

		for _, fs := range br.News {
			tUrl, err := url.Parse(fs.Link)
			if err != nil {
				continue
			}

			if tUrl.Host == parsedCheckUrl.Host {
				featuresForCheckurl = append(featuresForCheckurl, "news")
				break
			}
		}
	}

	//check for Local Services Ads
	if len(br.LocalServicesAds.Items) > 0 {
		features = append(features, "local_services_ads")

		for _, fs := range br.LocalServicesAds.Items {
			tUrl, err := url.Parse(fs.Link)
			if err != nil {
				continue
			}

			if tUrl.Host == parsedCheckUrl.Host {
				featuresForCheckurl = append(featuresForCheckurl, "local_services_ads")
				break
			}
		}
	}

	return features, featuresForCheckurl, nil
}
