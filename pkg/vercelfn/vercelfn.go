package vercelfn

import (
	"app/pkg/entdb"
	"app/pkg/vercelapi"
	"app/pkg/vercelutil"
	"context"
)

// Sycn user and team to database from auth.json file
func SyncAuthJson() error {
	authJsonFile, err := vercelutil.AuthJsonFile()
	if err != nil {
		return err
	}

	authConfig, err := vercelutil.ParseAuthFile(authJsonFile)
	if err != nil {
		return err
	}

	user, err := vercelapi.GetUser(authConfig.Token)
	if err != nil {
		return err
	}

	ctx := context.Background()

	userID, err := entdb.Client().User.Create().
		SetID(user.ID).
		SetEmail(user.Email).
		SetName(user.Name).
		SetUsername(user.Username).
		SetToken(authConfig.Token).
		OnConflict().
		UpdateNewValues().
		ID(ctx)
	if err != nil {
		return err
	}

	teams, err := vercelapi.GetTeams(authConfig.Token)
	if err != nil {
		return err
	}

	for _, team := range teams {
		_, err := entdb.Client().Team.Create().
			SetID(team.ID).
			SetUserID(userID).
			SetName(team.Name).
			SetSlug(team.Slug).
			OnConflict().
			UpdateNewValues().
			ID(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
