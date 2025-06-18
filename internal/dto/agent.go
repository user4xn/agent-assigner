package dto

type (
	CommonResponse struct {
		Data any `json:"data"`
		Meta any `json:"meta"`
	}

	ResponseAPICountUnserved struct {
		TotalUnresolved int `json:"total_unresolved"`
		TotalUnserved   int `json:"total_unserved"`
	}

	BodyAPIChatRoom struct {
		Channels       []Channel `json:"channels"`
		UserIds        []any     `json:"user_ids"`
		TagIds         []any     `json:"tag_ids"`
		CursorAfter    *string   `json:"cursor_after"`
		CursorBefore   *string   `json:"cursor_before"`
		IsHandledByBot *bool     `json:"is_handled_by_bot"`
		Limit          *int64    `json:"limit"`
		Name           *string   `json:"name"`
		Order          *string   `json:"order"`
		ServeStatus    *string   `json:"serve_status"`
		Status         *string   `json:"status"`
	}

	Channel struct {
		ChannelID int64  `json:"channel_id"`
		Source    string `json:"source"`
	}

	ResponseAPIChatRoom struct {
		CustomerRooms []CustomerRoom `json:"customer_rooms"`
	}

	CustomerRoom struct {
		ChannelID               *int64      `json:"channel_id,omitempty"`
		ContactID               *int64      `json:"contact_id,omitempty"`
		ID                      *int64      `json:"id,omitempty"`
		IsCalling               *bool       `json:"is_calling,omitempty"`
		IsHandledByBot          *bool       `json:"is_handled_by_bot,omitempty"`
		IsResolved              *bool       `json:"is_resolved,omitempty"`
		IsWaiting               *bool       `json:"is_waiting,omitempty"`
		LastCommentSender       *string     `json:"last_comment_sender,omitempty"`
		LastCommentSenderType   *string     `json:"last_comment_sender_type,omitempty"`
		LastCommentText         *string     `json:"last_comment_text,omitempty"`
		LastCommentTimestamp    *string     `json:"last_comment_timestamp,omitempty"`
		LastCustomerCommentText *string     `json:"last_customer_comment_text,omitempty"`
		LastCustomerTimestamp   *string     `json:"last_customer_timestamp,omitempty"`
		Name                    *string     `json:"name,omitempty"`
		RoomBadge               interface{} `json:"room_badge"`
		RoomID                  *string     `json:"room_id,omitempty"`
		RoomType                *string     `json:"room_type,omitempty"`
		Source                  *string     `json:"source,omitempty"`
		UserAvatarURL           *string     `json:"user_avatar_url,omitempty"`
		UserID                  *string     `json:"user_id,omitempty"`
	}
)
