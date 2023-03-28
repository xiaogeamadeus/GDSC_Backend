package service

import (
	"errors"
	"hello-run/dao"
)

func GetRoomInfoByRoomId(roomId int64) (*FullRoomVo, error) {
	//first get room info
	room, err := dao.GetRoomInfoById(roomId)
	if err != nil {
		return nil, err
	}
	//then get household info
	var householdVoList []HouseholdVo
	householdList, err := dao.GetHouseholdListByRoomId(roomId)
	if err != nil {
		return nil, err
	}
	householdVoList = GetHouseholdVoList(householdList)

	shortRoomVo := ShortRoomVo{
		Id:       room.Id,
		RoomName: room.RoomName,
	}

	return &FullRoomVo{
		ShortRoomVo: shortRoomVo,
		City:        room.City,
		Household:   householdVoList,
	}, nil
}

func DeleteRoomByRoomId(userId string, roomId int64) error {
	//first test if the room exists
	_, err := dao.GetRoomInfoById(roomId)
	if err != nil {
		return err
	}
	//check if it is the last room of the user
	count, err := dao.GetRoomCountByUser(userId)
	if err != nil {
		return err
	}
	if count == 1 {
		return errors.New("could not delete the last room")
	}
	return dao.DeleteRoomById(roomId)
}

func CreateRoom(userId string, roomName string, city string) (*FullRoomVo, error) {
	roomId, err := dao.CreateRoom(userId, roomName, city)
	if err != nil {
		return nil, err
	}
	//return new room info
	return GetRoomInfoByRoomId(roomId)
}

func UpdateRoomByRoomId(roomId int64, roomName string, city string) (*FullRoomVo, error) {
	//first test if the room exists
	_, err := dao.GetRoomInfoById(roomId)
	if err != nil {
		return nil, err
	}
	roomId, err = dao.UpdateRoom(roomId, roomName, city)
	if err != nil {
		return nil, err
	}
	//return updated room info
	return GetRoomInfoByRoomId(roomId)
}
