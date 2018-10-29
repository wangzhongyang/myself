package main

import (
	"bindolabs/bindocommon/helpers/logger"
	"bindolabs/bindoio/models"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	getTokenUrl            = "http://devmyflightext.mtel.ws/Hkia/api/getToken"
	directSendPushUrl      = "http://devmyflightext.mtel.ws/Hkia/api/directSendPush"
	batchDirectSendPushUrl = "http://devmyflightext.mtel.ws/Hkia/api/batchDirectSendPush"
)

type MyFlight struct {
	MyFlightToken string
	Message       string
	StoreId       int
	UserIds       []int
	PushUsers     []PushUserInfo
}

type PushUserInfo struct {
	UserId         string
	DeviceId       string
	ExternalUserId string
	DeviceType     string
	PushToken      string
}

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetTokenResponse struct {
	Token      string `json:"token"`
	ExpiryTime string `json:"expiry_time"`
}

type DirectSendPushRequest struct {
	Token          string            `json:"token"`
	UserId         string            `json:"user_id"`
	DeviceId       string            `json:"device_id"`
	ExternalUserId string            `json:"external_user_id"`
	PushToken      string            `json:"push_token"`
	DeviceType     string            `json:"device_type"`
	Language       string            `json:"language"`
	Message        string            `json:"message"`
	Extra          map[string]string `json:"extra"`
}

type DirectSendPushResponse struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	RefId string `json:"refId"`
}

type BatchDirectSendPushRequest struct {
	Token    string                           `json:"token"`
	Users    []BatchDirectSendPushRequestUser `json:"users"`
	Language string                           `json:"language"`
	Message  string                           `json:"message"`
	Extra    map[string]string                `json:"extra"`
}

type BatchDirectSendPushRequestUser struct {
	UserId         string `json:"user_id"`
	DeviceId       string `json:"device_id"`
	ExternalUserId string `json:"external_user_id"`
	PushToken      string `json:"push_token"`
	DeviceType     string `json:"device_type"`
}

//NewClient token 由外部传入
func NewClient(myFlightToken string, storeId int) *MyFlight {
	return &MyFlight{
		MyFlightToken: myFlightToken,
		StoreId:       storeId,
	}
}

//BuildMessage 主要为了构建 PushUserInfo
// Param: message,userIds
func (m *MyFlight) BuildMessage(args ...interface{}) error {
	if len(args) != 2 {
		return errors.New("my flight build message error,len(args)!=1")
	}
	m.Message = args[0].(string)
	m.UserIds = args[1].([]int)
	// 此处需保证查询结果的顺序与传入的 userIds 一致
	// get notification platform app info
	sqlOrderBy := m.setOrderByStr(m.UserIds)
	var customers []models.Customer
	if err := models.DB(models.Customer{}).Where("store_id = ? and linked_source_type = 'User' and linked_source_id in (?)", m.UserIds).Order(sqlOrderBy).Find(&customers).Error; err != nil {
		return errors.New("get customer error:" + err.Error())
	}
	if len(customers) != len(m.UserIds) {
		return errors.New("have user not in this store")
	}
	customerIds := make([]int, 0)
	for _, v := range customers {
		customerIds = append(customerIds, v.ID.V())
	}

	sqlOrderBy = m.setOrderByStr(customerIds)
	var apps []models.NotificationPlatformMobileApp
	if err := models.DB(models.NotificationPlatformMobileApp{}).Where("customer_id in (?)", customerIds).Order(sqlOrderBy).Find(&apps).Error; err != nil {
		return errors.New("cant get apps info: " + err.Error())
	}
	// set PushUserInfo array
	if len(apps) != len(m.UserIds) {
		return errors.New("have user not login")
	}
	pushUsers := make([]PushUserInfo, 0)
	for k, v := range apps {
		temp := PushUserInfo{
			UserId:         strconv.Itoa(m.UserIds[k]),
			DeviceId:       v.DeviceID.V(),
			ExternalUserId: "",
			DeviceType:     v.MobileType.V(),
			PushToken:      v.PushToken.V(),
		}
		pushUsers = append(pushUsers, temp)
	}
	m.PushUsers = pushUsers
	return nil
}

func (m *MyFlight) PushToOne() error {
	if len(m.PushUsers) != 1 {
		return errors.New("push to one,the number of push messages is incorrect")
	}
	pushUser := m.PushUsers[0]
	directSendPush := DirectSendPushRequest{
		Token:          m.MyFlightToken,
		UserId:         pushUser.UserId,
		DeviceId:       pushUser.DeviceId,
		ExternalUserId: "external_user_id",
		PushToken:      pushUser.PushToken,
		DeviceType:     pushUser.DeviceType,
		Language:       "EN",
		Message:        m.Message,
		Extra:          nil,
	}

	// send
	bodyStr, err := json.Marshal(directSendPush)
	if err != nil {
		return errors.New("json marshal error:" + err.Error())
	}
	respByte, err := m.postRequest(directSendPushUrl, string(bodyStr))
	if err != nil {
		return errors.New("post request error:" + err.Error())
	}
	var resp DirectSendPushResponse
	if err := json.Unmarshal(respByte, &resp); err != nil {
		return errors.New(fmt.Sprintf("json unmarshal error:%s,resp:%s", err.Error(), string(respByte)))
	}
	logger.NewErrorf("my flight, push to one message success:%v", resp)
	return nil
}

func (m *MyFlight) PushToGroup() error {
	if len(m.PushUsers) <= 1 {
		return errors.New("push to group,the number of push messages is incorrect")
	}
	items := make([]BatchDirectSendPushRequestUser, 0)
	for _, v := range m.PushUsers {
		item := BatchDirectSendPushRequestUser{
			UserId:         v.UserId,
			DeviceId:       v.DeviceId,
			ExternalUserId: "",
			PushToken:      v.PushToken,
			DeviceType:     v.DeviceType,
		}
		items = append(items, item)
	}
	batchDirectSendPush := BatchDirectSendPushRequest{
		Token:    m.MyFlightToken,
		Users:    items,
		Language: "EN",
		Message:  m.Message,
		Extra:    nil,
	}

	// send
	bodyStr, err := json.Marshal(batchDirectSendPush)
	if err != nil {
		return errors.New("json marshal error:" + err.Error())
	}
	respByte, err := m.postRequest(batchDirectSendPushUrl, string(bodyStr))
	if err != nil {
		return errors.New("post request error:" + err.Error())
	}
	var resp DirectSendPushResponse
	if err := json.Unmarshal(respByte, &resp); err != nil {
		return errors.New(fmt.Sprintf("json unmarshal error:%s,resp:%s", err.Error(), string(respByte)))
	}
	logger.NewErrorf("my flight, push to group message success:%v", resp)
	return nil
}

func (m *MyFlight) BuildCheckMessage(args ...interface{}) error {
	return errors.New("my flight no such feature:BuildCheckMessage")
}

func (m *MyFlight) Check() error {
	return errors.New("my flight no such feature:Check")
}

func (m *MyFlight) GetToken(userName, password string) (*GetTokenResponse, error) {
	user := UserInfo{
		Username: userName,
		Password: password,
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	bodyByte, err := m.postRequest(getTokenUrl, string(userJson))
	if err != nil {
		return nil, errors.New("post request error:" + err.Error())
	}
	var resp GetTokenResponse
	if err := json.Unmarshal(bodyByte, &resp); err != nil {
		return nil, errors.New("json unmarshal error:" + err.Error())
	}
	return &resp, nil
}

func (m *MyFlight) setOrderByStr(ids []int) string {
	sqlOrderBy := ""
	for _, v := range ids {
		sqlOrderBy = fmt.Sprintf(sqlOrderBy+",%d", v)
	}
	sqlOrderBy = "field(id" + sqlOrderBy + ")"
	return sqlOrderBy
}

func (m *MyFlight) postRequest(url, postBody string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(postBody)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	reap, err := client.Do(req)
	defer reap.Body.Close()
	body, err := ioutil.ReadAll(reap.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
