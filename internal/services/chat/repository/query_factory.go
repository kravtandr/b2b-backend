package repository

import (
	"b2b/m/internal/services/chat/models"
	"b2b/m/pkg/query"
)

type QueryFactory interface {
	CreateCheckIfUniqChat(productId int64, userId int64) *query.Query
	CreateNewChat(newChat *models.Chat) *query.Query
	CreateWriteNewMsg(newMsg *models.Msg) *query.Query
	CreateGetMsgsFromChat(chatId int64) *query.Query
	CreateGetAllUserChats(userId int64) *query.Query
	CreateGetUserLastMsgs(userId int64) *query.Query
}

type queryFactory struct{}

func (q *queryFactory) CreateCheckIfUniqChat(productId int64, userId int64) *query.Query {
	return &query.Query{
		Request: createCheckIfUniqChat,
		Params:  []interface{}{productId, userId},
	}
}

func (q *queryFactory) CreateNewChat(newChat *models.Chat) *query.Query {
	return &query.Query{
		Request: createNewChat,
		Params:  []interface{}{newChat.Name, newChat.CreatorId, newChat.ProductId},
	}
}

func (q *queryFactory) CreateWriteNewMsg(newMsg *models.Msg) *query.Query {
	return &query.Query{
		Request: createWriteNewMsg,
		Params:  []interface{}{newMsg.Id, newMsg.ChatId, newMsg.Checked, newMsg.Text, newMsg.Type, newMsg.Time},
	}
}

func (q *queryFactory) CreateGetMsgsFromChat(chatId int64) *query.Query {
	return &query.Query{
		Request: createGetMsgsFromChat,
		Params:  []interface{}{chatId},
	}
}

func (q *queryFactory) CreateGetAllUserChats(userId int64) *query.Query {
	return &query.Query{
		Request: createGetAllUserChats,
		Params:  []interface{}{userId},
	}
}

func (q *queryFactory) CreateGetUserLastMsgs(userId int64) *query.Query {
	return &query.Query{
		Request: createGetLastMsgsFromAllUserChats,
		Params:  []interface{}{userId},
	}
}

func NewQueryFactory() QueryFactory {
	return &queryFactory{}
}
