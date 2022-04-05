package f1fantasy

import "fmt"

// LeaderboardEntry is a particular entry in the league leaderboard.
type LeaderboardEntry struct {
	UserId            int     `json:"user_id"`
	UserGlobalId      string  `json:"user_global_id"`
	UserExternalId    int     `json:"user_external_id"`
	IsVerifiedEntrant bool    `json:"is_verified_entrant"`
	Country           string  `json:"bool"`
	Score             float64 `json:"score"`
	TeamName          string  `json:"team_name"`
	Rank              int     `json:"rank"`
	UserName          string  `json:"username"`
	Slot              int     `json:"slot"`
	UsedBoosterIds    []int   `json:"overall_used_booster_ids"` // todo: Figure out format!
}

// Leaderboard contains information for the leaderboard for a particular league.
type Leaderboard struct {
	NumEntrants            int                `json:"entrants_count"`
	LeagueName             string             `json:"league_name"`
	MaxPoints              float64            `json:"max_points"`
	MinPoints              float64            `json:"min_points"`
	LeagueMappingType      *string            `json:"league_mapping_type,omitempty"`       // todo: figure out format
	LeagueMappingTypeParam *string            `json:"league_mapping_type_param,omitempty"` // todo: figure out format
	Entries                []LeaderboardEntry `json:"leaderboard_entrants"`
}

// LeagueLeaderboard is the parent container for all of the league leaderboard information.
type LeagueLeaderboard struct {
	Leaderboard Leaderboard `json:"leaderboard"`
}

// GetLeagueLeaderboard is an Authenticated API that retrieves a particular leagues leaderboard.
func (api *AuthenticatedApi) GetLeagueLeaderboard(leagueId int) (*LeagueLeaderboard, error) {
	const PAGE = "/leaderboards/leagues?game_period_id=&league_id=%d"

	var leaderboard LeagueLeaderboard
	err := api.getAndDecode(fmt.Sprintf(PAGE, leagueId), &leaderboard)
	if err != nil {
		return nil, err
	}
	return &leaderboard, nil
}
