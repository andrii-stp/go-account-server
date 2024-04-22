package main

import (
	"github.com/golang-jwt/jwt/v4"
)

// Token: {"token_type":"Bearer","expires_at":1594843861,"expires_in":19430,"refresh_token":"9ec0cedf070a8b3832f7be43b5acca8bc976deae","access_token":"a16bffde2871928071e98fbff5674b12d4fd0198","athlete":{"id":24336962,"username":null,"resource_state":2,"firstname":"Eric","lastname":"Lang","city":"Seattle","state":"Washington","country":"United States","sex":"M","premium":false,"summit":false,"created_at":"2017-08-19T18:38:37Z","updated_at":"2020-07-07T20:09:26Z","badge_type_id":0,"profile_medium":"https://dgalywyr863hv.cloudfront.net/pictures/athletes/23376952/7221532/1/medium.jpg","profile":"https://dgalywyr863hv.cloudfront.net/pictures/athletes/23376952/7221532/1/large.jpg","friend":null,"follower":null}}

type Payload struct {
	ID    string
	Name  string
	Email string
	Host  string

	ExpiresAt int

	Claims []Claims
}

type Claims struct {
	Role string

	jwt.RegisteredClaims
}
