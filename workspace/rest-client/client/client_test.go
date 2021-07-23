package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadfromEnv(t *testing.T) {
	got := loadfromEnv()
	want := 8
	assert.Equal(t, len(got), want)
}

func TestGetAccesToken(t *testing.T) {
	t.Run("Checking access token , token string length should not be 0 ", func(t *testing.T) {
		got_token, _ := getAccessToken()
		expectedToken := ""
		assert.NotEqual(t, len(got_token), len(expectedToken))
	})
	t.Run("Checking token type", func(t *testing.T) {
		_, got_tokenType := getAccessToken()
		expectedTokenType := "Bearer"
		assert.Equal(t, got_tokenType, expectedTokenType)

	})
}


// func newServer(h func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
// 	return httptest.NewServer(http.HandlerFunc(h))
// }

// func TestPasswordCredentialsTokenRequest(t *testing.T) {
// 	ts := newServer(func(w http.ResponseWriter, r *http.Request) {
// 		defer r.Body.Close()
// 		expected := "/token"
// 		if r.URL.String() != expected {
// 			t.Errorf("URL = %q; want %q", r.URL, expected)
// 		}

// 		headerAuth := r.Header.Get("Authorization")
// 		expected = "Basic Q0xJRU5UX0lEOkNMSUVOVF9TRUNSRVQ="
// 		if headerAuth != expected {
// 			t.Errorf("Authorization header = %q; want %q", headerAuth, expected)
// 		}

// 		headerContentType := r.Header.Get("Content-Type")
// 		expected = "application/x-www-form-urlencoded"
// 		if headerContentType != expected {
// 			t.Errorf("Content-Type header = %q; want %q", headerContentType, expected)
// 		}

// 		body, err := ioutil.ReadAll(r.Body)
// 		if err != nil {
// 			t.Errorf("Failed reading request body: %s.", err)
// 		}

// 		expected = "grant_type=password&password=password1&scope=scope1+scope2&username=user1"
// 		if string(body) != expected {
// 			t.Errorf("res.Body = %q; want %q", string(body), expected)
// 		}

// 		w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
// 		_, _ = w.Write([]byte("access_token=ProperToken&scope=user&token_type=bearer"))
// 	})
// 	defer ts.Close()

// 	//client := newClient(ts.URL)
// 	// tok, err := client.CredentialsToken(context.Background(), "user1", "password1")
// 	// if err != nil {
// 	// 	t.Error(err)
// 	// }
// 	User = "user1"
// 	Pass = "password1"
// 	got1, _ := setonOAuthConfig()
// 	fmt.Print(got1)

// 	// if !tok.Valid() {
// 	// 	t.Fatalf("Token invalid. Got: %#v", tok)
// 	// }

// 	expected := "ProperToken"
// 	if got1 != expected {
// 		t.Errorf("AccessToken = %q; want %q", got1, expected)
// 	}

// 	// expected = "bearer"
// 	// if tok.TokenType != expected {
// 	// 	t.Errorf("TokenType = %q; want %q", tok.TokenType, expected)
// 	// }
// }

// type Config struct {
// 	ClientID     string // ClientID is the application's ID.
// 	ClientSecret string // ClientSecret is the application's secret.
// 	AuthURL      string // AuthURL is a URL for authentication.
// 	TokenURL     string // TokenURL is a URL for retrieving a token.
// 	//Mode         Mode     // Mode represents how tokens are represented in requests.
// 	RedirectURL string   // RedirectURL is the URL to redirect users going through the OAuth flow.
// 	Scopes      []string // Scope specifies optional requested permissions.
// }
// type Client struct {
// 	client *http.Client
// 	config Config
// }

// func NewClient(client *http.Client, config Config) *Client {
// 	c := &Client{
// 		client: client,
// 		config: config,
// 	}
// 	return c
// }

// func newClient(url string) *Client {
// 	cfg := Config{
// 		ClientID:     "CLIENT_ID",
// 		ClientSecret: "CLIENT_SECRET",
// 		AuthURL:      url + "/auth",
// 		TokenURL:     url + "/token",
// 		//Mode:         AutoDetectMode,
// 		RedirectURL: "REDIRECT_URL",
// 		Scopes:      []string{"scope1", "scope2"},
// 	}
// 	return NewClient(http.DefaultClient, cfg)
// }
