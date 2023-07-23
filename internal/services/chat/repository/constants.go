package repository

// add skip limit
const (
	createCheckIfUniqChat = "SELECT * FROM Chats WHERE product_id = $1 AND creater_id = $2 "
	createNewChat         = "INSERT INTO Chats (name, creator_id, product_id) VALUES ($1,$2,$3) RETURNING id "
	createWriteNewMsg     = "INSERT INTO Msgs (chat_id,sender_id, receiver_id, checked, text, type) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id, time"
	createGetMsgsFromChat = "SELECT * FROM Msgs WHERE chat_id = $1"
	//createGetAllChatsAndLastMsg       = "SELECT chat.id, chat.name, chat.creator_id, chat.product_id,msg.id, msg.chat_id, msg.sender_id, msg.receiver_id, msg.checked,msg.text,msg.type, MIN(msg.time) FROM Chats as chat JOIN Msgs as msg ON chat.id=msg.chat_id ORDER BY msg.time"
	createGetAllUserChats             = "SELECT * FROM Chats WHERE sender_id = $1"
	createGetLastMsgFromChat          = "SELECT * FROM Msgs WHERE chat_id = $1 ORDER BY time LIMIT 1"
	createGetLastMsgsFromAllUserChats = "SELECT id, sender_id, receiver_id, checked, text, type, MAX(time) FROM Msgs WHERE sender_id = $1 OR receiver_id = $1"
)
