package visual

import (
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/service/visual/model"
	"testing"
)

const (
	testAk = "ak"
	testSk = "sk"
)

func TestVisual_ConvertPhotoV2(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	reqBody := &model.ConvertPhotoV2Request{
		ReqKey:           "lens_opr", // 固定值
		BinaryDataBase64: []string{"image_base64"},
		//IsColor: false,
	}

	resp, status, err := DefaultInstance.ConvertPhotoV2(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVisual_LensVidaVideoSubmitTaskV2(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	reqBody := &model.LensVidaVideoSubmitTaskV2Request{
		ReqKey:   "lens_vida_video", // 固定值
		Url:      "https://xxxx",
		VidaMode: "vida_simple",
	}

	resp, status, err := DefaultInstance.LensVidaVideoSubmitTaskV2(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVisual_LensVidaVideoGetResultV2(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)

	reqBody := &model.LensVidaVideoGetResultV2Request{
		ReqKey: "lens_vida_video", // 固定值
		TaskId: "xxx",
	}

	resp, status, err := DefaultInstance.LensVidaVideoGetResultV2(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
