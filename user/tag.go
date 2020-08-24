package user

import (
	"encoding/json"
	"fmt"
	"github.com/tboss-maker/wechat/context"
	"github.com/tboss-maker/wechat/util"
)

const (
	userCreateTagURL = "https://api.weixin.qq.com/cgi-bin/tags/create"
	userGetTagURL = "https://api.weixin.qq.com/cgi-bin/tags/get"
	userUpdateTagURL = "https://api.weixin.qq.com/cgi-bin/tags/update"
	userDeleteTagURL = "https://api.weixin.qq.com/cgi-bin/tags/delete"


	userBatchTaggingURL = "https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging"
	userBatchUntaggingURL = "https://api.weixin.qq.com/cgi-bin/tags/members/batchuntagging"
	userGetTagListURL = "https://api.weixin.qq.com/cgi-bin/tags/getidlist"
)

type Tag struct {
	*context.Context
}


type ReqCreateTag struct {
	Tag createTagInfo `json:"tag"`
}

type createTagInfo struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
}

type ResCreateTag struct {
	util.CommonError
	Tag createTagInfo `json:"tag"`
}

func NewTag(context *context.Context) *Tag {
	tag := new(Tag)
	tag.Context = context
	return tag
}

func (t *Tag) CreateTag(name string) (createTagInfo, error) {
	accessToken, err := t.GetAccessToken()
	if err != nil {
		return createTagInfo{}, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", userCreateTagURL, accessToken)
	reqMenu := &ReqCreateTag{
		Tag: createTagInfo{Name: name},
	}

	response, err := util.PostJSON(uri, reqMenu)
	if err != nil {
		return createTagInfo{}, err
	}

	var ResCreateTag ResCreateTag
	err = json.Unmarshal(response, &ResCreateTag)
	if err != nil {
		return createTagInfo{}, err
	}

	if ResCreateTag.ErrCode != 0 {
		err = fmt.Errorf("CreateTag Error , errcode=%d , errmsg=%s", ResCreateTag.ErrCode, ResCreateTag.ErrMsg)
		return createTagInfo{}, err
	}

	return ResCreateTag.Tag, nil
}

type updateTagInfo struct {
	Id int64 `json:"id,omitempty"`
	Name string `json:"name"`
}

type reqUpdateTag struct {
	Tag updateTagInfo `json:"tag"`
}

func (t *Tag) UpdateTag(id int64, name string) error {
	accessToken, err := t.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", userUpdateTagURL, accessToken)
	reqUpdateTag := &reqUpdateTag{
		Tag: updateTagInfo{Id: id, Name: name},
	}

	response, err := util.PostJSON(uri, reqUpdateTag)
	if err != nil {
		return err
	}
	return util.DecodeWithCommonError(response, "UpdateTag")
}

type tagListInfo struct {
	Id int64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Count int64 `json:"count,omitempty"`
}

type resGetTagList struct {
	util.CommonError
	Tags []tagListInfo `json:"tags"`
}

func (t *Tag) GetTags() ([]tagListInfo, error) {
	accessToken, err := t.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("%s?access_token=%s", userGetTagURL, accessToken)
	response, err := util.HTTPGet(uri)
	if err != nil {
		return nil, err
	}
	resGetTagList := &resGetTagList{}
	err = json.Unmarshal(response, resGetTagList)
	if err != nil {
		return nil, err
	}
	if resGetTagList.ErrCode != 0 {
		err = fmt.Errorf("GetTags Error , errcode=%d , errmsg=%s", resGetTagList.ErrCode, resGetTagList.ErrMsg)
		return nil, err
	}
	return nil, err
}

type deleteTagInfo struct {
	Id int64 `json:"id,omitempty"`
}

type reqDeleteTag struct {
	Tag deleteTagInfo `json:"tag"`
}

func (t *Tag) DeleteTag(id int64) error {
	accessToken, err := t.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", userDeleteTagURL, accessToken)
	reqDeleteTag := &reqDeleteTag{
		Tag: deleteTagInfo{Id: id},
	}

	response, err := util.PostJSON(uri, reqDeleteTag)
	if err != nil {
		return err
	}
	return util.DecodeWithCommonError(response, "UpdateTag")
}

type reqBatchTagging struct {
	OpenidList []string `json:"openid_list"`
	TagId int64 `json:"tagid"`
}

func (t *Tag) BatchTagging(openIds []string, tagId int64) error {
	accessToken, err := t.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", userBatchTaggingURL, accessToken)
	reqBatchTagging := &reqBatchTagging{
		OpenidList: openIds,
		TagId: tagId,
	}

	response, err := util.PostJSON(uri, reqBatchTagging)
	if err != nil {
		return err
	}

	return util.DecodeWithCommonError(response, "BatchTagging")
}

type reqBatchUntagging struct {
	OpenidList []string `json:"openid_list"`
	TagId int64 `json:"tagid"`
}

func (t *Tag) BatchUntagging(openIds []string, tagId int64) error {
	accessToken, err := t.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", userBatchUntaggingURL, accessToken)
	reqBatchUntagging := &reqBatchUntagging{
		OpenidList: openIds,
		TagId: tagId,
	}

	response, err := util.PostJSON(uri, reqBatchUntagging)
	if err != nil {
		return err
	}

	return util.DecodeWithCommonError(response, "BatchUntagging")
}