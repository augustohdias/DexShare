package parser

import (
	"bytes"
	"dexshare/src/core/entity"
	"dexshare/src/core/parser/decoder"
	"dexshare/src/core/parser/schema"
)

type SaveFile struct {
	SaveA schema.SaveData
	SaveB schema.SaveData
}

func LoadSaveFile(file []byte) schema.SaveData {
	saveACh := make(chan schema.SaveData)
	saveBCh := make(chan schema.SaveData)
	go parseSave("A", file, saveACh)
	go parseSave("B", file, saveBCh)
	saveFile := SaveFile{
		SaveA: <-saveACh,
		SaveB: <-saveBCh,
	}
	if saveFile.SaveA.SaveCount > saveFile.SaveB.SaveCount {
		return saveFile.SaveA
	}
	return saveFile.SaveB
}

func parseSave(section string, file []byte, replyChannel chan<- schema.SaveData) {
	gameSaveBytes := loadBytes(section, file)
	signature := []byte{0x25, 0x20, 0x01, 0x08}
	// Iterator?
	sections := make(map[int][]byte)
	startingOffset := 0
	saveCount := 0
	for {
		relativeOffset := bytes.Index(gameSaveBytes[startingOffset:], signature)
		if relativeOffset == -1 {
			break
		}
		endingOffset := startingOffset + relativeOffset
		footer := gameSaveBytes[startingOffset+relativeOffset-4 : startingOffset+relativeOffset+8]
		sectionNumber := int(footer[0])
		saveCount = readInt(footer[8:12]).(int)
		sections[sectionNumber] = gameSaveBytes[startingOffset : endingOffset+8]
		startingOffset = endingOffset + 8
	}
	var pcBuffer [][]byte
	for i := 5; i < 14; i++ {
		pcBuffer = append(pcBuffer, sections[i])
	}
	pc := readPCSection(pcBuffer)
	trainerInfo := readTrainerInfo(sections[0])
	team := readTeamSection(trainerInfo.Game, sections[1])
	saveData := schema.SaveData{
		SaveCount:   saveCount,
		TrainerInfo: trainerInfo,
		Team:        team,
		PC:          pc,
	}
	replyChannel <- saveData
}

func readPCSection(pcBuffer [][]byte) schema.PCSection {
	const MaxBoxSize = 30
	const PokemonSize = 80
	var pokemons []entity.PokemonEntity
	for i, buffer := range pcBuffer {
		for j := 0; j < MaxBoxSize; j++ {
			box := buffer
			if i == 0 {
				box = buffer[4:]
			}
			pokemonDecoder := decoder.PokemonDecoder{}
			pokemon := pokemonDecoder.Decode(box[j*PokemonSize:])
			if (pokemon != entity.PokemonEntity{}) {
				pokemon.Level = 0 // Pokemons on PC needs too much to calculate level
				pokemons = append(pokemons, pokemon)
			}
		}
	}
	return schema.PCSection{Pokemons: pokemons}
}

func readTeamSection(game schema.GameType, data []byte) schema.TeamSection {
	const PokemonSize = 100
	if game == schema.FireRedLeafGreen {
		teamSize := readInt(data[0x034 : 0x034+4]).(int)
		partyData := data[0x038 : 0x038+600]
		var pokemons []entity.PokemonEntity
		for i := 0; i < teamSize; i++ {
			pokemonDecoder := decoder.PokemonDecoder{}
			pokemons = append(pokemons, pokemonDecoder.Decode(partyData[i*PokemonSize:]))
		}
		return schema.TeamSection{Size: teamSize, Pokemons: pokemons}
	}
	return schema.TeamSection{}
}

func readTrainerInfo(data []byte) schema.TrainerInfoSection {
	gameCode := readInt(data[0xAC : 0xAC+4]).(int)
	game := schema.FireRedLeafGreen
	if gameCode == 0 {
		game = schema.RubySapphire
	} else if gameCode == 1 {
		game = schema.FireRedLeafGreen
	} else {
		game = schema.Emerald
	}
	return schema.TrainerInfoSection{
		TrainerName: decoder.StringDecoder{}.Decode(data[:7]).(string),
		TrainerID:   readInt(data[10:14]).(int),
		Game:        game,
	}
}

func readInt(bytes []byte) interface{} {
	acc := 0
	for i, b := range bytes {
		acc += int(b) << (i * 8)
	}
	return acc
}

func loadBytes(section string, file []byte) []byte {
	var saveBytes []byte
	if section == "A" {
		saveBytes = make([]byte, 57344)
		copy(saveBytes, file[:57344])
		return saveBytes
	}
	saveBytes = make([]byte, 57344)
	copy(saveBytes, file[57344:])
	return saveBytes
}
