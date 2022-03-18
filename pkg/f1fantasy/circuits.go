package f1fantasy

import "time"

// CircuitInfo encodes detailed information for an F1 Circuit.
type CircuitInfo struct {
	Id             int       `json:"id"`
	FirstGrandPrix string    `json:"first_grand_prix"`
	TotalLaps      string    `json:"laps_total"`
	Length         string    `json:"length"`
	Distance       string    `json:"distance"`
	LapRecord      string    `json:"lap_record"`
	Practice1      string    `json:"practice_one"`
	Practice2      string    `json:"practice_two"`
	Practice3      string    `json:"practice_three"`
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
