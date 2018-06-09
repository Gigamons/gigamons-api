package Router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Gigamons/common/consts"

	"github.com/Gigamons/common/tools/usertools"
	"github.com/gorilla/mux"
)

type Silenced struct {
	IsSilenced bool
	Until      time.Time
	Because    string
}

type Banned struct {
	IsBanned bool
	Until    time.Time
	Because  string
}

type UserName struct {
	Normal string
	Safe   string
}

type User struct {
	ID          int32
	Privileges  int32
	Country     int16
	UserName    UserName
	Banned      Banned
	Silenced    Silenced
	Leaderboard consts.Leaderboard
}

// *Biep* *boop*
func UserRouter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestedUser := fmt.Sprintf("%s", vars["user"])

	if requestedUser == "" {
		JsonException(404, "No Username/UserID applied! please use /api/user/(userid or username)", w, r)
		return
	}

	isInt := true
	str, err := strconv.Atoi(requestedUser)
	if err != nil {
		isInt = false
	}

	var u *consts.User
	if isInt {
		u = usertools.GetUser(str)
	} else {
		u = usertools.GetUser((usertools.GetUserID(requestedUser)))
	}

	// Always set important data to nil/empty string
	u.EMail = ""
	u.BCryptPassword = ""

	if u.ID == 0 {
		JsonException(404, "User not found!", w, r)
		return
	}
	b, err := json.MarshalIndent(User{
		ID: u.ID,
		UserName: UserName{
			Normal: u.UserName,
			Safe:   u.UserNameSafe,
		},
		Privileges: u.Privileges,
		Banned: Banned{
			IsBanned: u.Status.Banned,
			Until:    u.Status.BannedUntil,
			Because:  u.Status.BannedReason,
		},
		Silenced: Silenced{
			IsSilenced: u.Status.Silenced,
			Until:      u.Status.SilencedUntil,
			Because:    u.Status.SilencedReason,
		},
		Country:     u.Status.Country,
		Leaderboard: usertools.GetLeaderboard(*u, int8(0)),
	}, "", " ")
	if err != nil {
		JsonException(500, "Server side Error!", w, r)
		return
	}
	w.WriteHeader(200)
	w.Write(b)
}
