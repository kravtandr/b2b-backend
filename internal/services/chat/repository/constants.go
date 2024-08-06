package repository

// add skip limit
const (
	createCheckIfUniqChat             = "SELECT COUNT(*) FROM Chats WHERE product_id = $1 AND creator_id = $2 "
	createGetChat                     = "SELECT id, name, creator_id, product_id, status, blured FROM Chats WHERE product_id = $1 AND creator_id = $2"
	createNewChat                     = "INSERT INTO Chats (name, creator_id, product_id) VALUES ($1,$2,$3) RETURNING id,name, creator_id, product_id, status, blured "
	createWriteNewMsg                 = "INSERT INTO Msgs (chat_id, sender_id, receiver_id, text, type) VALUES ($1,$2,$3,$4,$5) RETURNING id"
	createGetMsgsFromChat             = "SELECT id, chat_id, sender_id, receiver_id, checked, text, type, created_at FROM Msgs WHERE chat_id = $1 AND (sender_id = $2 OR receiver_id = $2) ORDER BY created_at"
	createGetLastMsgFromChat          = "SELECT * FROM Msgs WHERE chat_id = $1 ORDER BY created_at LIMIT 1"
	createGetLastMsgsFromAllUserChats = "SELECT msg2.chat_id, chats.name, chats.creator_id, product_id, chats.status, chats.blured, msg2.id, msg2.sender_id, msg2.receiver_id, msg2.checked, msg2.text, msg2.type, msg2.created_at FROM Msgs as msg2 JOIN (SELECT max(msg1.id), min(msg1.created_at) as time, msg1.chat_id FROM Msgs as msg1 WHERE msg1.sender_id = $1 OR msg1.receiver_id = $1 GROUP BY msg1.chat_id) as msg3 on msg3.max = msg2.id JOIN Chats as chats on msg2.chat_id = chats.id"
	createGetAmountOfUserChats        = "SELECT Count(*) FROM Msgs as msg2 JOIN (SELECT max(msg1.id), min(msg1.created_at) as time, msg1.chat_id FROM Msgs as msg1 WHERE msg1.sender_id = $1 OR msg1.receiver_id = $1 GROUP BY msg1.chat_id) as msg3 on msg3.max = msg2.id JOIN Chats as chats on msg2.chat_id = chats.id"
	createUserCreatedChats            = "SELECT  chat.id, chat.name, chat.creator_id, chat.product_id, chat.status, chat.blured FROM Chats as chat WHERE chat.creator_id = $1"
	createDeleteChat                  = "DELETE FROM Chats WHERE id = $1"
	createUpdateChat                  = "UPDATE Chats SET name = $1, status = $2, blured = $3 WHERE id = $4"
	createGetChatById                 = "SELECT id, name, creator_id, product_id, status, blured FROM Chats WHERE id = $1"
)
