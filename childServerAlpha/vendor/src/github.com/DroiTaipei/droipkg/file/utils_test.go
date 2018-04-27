package file

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/url"
	"os"
	"testing"
)

const (
	appid    = "497umbzhcHN44alFPsXyQ4cZdUyo31PMlQAwT0AO"
	fid      = "861874359009153024"
	cdnFqdn  = "newmarket1.oo523.com"
	filename = "/foo/bar.jpg"
)

var (
	cdnURL = fmt.Sprintf("http://%s/droi/%s/%s/%s", cdnFqdn, appid, fid, url.QueryEscape(GetFileBasename(fid, filename)))
)

func BeforeTest() {

}

func Test_GetCdnURL(t *testing.T) {
	assert.Equal(t, cdnURL, GetCdnURL(cdnFqdn, appid, fid, filename))
}

func Test_ParseCdnURL(t *testing.T) {
	actFqdn, actAppid, actFid, actFilename, err := ParseCdnURL(cdnURL)
	assert.NoError(t, err)
	assert.Equal(t, cdnFqdn, actFqdn)
	assert.Equal(t, appid, actAppid)
	assert.Equal(t, fid, actFid)
	assert.Equal(t, GetFileBasename(fid, filename), actFilename)

	actFqdn, actAppid, actFid, actFilename, err = ParseCdnURL("I am fake url")
	assert.Error(t, err)
	assert.Equal(t, "", actFqdn)
	assert.Equal(t, "", actAppid)
	assert.Equal(t, "", actFid)
	assert.Equal(t, "", actFilename)
}

func Test_GetFileBase(t *testing.T) {
	testCases := map[string]string{
		// linux
		"":                           fid,
		".":                          fid,
		".jpg":                       fid + ".jpg",
		".foo.jpg":                   fid + ".jpg",
		"foo.jpg":                    "foo.jpg",
		"FoO.jpg":                    "FoO.jpg",
		"foo":                        "foo",
		"123.jpg":                    "123.jpg",
		"56. pikachu":                "56.pikachu",
		"/foo/bar/foo.jpg":           "foo.jpg",
		"/foo/ foo .jpg":             "foo.jpg",
		"/foo/bar/foo":               "foo",
		"-foo.jpg":                   "-foo.jpg",
		"foo-bar.jpg":                "foo-bar.jpg",
		"foo_bar.jpg":                "foo_bar.jpg",
		"http://foo.com/bar/foo.png": "foo.png",
		"/Users/Droi/foo.jpg":        "foo.jpg",
		"/Users/../foo.jpg":          "foo.jpg",
		"/foo.jpg":                   "foo.jpg",
		"getFile?id=703fc02c-43c7-bcc9-f21a-8527640eca68": "getFile3Fid-703fc02c-43c7-bcc9-f21a-8527640eca68",
		// windows
		`C:\Windows Progra\Droi\foo.jpg`: "foo.jpg",
		`c:\windows\droi\foo.dat`:        "foo.dat",
		`c:\windows\droi\foo`:            "foo",
		`D:\windows\droi\foo.dat`:        "foo.dat",
		// multibyte filename
		"/云存储/卓易.jpg": fid + ".jpg",
		"您可能不知道DeepMind是谁，但一定听说过半年前那场轰动全球的人机大战，有八个世界围棋冠军头衔的李世石以1：4不敌人工智能AlphaGo": "DeepMind14AlphaGo",
		"1你2好.jpg.jpg": "12.jpg.jpg",
		// special characters
		`!@#$%^&*()_+-={}:"<>?[];',.|.jpg`: "21-23-255E-2A2829_---7B7D:223C3E3F5B5D3B272C.7C.jpg",
	}
	for input, expected := range testCases {
		assert.Equal(t, expected, GetFileBasename(fid, input))
	}
}

func Test_IsGosunURL(t *testing.T) {
	url := "http://" + VendorGosunDomain + "/foo"
	assert.True(t, IsGosunURL(url))
	assert.False(t, IsGosunURL("http://foo.bar"))
}

func Test_IsWangsuURL(t *testing.T) {
	url := "http://" + VendorWangsuDomain + "/foo"
	assert.True(t, IsWangsuURL(url))
	assert.False(t, IsWangsuURL("http://foo.bar"))
}

func Test_IsQiniuURL(t *testing.T) {
	url := "http://" + VendorQiniuDomain + "/foo"
	assert.True(t, IsQiniuURL(url))
	assert.False(t, IsQiniuURL("http://foo.bar"))
}

// Do somethings after all test cases
func AfterTest() {

}

func TestMain(m *testing.M) {
	BeforeTest()
	retCode := m.Run()
	AfterTest()
	os.Exit(retCode)
}
