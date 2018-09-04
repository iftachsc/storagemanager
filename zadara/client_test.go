package zadara

import (
	"fmt"
	"net/http"
	"testing"

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

func getClientAndApiUrl(apiPath string) (*ZadaraClient, string) {
	return NewClient(ZadaraCredentials{VpsaHost, VpsaUser,
		VpsaPassword, VpsaToken}), fmt.Sprintf("http://%s/%s", VpsaHost, apiPath)
}

//API Tests
func TestGetVolumesOK(t *testing.T) {

	c, url := getClientAndApiUrl(
		VolumesPath)

	httpmock.Activate()

	httpmock.RegisterResponder(http.MethodGet, url,
		httpmock.NewStringResponder(200, RootVolumeResponseJson))

	defer httpmock.DeactivateAndReset()

	vols, err := c.GetVolumes()
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, 2, len(vols), "volume count should be equal")

	assert.Equal(t, mockedVol1.Name, vols[0].Name, "volume name should be equal")
	assert.Equal(t, mockedVol2.Name, vols[1].Name, "volume name should be equal")

	assert.Equal(t, mockedVol1.PoolName, vols[0].PoolName, "pool name should be equal")
	assert.Equal(t, mockedVol2.PoolName, vols[1].PoolName, "pool name should be equal")

	assert.Equal(t, mockedVol1.Encryption, vols[0].Encryption, "vol encryption should be equal")
	assert.Equal(t, mockedVol1.VirtualCapacity, vols[0].VirtualCapacity, "vol virtual capacity should be equal")
}

func TestGetMirrorJobsOK(t *testing.T) {

	c, url := getClientAndApiUrl(
		VolumesPath)

	httpmock.Activate()

	httpmock.RegisterResponder(http.MethodGet, url,
		httpmock.NewStringResponder(200, RootVolumeResponseJson))

	defer httpmock.DeactivateAndReset()

	vols, err := c.GetVolumes()
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, 2, len(vols), "volume count should be equal")

	assert.Equal(t, mockedVol1.Name, vols[0].Name, "volume name should be equal")
	assert.Equal(t, mockedVol2.Name, vols[1].Name, "volume name should be equal")

	assert.Equal(t, mockedVol1.PoolName, vols[0].PoolName, "pool name should be equal")
	assert.Equal(t, mockedVol2.PoolName, vols[1].PoolName, "pool name should be equal")

	assert.Equal(t, mockedVol1.Encryption, vols[0].Encryption, "vol encryption should be equal")
	assert.Equal(t, mockedVol1.VirtualCapacity, vols[0].VirtualCapacity, "vol virtual capacity should be equal")
}
