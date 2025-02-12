package generates

import (
	"bytes"
	"encoding/base64"
	"strconv"
	"strings"

	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/utils/uuid"
)

// NewAccessGenerate create to generate the access token instance
func NewAccessGenerate() *AccessGenerate {
	return &AccessGenerate{}
}

// AccessGenerate generate the access token
type AccessGenerate struct {
}

// Token based on the UUID generated token
func (ag *AccessGenerate) Token(data *oauth2.GenerateBasic, isGenRefresh bool) (string, string, error) {
	buf := bytes.NewBufferString(data.Client.GetID())
	buf.WriteString(data.UserID)
	buf.WriteString(strconv.FormatInt(data.CreateAt.UnixNano(), 10))

	access := base64.URLEncoding.EncodeToString(uuid.NewMD5(uuid.Must(uuid.NewRandom()), buf.Bytes()).Bytes())
	access = strings.ToUpper(strings.TrimRight(access, "="))
	refresh := ""
	if isGenRefresh {
		refresh = base64.URLEncoding.EncodeToString(uuid.NewSHA1(uuid.Must(uuid.NewRandom()), buf.Bytes()).Bytes())
		refresh = strings.ToUpper(strings.TrimRight(refresh, "="))
	}

	return access, refresh, nil
}
