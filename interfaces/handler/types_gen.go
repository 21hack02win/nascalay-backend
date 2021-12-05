// Package handler provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package handler

import "github.com/gofrs/uuid"

// Defines values for NextShowStatus.
const (
	NextShowStatusAnswer NextShowStatus = "answer"

	NextShowStatusCanvas NextShowStatus = "canvas"

	NextShowStatusEnd NextShowStatus = "end"

	NextShowStatusOdai NextShowStatus = "odai"
)

// Defines values for WsEvent.
const (
	WsEventANSWERCANCEL WsEvent = "ANSWER_CANCEL"

	WsEventANSWERFINISH WsEvent = "ANSWER_FINISH"

	WsEventANSWERREADY WsEvent = "ANSWER_READY"

	WsEventANSWERSEND WsEvent = "ANSWER_SEND"

	WsEventANSWERSTART WsEvent = "ANSWER_START"

	WsEventBREAKROOM WsEvent = "BREAK_ROOM"

	WsEventCHANGEHOST WsEvent = "CHANGE_HOST"

	WsEventDRAWCANCEL WsEvent = "DRAW_CANCEL"

	WsEventDRAWFINISH WsEvent = "DRAW_FINISH"

	WsEventDRAWREADY WsEvent = "DRAW_READY"

	WsEventDRAWSEND WsEvent = "DRAW_SEND"

	WsEventDRAWSTART WsEvent = "DRAW_START"

	WsEventGAMESTART WsEvent = "GAME_START"

	WsEventNEXTROOM WsEvent = "NEXT_ROOM"

	WsEventODAICANCEL WsEvent = "ODAI_CANCEL"

	WsEventODAIFINISH WsEvent = "ODAI_FINISH"

	WsEventODAIREADY WsEvent = "ODAI_READY"

	WsEventODAISEND WsEvent = "ODAI_SEND"

	WsEventREQUESTGAMESTART WsEvent = "REQUEST_GAME_START"

	WsEventRETURNROOM WsEvent = "RETURN_ROOM"

	WsEventROOMNEWMEMBER WsEvent = "ROOM_NEW_MEMBER"

	WsEventROOMSETOPTION WsEvent = "ROOM_SET_OPTION"

	WsEventROOMUPDATEOPTION WsEvent = "ROOM_UPDATE_OPTION"

	WsEventSHOWANSWER WsEvent = "SHOW_ANSWER"

	WsEventSHOWCANVAS WsEvent = "SHOW_CANVAS"

	WsEventSHOWNEXT WsEvent = "SHOW_NEXT"

	WsEventSHOWODAI WsEvent = "SHOW_ODAI"

	WsEventSHOWSTART WsEvent = "SHOW_START"
)

// 回答の入力が完了していることを通知する (ルームの各員 -> サーバー)
type AnswerReadyEvent struct {
	Answer string `json:"answer"`
}

// 回答を送信する (ルームの各員 -> サーバー)
type AnswerSendEvent struct {
	Answer string `json:"answer"`
}

// 絵が飛んできて，回答する (サーバー -> ルーム各員)
type AnswerStartEvent struct {
	Img string `json:"img"`
}

// 最後の回答を受信する (サーバー -> ルーム全員)
type ChangeHostEvent struct {
	HostId string `json:"hostId"`
}

// 絵を送信する (ルームの各員 -> サーバー)
//
// -> (DRAWフェーズが終わってなかったら) また，DRAW_START が飛んでくる
type DrawSendEvent struct {
	Img string `json:"img"`
}

// キャンバス情報とお題を送信する (サーバー -> ルーム各員)
type DrawStartEvent struct {
	AllDrawPhaseNum float32 `json:"allDrawPhaseNum"`
	DrawPhaseNum    int     `json:"drawPhaseNum"`
	Img             string  `json:"img"`
	Odai            string  `json:"odai"`
	TimeLimit       int     `json:"timeLimit"`
}

// ゲームの開始を通知する (サーバー -> ルーム全員)
type GameStartEvent struct {
	OdaiHint  string `json:"odaiHint"`
	TimeLimit int    `json:"timeLimit"`
}

// ルーム参加リクエスト
type JoinRoomRequest struct {
	Avatar int    `json:"avatar"`
	Name   string `json:"name"`
	RoomId string `json:"roomId"`
}

// 新規ルーム情報
type NewRoom struct {
	// ルームID
	RoomId string `json:"roomId"`

	// ユーザーUUID
	UserId uuid.UUID `json:"userId"`
}

// NewRoomRequest defines model for NewRoomRequest.
type NewRoomRequest struct {
	Capacity int    `json:"capacity"`
	Name     string `json:"name"`
}

// 次のWebsocketイベントのリスト
type NextShowStatus string

// お題を送信する (ルームの各員 -> サーバー)
type OdaiSendEvent struct {
	Odai string `json:"odai"`
}

// ルーム情報
type Room struct {
	Capacity int       `json:"capacity"`
	HostId   uuid.UUID `json:"hostId"`
	Members  []User    `json:"members"`
	RoomId   string    `json:"roomId"`
	UserId   uuid.UUID `json:"userId"`
}

// 部屋に追加のメンバーが来たことを通知する (サーバー -> ルーム全員)
type RoomNewMemberEvent map[string]interface{}

// ゲームのオプションを設定する (ホスト -> サーバー)
type RoomSetOptionEvent map[string]interface{}

// RoomUpdateOptionEvent defines model for RoomUpdateOptionEvent.
type RoomUpdateOptionEvent map[string]interface{}

// 最後の回答を受信する (サーバー -> ルーム全員)
type ShowAnswerEvent struct {
	Answer string `json:"answer"`

	// 次のWebsocketイベントのリスト
	Next NextShowStatus `json:"next"`
}

// 次のキャンバスを受信する (サーバー -> ルーム全員)
type ShowCanvasEvent struct {
	Img string `json:"img"`

	// 次のWebsocketイベントのリスト
	Next NextShowStatus `json:"next"`
}

// 最初のお題を受信する (サーバー -> ルーム全員)
type ShowOdaiEvent struct {
	// 次のWebsocketイベントのリスト
	Next NextShowStatus `json:"next"`
	Odai string         `json:"odai"`
}

// User defines model for User.
type User struct {
	Avatar int       `json:"avatar"`
	Name   string    `json:"name"`
	UserId uuid.UUID `json:"userId"`
}

// Websocketイベントのリスト
type WsEvent string

// RoomId defines model for roomId.
type RoomId string

// JoinRoomJSONBody defines parameters for JoinRoom.
type JoinRoomJSONBody JoinRoomRequest

// CreateRoomJSONBody defines parameters for CreateRoom.
type CreateRoomJSONBody NewRoomRequest

// WsJSONBody defines parameters for Ws.
type WsJSONBody struct {
	Body *interface{} `json:"body,omitempty"`

	// Websocketイベントのリスト
	Type *WsEvent `json:"type,omitempty"`
}

// JoinRoomJSONRequestBody defines body for JoinRoom for application/json ContentType.
type JoinRoomJSONRequestBody JoinRoomJSONBody

// CreateRoomJSONRequestBody defines body for CreateRoom for application/json ContentType.
type CreateRoomJSONRequestBody CreateRoomJSONBody

// WsJSONRequestBody defines body for Ws for application/json ContentType.
type WsJSONRequestBody WsJSONBody
