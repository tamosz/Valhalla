package constants

const (
	MAPLE_VERSION           = 28
	CLIENT_HEADER_SIZE      = 4
	INTERSERVER_HEADER_SIZE = 4
	OPCODE_LENGTH           = 1

	// Opcodes Server -> Client
	SEND_LOGIN_RESPONCE           = 0x01
	SEND_LOGIN_WORLD_META         = 0x03
	SEND_LOGIN_PIN_REGISTER       = 0x07 // Add 1 byte, 1 = register a pin
	SEND_LOGIN_PIN_STUFF          = 0x08 // Setting pin good
	SEND_LOGIN_SEND_WORLD_LIST    = 0x09
	SEND_LOGIN_RESTARTER          = 0x15
	SEND_LOGIN_CHARACTER_DATA     = 0x0A
	SEND_LOGIN_CHARACTER_MIGRATE  = 0x0B
	SEND_LOGIN_NAME_CHECK_RESULT  = 0x0C
	SEND_LOGIN_NEW_CHARACTER_GOOD = 0x0D
	SEND_LOGIN_DELETE_CHARACTER   = 0x0E

	SEND_CHANNEL_INVENTORY_OPERATION  = 0x18
	SEND_CHANNEL_STAT_CHANGE          = 0x1A
	SEND_CHANNEL_SKILL_RECORD_UPDATE  = 0x1D
	SEND_CHANNEL_LIE_DETECTOR_TEST    = 0x23
	SEND_CHANNEL_PARTY_INFO           = 0x2D
	SEND_CHANNEL_BROADCAST_MESSAGE    = 0x32
	SEND_CHANNEL_WARP_TO_MAP          = 0x36
	SEND_CHANNEL_PORTAL_CLOSED        = 0x3A
	SEND_CHANNEL_BUBBLESS_CHAT        = 0x3D
	SEND_CHANNEL_WHISPER              = 0x3E
	SEND_CHANNEL_QUIZ_Q_AND_A         = 0x44
	SEND_CHANNEL_CHARCTER_ENTER_FIELD = 0x4E
	SEND_CHANNEL_CHARCTER_LEAVE_FIELD = 0x4F
	SEND_CHANNEL_ALL_CHAT_MSG         = 0x51
	SEND_CHANNEL_PLAYER_MOVEMENT      = 0x65
	SEND_CHANNEL_PLAYER_USE_SKILL     = 0x66
	SEND_CHANNEL_PLAYER_EMOTION       = 0x6C
	SEND_CHANNEL_PLAYER_ANIMATION     = 0x70
	SEND_CHANNEL_LEVEL_UP_ANIMATION   = 0x79
	SEND_CHANNEL_SHOW_MOB             = 0x86
	SEND_CHANNEL_CONTROL_MOB          = 0x88
	SEND_CHANNE_MOVE_MOB              = 0x8A
	SEND_CHANNEL_CONTROL_MOB_ACK      = 0x8B
	SEND_CHANNEL_NPC_SPAWN_1          = 0x97
	SEND_CHANNEL_NPC_SPAWN_2          = 0x9B
	SEND_CHANNEL_SPAWN_DOOR           = 0xB1
	SEND_CHANNEL_REMOVE_DOOR          = 0xB2

	// Opcodes Client -> Server
	RECV_LOGIN_REQUEST          = 0x01
	RECV_LOGIN_CHANNEL_SELECT   = 0x04
	RECV_LOGIN_WORLD_SELECT     = 0x05
	RECV_LOGIN_CHECK_LOGIN      = 0x08
	RECV_LOGIN_CREATE_CHARACTER = 0x09
	RECV_LOGIN_SELECT_CHARACTER = 0x0B
	RECV_LOGIN_NAME_CHECK       = 0x0D
	RECV_LOGIN_NEW_CHARACTER    = 0x0E
	RECV_LOGIN_DELETE_CHAR      = 0x0F
	RECV_PING                   = 0x12
	RECV_RETURN_TO_LOGIN_SCREEN = 0x14

	RECV_CHANNEL_PLAYER_LOAD = 0x0C

	RECV_CHANNEL_USE_PORTAL                 = 0x17
	RECV_CHANNEL_REQUEST_TO_ENTER_CASH_SHOP = 0x19
	RECV_CHANNEL_MOVEMENT                   = 0x1A
	RECV_CHANNEL_STANDARD_SKILL             = 0x1D
	RECV_CHANNEL_RANGED_SKILL               = 0x1F
	RECV_CHANNEL_DMG_RECV                   = 0x21
	RECV_CHANNEL_PLAYER_SEND_ALL_CHAT       = 0x22
	RECV_CHANNEL_EMOTION                    = 0x23
	RECV_CHANNEL_NPC_DIALOGUE               = 0x27
	RECV_CHANNEL_CHANGE_STAT                = 0x36
	RECV_CHANNEL_PASSIVE_REGEN              = 0x37
	RECV_CHANNEL_SKILL_UPDATE               = 0x38
	RECV_CHANNEL_SPECIAL_SKILL_USAGE        = 0x39
	RECV_CHANNEL_DOUBLE_CLICK_CHARACTER     = 0x3F
	RECV_CHANNEL_LIE_DETECTOR_RESULT        = 0x45
	RECV_CHANNEL_PARTY_INFO                 = 0x4F
	RECV_CHANNEL_GUILD_MANAGEMENT           = 0x51
	RECV_CHANNEL_GUILD_REJECT               = 0x52
	RECV_CHANNEL_ADD_BUDDY                  = 0x55
	RECV_CHANNEL_MOB_MOVEMENT               = 0x6A
)
