package nx

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/rand"
)

type Life struct {
	ID      uint32
	Cy      int64
	F       int64
	Fh      int16
	Hide    bool
	MobTime int64
	Rx0     int16
	Rx1     int16
	Npc     bool
	X       int16
	Y       int16
}

type Portal struct {
	ID      byte
	Tm      uint32
	Pt      byte
	IsSpawn bool
	X       int16
	Y       int16
}

type Stage struct {
	Life         []Life
	ForcedReturn uint32
	ReturnMap    uint32
	MobRate      float64
	Town         bool
	Portals      []Portal
}

var Maps map[uint32]Stage

func GetRandomSpawnPortal(mapID uint32) byte {
	var portals []byte
	for _, v := range Maps[mapID].Portals {
		if v.IsSpawn {
			portals = append(portals, v.ID)
		}
	}

	rand.Seed(uint64(time.Now().Unix()))

	return portals[rand.Int()%len(portals)]
}

func getMapInfo() {
	Maps = make(map[uint32]Stage)
	var maps []string

	// Get the list of maps
	for _, mapSet := range []string{"0", "1", "2", "9"} {
		path := "Map/Map/Map"
		result := searchNode(path+mapSet, func(cursor *node) {
			list := make([]string, int(cursor.ChildCount))

			for i := uint32(0); i < uint32(cursor.ChildCount); i++ {
				n := nodes[cursor.ChildID+i]
				list[i] = path + mapSet + "/" + strLookup[n.NameID]
			}

			maps = append(maps, list...)
		})

		if !result {
			panic("Bad search: Map/Map/Map" + mapSet)
		}
	}
	// Populate the Maps object - Refactor
	for _, mapPath := range maps {
		result := searchNode(mapPath, func(cursor *node) {
			mapStr := strings.Split(mapPath, "/")
			val, err := strconv.Atoi(strings.Split(mapStr[len(mapStr)-1], ".")[0])

			if err != nil {
				panic(err)
			}

			mapID := uint32(val)
			var lifes node
			var info node
			var portals node

			for i := uint32(0); i < uint32(cursor.ChildCount); i++ {
				mapChild := nodes[cursor.ChildID+i]
				switch strLookup[mapChild.NameID] {
				case "life":
					lifes = mapChild
				case "info":
					info = mapChild
				case "portal":
					portals = mapChild
				}
			}
			mapItem := Stage{Life: make([]Life, lifes.ChildCount)}

			// Portal handling
			mapItem.Portals = getPortalItem(portals)

			// Info handling
			for i := uint32(0); i < uint32(info.ChildCount); i++ {
				n := nodes[info.ChildID+i]

				for j := uint32(0); j < uint32(n.ChildCount); j++ {
					infoNode := nodes[n.ChildID+j]
					switch strLookup[infoNode.NameID] {
					case "forcedReturn":
						mapItem.ForcedReturn = uint32(dataToInt64(infoNode.Data))
					case "mobRate":
						mapItem.MobRate = math.Float64frombits(dataToUint64(infoNode.Data))
					case "returnMap":
						mapItem.ReturnMap = uint32(dataToInt64(infoNode.Data))
					case "town":
						mapItem.Town = bool(dataToInt64(infoNode.Data) == 1)
					}
				}
			}

			// Life handling
			for i := uint32(0); i < uint32(lifes.ChildCount); i++ {
				mapItem.Life[i] = getLifeItem(nodes[lifes.ChildID+i])
			}

			Maps[mapID] = mapItem
		})
		if !result {
			panic("Bad search:" + mapPath)
		}
	}
}

func getPortalItem(n node) []Portal {
	portals := make([]Portal, n.ChildCount)

	for i := uint32(0); i < uint32(n.ChildCount); i++ {
		p := nodes[n.ChildID+i]
		portal := Portal{}

		portalNumber, err := strconv.Atoi(strLookup[p.NameID])

		if err != nil {
			panic(err)
		}

		portal.ID = byte(portalNumber)

		for j := uint32(0); j < uint32(p.ChildCount); j++ {
			options := nodes[p.ChildID+j]

			switch strLookup[options.NameID] {
			case "pt":
				portal.Pt = options.Data[0]
			case "pn":
				portal.IsSpawn = bool(strLookup[dataToUint32(options.Data)] == "sp")
			case "tm":
				portal.Tm = dataToUint32(options.Data)
			case "x":
				portal.X = dataToInt16(options.Data)
			case "y":
				portal.Y = dataToInt16(options.Data)
			default:
			}
		}

		portals[i] = portal
	}

	return portals
}

func getLifeItem(n node) Life {
	lifeItem := Life{}
	for i := uint32(0); i < uint32(n.ChildCount); i++ {
		lifeNode := nodes[n.ChildID+i]

		switch strLookup[lifeNode.NameID] {
		case "id":
			val, err := strconv.Atoi(strLookup[dataToUint32(lifeNode.Data)])

			if err != nil {
				panic(err)
			}

			lifeItem.ID = uint32(val)
		case "cy":
			lifeItem.Cy = dataToInt64(lifeNode.Data)
		case "f":
			lifeItem.F = dataToInt64(lifeNode.Data)
		case "fh":
			lifeItem.Fh = dataToInt16(lifeNode.Data)
		case "hide":
			lifeItem.Hide = bool(dataToInt64(lifeNode.Data) == 1)
		case "mobTime":
			lifeItem.MobTime = dataToInt64(lifeNode.Data)
		case "rx0":
			lifeItem.Rx0 = dataToInt16(lifeNode.Data)
		case "rx1":
			lifeItem.Rx1 = dataToInt16(lifeNode.Data)
		case "type":
			lifeItem.Npc = bool(strLookup[dataToUint32(lifeNode.Data)] == "n")
		case "x":
			lifeItem.X = dataToInt16(lifeNode.Data)
		case "y":
			lifeItem.Y = dataToInt16(lifeNode.Data)
		case "info":
			// Don't think this is needed for anythng?
		default:
			fmt.Println("Unkown life type from nx file:", strLookup[lifeNode.NameID], "->", lifeNode.Data)
		}
	}
	return lifeItem
}