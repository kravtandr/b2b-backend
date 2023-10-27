package repository

import (
	"b2b/m/internal/services/chat/models"
	"b2b/m/pkg/errors"
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

type ChatRepository interface {
	CheckIfUniqChat(ctx context.Context, productId int64, userId int64) (bool, error)
	NewChat(ctx context.Context, newChat *models.Chat) (*models.Chat, error)
	GetChat(ctx context.Context, chat *models.Chat) (*models.Chat, error)
	WriteNewMsg(ctx context.Context, newMsg *models.Msg) (int64, error)
	GetMsgsFromChat(ctx context.Context, chatId int64, userId int64) (*models.Msgs, error)
	GetAllChatsAndLastMsg(ctx context.Context, userId int64) (*models.ChatsAndLastMsg, error)
	GetAllUserChats(ctx context.Context, userId int64) (*models.Chats, error)
	GetUserLastMsgs(ctx context.Context, userId int64) (*models.Msgs, error)
}

type chatRepository struct {
	queryFactory QueryFactory
	conn         *pgxpool.Pool
}

func (a *chatRepository) CheckIfUniqChat(ctx context.Context, productId int64, userId int64) (bool, error) {
	query := a.queryFactory.CreateCheckIfUniqChat(productId, userId)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)
	count := 0
	unique := false
	if err := row.Scan(&count); err != nil {
		if err == pgx.ErrNoRows {
			return false, errors.ChatDoesNotExist
		}

		return false, err
	}
	if count == 0 {
		unique = true
	}

	return unique, nil
}

func (a *chatRepository) NewChat(ctx context.Context, newChat *models.Chat) (*models.Chat, error) {
	query := a.queryFactory.CreateNewChat(newChat)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)
	chat := &models.Chat{}
	if err := row.Scan(&chat.Id, &chat.Name, &chat.CreatorId, &chat.ProductId, &chat.Status); err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.ChatDoesNotExist
		}

		return nil, err
	}

	return chat, nil
}

func (a *chatRepository) GetChat(ctx context.Context, chat *models.Chat) (*models.Chat, error) {
	query := a.queryFactory.CreateGetChat(chat)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)
	getChat := &models.Chat{}
	if err := row.Scan(&getChat.Id, &getChat.Name, &getChat.CreatorId, &getChat.ProductId, &chat.Status); err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.ChatDoesNotExist
		}

		return nil, err
	}

	return getChat, nil
}

func (a *chatRepository) WriteNewMsg(ctx context.Context, newMsg *models.Msg) (int64, error) {
	query := a.queryFactory.CreateWriteNewMsg(newMsg)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)
	var id int64
	if err := row.Scan(&id); err != nil {
		if err == pgx.ErrNoRows {
			return -1, errors.MsgDoesNotExist
		}
		log.Panicln("Error: WriteNewMsg", err)
		return -1, err
	}
	return id, nil
}

func (a *chatRepository) GetMsgsFromChat(ctx context.Context, chatId int64, userId int64) (*models.Msgs, error) {
	count := 0 // debug
	query := a.queryFactory.CreateGetMsgsFromChat(chatId, userId)
	var msg models.Msg
	var msgs models.Msgs
	rows, err := a.conn.Query(ctx, query.Request, query.Params...)
	if err != nil {
		log.Println("ERROR ", err)
		return &msgs, err
	}
	defer rows.Close()
	for rows.Next() {
		count += 1
		err = rows.Scan(&msg.Id, &msg.ChatId, &msg.SenderId, &msg.ReceiverId, &msg.Checked, &msg.Text, &msg.Type, &msg.Time)
		msgs = append(msgs, msg)
	}
	if rows.Err() != nil {
		log.Println("ERROR ", err)
		return &msgs, err
	}
	log.Println("ROWS COUNT", count)
	return &msgs, err
}

func (a *chatRepository) GetUserLastMsgs(ctx context.Context, userId int64) (*models.Msgs, error) {
	query := a.queryFactory.CreateGetUserLastMsgs(userId)
	var msg models.Msg
	var msgs models.Msgs
	rows, err := a.conn.Query(ctx, query.Request, query.Params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&msg.Id, &msg.ChatId, &msg.SenderId, &msg.ReceiverId, &msg.Checked, &msg.Text, &msg.Type, &msg.Time)
		msgs = append(msgs, msg)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return &msgs, err
}

func (a *chatRepository) GetAllChatsAndLastMsg(ctx context.Context, userId int64) (*models.ChatsAndLastMsg, error) {
	//var chat models.ChatAndLastMsg

	var chats models.ChatsAndLastMsg

	//if chat with no msgs
	onlyChats, err := a.GetAllUserChats(ctx, userId)

	if err != nil {
		return &chats, err
	}
	//for onlyChats
	onlyChatsLM := make(models.ChatsAndLastMsg, 100)
	chats_count := 0
	for _, chat := range *onlyChats {
		chats_count += 1
		onlyChatsLM = append(onlyChatsLM,
			models.ChatAndLastMsg{
				Chat: models.Chat{
					Id:        chat.Id,
					Name:      chat.Name,
					CreatorId: chat.CreatorId,
					ProductId: chat.ProductId,
					Status:    chat.Status,
				},
				LastMsg: models.Msg{},
			})
	}
	query := a.queryFactory.CreateGetAllUserChatsAndLastMsgs(userId)
	rows, err := a.conn.Query(ctx, query.Request, query.Params...)
	if err != nil {
		return &chats, err
	}
	defer rows.Close()
	rows_count := 0
	for rows.Next() {
		err = rows.Scan(&onlyChatsLM[rows_count].Chat.Id, &onlyChatsLM[rows_count].Chat.Name, &onlyChatsLM[rows_count].Chat.CreatorId, &onlyChatsLM[rows_count].Chat.ProductId, &onlyChatsLM[rows_count].Chat.Status, &onlyChatsLM[rows_count].LastMsg.Id, &onlyChatsLM[rows_count].LastMsg.SenderId, &onlyChatsLM[rows_count].LastMsg.ReceiverId, &onlyChatsLM[rows_count].LastMsg.Checked, &onlyChatsLM[rows_count].LastMsg.Text, &onlyChatsLM[rows_count].LastMsg.Type, &onlyChatsLM[rows_count].LastMsg.Time)
		//chat.LastMsg.ChatId = chat.Chat.Id
		log.Println("len(onlyChatsLM)", len(onlyChatsLM))
		if len(onlyChatsLM) < rows_count {
			onlyChatsLM[rows_count].LastMsg.ChatId = onlyChatsLM[rows_count].Chat.Id
		}

		//chats = append(chats, chat)
		rows_count += 1
	}
	if rows.Err() != nil {
		return &chats, err
	}
	//log.Println(chats)
	// if rows_count < chats_count{
	// 	chats = append(chats, chat)
	// }
	// if rows_count == 0 {
	// 	chats = onlyChatsLM
	// }

	return &onlyChatsLM, err

}

func (a *chatRepository) GetAllUserChats(ctx context.Context, userId int64) (*models.Chats, error) {
	query := a.queryFactory.CreateGetAllUserChats(userId)
	var chat models.Chat
	var chats models.Chats
	rows, err := a.conn.Query(ctx, query.Request, query.Params...)
	if err != nil {
		return &chats, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&chat.Id, &chat.Name, &chat.CreatorId, &chat.ProductId, &chat.Status)
		chats = append(chats, chat)
	}
	if rows.Err() != nil {
		return &chats, err
	}
	return &chats, err
}

func NewChatRepository(
	queryFactory QueryFactory,
	conn *pgxpool.Pool,
) ChatRepository {
	return &chatRepository{
		queryFactory: queryFactory,
		conn:         conn,
	}
}
