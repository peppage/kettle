package steam

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"kettle/errors"

	"github.com/ChimeraCoder/tokenbucket"
	log "github.com/Sirupsen/logrus"
)

const (
	BaseUrl = "http://api.steampowered.com"
)

type Steam struct {
	queryQueue chan query
	HttpClient *http.Client
	bucket     *tokenbucket.Bucket
	baseUrl    string
	key        string
}

type query struct {
	url         string
	form        url.Values
	data        interface{}
	response_ch chan response
}

type response struct {
	data interface{}
	err  error
}

func New(key string) *Steam {
	queue := make(chan query)
	api := &Steam{
		queryQueue: queue,
		bucket:     nil,
		HttpClient: http.DefaultClient,
		baseUrl:    BaseUrl,
		key:        key,
	}
	log.SetFormatter(&log.JSONFormatter{})
	go api.throttledQuery()
	return api
}

func (api *Steam) EnableThrottling(rate time.Duration, capacity int64) {
	api.bucket = tokenbucket.NewBucket(rate, capacity)
}

func (api *Steam) DisableThrottling() {
	api.bucket = nil
}

func (api *Steam) SetDelay(rate time.Duration) {
	api.bucket.SetRate(rate)
}

func (api *Steam) GetDelay() time.Duration {
	return api.bucket.GetRate()
}

func cleanValues(v url.Values) url.Values {
	if v == nil {
		return url.Values{}
	}
	return v
}

func (api Steam) execQuery(urlStr string, form url.Values, data interface{}) error {
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return err
	}
	form.Set("key", api.key)
	req.URL.RawQuery = form.Encode()

	resp, err := api.HttpClient.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	log.WithFields(log.Fields{
		"url":    req.URL.String(),
		"Status": resp.Status,
	}).Debug("Executed Query")
	return decodeResponse(resp, data)
}

func decodeResponse(resp *http.Response, data interface{}) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return NewApiError(resp)
	}

	return json.Unmarshal(body, &data)
}

func NewApiError(resp *http.Response) *errors.ApiError {
	body, _ := ioutil.ReadAll(resp.Body)

	return &errors.ApiError{
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Body:       string(body),
		URL:        resp.Request.URL,
	}
}

func (api *Steam) throttledQuery() {
	for q := range api.queryQueue {
		url := q.url
		form := q.form
		data := q.data

		response_ch := q.response_ch

		if api.bucket != nil {
			<-api.bucket.SpendToken(1)
		}

		err := api.execQuery(url, form, data)

		if err != nil {
			if apiErr, ok := err.(*errors.ApiError); ok {
				limited, nextWindow := apiErr.RateLimitCheck()
				if limited {
					go func() {
						api.queryQueue <- q
					}()

					delay := nextWindow.Sub(time.Now())
					<-time.After(delay)

					// Drain the bucket (start over fresh)
					if api.bucket != nil {
						api.bucket.Drain()
					}

					continue
				}
				log.WithFields(log.Fields{
					"Status Code": apiErr.StatusCode,
					"URL":         apiErr.URL,
				}).Error("Api error")
			}
		}

		response_ch <- response{data, err}
	}
}
