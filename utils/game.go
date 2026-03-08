package utils

import (
	"gamecheck-backend/db"
	"math"
	"strconv"

	"github.com/dimuska139/rawg-sdk-go/v3"
	"github.com/jackc/pgx/v5/pgtype"
)

func convertFloat32ToPgNumeric(f float64) (pgtype.Numeric, error) {
	if math.IsNaN(f) {
		return pgtype.Numeric{NaN: true, Valid: true}, nil
	}

	if math.IsInf(f, 1) {
		return pgtype.Numeric{InfinityModifier: pgtype.Infinity, Valid: true}, nil
	}

	if math.IsInf(f, -1) {
		return pgtype.Numeric{InfinityModifier: pgtype.NegativeInfinity, Valid: true}, nil
	}

	// Convert float -> decimal string (important)
	s := strconv.FormatFloat(f, 'f', -1, 64)

	var n pgtype.Numeric
	err := n.Scan(s)
	return n, err
}

func ConvertRAWGGameTypeToRecordGameType(rawgGame *rawg.Game) (*db.UpsertGamesParams, error) {
	ratingParam, err := convertFloat32ToPgNumeric(float64(rawgGame.Rating))

	if err != nil {
		return nil, err
	}

	params := &db.UpsertGamesParams{
		ID:   int32(rawgGame.ID),
		Name: rawgGame.Name,
		Released: pgtype.Date{
			Time:             rawgGame.Released.Time,
			InfinityModifier: 0,
			Valid:            true,
		},
		BackgroundImage: pgtype.Text{
			String: rawgGame.ImageBackground,
			Valid:  true,
		},
		Rating: ratingParam,
	}

	return params, nil
}
