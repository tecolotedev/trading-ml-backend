package twelve_data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tecolotedev/trading-ml-backend/config"
	"github.com/tecolotedev/trading-ml-backend/utils"
)

// func FetchIndicator(outputsize int, symbol, interval, tz ,string) {

// }

const MA = "ma"

type MAResponse struct {
	Values []utils.MAValueResponse
}

func FetchMA(outputSize, timePeriod int, symbol, interval, tz, startDate, endDate, maType, seriesType string) (values []utils.MAValueParsed, err error) {
	var key = config.EnvVars.TWELVE_DATA_KEY

	url := fmt.Sprintf(
		"%sma?symbol=%s&interval=%s&outputsize=%d&timezone=%s&start_date=%s&end_date=%s&ma_type=%s&series_type=%s&time_period=%d&apikey=%s",
		twelveDataUrl,
		utils.Symbols[symbol],
		utils.Intervals[interval],
		outputSize,
		tz,
		startDate,
		endDate,
		maType,
		seriesType,
		timePeriod,
		key,
	)

	res, err := http.Get(url)
	if err != nil {
		utils.Log.ErrorLog(err, pack)
		err = fmt.Errorf("error fetching financial data, please try it later")
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		utils.Log.ErrorLog(err, pack)
		err = fmt.Errorf("error fetching financial data, please try it later")
		return
	}

	var maResponse = MAResponse{}
	err = json.Unmarshal(body, &maResponse)

	if err != nil {
		utils.Log.ErrorLog(err, pack)
		err = fmt.Errorf("error fetching financial data, please try it later")
		return
	}

	// parse values from string to float
	values = utils.ParseMAValues(maResponse.Values)

	return

}
