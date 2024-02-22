package twelve_data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tecolotedev/trading-ml-backend/config"
	"github.com/tecolotedev/trading-ml-backend/utils"
)

const pack = "twelve_data"

const twelveDataUrl = "https://api.twelvedata.com/"

type TimeSeriesResponse struct {
	Values []utils.ValueResponse
}

func FetchTimeSeries(outputSize int, symbol, interval, tz, startDate, endDate string) (values []utils.ValueParsed, err error) {
	var key = config.EnvVars.TWELVE_DATA_KEY

	// build url
	url := fmt.Sprintf(
		"%stime_series?symbol=%s&interval=%s&outputsize=%d&timezone=%s&start_date=%s&end_date=%s&apikey=%s",
		twelveDataUrl,
		utils.Symbols[symbol],
		utils.Intervals[interval],
		outputSize,
		tz,
		startDate,
		endDate,
		key,
	)

	res, err := http.Get(url)
	if err != nil {
		utils.Log.ErrorLog(err, pack)
		err = fmt.Errorf("error fetching financial data, please try it later")
		return
	}

	body, err := io.ReadAll(res.Body)
	fmt.Println(string(body))
	if err != nil {
		utils.Log.ErrorLog(err, pack)
		err = fmt.Errorf("error fetching financial data, please try it later")
		return
	}

	var timeSeriesRes = TimeSeriesResponse{}
	err = json.Unmarshal(body, &timeSeriesRes)

	if err != nil {
		utils.Log.ErrorLog(err, pack)
		err = fmt.Errorf("error fetching financial data, please try it later")
		return
	}

	// parse values from string to float
	values = utils.ParseValuesToFloat(timeSeriesRes.Values)

	return
}
