package resources

const (
	ServiceHost     = "http://localhost:8004"
	BaseUserStorage = "http://localhost:8002/users-storage/v1"
	BaseInitiative  = "http://localhost:8003/initiative-tracker/v1"
)

type Endpoint struct {
	Public  string
	Private string
}

var ReactionsEndpoints = struct {
	Likes   Endpoint
	Reposts Endpoint
	Reports Endpoint
}{
	Likes: Endpoint{
		Public:  "/public/reactions/likes/",
		Private: "/private/reactions/likes/",
	},
	Reposts: Endpoint{
		Public:  "/public/reactions/reposts/",
		Private: "/private/reactions/reposts/",
	},
	Reports: Endpoint{
		Public:  "/public/reactions/reports/",
		Private: "/private/reactions/reports/",
	},
}

var UserStorageEndpoints = struct {
	Base Endpoint
}{
	Base: Endpoint{
		Public:  "/public/users/",
		Private: "/private/users/",
	},
}

var InitiativeEndpoints = struct {
	Base         Endpoint
	Participants Endpoint
	Points       Endpoint
}{
	Base: Endpoint{
		Public:  "/public/initiatives/",
		Private: "/private/initiatives/",
	},
	Participants: Endpoint{
		Public:  "/participants",
		Private: "/participants",
	},
	Points: Endpoint{
		Public:  "/points",
		Private: "/points",
	},
}

var ChatEndpoints = Endpoint{
	Public:  "/public/chat/",
	Private: "/private/chat/",
}

const Pagination10EndLink = "?page[size]=10&page[number]=10"
