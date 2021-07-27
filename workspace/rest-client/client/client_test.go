package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LoadfromEnv(t *testing.T) {
	t.Run("Checking error", func(t *testing.T) {
		_, err := loadfromEnv("/data")
		if err != nil {
			assert.Error(t, err)
		}
	})
	t.Run("Checking loaded data", func(t *testing.T) {
		got, _ := loadfromEnv("../configfile")
		assert.NotNil(t, got)
	})

}

func Test_GetToken(t *testing.T) {
	t.Run("Checking Error ", func(t *testing.T) {
		congdata := configuration{
			ClientID:     "124",
			ClientSecret: "wer3d",
			UserID:       "gnnsn@dio",
			UserPwd:      "slnnc#98",
			TokenURL:     "token/ncksnckckd@url.in",
			ProductURL:   "product/nckdncdkc@url.in",
		}
		got, err := getToken(congdata)
		if err != nil {
			assert.Error(t, err)
		}
		assert.Nil(t, got)
	})
	t.Run("Checking token", func(t *testing.T) {
		configData, _ := loadfromEnv("../configfile")
		got, err := getToken(configData)
		if err != nil {
			assert.Error(t, err)
		}
		assert.NotNil(t, got)
	})
}

func Test_SetHeader(t *testing.T) {
	//to do mock the header to check func.
	t.Run("", func(t *testing.T) {
		configData, _ := loadfromEnv("../configfile")
		token, err := getToken(configData)
		if err != nil {
			t.Fatal(err)
		}
		got := setHeader(token)
		assert.NotNil(t, got)
		assert.Contains(t, got, "Authorization")

	})

}

type mockClient struct {
}

var GetDoFunc func(req *http.Request) (*http.Response, error)

func (m *mockClient) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}
func Test_GetHttpRequest(t *testing.T) {
	json := `{"name":"Test Name","full_name":"test full name","owner":{"login": "octocat"}}`
	// create a new reader with that JSON
	reader := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       reader,
		}, nil
	}
	t.Run("Checking fetch valid data from restapi", func(t *testing.T) {
		client = &mockClient{}
		resp, err := getHttpRequest(nil, "@com.in")
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		bodyString := string(bodyBytes)
		assert.NotNil(t, resp)
		assert.Nil(t, err)
		assert.EqualValues(t, 200, resp.StatusCode)
		assert.Contains(t, bodyString, "Test Name")
	})

}

func Test_getDatafromRestapi(t *testing.T) {
	json := `{"hasUnlimitedLicenses":false,"lastModifiedBy":"test full name","createdBy":"login"}`
	// create a new reader with that JSON
	reader := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       reader,
		}, nil
	}
	t.Run("Checking fetch valid data from restapi", func(t *testing.T) {
		client = &mockClient{}
		resp, err := getHttpRequest(nil, "@com.in")
		if err != nil {
			t.Fatal(err)
		}
		_, err, got := getDatafromRestapi(resp)
		assert.NotNil(t, got)
		assert.Nil(t, err)
		assert.EqualValues(t, 200, resp.StatusCode)
		assert.Contains(t, got, "hasUnlimitedLicenses")
	})

}
