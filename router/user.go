package router

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/pquerna/ffjson/ffjson"

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
	ID           int32
	Privileges   int32
	Achievements int32
	Country      string
	UserName     UserName
	Banned       Banned
	Silenced     Silenced
	Leaderboard  consts.Leaderboard
}

// *Biep* *boop*
func UserRouter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestedUser := fmt.Sprintf("%s", vars["user"])
	p := r.URL.Query().Get("p")
	pmode := 0
	if p != "" {
		var err error
		pmode, err = strconv.Atoi(p)
		if err != nil {
			JSONAnswer(500, "Serverside Exception!", w, r)
			return
		}
	}

	if requestedUser == "" {
		JSONAnswer(404, "No Username/UserID applied! please use /api/v1/user/(userid or username)", w, r)
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
		JSONAnswer(404, "User not found!", w, r)
		return
	}
	b, err := ffjson.Marshal(User{
		ID: u.ID,
		UserName: UserName{
			Normal: u.UserName,
			Safe:   u.UserNameSafe,
		},
		Privileges:   u.Privileges,
		Achievements: u.Achievements,
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
		Country:     consts.ToCountryCode(uint8(u.Status.Country)),
		Leaderboard: *usertools.GetLeaderboard(u, int8(pmode)),
	})
	if err != nil {
		JSONAnswer(500, "Server side Error!", w, r)
		return
	}
	JSONAnswer(200, b, w, r)
}
