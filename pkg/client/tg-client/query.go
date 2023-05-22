package tg_client

import "strconv"

type Query interface {
	params() QueryParams
}

type UpdateQuery struct {
	Offset  int
	Limit   int
	Timeout int
}

func (query UpdateQuery) params() QueryParams {
	return QueryParams{
		"offset":  addNonZero(query.Offset),
		"limit":   addNonZero(query.Limit),
		"timeout": addNonZero(query.Timeout),
	}
}

type SendMessageQuery struct {
	ChatId int
	Text   string
}

func (query SendMessageQuery) params() QueryParams {
	return QueryParams{
		"chat_id": strconv.Itoa(query.ChatId),
		"text":    query.Text,
	}
}

func addNonZero(digit int) string {
	if digit == 0 {
		return "10"
	}

	return strconv.Itoa(digit)
}
