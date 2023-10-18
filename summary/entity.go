package summary

type Summary struct {
	UserId       int64  `json:"userId"`
	Date         string `json:"date"`
	Post         int64  `json:"post,omitempty"`
	LikeReceived int64  `json:"likeReceived,omitempty"`
	Comment      int64  `json:"comment,omitempty"`
}
