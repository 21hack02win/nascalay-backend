// Package oapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package oapi

import "github.com/gofrs/uuid"

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

// Defines values for WsNextShowStatus.
const (
	WsNextShowStatusAnswer WsNextShowStatus = "answer"

	WsNextShowStatusCanvas WsNextShowStatus = "canvas"

	WsNextShowStatusEnd WsNextShowStatus = "end"

	WsNextShowStatusOdai WsNextShowStatus = "odai"
)

// ユーザーが描画するキャンバスの分割情報・描画位置
type Canvas struct {
	// ボードの座標ID
	AreaId int `json:"areaId"`

	// ボード名
	BoardName string `json:"boardName"`
}

// 新規ルーム作成リクエスト
type CreateRoomRequest struct {
	// アバターの種類
	Avatar int `json:"avatar"`

	// ルームの最大収容人数
	Capacity int `json:"capacity"`

	// ユーザー名
	Username string `json:"username"`
}

// ルーム参加リクエスト
type JoinRoomRequest struct {
	// アバターの種類
	Avatar int `json:"avatar"`

	// ルームID
	RoomId string `json:"roomId"`

	// ユーザー名
	Username string `json:"username"`
}

// ルーム情報
type Room struct {
	// ルームの最大収容人数
	Capacity int `json:"capacity"`

	// ホストのユーザーUUID
	HostId uuid.UUID `json:"hostId"`

	// 現在ルームにいるメンバーの情報
	Members []User `json:"members"`

	// ルームID
	RoomId string `json:"roomId"`

	// ユーザーUUID
	UserId uuid.UUID `json:"userId"`
}

// ユーザー情報
type User struct {
	// アバターの種類
	Avatar int `json:"avatar"`

	// ユーザーUUID
	UserId uuid.UUID `json:"userId"`

	// ユーザー名
	Username string `json:"username"`
}

// 回答を送信する (ルームの各員 -> サーバー)
type WsAnswerSendEventBody struct {
	// 回答
	Answer string `json:"answer"`
}

// 絵が飛んできて，回答する (サーバー -> ルーム各員)
type WsAnswerStartEventBody struct {
	// 画像ID
	Img string `json:"img"`

	// 制限時間
	TimeLimit int `json:"timeLimit"`
}

// 最後の回答を受信する (サーバー -> ルーム全員)
type WsChangeHostEventBody struct {
	// ホストのユーザーUUID
	HostId string `json:"hostId"`
}

// 絵を送信する (ルームの各員 -> サーバー)
//
// -> (DRAWフェーズが終わってなかったら) また，DRAW_START が飛んでくる
type WsDrawSendEventBody struct {
	// 画像ID
	Img string `json:"img"`
}

// キャンバス情報とお題を送信する (サーバー -> ルーム各員)
type WsDrawStartEventBody struct {
	// 全DRAWフェーズ数
	AllDrawPhaseNum float32 `json:"allDrawPhaseNum"`

	// ユーザーが描画するキャンバスの分割情報・描画位置
	Canvas Canvas `json:"canvas"`

	// 現在のDRAWフェーズの番号
	DrawPhaseNum int `json:"drawPhaseNum"`

	// 画像ID
	Img string `json:"img"`

	// お題
	Odai string `json:"odai"`

	// 制限時間
	TimeLimit int `json:"timeLimit"`
}

// Websocketイベントのリスト
type WsEvent string

// ゲームの開始を通知する (サーバー -> ルーム全員)
type WsGameStartEventBody struct {
	// お題のサジェスト
	OdaiHint string `json:"odaiHint"`

	// 制限時間
	TimeLimit int `json:"timeLimit"`
}

// 次のWebsocketイベントのリスト
type WsNextShowStatus string

// お題を送信する (ルームの各員 -> サーバー)
type WsOdaiSendEventBody struct {
	// お題
	Odai string `json:"odai"`
}

// 部屋に追加のメンバーが来たことを通知する (サーバー -> ルーム全員)
type WsRoomNewMemberEventBody struct {
	// ルームの最大収容人数
	Capacity int `json:"capacity"`

	// ホストのユーザーUUID
	HostId uuid.UUID `json:"hostId"`

	// 現在ルームにいるメンバーの情報
	Members []User `json:"members"`
}

// ゲームのオプションを設定する (ホスト -> サーバー)
type WsRoomSetOptionEventBody struct {
	// 何か
	Something string `json:"something"`
}

// ゲームの設定を更新する (サーバー -> ルーム全員)
type WsRoomUpdateOptionEventBody struct {
	// 何か
	Something string `json:"something"`
}

// 最後の回答を受信する (サーバー -> ルーム全員)
type WsShowAnswerEventBody struct {
	// 回答
	Answer string `json:"answer"`

	// 次のWebsocketイベントのリスト
	Next WsNextShowStatus `json:"next"`
}

// 次のキャンバスを受信する (サーバー -> ルーム全員)
type WsShowCanvasEventBody struct {
	// 画像ID
	Img string `json:"img"`

	// 次のWebsocketイベントのリスト
	Next WsNextShowStatus `json:"next"`
}

// 最初のお題を受信する (サーバー -> ルーム全員)
type WsShowOdaiEventBody struct {
	// 次のWebsocketイベントのリスト
	Next WsNextShowStatus `json:"next"`

	// お題
	Odai string `json:"odai"`
}

// ルームID
type RoomIdInPath string

// UserIdInQuery defines model for userIdInQuery.
type UserIdInQuery string

// JoinRoomJSONBody defines parameters for JoinRoom.
type JoinRoomJSONBody JoinRoomRequest

// CreateRoomJSONBody defines parameters for CreateRoom.
type CreateRoomJSONBody CreateRoomRequest

// WsJSONBody defines parameters for Ws.
type WsJSONBody struct {
	// Embedded fields due to inline allOf schema
	Body interface{} `json:"body"`

	// Websocketイベントのリスト
	Type WsEvent `json:"type"`
}

// WsParams defines parameters for Ws.
type WsParams struct {
	// ユーザーUUID
	User UserIdInQuery `json:"user"`
}

// JoinRoomJSONRequestBody defines body for JoinRoom for application/json ContentType.
type JoinRoomJSONRequestBody JoinRoomJSONBody

// CreateRoomJSONRequestBody defines body for CreateRoom for application/json ContentType.
type CreateRoomJSONRequestBody CreateRoomJSONBody

// WsJSONRequestBody defines body for Ws for application/json ContentType.
type WsJSONRequestBody WsJSONBody
