package brightdata

import "net/url"

func (bdc *BrightDataClient) getRanksForDomain(br *brightDataResponse, checkUrl string) (Domain, error) {

	var domainData Domain

	parsedCheckUrl, err := url.Parse(checkUrl)
	if err != nil {
		return domainData, err
	}

	var first = true
	for _, orr := range br.Organic {
		parsedLink, err := url.Parse(orr.Link)
		if err != nil {
			continue
		}

		if parsedLink.Host == parsedCheckUrl.Host && first == true {

			domainData.Query.Rank = orr.Rank
			domainData.Query.Link = orr.Link
			domainData.Query.Description = orr.Description

			first = false
		}

		if parsedLink.Host == parsedCheckUrl.Host {

			var tempQuery Query

			tempQuery.Rank = orr.Rank
			tempQuery.Link = orr.Link
			tempQuery.Description = orr.Description

			domainData.Results = append(domainData.Results, tempQuery)
		}
	}

	return domainData, nil
}
