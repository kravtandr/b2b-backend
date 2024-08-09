package repository

import (
	"b2b/m/internal/services/chat/models"
	"b2b/m/pkg/query"
)

type QueryFactory interface {
	CreateCheckIfUniqChat(productId int64, userId int64) *query.Query
	CreateNewChat(newChat *models.Chat) *query.Query
	CreateUpdateChat(chat *models.Chat) *query.Query
	CreateDeleteChat(chat_id int64) *query.Query
	CreateGetChat(newChat *models.Chat) *query.Query
	CreateGetChatById(chatId int64) *query.Query
	CreateWriteNewMsg(newMsg *models.Msg) *query.Query
	CreateGetMsgsFromChat(chatId int64, userId int64) *query.Query
	CreateGetUserLastMsgs(userId int64) *query.Query
	CreateGetAllUserChatsAndLastMsgs(userId int64) *query.Query
	CreateGetAmountOfUserChats(userId int64) *query.Query
	CreateUserCreatedChats(userId int64) *query.Query
	CreateUpdateMsg(msg *models.Msg) *query.Query
	CreateGetMsgById(msgId int64) *query.Query
}

type queryFactory struct{}

func (q *queryFactory) CreateUpdateMsg(msg *models.Msg) *query.Query {
	return &query.Query{
		Request: createUpdateMsg,
		Params:  []interface{}{msg.Checked, msg.Text, msg.Type, msg.Id},
	}
}

func (q *queryFactory) CreateGetMsgById(msgId int64) *query.Query {
	return &query.Query{
		Request: createUpdateMsg,
		Params:  []interface{}{msgId},
	}
}

func (q *queryFactory) CreateUpdateChat(chat *models.Chat) *query.Query {
	return &query.Query{
		Request: createUpdateChat,
		Params:  []interface{}{chat.Name, chat.Status, chat.Blured, chat.Id},
	}
}

func (q *queryFactory) CreateDeleteChat(chat_id int64) *query.Query {
	return &query.Query{
		Request: createDeleteChat,
		Params:  []interface{}{chat_id},
	}
}
func (q *queryFactory) CreateCheckIfUniqChat(productId int64, userId int64) *query.Query {
	return &query.Query{
		Request: createCheckIfUniqChat,
		Params:  []interface{}{productId, userId},
	}
}

func (q *queryFactory) CreateUserCreatedChats(userId int64) *query.Query {
	return &query.Query{
		Request: createUserCreatedChats,
		Params:  []interface{}{userId},
	}
}

func (q *queryFactory) CreateNewChat(newChat *models.Chat) *query.Query {
	return &query.Query{
		Request: createNewChat,
		Params:  []interface{}{newChat.Name, newChat.CreatorId, newChat.ProductId},
	}
}

func (q *queryFactory) CreateGetChat(chat *models.Chat) *query.Query {
	return &query.Query{
		Request: createGetChat,
		Params:  []interface{}{chat.ProductId, chat.CreatorId},
	}
}

func (q *queryFactory) CreateGetChatById(chatId int64) *query.Query {
	return &query.Query{
		Request: createGetChatById,
		Params:  []interface{}{chatId},
	}
}

func (q *queryFactory) CreateWriteNewMsg(newMsg *models.Msg) *query.Query {
	return &query.Query{
		Request: createWriteNewMsg,
		Params:  []interface{}{newMsg.ChatId, newMsg.SenderId, newMsg.ReceiverId, newMsg.Text, newMsg.Type},
	}
}

func (q *queryFactory) CreateGetMsgsFromChat(chatId int64, userId int64) *query.Query {
	return &query.Query{
		Request: createGetMsgsFromChat,
		Params:  []interface{}{chatId, userId},
	}
}

func (q *queryFactory) CreateGetAmountOfUserChats(userId int64) *query.Query {
	return &query.Query{
		Request: createGetAmountOfUserChats,
		Params:  []interface{}{userId},
	}
}

func (q *queryFactory) CreateGetAllUserChatsAndLastMsgs(userId int64) *query.Query {
	return &query.Query{
		Request: createGetLastMsgsFromAllUserChats,
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
