// Code generated by protoc-gen-volcengine-sdk
// source: VodUploadService
// DO NOT EDIT!

package vod

import (
	"fmt"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/business"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
)

func Test_UploadMediaByUrl(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	urlSets := make([]*business.VodUrlUploadURLSet, 0)
	urlSet := &business.VodUrlUploadURLSet{
		SourceUrl:        "", // 源视频URL
		CallbackArgs:     "", // 希望回调时会透传返回的信息，最大长度 512字节，非必须字段。
		Md5:              "", // 视频文件的 Md5 值，如添加则会进行 Md5 的校验，非必须字段
		TemplateId:       "", // 如希望上传后自动启动工作流，则填写工作流 Id，非必须字段
		Title:            "", // 标题，不超过 128个字符，非必须字段
		Description:      "", // 描述，不超过 256个字符，非必须字段
		Tags:             "", // 标签，多个以“,”分隔，不超过 36个字符，非必须字段
		FileName:         "", //
		ClassificationId: 0,  // 分类 ID，可上传时指定分类，非必须字段
	}
	urlSets = append(urlSets, urlSet)

	query := &request.VodUrlUploadRequest{
		SpaceName: "your SpaceName",
		URLSets:   urlSets,
	}

	resp, status, err := instance.UploadMediaByUrl(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())

	fmt.Println(resp.GetResponseMetadata().GetRequestId())    // requestId 用于辅助查询
	fmt.Println(resp.GetResult().GetData()[0].GetSourceUrl()) // 原 url
	fmt.Println(resp.GetResult().GetData()[0].GetJobId())     // url 对应的 JobId，可用来临时查询视频状态信息
}

func Test_QueryUploadTaskInfo(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodQueryUploadTaskInfoRequest{
		JobIds: "your JobIds",
	}

	resp, status, err := instance.QueryUploadTaskInfo(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
	// 字段解释均可在火山引擎点播官网 开发者API-媒资上传-查询URL批量上传任务状态中找到相关解释

	fmt.Println(resp.GetResult().GetData().GetMediaInfoList()[0].GetJobId())                    // 可以查看到此记录的 JobId
	fmt.Println(resp.GetResult().GetData().GetMediaInfoList()[0].GetVid())                      // 此 JobId 对应的视频 Vid
	fmt.Println(resp.GetResult().GetData().GetMediaInfoList()[0].GetSourceInfo().GetStoreUri()) // 此视频存储 uri
	fmt.Println(resp.GetResult().GetData().GetMediaInfoList()[0].GetSourceInfo().GetSize())     // 此视频的大小，单位字节

}
