package auth

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

func (r Repository) GetUserBySessionID(ctx context.Context, sessionID string) (model.GetUserBySessionID, error) {
	query := `
		local arg = {...} 
		local sessionID = arg[1]

		session = box.space.sessions.index.id:select(sessionID)
        if session then
			user = box.space.users.index.id:select(session.user_id)
			if user then
				return user
			else
				return {}
			end
    	else
        	return {}
    	end	
	`

	resp, err := r.conn.Eval(query, []interface{}{
		sessionID,
	})
	if err != nil {
		return model.GetUserBySessionID{}, fmt.Errorf("select user by sessiond id from tarantool storage: %w", err)
	}

	user := toGetUserBySessionID(resp)
	if user.ID == "" {
		return model.GetUserBySessionID{}, model.ErrNotFound
	}

	return user, nil
}
