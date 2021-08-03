package client

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("GO_ENV", "test")
}

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
		//mocking to get token
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
	t.Run("Checking Request Header on passing token data", func(t *testing.T) {
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
	t.Run("Checking fetch valid data from restapi", func(t *testing.T) {
		json := `{"name":"Test Name","full_name":"test full name","owner":{"login": "octocat"}}`
		// create a new reader with that JSON
		reader := ioutil.NopCloser(bytes.NewReader([]byte(json)))
		GetDoFunc = func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       reader,
			}, nil
		}
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
		_, got, err := getDatafromRestapi(resp)
		assert.NotNil(t, got)
		assert.Nil(t, err)
		assert.EqualValues(t, 200, resp.StatusCode)
		assert.Contains(t, got, "hasUnlimitedLicenses")
	})
	t.Run("Checking error on passing invalid response ", func(t *testing.T) {
		inputs := `{"hasUnlimitedLicenses":false,"lastModifiedBy":"test full name","createdBy":"login"}`
		// create a new reader with that JSON
		readerInputs := ioutil.NopCloser(bytes.NewReader([]byte(inputs)))
		GetDoFunc = func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       readerInputs,
			}, nil
		}
		client = &mockClient{}
		resp, err := getHttpRequest(nil, "@com.in")
		if err != nil {
			t.Fatal(err)
		}
		// Mocked function for os.ReadFile
		file_read := func(io.Reader) ([]byte, error) {
			myErr := errors.New("Simulated error")
			return nil, myErr
		}
		read_All = file_read
		//assert.Error(t, err)
		_, got, err := getDatafromRestapi(resp)
		assert.Error(t, err)
		assert.EqualValues(t, got, "invalid data")
	})

}

func Test_WriteDataintoJson(t *testing.T) {
	t.Run("checking writing data to json file", func(t *testing.T) {
		test_data := []data{
			{HasUnlimited: false},
			{HasUnlimited: true},
		}
		var fileRead = ioutil.ReadFile
		err := writeDataintoJson(test_data, "test.json")
		got, _ := fileRead("test.json")
		assert.Contains(t, string(got), "hasUnlimitedLicenses")
		assert.NoError(t, err)
	})
	t.Run("Checking error on passing invalid file path", func(t *testing.T) {
		test_data := []data{
			{HasUnlimited: true},
		}
		err := writeDataintoJson(test_data, "../test/client.json")
		assert.Error(t, err)

	})
}

func Test_SetupforDbConnection(t *testing.T) {
	err := setupforDbConnection(nil)
	assert.Error(t, err)

}

func Test_WriteintoDatabase(t *testing.T) {
	got := writeintoDatabase()
	assert.NoError(t, got)
}

func Test_GetHttpRequestWithError(t *testing.T) {
	t.Run("Checking error on passing invalid url", func(t *testing.T) {
		file_read := func(string, string, io.Reader) (*http.Request, error) {
			myErr := errors.New("Request error")
			return nil, myErr
		}
		httpRequest = file_read
		_, err := getHttpRequest(nil, "dummyURL@dummy.com.in")
		assert.Error(t, err)
	})

}
