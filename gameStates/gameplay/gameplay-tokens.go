package gameplay

import "go-game/data"

type Tokens struct {
	idByToken map[string]int
	tokenById map[int]string
}

func NewTokens(tokens map[int]string) Tokens {
	idByToken := make(map[string]int)
	for id, token := range tokens {
		idByToken[token] = id
	}
	return Tokens{
		idByToken: idByToken,
		tokenById: tokens,
	}
}

func (t *Tokens) GetId(token string) int {
	if _, ok := t.idByToken[token]; !ok {
		return data.UnknownId
	}
	return t.idByToken[token]
}

func (t *Tokens) GetToken(id int) string {
	if _, ok := t.tokenById[id]; !ok {
		return t.tokenById[data.UnknownId]
	}
	return t.tokenById[id]
}
