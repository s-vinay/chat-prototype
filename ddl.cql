
-- Keyspace creation
CREATE KEYSPACE IF NOT EXISTS chat
WITH REPLICATION = {
	'class' : 'NetworkTopologyStrategy',
	'datacenter1' : 1
};


-- Table creation (conversation)
CREATE TABLE IF NOT EXISTS chat.conversation (
	initiator_id text,
	conversation_id uuid,
	created_ts timestamp,
	conversation_status text,
	PRIMARY KEY (initiator_id, conversation_id)
) WITH comment = 'Conversation initiator, status and other details';


-- Table creation (message)
CREATE TABLE IF NOT EXISTS chat.message (
	conversation_id uuid,
	message_id uuid,
	message_content varchar,
	initiator_id text,
	created_ts timestamp,
	message_status text,
	parent_message_id uuid,
	PRIMARY KEY (conversation_id, message_id)
) WITH comment = 'Messages within a conversation with initiator, status and other details';


-- Table creation (Unread messages)
CREATE TABLE IF NOT EXISTS chat.unread (
	initiator_id text,
	conversation_id uuid,
	message_id uuid,
	recipient_id text,
	message_content varchar,
	PRIMARY KEY (initiator_id, conversation_id, message_id, recipient_id)
) WITH comment = 'Unread messages';


-- Table creation (user-conversation relationship)
CREATE TABLE IF NOT EXISTS chat.user_conversation (
	user_id text,
	conversation_id uuid,
	PRIMARY KEY (user_id, conversation_id)
) WITH comment = 'Participants in a conversation';


-- Table creation (user last seen)
CREATE TABLE IF NOT EXISTS chat.lastseen (
	user_id text PRIMARY KEY,
	last_ts timestamp
) WITH comment = 'Last seen timestamp of a user';