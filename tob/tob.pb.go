// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tob.proto

/*
Package tob is a generated protocol buffer package.

It is generated from these files:
	tob.proto

It has these top-level messages:
	Empty
	ServerChangeEvent
	ServerEvent
	Vector
	Quaternion
	PlayerMoveEvent
	PlayerCastEvent
	PlayerAppearance
	PlayerEquiped
	PlayerAnimationEvent
	PlayerEvent
	MonsterSpawnEvent
	MonsterMoveEvent
	MonsterLootEvent
	MonsterDieEvent
	MonsterEvent
	Event
*/
package tob

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EventTopic int32

const (
	EventTopic_SERVER_EVENT  EventTopic = 0
	EventTopic_PLAYER_EVENT  EventTopic = 1
	EventTopic_MONSTER_EVENT EventTopic = 2
)

var EventTopic_name = map[int32]string{
	0: "SERVER_EVENT",
	1: "PLAYER_EVENT",
	2: "MONSTER_EVENT",
}
var EventTopic_value = map[string]int32{
	"SERVER_EVENT":  0,
	"PLAYER_EVENT":  1,
	"MONSTER_EVENT": 2,
}

func (x EventTopic) String() string {
	return proto.EnumName(EventTopic_name, int32(x))
}
func (EventTopic) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ServerEventType int32

const (
	ServerEventType_SERVER_CHANGE ServerEventType = 0
	ServerEventType_SERVER_YIELD  ServerEventType = 1
)

var ServerEventType_name = map[int32]string{
	0: "SERVER_CHANGE",
	1: "SERVER_YIELD",
}
var ServerEventType_value = map[string]int32{
	"SERVER_CHANGE": 0,
	"SERVER_YIELD":  1,
}

func (x ServerEventType) String() string {
	return proto.EnumName(ServerEventType_name, int32(x))
}
func (ServerEventType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type PlayerEventType int32

const (
	PlayerEventType_PLAYER_ENTER     PlayerEventType = 0
	PlayerEventType_PLAYER_EXIT      PlayerEventType = 1
	PlayerEventType_PLAYER_MOVE      PlayerEventType = 2
	PlayerEventType_PLAYER_CAST      PlayerEventType = 3
	PlayerEventType_PLAYER_DAMAGED   PlayerEventType = 4
	PlayerEventType_PLAYER_DIE       PlayerEventType = 5
	PlayerEventType_PLAYER_JUMP      PlayerEventType = 6
	PlayerEventType_PLAYER_CROUCH    PlayerEventType = 7
	PlayerEventType_PLAYER_POSITION  PlayerEventType = 8
	PlayerEventType_PLAYER_EQUIPPED  PlayerEventType = 9
	PlayerEventType_PLAYER_ANIMATION PlayerEventType = 10
)

var PlayerEventType_name = map[int32]string{
	0:  "PLAYER_ENTER",
	1:  "PLAYER_EXIT",
	2:  "PLAYER_MOVE",
	3:  "PLAYER_CAST",
	4:  "PLAYER_DAMAGED",
	5:  "PLAYER_DIE",
	6:  "PLAYER_JUMP",
	7:  "PLAYER_CROUCH",
	8:  "PLAYER_POSITION",
	9:  "PLAYER_EQUIPPED",
	10: "PLAYER_ANIMATION",
}
var PlayerEventType_value = map[string]int32{
	"PLAYER_ENTER":     0,
	"PLAYER_EXIT":      1,
	"PLAYER_MOVE":      2,
	"PLAYER_CAST":      3,
	"PLAYER_DAMAGED":   4,
	"PLAYER_DIE":       5,
	"PLAYER_JUMP":      6,
	"PLAYER_CROUCH":    7,
	"PLAYER_POSITION":  8,
	"PLAYER_EQUIPPED":  9,
	"PLAYER_ANIMATION": 10,
}

func (x PlayerEventType) String() string {
	return proto.EnumName(PlayerEventType_name, int32(x))
}
func (PlayerEventType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Gender int32

const (
	Gender_MALE   Gender = 0
	Gender_FEMALE Gender = 1
)

var Gender_name = map[int32]string{
	0: "MALE",
	1: "FEMALE",
}
var Gender_value = map[string]int32{
	"MALE":   0,
	"FEMALE": 1,
}

func (x Gender) String() string {
	return proto.EnumName(Gender_name, int32(x))
}
func (Gender) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type MonsterEventType int32

const (
	MonsterEventType_MONSTER_SPAWN       MonsterEventType = 0
	MonsterEventType_MONSTER_MOVE        MonsterEventType = 1
	MonsterEventType_MONSTER_ATTACK      MonsterEventType = 2
	MonsterEventType_MONSTER_DIE         MonsterEventType = 3
	MonsterEventType_MONSTER_DESTROY     MonsterEventType = 4
	MonsterEventType_MONSTER_LOOT        MonsterEventType = 5
	MonsterEventType_MONSTER_LOOT_RESULT MonsterEventType = 6
	MonsterEventType_MONSTER_BACK        MonsterEventType = 7
)

var MonsterEventType_name = map[int32]string{
	0: "MONSTER_SPAWN",
	1: "MONSTER_MOVE",
	2: "MONSTER_ATTACK",
	3: "MONSTER_DIE",
	4: "MONSTER_DESTROY",
	5: "MONSTER_LOOT",
	6: "MONSTER_LOOT_RESULT",
	7: "MONSTER_BACK",
}
var MonsterEventType_value = map[string]int32{
	"MONSTER_SPAWN":       0,
	"MONSTER_MOVE":        1,
	"MONSTER_ATTACK":      2,
	"MONSTER_DIE":         3,
	"MONSTER_DESTROY":     4,
	"MONSTER_LOOT":        5,
	"MONSTER_LOOT_RESULT": 6,
	"MONSTER_BACK":        7,
}

func (x MonsterEventType) String() string {
	return proto.EnumName(MonsterEventType_name, int32(x))
}
func (MonsterEventType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ServerChangeEvent struct {
	Previous string `protobuf:"bytes,1,opt,name=previous" json:"previous,omitempty"`
	Current  string `protobuf:"bytes,2,opt,name=current" json:"current,omitempty"`
}

func (m *ServerChangeEvent) Reset()                    { *m = ServerChangeEvent{} }
func (m *ServerChangeEvent) String() string            { return proto.CompactTextString(m) }
func (*ServerChangeEvent) ProtoMessage()               {}
func (*ServerChangeEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ServerChangeEvent) GetPrevious() string {
	if m != nil {
		return m.Previous
	}
	return ""
}

func (m *ServerChangeEvent) GetCurrent() string {
	if m != nil {
		return m.Current
	}
	return ""
}

type ServerEvent struct {
	Id   string          `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type ServerEventType `protobuf:"varint,2,opt,name=type,enum=tob.ServerEventType" json:"type,omitempty"`
}

func (m *ServerEvent) Reset()                    { *m = ServerEvent{} }
func (m *ServerEvent) String() string            { return proto.CompactTextString(m) }
func (*ServerEvent) ProtoMessage()               {}
func (*ServerEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ServerEvent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ServerEvent) GetType() ServerEventType {
	if m != nil {
		return m.Type
	}
	return ServerEventType_SERVER_CHANGE
}

type Vector struct {
	X float32 `protobuf:"fixed32,1,opt,name=x" json:"x,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=y" json:"y,omitempty"`
	Z float32 `protobuf:"fixed32,3,opt,name=z" json:"z,omitempty"`
}

func (m *Vector) Reset()                    { *m = Vector{} }
func (m *Vector) String() string            { return proto.CompactTextString(m) }
func (*Vector) ProtoMessage()               {}
func (*Vector) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Vector) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Vector) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Vector) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

type Quaternion struct {
	A float32 `protobuf:"fixed32,1,opt,name=a" json:"a,omitempty"`
	B float32 `protobuf:"fixed32,2,opt,name=b" json:"b,omitempty"`
	C float32 `protobuf:"fixed32,3,opt,name=c" json:"c,omitempty"`
	D float32 `protobuf:"fixed32,4,opt,name=d" json:"d,omitempty"`
}

func (m *Quaternion) Reset()                    { *m = Quaternion{} }
func (m *Quaternion) String() string            { return proto.CompactTextString(m) }
func (*Quaternion) ProtoMessage()               {}
func (*Quaternion) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Quaternion) GetA() float32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *Quaternion) GetB() float32 {
	if m != nil {
		return m.B
	}
	return 0
}

func (m *Quaternion) GetC() float32 {
	if m != nil {
		return m.C
	}
	return 0
}

func (m *Quaternion) GetD() float32 {
	if m != nil {
		return m.D
	}
	return 0
}

type PlayerMoveEvent struct {
	Target    *Vector     `protobuf:"bytes,1,opt,name=target" json:"target,omitempty"`
	Direction *Quaternion `protobuf:"bytes,2,opt,name=direction" json:"direction,omitempty"`
}

func (m *PlayerMoveEvent) Reset()                    { *m = PlayerMoveEvent{} }
func (m *PlayerMoveEvent) String() string            { return proto.CompactTextString(m) }
func (*PlayerMoveEvent) ProtoMessage()               {}
func (*PlayerMoveEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PlayerMoveEvent) GetTarget() *Vector {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *PlayerMoveEvent) GetDirection() *Quaternion {
	if m != nil {
		return m.Direction
	}
	return nil
}

type PlayerCastEvent struct {
	Id             string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	TargetId       string  `protobuf:"bytes,2,opt,name=targetId" json:"targetId,omitempty"`
	TargetPosition *Vector `protobuf:"bytes,3,opt,name=targetPosition" json:"targetPosition,omitempty"`
}

func (m *PlayerCastEvent) Reset()                    { *m = PlayerCastEvent{} }
func (m *PlayerCastEvent) String() string            { return proto.CompactTextString(m) }
func (*PlayerCastEvent) ProtoMessage()               {}
func (*PlayerCastEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PlayerCastEvent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PlayerCastEvent) GetTargetId() string {
	if m != nil {
		return m.TargetId
	}
	return ""
}

func (m *PlayerCastEvent) GetTargetPosition() *Vector {
	if m != nil {
		return m.TargetPosition
	}
	return nil
}

type PlayerAppearance struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Gender    Gender `protobuf:"varint,2,opt,name=gender,enum=tob.Gender" json:"gender,omitempty"`
	HairColor int32  `protobuf:"varint,3,opt,name=hairColor" json:"hairColor,omitempty"`
}

func (m *PlayerAppearance) Reset()                    { *m = PlayerAppearance{} }
func (m *PlayerAppearance) String() string            { return proto.CompactTextString(m) }
func (*PlayerAppearance) ProtoMessage()               {}
func (*PlayerAppearance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PlayerAppearance) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PlayerAppearance) GetGender() Gender {
	if m != nil {
		return m.Gender
	}
	return Gender_MALE
}

func (m *PlayerAppearance) GetHairColor() int32 {
	if m != nil {
		return m.HairColor
	}
	return 0
}

type PlayerEquiped struct {
	Weapon string `protobuf:"bytes,1,opt,name=weapon" json:"weapon,omitempty"`
	Head   string `protobuf:"bytes,2,opt,name=head" json:"head,omitempty"`
	Chest  string `protobuf:"bytes,3,opt,name=chest" json:"chest,omitempty"`
	Legs   string `protobuf:"bytes,4,opt,name=legs" json:"legs,omitempty"`
	Shoes  string `protobuf:"bytes,5,opt,name=shoes" json:"shoes,omitempty"`
	Shield string `protobuf:"bytes,6,opt,name=shield" json:"shield,omitempty"`
}

func (m *PlayerEquiped) Reset()                    { *m = PlayerEquiped{} }
func (m *PlayerEquiped) String() string            { return proto.CompactTextString(m) }
func (*PlayerEquiped) ProtoMessage()               {}
func (*PlayerEquiped) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *PlayerEquiped) GetWeapon() string {
	if m != nil {
		return m.Weapon
	}
	return ""
}

func (m *PlayerEquiped) GetHead() string {
	if m != nil {
		return m.Head
	}
	return ""
}

func (m *PlayerEquiped) GetChest() string {
	if m != nil {
		return m.Chest
	}
	return ""
}

func (m *PlayerEquiped) GetLegs() string {
	if m != nil {
		return m.Legs
	}
	return ""
}

func (m *PlayerEquiped) GetShoes() string {
	if m != nil {
		return m.Shoes
	}
	return ""
}

func (m *PlayerEquiped) GetShield() string {
	if m != nil {
		return m.Shield
	}
	return ""
}

type PlayerAnimationEvent struct {
	StateHash      []int32   `protobuf:"varint,1,rep,packed,name=stateHash" json:"stateHash,omitempty"`
	NormalizedTime []float32 `protobuf:"fixed32,2,rep,packed,name=normalizedTime" json:"normalizedTime,omitempty"`
	IntParams      []int32   `protobuf:"varint,3,rep,packed,name=intParams" json:"intParams,omitempty"`
	FloatParams    []float32 `protobuf:"fixed32,4,rep,packed,name=floatParams" json:"floatParams,omitempty"`
	BoolParams     []bool    `protobuf:"varint,5,rep,packed,name=boolParams" json:"boolParams,omitempty"`
}

func (m *PlayerAnimationEvent) Reset()                    { *m = PlayerAnimationEvent{} }
func (m *PlayerAnimationEvent) String() string            { return proto.CompactTextString(m) }
func (*PlayerAnimationEvent) ProtoMessage()               {}
func (*PlayerAnimationEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *PlayerAnimationEvent) GetStateHash() []int32 {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *PlayerAnimationEvent) GetNormalizedTime() []float32 {
	if m != nil {
		return m.NormalizedTime
	}
	return nil
}

func (m *PlayerAnimationEvent) GetIntParams() []int32 {
	if m != nil {
		return m.IntParams
	}
	return nil
}

func (m *PlayerAnimationEvent) GetFloatParams() []float32 {
	if m != nil {
		return m.FloatParams
	}
	return nil
}

func (m *PlayerAnimationEvent) GetBoolParams() []bool {
	if m != nil {
		return m.BoolParams
	}
	return nil
}

type PlayerEvent struct {
	Id         string                `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type       PlayerEventType       `protobuf:"varint,2,opt,name=type,enum=tob.PlayerEventType" json:"type,omitempty"`
	Move       *PlayerMoveEvent      `protobuf:"bytes,3,opt,name=move" json:"move,omitempty"`
	Position   *Vector               `protobuf:"bytes,4,opt,name=position" json:"position,omitempty"`
	Cast       *PlayerCastEvent      `protobuf:"bytes,5,opt,name=cast" json:"cast,omitempty"`
	Damage     float32               `protobuf:"fixed32,6,opt,name=damage" json:"damage,omitempty"`
	Appearance *PlayerAppearance     `protobuf:"bytes,7,opt,name=appearance" json:"appearance,omitempty"`
	Equiped    *PlayerEquiped        `protobuf:"bytes,8,opt,name=equiped" json:"equiped,omitempty"`
	Animation  *PlayerAnimationEvent `protobuf:"bytes,9,opt,name=animation" json:"animation,omitempty"`
}

func (m *PlayerEvent) Reset()                    { *m = PlayerEvent{} }
func (m *PlayerEvent) String() string            { return proto.CompactTextString(m) }
func (*PlayerEvent) ProtoMessage()               {}
func (*PlayerEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *PlayerEvent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PlayerEvent) GetType() PlayerEventType {
	if m != nil {
		return m.Type
	}
	return PlayerEventType_PLAYER_ENTER
}

func (m *PlayerEvent) GetMove() *PlayerMoveEvent {
	if m != nil {
		return m.Move
	}
	return nil
}

func (m *PlayerEvent) GetPosition() *Vector {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *PlayerEvent) GetCast() *PlayerCastEvent {
	if m != nil {
		return m.Cast
	}
	return nil
}

func (m *PlayerEvent) GetDamage() float32 {
	if m != nil {
		return m.Damage
	}
	return 0
}

func (m *PlayerEvent) GetAppearance() *PlayerAppearance {
	if m != nil {
		return m.Appearance
	}
	return nil
}

func (m *PlayerEvent) GetEquiped() *PlayerEquiped {
	if m != nil {
		return m.Equiped
	}
	return nil
}

func (m *PlayerEvent) GetAnimation() *PlayerAnimationEvent {
	if m != nil {
		return m.Animation
	}
	return nil
}

type MonsterSpawnEvent struct {
	Id         string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	DemonType  int32   `protobuf:"varint,2,opt,name=demonType" json:"demonType,omitempty"`
	DemonSkin  int32   `protobuf:"varint,3,opt,name=demonSkin" json:"demonSkin,omitempty"`
	WeaponType int32   `protobuf:"varint,4,opt,name=weaponType" json:"weaponType,omitempty"`
	Position   *Vector `protobuf:"bytes,5,opt,name=position" json:"position,omitempty"`
}

func (m *MonsterSpawnEvent) Reset()                    { *m = MonsterSpawnEvent{} }
func (m *MonsterSpawnEvent) String() string            { return proto.CompactTextString(m) }
func (*MonsterSpawnEvent) ProtoMessage()               {}
func (*MonsterSpawnEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *MonsterSpawnEvent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MonsterSpawnEvent) GetDemonType() int32 {
	if m != nil {
		return m.DemonType
	}
	return 0
}

func (m *MonsterSpawnEvent) GetDemonSkin() int32 {
	if m != nil {
		return m.DemonSkin
	}
	return 0
}

func (m *MonsterSpawnEvent) GetWeaponType() int32 {
	if m != nil {
		return m.WeaponType
	}
	return 0
}

func (m *MonsterSpawnEvent) GetPosition() *Vector {
	if m != nil {
		return m.Position
	}
	return nil
}

type MonsterMoveEvent struct {
	Position *Vector `protobuf:"bytes,1,opt,name=position" json:"position,omitempty"`
	Target   *Vector `protobuf:"bytes,2,opt,name=target" json:"target,omitempty"`
}

func (m *MonsterMoveEvent) Reset()                    { *m = MonsterMoveEvent{} }
func (m *MonsterMoveEvent) String() string            { return proto.CompactTextString(m) }
func (*MonsterMoveEvent) ProtoMessage()               {}
func (*MonsterMoveEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *MonsterMoveEvent) GetPosition() *Vector {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *MonsterMoveEvent) GetTarget() *Vector {
	if m != nil {
		return m.Target
	}
	return nil
}

type MonsterLootEvent struct {
	PlayerId  string `protobuf:"bytes,1,opt,name=playerId" json:"playerId,omitempty"`
	ItemId    string `protobuf:"bytes,2,opt,name=itemId" json:"itemId,omitempty"`
	MonsterId string `protobuf:"bytes,3,opt,name=monsterId" json:"monsterId,omitempty"`
}

func (m *MonsterLootEvent) Reset()                    { *m = MonsterLootEvent{} }
func (m *MonsterLootEvent) String() string            { return proto.CompactTextString(m) }
func (*MonsterLootEvent) ProtoMessage()               {}
func (*MonsterLootEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *MonsterLootEvent) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

func (m *MonsterLootEvent) GetItemId() string {
	if m != nil {
		return m.ItemId
	}
	return ""
}

func (m *MonsterLootEvent) GetMonsterId() string {
	if m != nil {
		return m.MonsterId
	}
	return ""
}

type MonsterDieEvent struct {
	Items []int32 `protobuf:"varint,1,rep,packed,name=items" json:"items,omitempty"`
}

func (m *MonsterDieEvent) Reset()                    { *m = MonsterDieEvent{} }
func (m *MonsterDieEvent) String() string            { return proto.CompactTextString(m) }
func (*MonsterDieEvent) ProtoMessage()               {}
func (*MonsterDieEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *MonsterDieEvent) GetItems() []int32 {
	if m != nil {
		return m.Items
	}
	return nil
}

type MonsterEvent struct {
	Id    string             `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type  MonsterEventType   `protobuf:"varint,2,opt,name=type,enum=tob.MonsterEventType" json:"type,omitempty"`
	Spawn *MonsterSpawnEvent `protobuf:"bytes,3,opt,name=spawn" json:"spawn,omitempty"`
	Loot  *MonsterLootEvent  `protobuf:"bytes,4,opt,name=loot" json:"loot,omitempty"`
	Move  *MonsterMoveEvent  `protobuf:"bytes,5,opt,name=move" json:"move,omitempty"`
	Die   *MonsterDieEvent   `protobuf:"bytes,6,opt,name=die" json:"die,omitempty"`
}

func (m *MonsterEvent) Reset()                    { *m = MonsterEvent{} }
func (m *MonsterEvent) String() string            { return proto.CompactTextString(m) }
func (*MonsterEvent) ProtoMessage()               {}
func (*MonsterEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *MonsterEvent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MonsterEvent) GetType() MonsterEventType {
	if m != nil {
		return m.Type
	}
	return MonsterEventType_MONSTER_SPAWN
}

func (m *MonsterEvent) GetSpawn() *MonsterSpawnEvent {
	if m != nil {
		return m.Spawn
	}
	return nil
}

func (m *MonsterEvent) GetLoot() *MonsterLootEvent {
	if m != nil {
		return m.Loot
	}
	return nil
}

func (m *MonsterEvent) GetMove() *MonsterMoveEvent {
	if m != nil {
		return m.Move
	}
	return nil
}

func (m *MonsterEvent) GetDie() *MonsterDieEvent {
	if m != nil {
		return m.Die
	}
	return nil
}

type Event struct {
	Topic EventTopic    `protobuf:"varint,1,opt,name=topic,enum=tob.EventTopic" json:"topic,omitempty"`
	S     *ServerEvent  `protobuf:"bytes,2,opt,name=s" json:"s,omitempty"`
	P     *PlayerEvent  `protobuf:"bytes,3,opt,name=p" json:"p,omitempty"`
	M     *MonsterEvent `protobuf:"bytes,4,opt,name=m" json:"m,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *Event) GetTopic() EventTopic {
	if m != nil {
		return m.Topic
	}
	return EventTopic_SERVER_EVENT
}

func (m *Event) GetS() *ServerEvent {
	if m != nil {
		return m.S
	}
	return nil
}

func (m *Event) GetP() *PlayerEvent {
	if m != nil {
		return m.P
	}
	return nil
}

func (m *Event) GetM() *MonsterEvent {
	if m != nil {
		return m.M
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "tob.Empty")
	proto.RegisterType((*ServerChangeEvent)(nil), "tob.ServerChangeEvent")
	proto.RegisterType((*ServerEvent)(nil), "tob.ServerEvent")
	proto.RegisterType((*Vector)(nil), "tob.Vector")
	proto.RegisterType((*Quaternion)(nil), "tob.Quaternion")
	proto.RegisterType((*PlayerMoveEvent)(nil), "tob.PlayerMoveEvent")
	proto.RegisterType((*PlayerCastEvent)(nil), "tob.PlayerCastEvent")
	proto.RegisterType((*PlayerAppearance)(nil), "tob.PlayerAppearance")
	proto.RegisterType((*PlayerEquiped)(nil), "tob.PlayerEquiped")
	proto.RegisterType((*PlayerAnimationEvent)(nil), "tob.PlayerAnimationEvent")
	proto.RegisterType((*PlayerEvent)(nil), "tob.PlayerEvent")
	proto.RegisterType((*MonsterSpawnEvent)(nil), "tob.MonsterSpawnEvent")
	proto.RegisterType((*MonsterMoveEvent)(nil), "tob.MonsterMoveEvent")
	proto.RegisterType((*MonsterLootEvent)(nil), "tob.MonsterLootEvent")
	proto.RegisterType((*MonsterDieEvent)(nil), "tob.MonsterDieEvent")
	proto.RegisterType((*MonsterEvent)(nil), "tob.MonsterEvent")
	proto.RegisterType((*Event)(nil), "tob.Event")
	proto.RegisterEnum("tob.EventTopic", EventTopic_name, EventTopic_value)
	proto.RegisterEnum("tob.ServerEventType", ServerEventType_name, ServerEventType_value)
	proto.RegisterEnum("tob.PlayerEventType", PlayerEventType_name, PlayerEventType_value)
	proto.RegisterEnum("tob.Gender", Gender_name, Gender_value)
	proto.RegisterEnum("tob.MonsterEventType", MonsterEventType_name, MonsterEventType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ToB service

type ToBClient interface {
	Subscribe(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ToB_SubscribeClient, error)
	Publish(ctx context.Context, opts ...grpc.CallOption) (ToB_PublishClient, error)
}

type toBClient struct {
	cc *grpc.ClientConn
}

func NewToBClient(cc *grpc.ClientConn) ToBClient {
	return &toBClient{cc}
}

func (c *toBClient) Subscribe(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ToB_SubscribeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ToB_serviceDesc.Streams[0], c.cc, "/tob.ToB/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &toBSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ToB_SubscribeClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type toBSubscribeClient struct {
	grpc.ClientStream
}

func (x *toBSubscribeClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *toBClient) Publish(ctx context.Context, opts ...grpc.CallOption) (ToB_PublishClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ToB_serviceDesc.Streams[1], c.cc, "/tob.ToB/Publish", opts...)
	if err != nil {
		return nil, err
	}
	x := &toBPublishClient{stream}
	return x, nil
}

type ToB_PublishClient interface {
	Send(*Event) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type toBPublishClient struct {
	grpc.ClientStream
}

func (x *toBPublishClient) Send(m *Event) error {
	return x.ClientStream.SendMsg(m)
}

func (x *toBPublishClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ToB service

type ToBServer interface {
	Subscribe(*Empty, ToB_SubscribeServer) error
	Publish(ToB_PublishServer) error
}

func RegisterToBServer(s *grpc.Server, srv ToBServer) {
	s.RegisterService(&_ToB_serviceDesc, srv)
}

func _ToB_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ToBServer).Subscribe(m, &toBSubscribeServer{stream})
}

type ToB_SubscribeServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type toBSubscribeServer struct {
	grpc.ServerStream
}

func (x *toBSubscribeServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _ToB_Publish_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ToBServer).Publish(&toBPublishServer{stream})
}

type ToB_PublishServer interface {
	SendAndClose(*Empty) error
	Recv() (*Event, error)
	grpc.ServerStream
}

type toBPublishServer struct {
	grpc.ServerStream
}

func (x *toBPublishServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *toBPublishServer) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ToB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tob.ToB",
	HandlerType: (*ToBServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _ToB_Subscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Publish",
			Handler:       _ToB_Publish_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "tob.proto",
}

func init() { proto.RegisterFile("tob.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xdd, 0x72, 0xdb, 0x44,
	0x14, 0xee, 0xda, 0x96, 0x6d, 0x1d, 0xb7, 0x8e, 0xb2, 0x0d, 0x45, 0x64, 0x3a, 0x25, 0xa3, 0x4e,
	0xdb, 0x90, 0x29, 0x1d, 0x26, 0x1d, 0xe0, 0x5a, 0x75, 0x44, 0x6a, 0x88, 0x7f, 0xba, 0x56, 0x02,
	0xbd, 0x2a, 0x6b, 0x6b, 0x89, 0x35, 0x58, 0x5a, 0x21, 0x29, 0x69, 0xd3, 0x97, 0xe0, 0x82, 0x97,
	0xe0, 0x8a, 0x27, 0xe0, 0x59, 0x78, 0x0b, 0x6e, 0xb8, 0x62, 0xf6, 0x47, 0x3f, 0x76, 0x42, 0x87,
	0x3b, 0x9d, 0xef, 0x7c, 0xe7, 0x9c, 0x3d, 0x7f, 0xbb, 0x02, 0x33, 0xe7, 0xf3, 0x67, 0x49, 0xca,
	0x73, 0x8e, 0x9b, 0x39, 0x9f, 0x3b, 0x1d, 0x30, 0xbc, 0x28, 0xc9, 0xaf, 0x9c, 0x21, 0x6c, 0xcf,
	0x58, 0x7a, 0xc9, 0xd2, 0xc1, 0x92, 0xc6, 0xe7, 0xcc, 0xbb, 0x64, 0x71, 0x8e, 0x77, 0xa1, 0x9b,
	0xa4, 0xec, 0x32, 0xe4, 0x17, 0x99, 0x8d, 0xf6, 0xd0, 0xbe, 0x49, 0x4a, 0x19, 0xdb, 0xd0, 0x59,
	0x5c, 0xa4, 0x29, 0x8b, 0x73, 0xbb, 0x21, 0x55, 0x85, 0xe8, 0x1c, 0x43, 0x4f, 0xb9, 0x52, 0x4e,
	0xfa, 0xd0, 0x08, 0x03, 0x6d, 0xde, 0x08, 0x03, 0xbc, 0x0f, 0xad, 0xfc, 0x2a, 0x61, 0xd2, 0xaa,
	0x7f, 0xb8, 0xf3, 0x4c, 0x9c, 0xa8, 0xc6, 0xf7, 0xaf, 0x12, 0x46, 0x24, 0xc3, 0x39, 0x84, 0xf6,
	0x19, 0x5b, 0xe4, 0x3c, 0xc5, 0xb7, 0x01, 0xbd, 0x93, 0x2e, 0x1a, 0x04, 0xbd, 0x13, 0xd2, 0x95,
	0x34, 0x6f, 0x10, 0x74, 0x25, 0xa4, 0xf7, 0x76, 0x53, 0x49, 0xef, 0x9d, 0x23, 0x80, 0x57, 0x17,
	0x34, 0x67, 0x69, 0x1c, 0xf2, 0x58, 0xe8, 0x68, 0x61, 0x47, 0x85, 0x34, 0x2f, 0xec, 0xe6, 0x42,
	0x5a, 0x14, 0x76, 0x0b, 0x21, 0x05, 0x76, 0x4b, 0x49, 0x81, 0xc3, 0x60, 0x6b, 0xba, 0xa2, 0x57,
	0x2c, 0x1d, 0xf1, 0x4b, 0x5d, 0x8b, 0x87, 0xd0, 0xce, 0x69, 0x7a, 0xce, 0x72, 0xe9, 0xaf, 0x77,
	0xd8, 0x93, 0x07, 0x57, 0xe7, 0x23, 0x5a, 0x85, 0x3f, 0x07, 0x33, 0x08, 0x53, 0xb6, 0xc8, 0x43,
	0x1e, 0xcb, 0x48, 0xbd, 0xc3, 0x2d, 0xc9, 0xab, 0xce, 0x44, 0x2a, 0x86, 0x93, 0x16, 0x61, 0x06,
	0x34, 0xcb, 0x6f, 0xae, 0xd6, 0x2e, 0x74, 0x95, 0xef, 0x61, 0xa0, 0xeb, 0x5c, 0xca, 0xf8, 0x39,
	0xf4, 0xd5, 0xf7, 0x94, 0x67, 0xa1, 0x0c, 0xd9, 0xbc, 0x7e, 0xb4, 0x0d, 0x8a, 0x13, 0x82, 0xa5,
	0x62, 0xba, 0x49, 0xc2, 0x68, 0x4a, 0xe3, 0x05, 0xc3, 0x18, 0x5a, 0x31, 0x8d, 0x98, 0x0e, 0x2b,
	0xbf, 0x45, 0xbe, 0xe7, 0x2c, 0x0e, 0x58, 0xaa, 0x1b, 0xa5, 0x9c, 0x1e, 0x4b, 0x88, 0x68, 0x15,
	0xbe, 0x0f, 0xe6, 0x92, 0x86, 0xe9, 0x80, 0xaf, 0x78, 0x2a, 0x83, 0x1b, 0xa4, 0x02, 0x9c, 0xdf,
	0x10, 0xdc, 0x51, 0xb1, 0xbc, 0x5f, 0x2e, 0xc2, 0x84, 0x05, 0xf8, 0x1e, 0xb4, 0xdf, 0x32, 0x9a,
	0xf0, 0x58, 0x87, 0xd2, 0x92, 0x38, 0xc0, 0x92, 0xd1, 0x22, 0x43, 0xf9, 0x8d, 0x77, 0xc0, 0x58,
	0x2c, 0x59, 0x96, 0x4b, 0xbf, 0x26, 0x51, 0x82, 0x60, 0xae, 0xd8, 0x79, 0x26, 0x5b, 0x65, 0x12,
	0xf9, 0x2d, 0x98, 0xd9, 0x92, 0xb3, 0xcc, 0x36, 0x14, 0x53, 0x0a, 0x22, 0x56, 0xb6, 0x0c, 0xd9,
	0x2a, 0xb0, 0xdb, 0x2a, 0x96, 0x92, 0x9c, 0x3f, 0x11, 0xec, 0xe8, 0x0a, 0xc4, 0x61, 0x44, 0x45,
	0x51, 0x54, 0xe9, 0xef, 0x83, 0x99, 0xe5, 0x34, 0x67, 0x2f, 0x69, 0xb6, 0xb4, 0xd1, 0x5e, 0x53,
	0x24, 0x53, 0x02, 0xf8, 0x31, 0xf4, 0x63, 0x9e, 0x46, 0x74, 0x15, 0xbe, 0x67, 0x81, 0x1f, 0x46,
	0x62, 0x80, 0x9b, 0xfb, 0x0d, 0xb2, 0x81, 0x0a, 0x2f, 0x61, 0x9c, 0x4f, 0x69, 0x4a, 0xa3, 0xcc,
	0x6e, 0x2a, 0x2f, 0x25, 0x80, 0xf7, 0xa0, 0xf7, 0xd3, 0x8a, 0xd3, 0x42, 0xdf, 0x92, 0x2e, 0xea,
	0x10, 0x7e, 0x00, 0x30, 0xe7, 0x7c, 0xa5, 0x09, 0xc6, 0x5e, 0x73, 0xbf, 0x4b, 0x6a, 0x88, 0xf3,
	0x4f, 0x03, 0x7a, 0xba, 0xa8, 0xff, 0x7b, 0xbd, 0x6a, 0xfc, 0x6a, 0xbd, 0x04, 0x33, 0xe2, 0x97,
	0x4c, 0x0f, 0x4d, 0x9d, 0x59, 0x4e, 0x3d, 0x91, 0x0c, 0xfc, 0x04, 0xba, 0x49, 0x31, 0x62, 0xad,
	0xeb, 0x23, 0x56, 0x2a, 0x85, 0xcb, 0x05, 0xcd, 0x72, 0xd9, 0x88, 0x75, 0x97, 0xe5, 0x84, 0x13,
	0xc9, 0x10, 0xdd, 0x09, 0x68, 0x44, 0xcf, 0x99, 0xec, 0x4e, 0x83, 0x68, 0x09, 0x7f, 0x09, 0x40,
	0xcb, 0xc1, 0xb4, 0x3b, 0xd2, 0xcf, 0x47, 0x35, 0x3f, 0xd5, 0xd4, 0x92, 0x1a, 0x11, 0x3f, 0x85,
	0x0e, 0x53, 0x33, 0x66, 0x77, 0xa5, 0x0d, 0xae, 0x27, 0xae, 0x34, 0xa4, 0xa0, 0xe0, 0xaf, 0xc1,
	0xa4, 0x45, 0xef, 0x6d, 0x53, 0xf2, 0x3f, 0xa9, 0xc7, 0x58, 0x9b, 0x0b, 0x52, 0x71, 0x9d, 0xdf,
	0x11, 0x6c, 0x8f, 0x78, 0x9c, 0xe5, 0x2c, 0x9d, 0x25, 0xf4, 0x6d, 0x7c, 0x73, 0x0b, 0xee, 0x83,
	0x19, 0xb0, 0x88, 0xc7, 0x7e, 0xd1, 0x07, 0x83, 0x54, 0x40, 0xa9, 0x9d, 0xfd, 0x1c, 0xc6, 0xc5,
	0xce, 0x94, 0x80, 0x68, 0xbf, 0xda, 0x09, 0x69, 0xdc, 0x92, 0xea, 0x1a, 0xb2, 0xd6, 0x0a, 0xe3,
	0x03, 0xad, 0x70, 0x7e, 0x04, 0x4b, 0x9f, 0xb4, 0xba, 0xc3, 0xea, 0xc6, 0xe8, 0x43, 0x7d, 0xac,
	0x2e, 0xbb, 0xc6, 0x7f, 0x5e, 0x76, 0x4e, 0x50, 0x46, 0x38, 0xe1, 0x3c, 0xaf, 0x5e, 0x0c, 0x59,
	0xc3, 0x61, 0x50, 0xbe, 0x18, 0x5a, 0x16, 0x2d, 0x0f, 0x73, 0x16, 0x95, 0x17, 0x99, 0x96, 0x44,
	0x41, 0x22, 0xe5, 0x67, 0x18, 0xe8, 0x65, 0xaf, 0x00, 0xe7, 0x09, 0x6c, 0xe9, 0x28, 0x47, 0xa1,
	0x4e, 0x63, 0x07, 0x0c, 0x61, 0x9a, 0xe9, 0x25, 0x55, 0x82, 0xf3, 0x37, 0x82, 0xdb, 0x9a, 0x79,
	0x73, 0x5b, 0x3e, 0x5b, 0xdb, 0x0c, 0x35, 0x54, 0x75, 0x83, 0xda, 0x6a, 0x3c, 0x05, 0x23, 0x13,
	0xfd, 0xd5, 0xbb, 0x71, 0xaf, 0xce, 0xad, 0x1a, 0x4f, 0x14, 0x49, 0x38, 0x5e, 0x71, 0x9e, 0xeb,
	0xd5, 0x58, 0x73, 0x5c, 0x56, 0x86, 0x48, 0x8a, 0xa0, 0xca, 0x9d, 0x33, 0xae, 0x53, 0x37, 0x97,
	0xee, 0x31, 0x34, 0x83, 0x50, 0xad, 0x47, 0xb1, 0x4a, 0x1b, 0x85, 0x20, 0x82, 0xe0, 0xfc, 0x8a,
	0xc0, 0x50, 0x09, 0x3f, 0x02, 0x23, 0xe7, 0x49, 0xb8, 0x90, 0x39, 0xf7, 0xf5, 0xcb, 0xa3, 0x52,
	0x13, 0x30, 0x51, 0x5a, 0xfc, 0x00, 0x50, 0xa6, 0xfb, 0x6a, 0x6d, 0xbe, 0xbe, 0x04, 0x89, 0x1b,
	0x08, 0x25, 0x3a, 0x71, 0x6b, 0xf3, 0xfa, 0x20, 0x28, 0xc1, 0x9f, 0x02, 0x8a, 0x74, 0xae, 0xdb,
	0xd7, 0x8a, 0x48, 0x50, 0x74, 0x30, 0x00, 0xa8, 0xa2, 0x62, 0x0b, 0x6e, 0xcf, 0x3c, 0x72, 0xe6,
	0x91, 0x37, 0xde, 0x99, 0x37, 0xf6, 0xad, 0x5b, 0x02, 0x99, 0x9e, 0xb8, 0xaf, 0x4b, 0x04, 0xe1,
	0x6d, 0xb8, 0x33, 0x9a, 0x8c, 0x67, 0x7e, 0x09, 0x35, 0x0e, 0xbe, 0x82, 0xad, 0x8d, 0xbf, 0x02,
	0xc1, 0xd2, 0x9e, 0x06, 0x2f, 0xdd, 0xf1, 0xb1, 0xa7, 0x5c, 0x69, 0xe8, 0xf5, 0xd0, 0x3b, 0x39,
	0xb2, 0xd0, 0xc1, 0x5f, 0xa8, 0x78, 0x54, 0x2b, 0xc3, 0x5a, 0xc0, 0xb1, 0xef, 0x11, 0xeb, 0x16,
	0xde, 0x82, 0x5e, 0x81, 0xfc, 0x30, 0x14, 0x27, 0xa8, 0x80, 0xd1, 0xe4, 0xcc, 0xb3, 0x1a, 0x35,
	0x60, 0xe0, 0xce, 0x7c, 0xab, 0x89, 0x31, 0xf4, 0x35, 0x70, 0xe4, 0x8e, 0xdc, 0x63, 0xef, 0xc8,
	0x6a, 0xe1, 0x3e, 0x40, 0x81, 0x0d, 0x3d, 0xcb, 0xa8, 0x19, 0x7d, 0x7b, 0x3a, 0x9a, 0x5a, 0x6d,
	0x71, 0xe4, 0xc2, 0x0b, 0x99, 0x9c, 0x0e, 0x5e, 0x5a, 0x1d, 0x7c, 0x17, 0xb6, 0x34, 0x34, 0x9d,
	0xcc, 0x86, 0xfe, 0x70, 0x32, 0xb6, 0xba, 0x35, 0xd0, 0x7b, 0x75, 0x3a, 0x9c, 0x4e, 0xbd, 0x23,
	0xcb, 0xc4, 0x3b, 0x60, 0x69, 0xd0, 0x1d, 0x0f, 0x47, 0xae, 0xa4, 0xc2, 0xc1, 0x03, 0x68, 0xab,
	0x57, 0x18, 0x77, 0xa1, 0x35, 0x72, 0x4f, 0x44, 0x19, 0x00, 0xda, 0xdf, 0x78, 0xf2, 0x1b, 0x1d,
	0xfc, 0x81, 0xca, 0xbd, 0x5c, 0x2b, 0x5d, 0x51, 0xe0, 0xd9, 0xd4, 0xfd, 0x7e, 0xac, 0x4a, 0x57,
	0x40, 0x32, 0x65, 0x24, 0x32, 0x2c, 0x10, 0xd7, 0xf7, 0xdd, 0xc1, 0x77, 0xaa, 0x0c, 0x05, 0x26,
	0x52, 0x6c, 0x8a, 0x93, 0x96, 0x80, 0x37, 0xf3, 0xc9, 0xe4, 0xb5, 0xd5, 0xaa, 0xfb, 0x3a, 0x99,
	0x4c, 0x7c, 0xcb, 0xc0, 0x1f, 0xc3, 0xdd, 0x3a, 0xf2, 0x86, 0x78, 0xb3, 0xd3, 0x13, 0xdf, 0x6a,
	0xd7, 0xa9, 0x2f, 0x44, 0x88, 0xce, 0xe1, 0x2b, 0x68, 0xfa, 0xfc, 0x05, 0x7e, 0x04, 0xe6, 0xec,
	0x62, 0x9e, 0x2d, 0xd2, 0x70, 0xce, 0x30, 0xa8, 0xd9, 0x15, 0xbf, 0xa6, 0xbb, 0x50, 0xcd, 0xf1,
	0x17, 0x08, 0x3f, 0x84, 0xce, 0xf4, 0x62, 0xbe, 0x0a, 0xb3, 0x25, 0xae, 0x29, 0x76, 0x6b, 0x06,
	0xfb, 0x68, 0xde, 0x96, 0xbf, 0xb8, 0xcf, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x4f, 0xc5,
	0x6c, 0xef, 0x0a, 0x00, 0x00,
}
