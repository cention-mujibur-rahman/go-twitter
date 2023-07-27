package twitter

import (
	"net/http"
	"strconv"
	"strings"
)

// DMField is used for
type DMField string

const (
	// DMFieldID is the unique identifier of the requested Tweet.
	DMFieldID DMField = "id"
	// DMFieldText is the actual UTF-8 text of the Tweet. See twitter-text for details on what characters are currently considered valid.
	DMFieldText DMField = "text"
	// DMFieldAttachments specifies the type of attachments (if any) present in this Tweet.
	DMFieldAttachments DMField = "attachments"
	// DMFieldSenderID is the unique identifier of the User who posted this Tweet
	DMFieldSenderID DMField = "sender_id"
	// DMFieldConversationID is the Tweet ID of the original Tweet of the conversation (which includes direct replies, replies of replies).
	DMFieldConversationID DMField = "dm_conversation_id"
	// DMFieldCreatedAt is the creation time of the Tweet.
	DMFieldCreatedAt DMField = "created_at"
)

// UserDMOpts ...
// https://api.twitter.com/2/dm_events?dm_event.fields=id,text,event_type,dm_conversation_id,created_at,sender_id,attachments,participant_ids,referenced_tweets&event_types=MessageCreate&max_results=100&media.fields=&user.fields=created_at,description,id,location,name,pinned_tweet_id,public_metrics,url,username&expansions=sender_id,referenced_tweets.id,attachments.media_keys,participant_ids
type UserDMOpts struct {
	DMExpansions    []DMExpansion
	MediaFields     []MediaField
	EventFields     []DMField
	UserFields      []UserField
	MaxResults      int
	PaginationToken string
}

func (t UserDMOpts) addQuery(req *http.Request) {
	q := req.URL.Query()
	if len(t.EventFields) > 0 {
		q.Add("dm_event.fields", strings.Join(dmEventStringArray(t.EventFields), ","))
	}
	if len(t.DMExpansions) > 0 {
		q.Add("expansions", strings.Join(dmEexpansionStringArray(t.DMExpansions), ","))
	}
	if len(t.MediaFields) > 0 {
		q.Add("media.fields", strings.Join(mediaFieldStringArray(t.MediaFields), ","))
	}

	if len(t.UserFields) > 0 {
		q.Add("user.fields", strings.Join(userFieldStringArray(t.UserFields), ","))
	}

	if t.MaxResults > 0 {
		q.Add("max_results", strconv.Itoa(t.MaxResults))
	}
	if len(t.PaginationToken) > 0 {
		q.Add("pagination_token", t.PaginationToken)
	}

	if len(q) > 0 {
		req.URL.RawQuery = q.Encode()
	}
}

// DMExpansion ...
type DMExpansion string

const (

	// DMExpansionAttachmentsMediaKeys returns a media object representing the images, videos, GIFs included in the Tweet
	DMExpansionAttachmentsMediaKeys DMExpansion = "attachments.media_keys"
	// DMExpansionSenderID returns a user object representing the DM's sender
	DMExpansionSenderID DMExpansion = "sender_id"
	// DMExpansionReferencedTweetsID returns a user object of the referenced Tweet
	DMExpansionReferencedTweetsID DMExpansion = "referenced_tweets.id"
)

func dmEexpansionStringArray(arr []DMExpansion) []string {
	strs := make([]string, len(arr))
	for i, expansion := range arr {
		strs[i] = string(expansion)
	}
	return strs
}

func dmEventStringArray(arr []DMField) []string {
	strs := make([]string, len(arr))
	for i, event := range arr {
		strs[i] = string(event)
	}
	return strs
}
