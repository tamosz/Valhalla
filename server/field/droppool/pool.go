package droppool

import (
	"math"
	"time"

	"github.com/Hucaru/Valhalla/mnet"
	"github.com/Hucaru/Valhalla/mpacket"
	"github.com/Hucaru/Valhalla/server/item"
	"github.com/Hucaru/Valhalla/server/pos"
)

const (
	SpawnDisappears      = 0
	SpawnNormal          = 1
	SpawnShow            = 2
	SpawnFadeAtTopOfDrop = 3
)

type field interface {
	Send(mpacket.Packet) error
	CalculateFinalDropPos(pos.Data) pos.Data
}

type controller interface {
	Send(mpacket.Packet)
	Conn() mnet.Client
}

// Data structure for the pool
type Data struct {
	instance field
	poolID   int32
	drops    []drop
}

// CreateNewPool for drops
func CreateNewPool(inst field) Data {
	return Data{instance: inst}
}

func (pool *Data) nextID() int32 {
	pool.poolID++

	if pool.poolID == math.MaxInt32-1 {
		pool.poolID = math.MaxInt32 / 2
	} else if pool.poolID == 0 {
		pool.poolID = 1
	}

	return pool.poolID
}

// CanClose the pool down
func (pool Data) CanClose() bool {
	return false
}

// PlayerShowDrops when entering instance
func (pool Data) PlayerShowDrops(plr controller) {
	for _, drop := range pool.drops {
		plr.Send(packetShowDrop(SpawnShow, drop))
	}
}

// PlayerAttemptPickup of item
func (pool *Data) PlayerAttemptPickup(dropID int32, position pos.Data) (bool, item.Data) {
	return false, item.Data{}
}

// CreateMobDrop from a mobID from a player at a given location
func (pool *Data) CreateMobDrop(mesos int32, itemID int32, dropFrom pos.Data) {

}

// CreatePlayerDrop into field
func (pool *Data) CreatePlayerDrop(spawnType byte, dropType byte, mesos int32, item item.Data, dropFrom pos.Data, expire bool, ownerID, partyID int32) {
	finalPos := pool.instance.CalculateFinalDropPos(dropFrom)

	drop := drop{
		ID:      pool.nextID(),
		ownerID: ownerID,
		partyID: partyID,
		mesos:   mesos,
		item:    item,

		expireTime:  0,
		timeoutTime: 0,
		neverExpire: false,

		originPos: dropFrom,
		finalPos:  finalPos,

		dropType: dropType,
	}

	pool.drops = append(pool.drops, drop)

	pool.instance.Send(packetShowDrop(spawnType, drop))
}

// Update logic for the pool e.g. drops disappear
func (pool *Data) Update(t time.Time) {
}