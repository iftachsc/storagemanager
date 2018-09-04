package zadara

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type ZadaraCredentials struct {
	Host     string
	User     string
	Password string
	Token    string
}
type ZadaraClient struct {
	credentials ZadaraCredentials
}

func NewClient(credentials ZadaraCredentials) *ZadaraClient {
	zc := new(ZadaraClient)
	zc.credentials = credentials

	return zc
}

func (c *ZadaraClient) invokeGet(path string) ([]byte, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	url := fmt.Sprintf("http://%s/%s", c.credentials.Host, path)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-Access-Key", c.credentials.Token)
	//req.Header.Add("Content-Type","application/json")
	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (c *ZadaraClient) GetVolumes() ([]Volume, error) {

	jsonBytes, err := c.invokeGet(VolumesPath) //This reads raw request body
	if err != nil {
		return nil, err
	}

	var resp RootVolumeResponse
	json.Unmarshal(jsonBytes, &resp)
	//println("volume:",zadara.Response.Volumes[0].Vol_name)
	return resp.Response.Volumes, nil
}
