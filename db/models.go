package db

import (
	"encoding/json"
	"fmt"
	"strings"
)

// The SessionInfo struct represents a session announcement fetched from the database
// It is also used to insert new entries.
// When inserting, the "Started" field is ignored and the current timestamp is used
type SessionInfo struct {
	Realm              string   `json:"realm"`
	Host               string   `json:"host"`
	Port               int      `json:"port"`
	Id                 string   `json:"id"`
	Protocol           string   `json:"protocol"`
	Title              string   `json:"title"`
	Users              int      `json:"users"`
	MaxUsers           int      `json:"maxusers,omitempty"`
	Usernames          []string `json:"usernames"`
	Password           bool     `json:"password"`
	Nsfm               bool     `json:"nsfm"`
	Owner              string   `json:"owner"`
	Started            string   `json:"started"`
	Private            bool     `json:"private,omitempty"`
	Closed             bool     `json:"closed,omitempty"`
	ActiveDrawingUsers int      `json:"activedrawingusers"`
	AllowWeb           bool     `json:"allowweb,omitempty"`
}

func (info SessionInfo) HostAddress() string {
	if strings.ContainsRune(info.Host, ':') {
		return fmt.Sprintf("[%s]:%d", info.Host, info.Port)
	} else {
		return fmt.Sprintf("%s:%d", info.Host, info.Port)
	}
}

func (info *SessionInfo) UnmarshalJSON(data []byte) error {
	info.ActiveDrawingUsers = -1
	type jsonSessionInfo SessionInfo
	return json.Unmarshal(data, (*jsonSessionInfo)(info))
}

// Minimum info needed to join a session
type JoinSessionInfo struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Id   string `json:"id"`
}

// Info about a newly inserted session
type NewSessionInfo struct {
	ListingId int64  `json:"id"`
	UpdateKey string `json:"key"`
}

// Session list querying options
type QueryOptions struct {
	Realm    string // filter by realm
	Title    string // filter by title
	Nsfm     bool   // show NSFM sessions
	Protocol string // filter by protocol version (comma separated list accepted)
}

type AdminSession struct {
	Id                 int64    `json:"id"`
	Realm              string   `json:"realm"`
	Host               string   `json:"host"`
	Port               int      `json:"port"`
	SessionId          string   `json:"sessionid"`
	Protocol           string   `json:"protocol"`
	Title              string   `json:"title"`
	Users              int      `json:"users"`
	MaxUsers           int      `json:"maxusers,omitempty"`
	Usernames          []string `json:"usernames"`
	Password           bool     `json:"password"`
	Nsfm               bool     `json:"nsfm"`
	Owner              string   `json:"owner"`
	Started            string   `json:"started"`
	LastActive         string   `json:"lastactive"`
	Unlisted           bool     `json:"unlisted"`
	UpdateKey          string   `json:"updatekey"`
	ClientIp           string   `json:"clientip"`
	Alias              string   `json:"alias,omitempty"`
	UnlistReason       string   `json:"unlistreason,omitempty"`
	Kicked             bool     `json:"kicked"`
	TimedOut           bool     `json:"timedout"`
	Closed             bool     `json:"closed"`
	Included           bool     `json:"included"`
	Error              string   `json:"error,omitempty"`
	ActiveDrawingUsers int      `json:"activedrawingusers"`
	AllowWeb           bool     `json:"allowweb,omitempty"`
}

type AdminHostBan struct {
	Id      int64  `json:"id"`
	Host    string `json:"host"`
	Expires string `json:"expires,omitempty"`
	Active  bool   `json:"active"`
	Notes   string `json:"notes,omitempty"`
}

type AdminRole struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	Admin          bool   `json:"admin"`
	AccessSessions int    `json:"accesssessions"`
	AccessHostBans int    `json:"accesshostbans"`
	AccessRoles    int    `json:"accessroles"`
	AccessUsers    int    `json:"accessusers"`
	Used           bool   `json:"used"`
}

type AdminUserDetail struct {
	Id           int64     `json:"id"`
	Name         string    `json:"name"`
	Role         AdminRole `json:"role"`
	PasswordHash string    `json:"-"`
}

type AdminUser struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}
