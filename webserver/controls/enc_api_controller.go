package controls

import (
	"github.com/mcmx73/vaulta/webserver/forest"
	"github.com/DeepForestTeam/mobiussign/components/log"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"github.com/mcmx73/vaulta/components/tapeencript"
)

type DataBlock struct {
	Result     string
	ResultCode int
	DataBlock  string `json:"data_block"`
}
type EncodedData struct {
	Result     string
	ResultCode int
	BlockIndex string `json:"block_index"`
	KeyIndex   string `json:"key_index"`
}
type EncoderController struct {
	forest.Control
}

func (this *EncoderController)Get() {
	defer this.ServeJSON()
	block_index := this.Context.UrlVars["data_link"]
	key_index := this.Context.UrlVars["data_key"]
	if block_index == "" || key_index == "" {
		this.Data = ErrorMessage{
			Result:"Request error",
			ResultCode:406,
		}
	}
	decoder := tapeencript.TapeBlock{}
	decoder.BlockId = block_index
	decoder.Decrypt(key_index)
	this.Data = DataBlock{DataBlock:string(decoder.BlockData)}
	log.Debug("Try view...")
}

func (this *EncoderController)Post() {
	log.Debug("Try save...")
	defer this.ServeJSON()
	this.Output.Header().Set("Access-Control-Allow-Origin", "*")
	body, err := ioutil.ReadAll(this.Input.Body)
	if err != nil {
		log.Error("Can not read request body:", err)
		this.Data = ErrorMessage{
			Result:"Request error",
			ResultCode:406,
		}
		return
	}
	if len(body) > 16384 {
		log.Error("Content too long:", len(body))
		this.Data = ErrorMessage{
			Result:"Content too long.",
			Note: fmt.Sprintf("Max sign request body size %d bytes", 16384),
			ResultCode:406,
		}
		return
	}
	data_bock := DataBlock{}
	err = json.Unmarshal(body, &data_bock)
	if err != nil {
		log.Error("Invalid JSON:", len(body))
		this.Data = ErrorMessage{
			Result:"Invalid JSON",
			ResultCode:406,
		}
		return
	}
	encoder := tapeencript.TapeBlock{}
	encoder.BlockData = []byte(data_bock.DataBlock)
	block_index, key_index, err := encoder.Encrypt()
	if err != nil {
		log.Error("Can not encode data:", err)
		this.Data = ErrorMessage{
			Result:"Server error",
			ResultCode:500,
		}
		return
	}
	result := EncodedData{}
	result.BlockIndex = block_index
	result.KeyIndex = key_index
	result.Result = "OK"
	result.ResultCode = 200
	this.Data = result
}
