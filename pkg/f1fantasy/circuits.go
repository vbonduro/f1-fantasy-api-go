package f1fantasy

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/tidwall/gjson"
)

// CircuitInfo encodes detailed information for an F1 Circuit.
// todo change practice and qualifying to time.time!
type CircuitInfo struct {
	StartDay       time.Time
	Id             int       `json:"id"`
	FirstGrandPrix string    `json:"first_grand_prix"`
	TotalLaps      string    `json:"laps_total"`
	Length         string    `json:"length"`
	Distance       string    `json:"distance"`
	LapRecord      string    `json:"lap_record"`
	Practice1      string    `json:"pratice_one"`
	Practice2      string    `json:"pratice_two"`
	Practice3      string    `json:"pratice_three"`
	Qualifying     string    `json:"qualifying"`
	Race           string    `json:"race"`
	Created        time.Time `json:"created_at"`
	Updated        time.Time `json:"updated_at"`
	Name           string    `json:"name"`
	ShortName      string    `json:"short_name"`
	CountryIso     string    `json:"country_iso"`
	CircuitImage   Image     `json:"circuit_image"`
	GmtOffset      string    `json:"gmt_offset"`
}

// Circuit encodes information for an F1 Circuit.
type Circuit struct {
	PeriodId int         `json:"game_period_id"`
	Name     string      `json:"game_period_name"`
	Info     CircuitInfo `json:"circuit"`
}

// GetCircuits is a public API that retrieves the season's circuit information.
func (api *Api) GetCircuits() ([]Circuit, error) {
	const PAGE = "/circuits"

	var circuits []Circuit
	err := api.getAndDecode(PAGE, &circuits)
	if err != nil {
		return nil, err
	}
	return circuits, nil
}

// CurrentCircuit is an authenticated API that retrieves the upcoming/current race information.
func (api *AuthenticatedApi) CurrentCircuit() (*CircuitInfo, error) {
	bytes, err := api.get("")
	if err != nil {
		return nil, err
	}
	var race CircuitInfo
	currentCircuit := gjson.Get(string(bytes),
		"partner_game.current_partner_season.current_game_period.circuit").String()
	err = json.Unmarshal([]byte(currentCircuit), &race)
	if err != nil {
		return nil, err
	}

	startQuery := fmt.Sprintf("partner_game.current_partner_season.game_periods.%d.starts_at", race.Id-1)
	raceStart := gjson.Get(string(bytes), startQuery).String()

	const LAYOUT = "2006-01-02T15:04:05.000Z"
	race.StartDay, err = time.Parse(LAYOUT, raceStart)
	if err != nil {
		return nil, err
	}

	return &race, nil
}
