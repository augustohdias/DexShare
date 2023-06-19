package decoder

import (
	"dexshare/src/core/entity"
	"encoding/binary"
	"log"
	"strings"
)

type PokemonDecoder struct{}

func (p *PokemonDecoder) Decode(data []byte) entity.PokemonEntity {
	personality := binary.LittleEndian.Uint32(data[0:4])
	originalTrainerID := binary.LittleEndian.Uint32(data[4:8])
	decryptionKey := originalTrainerID ^ personality
	template := findShufflingTemplate(int(personality % 24))
	growthDataOffset := (strings.Index(template, "G") * 12) + 32
	dataToDecrypt := make([]byte, len(data[growthDataOffset:growthDataOffset+12]))
	copy(dataToDecrypt, data[growthDataOffset:growthDataOffset+12])
	for i := 0; i < len(dataToDecrypt); i += 4 {
		subData := binary.LittleEndian.Uint32(dataToDecrypt[i : i+4])
		decryptedData := subData ^ decryptionKey
		binary.LittleEndian.PutUint32(dataToDecrypt[i:i+4], decryptedData)
	}
	species := int(binary.LittleEndian.Uint16(dataToDecrypt[:2]))
	name := StringDecoder{}.Decode(data[8:18]).(string)
	level := readInt(data[84:85]).(int)
	if species == 0 || species >= 440 || (species >= 252 && species <= 276) {
		return entity.PokemonEntity{}
	}
	log.Printf("#%d: %s Lv.%d\n", species, name, level)
	return entity.PokemonEntity{
		Name:              name,
		NationalDexNumber: species,
		Level:             level,
	}
}

func findShufflingTemplate(x int) string {
	shuffleTemplates := []string{
		"GAEM", "GAME", "GEAM", "GEMA", "GMAE", "GMEA",
		"AGEM", "AGME", "AEGM", "AEMG", "AMGE", "AMEG",
		"EGAM", "EGMA", "EAGM", "EAMG", "EMGA", "EMAG",
		"MGAE", "MGEA", "MAGE", "MAEG", "MEGA", "MEAG",
	}
	if x >= 0 && x < len(shuffleTemplates) {
		return shuffleTemplates[x]
	}
	return ""
}

func readInt(bytes []byte) interface{} {
	acc := 0
	for i, b := range bytes {
		acc += int(b) << (i * 8)
	}
	return acc
}
