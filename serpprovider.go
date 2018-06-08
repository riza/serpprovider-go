package serpprovider

import (
	"log"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"strings"
	"strconv"
)

var (
	BaseUrl = "http://api.serpprovider.com/%s/%s/%s/%s/%s"
)

type SearchRequest struct {
	Query string
	Engine string
	Country string
	Language string
	Device string
	Key string
	Mode string // live, prefetch, mixed (default mixed)
	Url string
	Limit int
}

type SerpReqResponse struct {
	Provider string `json:"provider"`
	Hl		 string `json:"hl"`
	Device	 string `json:"device"`
	Query	 string `json:"q"`
	Offset   int  `json:"offset"`
}

type SerpResponse struct {
	Url		string `json:"url"`
	Position int64 `json:"pos"`
	Title	string `json:"title"`
	Cite	string `json:"cite"`
	Description	string `json:"desc"`
}

type SerpUserResponse struct {
	Credits int32 `json:"credits"`
	Url		string `json:"url"`
}

type SerpMainResponse struct {
	Req	SerpReqResponse
	Res []SerpResponse
	Status string `json:"status"`
	User SerpUserResponse
}

func New(APIKey string) *SearchRequest {

	if APIKey == "" {
		log.Fatal("[!] APIKey cannot be empty.")
	}

	req := &SearchRequest{
		Device: "desktop",
		Engine: "google",
		Country: "us",
		Language: "en",
		Mode:	"mixed",
		Key:	APIKey,
		Limit:	50,
	}
	return req
}

func (req *SearchRequest) Request(reqVals *SearchRequest) SerpMainResponse {

	var response SerpMainResponse;

	if reqVals.Query == "" {
		log.Fatal("[!] Query cannot be empty.")
	}

	if reqVals.Country != "" {
		req.Country = reqVals.Country;
	}

	if reqVals.Device != "" {
		req.Device = reqVals.Device;
	}

	if reqVals.Engine != "" {
		req.Engine = reqVals.Engine;
	}

	if reqVals.Language != "" {
		req.Language = reqVals.Language;
	}



	if reqVals.Mode != "" {
		req.Mode = reqVals.Mode;
	}

	req.Query = reqVals.Query;

	url := fmt.Sprintf(
		BaseUrl,
		req.Key,
		req.Engine,
		fmt.Sprintf("%s-%s", req.Language, req.Country),
		req.Device,
		strings.Replace(req.Query," ","+",-1),
	);

	request := gorequest.New().Get(url)

	if reqVals.Limit != 0 {
		request.Param("limit",strconv.Itoa(req.Limit))
	}

	if reqVals.Url != "" {
		request.Param("match",req.Url)
	}

	_,_,errs := request.EndStruct(&response)

	if errs != nil {
		log.Fatal(errs)
	}

	return response;
}
