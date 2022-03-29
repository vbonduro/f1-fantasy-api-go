package f1fantasy

type Headshot struct {
	Profile    string `json:"profile"`
	PitchView  string `json:"pitch_view"`
	PlayerList string `json:"player_list"`
}

type Image struct {
	Url *string `json:"url,omitempty"`
}

type DriverData struct {
	Wins                int    `json:"wins"`
	Podiums             int    `json:"podiums"`
	Poles               int    `json:"poles"`
	FastestLaps         int    `json:"fastest_laps"`
	TotalGrandPrix      int    `json:"grands_prix_entered"`
	Titles              int    `json:"titles"`
	ChampionshipPoints  int    `json:"championship_points"`
	BestFinish          int    `json:"best_finish"`
	BestFinishCount     int    `json:"best_finish_count"`
	BestGrid            int    `json:"best_grid"`
	BestGridCount       int    `json:"best_grid_count"`
	HighestRaceFinished string `json:"highest_race_finished"`
	PlaceOfBirth        string `json:"place_of_birth"`
}

type ConstructorData struct {
	BestFinish          int     `json:"best_finish"`
	BestFinishCount     int     `json:"best_finish_count"`
	BestGrid            int     `json:"best_grid"`
	BestGridCount       int     `json:"best_grid_count"`
	Titles              int     `json:"titles"`
	ChampionshipPoints  float64 `json:"championship_points"`
	FirstSeason         string  `json:"first_season"`
	Poles               int     `json:"poles"`
	FastestLaps         int     `json:"fastest_laps"`
	Country             string  `json:"country"`
	HighestRaceFinished string  `json:"highest_race_finished"`
}

type Player struct {
	Id                          int              `json:"id"`
	FirstName                   string           `json:"first_name"`
	LastName                    string           `json:"last_name"`
	TeamName                    string           `json:"team_name"`
	Position                    string           `json:"position"`
	PositionId                  int              `json:"position_id"`
	PositionAbbreviation        string           `json:"position_abbreviation"`
	Price                       float64          `json:"price"`
	PriceChangeInfo             *string          `json:"current_price_change_info,omitempty"` //todo: Figure out format.
	Status                      *string          `json:"status,omitempty"`                    //todo: Figure out format.
	Injured                     bool             `json:"injured"`
	InjuryType                  *string          `json:"injury_type,omitempty"` //todo: Figure out format.
	Banned                      bool             `json:"banned"`
	BanType                     *string          `json:"ban_type,omitempty"` //todo: Figure out format.
	ChanceOfPlaying             float64          `json:"chance_of_playing"`
	TeamAbbreviation            string           `json:"team_abbreviation"`
	WeeklyPriceChange           float64          `json:"weekly_price_change"`
	WeeklyPriceChangePercentage int64            `json:"weekly_price_change_percentage"`
	TeamId                      int              `json:"team_id"`
	KnownName                   *string          `json:"known_name,omitempty"` //todo: Figure out format.
	HeadshotImages              Headshot         `json:"headshot"`
	JerseyImage                 Image            `json:"jersey_image"`
	ProfileImage                Image            `json:"profile_image"`
	MiscImage                   Image            `json:"misc_image"`
	Score                       int64            `json:"score"`
	HumanizeStatus              *string          `json:"humanize_status,omitempty"` //todo: Figure out format.
	ShirtNumber                 *int64           `json:"shirt_number,omitempty"`
	Country                     *string          `json:"country,omitempty"`
	CountryIso                  *string          `json:"country_iso,omitempty"`
	IsConstructor               bool             `json:"is_constructor"`
	SeasonScore                 int64            `json:"season_score"`
	DriverStats                 *DriverData      `json:"driver_data,omitempty"`
	ConstructorStats            *ConstructorData `json:"constructor_data,omitempty"`
	Born                        *string          `json:"born_at,omitempty"`
	SeasonPrices                *string          `json:"season_prices,omitempty"`
	NumFixturesInGameweek       int              `json:"num_fixtures_in_gameweek"`
	DeletedInFeed               bool             `json:"deleted_in_feed"`
	HasFixture                  bool             `json:"has_fixture"`
	DisplayName                 string           `json:"display_name"`
	ExternalId                  string           `json:"external_id"`

	//StreakEventsProgress        []string         `json:"streak_events_progress"` //todo: Figure out format.
}

type Meta struct {
	Total int `json:"total"`
}

type Players struct {
	PlayerList []Player `json:"players"`
	MetaInfo   Meta     `json:"meta"`
}

// GetPlayers is a public API that retrieves player list and statistics.
func (api *Api) GetPlayers() (*Players, error) {
	const PAGE = "/players"

	var players Players
	err := api.getAndDecode(PAGE, &players)
	if err != nil {
		return nil, err
	}
	return &players, nil
}
