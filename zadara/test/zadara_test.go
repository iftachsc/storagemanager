package zadara_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/iftachsc/storagemanager/zadara"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func initHttpMock(url string, method string, json string) {
	httpmock.Activate()

	//url e.g. "https://api.mybiz.com/articles.json"
	httpmock.RegisterResponder(method, url,
		httpmock.NewStringResponder(200, json))
}

const (
	VpsaHost     = "my.vpsa.com"
	VpsaUser     = "u"
	VpsaPassword = "p"
	VpsaToken    = "token"
)

//API Tests
func TestGetVolumesOK(t *testing.T) {

	c := zadara.NewClient(zadara.ZadaraCredentials{VpsaHost, VpsaUser, VpsaPassword, VpsaToken})
	url := fmt.Sprintf("http://%s/%s", VpsaHost, zadara.VolumesPath)
	println(fmt.Sprintf("Mocking json on %s", url))
	//initHttpMock(http.MethodGet, url, RootVolumeResponseJson)

	httpmock.Activate()

	//url e.g. "https://api.mybiz.com/articles.json"
	httpmock.RegisterResponder(http.MethodGet, url,
		httpmock.NewStringResponder(200, RootVolumeResponseJson))

	defer httpmock.DeactivateAndReset()

	vols, err := c.GetVolumes()
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(vols) != 2 {
		t.Errorf("Expected 2 volume but got %d", len(vols))
	}

	assert.Equal(t, MockVolume1.Name, vols[0].Name, "volume name should be equal")
	assert.Equal(t, MockVolume1.PoolName, vols[0].PoolName, "pool name should be equal")
	assert.Equal(t, MockVolume1.Encryption, vols[0].Encryption, "vol names should be equal")

}

func Get(apiPath string, host string) error {
	url := fmt.Sprintf("http://%s/%s", host, apiPath)
	payload := fmt.Sprintf(`
			{
				"meta": {
					"lifeMeaning": "h"
				},
				"data": {
					" “message”: "j"
				}
			}`)
	resp, err := http.Post(url, "application/json", strings.NewReader(payload))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("nsqd didn’t respond 200 OK: %s", resp.Status)
	}
	return nil
}
