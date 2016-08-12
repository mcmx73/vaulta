package tapeencript

import (
	"github.com/mcmx73/vaulta/components/store"
	"github.com/DeepForestTeam/mobiussign/components/config"
	"github.com/DeepForestTeam/mobiussign/components/log"
	"crypto/rand"
	"errors"
	"fmt"
	"bytes"
	"encoding/binary"
)

const (
	TapeStore = "Blocks"
)

var (
	ErrDuplicateId = errors.New("duplicate block id. can not save")
	ErrTooLongBlock = errors.New("can not encode data: block too long")
)

type TapeBlock struct {
	BlockId   string
	BlockData []byte
}

func (this *TapeBlock)Save() (err error) {
	dumb := TapeBlock{}
	err = store.Get(TapeStore, this.BlockId, &dumb)
	if err != store.ErrKeyNotFound && err != store.ErrSectionNotFound {
		return ErrDuplicateId
	}
	_, err = store.Set(TapeStore, this.BlockId, this)
	return
}

func (this *TapeBlock) Load() (err error) {
	err = store.Get(TapeStore, this.BlockId, this)
	return
}

func (this *TapeBlock)Decrypt(key_index string) (err error) {
	err = this.Load()
	if err == store.ErrKeyNotFound {
		log.Warning("Data Block not found")
		generateRandBlock(&this.BlockData)
		this.Save()
	}
	max_block_size, err := config.GetInt64("ENC_BLOCK_SIZE")
	if err != nil {
		max_block_size = 16374
	}
	key_block := TapeBlock{}
	key_block.BlockId = key_index
	err = key_block.Load()
	if err == store.ErrKeyNotFound {
		//generate fake key and save
		log.Warning("Key Block not found")
		generateRandBlock(&key_block.BlockData)
		key_block.Save()
	}
	decrypted_data := make([]byte, max_block_size + 10)
	for i := 0; i < int(max_block_size); i++ {
		decrypted_data[i] = this.BlockData[i] ^ key_block.BlockData[i]
	}
	data_len_bytes := make([]byte, 2)
	data_len_bytes[0] = decrypted_data[int(max_block_size - 8)]
	data_len_bytes[1] = decrypted_data[int(max_block_size - 9)]
	data_len := bytesToInt16(data_len_bytes)
	log.Debug("DL:", data_len)
	return
}
func (this *TapeBlock)Encrypt() (block_index, key_index string, err error) {
	max_block_size, err := config.GetInt64("ENC_BLOCK_SIZE")
	if err != nil {
		max_block_size = 16374
	}
	if len(this.BlockData) > int(max_block_size) {
		return block_index, key_index, ErrTooLongBlock
	}
	key_block := TapeBlock{}
	key_index, err = store.GetRandomBlock(TapeStore, &key_block)
	source_data := this.BlockData
	data_len := int64(len(source_data))
	this.BlockData = make([]byte, max_block_size + 10)
	for index, val := range source_data {
		this.BlockData[index] = val
	}
	if data_len < max_block_size {
		fill_from := data_len
		fill_len := max_block_size - data_len
		fill_bytes := make([]byte, fill_len)
		rand.Read(fill_bytes)
		for index, val := range fill_bytes {
			this.BlockData[index + int(fill_from)] = val
		}
		data_len_bytes := int16ToBytes(int16(data_len))
		log.Debug(len(data_len_bytes))
		this.BlockData[max_block_size + 8] = data_len_bytes[0]
		this.BlockData[max_block_size + 9] = data_len_bytes[1]
	}
	for index, value := range this.BlockData {
		key_byte := key_block.BlockData[index]
		this.BlockData[index] = value ^ key_byte
	}

	block_index = generateRandKey()
	this.BlockId = block_index
	err = this.Save()
	return
}

func (this *TapeBlock)FillRandom(count int) {
	for {
		this.BlockId = generateRandKey()
		generateRandBlock(&this.BlockData)
		err := this.Save()
		if err != nil {
			log.Error(err)
		}
		count = count - 1
		if count == 0 {
			return
		}
	}
}

func (this *TapeBlock)TotalCoutunt() (count int64) {
	count, _ = store.Count(TapeStore)
	return
}

func generateRandKey() (pepper string) {
	pepper_bytes := make([]byte, 8)
	rand.Read(pepper_bytes)
	pepper = fmt.Sprintf("%x", pepper_bytes)
	return
}
func generateRandBlock(block *[]byte) {
	max_block_size, err := config.GetInt64("ENC_BLOCK_SIZE")
	if err != nil {
		max_block_size = 16374
	}
	*block = make([]byte, max_block_size)
	rand.Read(*block)

}

func int16ToBytes(num int16) []byte {
	buf := new(bytes.Buffer)
	_ = binary.Write(buf, binary.BigEndian, num)
	return buf.Bytes()
}
func bytesToInt16(v []byte) (num int16) {
	buf := bytes.NewReader(v)
	err := binary.Read(buf, binary.BigEndian, &num)
	log.Debug(err)
	return
}