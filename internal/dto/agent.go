package dto

type (
	CommonResponse struct {
		Data any `json:"data"`
		Meta any `json:"meta"`
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

	ResponseOtherAgent struct {
		Agents []Agent `json:"agents"`
	}

	Agent struct {
		AvatarURL            string        `json:"avatar_url"`
		CreatedAt            string        `json:"created_at"`
		CurrentCustomerCount int64         `json:"current_customer_count"`
		Email                string        `json:"email"`
		ForceOffline         bool          `json:"force_offline"`
		ID                   int64         `json:"id"`
		IsAvailable          bool          `json:"is_available"`
		IsReqOtpReset        any           `json:"is_req_otp_reset"`
		LastLogin            *string       `json:"last_login"`
		Name                 string        `json:"name"`
		SDKEmail             string        `json:"sdk_email"`
		SDKKey               string        `json:"sdk_key"`
		Type                 int64         `json:"type"`
		TypeAsString         string        `json:"type_as_string"`
		UserChannels         []UserChannel `json:"user_channels"`
		UserRoles            []UserRole    `json:"user_roles"`
	}

	UserChannel struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}

	UserRole struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}

	BodyAssignAgent struct {
		AgentID            int64   `json:"agent_id"`
		RoomID             int64   `json:"room_id"`
		ReplaceLatestAgent *string `json:"replace_latest_agent"`
		MaxAgent           *string `json:"max_agent"`
	}
)
