package state

import "go-telegram-tools/model"

type UserState struct {
	Mode  model.ChatMode
	Files [][]byte
}

func GetUserState(userStates map[int64]*UserState, chatID int64) *UserState {
	state, exists := userStates[chatID]
	if !exists {
		state = &UserState{Mode: model.MainMode, Files: [][]byte{}}
		userStates[chatID] = state
	}
	return state
}

func SetUserMode(userState *UserState, mode model.ChatMode) {
	userState.Mode = mode
}
