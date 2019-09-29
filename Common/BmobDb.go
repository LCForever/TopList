package Common

import(
	"github.com/bmob/bmob-go-sdk"
	"encoding/json"
	"fmt"
)

func insert(strTable string, itemData map[string]interface{}, appConfig bmob.RestConfig) bool{
	jsonData, err := json.Marshal(itemData)
	if nil != err{
		fmt.Println("json marshal ", itemData, " error: ", err)
		return false
	}
	var respDst interface{}
	_, err = bmob.DoRestReq(appConfig, 
		bmob.RestRequest{
			BaseReq: bmob.BaseReq{
				Method: "POST",
				Path: bmob.ApiRestURL(strTable) + "/",
				Token: ""},
			Type: "application/json",
			Body: jsonData},
			&respDst)
	if nil == err {
		return true
	} else {
		fmt.Println("DoRestReq error: ", err)
	}
	return false
}

func InsertData(strTable string, data []map[string]interface{}, strAppID string, strRestKey string) bool {
	appConfig := bmob.RestConfig{AppID: strAppID, RestKey: strRestKey}

	for _, itemData := range data{
		insert(strTable, itemData, appConfig)
	}
	return true
}
