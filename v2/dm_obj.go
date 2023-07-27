package twitter

// DMRaw represents the direct message structure
type DMRaw struct {
	Data     []*DMObj    `json:"data"`
	Includes *DMIncludes `json:"includes,omitempty"`
	Errors   []*ErrorObj `json:"errors,omitempty"`
}

// DirectMessageResponse contains the information from the user direct message callout
type DirectMessageResponse struct {
	Raw       *DMRaw
	RateLimit *RateLimit
}

// DMObj is the primary object on the DM endpoints
type DMObj struct {
	ID               string            `json:"id"`
	Text             string            `json:"text"`
	Attachments      *DMAttachmentsObj `json:"attachments,omitempty"`
	SenderID         string            `json:"sender_id,omitempty"`
	DMConversationID string            `json:"dm_conversation_id,omitempty"`
	CreatedAt        string            `json:"created_at,omitempty"`
	EventType        string            `json:"event_type,omitempty"`
}

// DMAttachmentsObj ...
type DMAttachmentsObj struct {
	MediaKeys []string `json:"media_keys"`
}

// DMIncludes ...
type DMIncludes struct {
	Users []*UserObj    `json:"users,omitempty"`
	Media []*DMMediaObj `json:"media,omitempty"`
}

// DMMediaObj ...
type DMMediaObj struct {
	MediaKey string `json:"media_key"`
	Type     string `json:"type"`
	URL      string `json:"url"`
}

// CreateDirectMessageResponse ...
type CreateDirectMessageResponse struct {
	Data struct {
		DMConversationID string `json:"dm_conversation_id"`
		DMEventID        string `json:"dm_event_id"`
	} `json:"data"`
	RateLimit *RateLimit
}
