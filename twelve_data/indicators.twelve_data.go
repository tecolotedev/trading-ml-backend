package twelve_data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tecolotedev/trading-ml-backend/config"
	"github.com/tecolotedev/trading-ml-backend/utils"
)

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

type MACDResponse struct {
	Values []utils.MACDValueResponse
}

func FetchMACD(outputSize, fastPeriod, signalPeriod, slowPeriod int, symbol, interval, tz, startDate, endDate, seriesType string) (values []utils.MACDValueParsed, err error) {
	var key = config.EnvVars.TWELVE_DATA_KEY

	url := fmt.Sprintf(
		"%smacd?symbol=%s&interval=%s&outputsize=%d&timezone=%s&start_date=%s&end_date=%s&series_type=%s&fast_period=%d&signal_period=%d&slow_period=%d&apikey=%s",
		twelveDataUrl,
		utils.Symbols[symbol],
		utils.Intervals[interval],
		outputSize,
		tz,
		startDate,
		endDate,
		seriesType,
		fastPeriod,
		signalPeriod,
		slowPeriod,
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

	var macdResponse = MACDResponse{}
	err = json.Unmarshal(body, &macdResponse)

	if err != nil {
		utils.Log.ErrorLog(err, pack)
		err = fmt.Errorf("error fetching financial data, please try it later")
		return
	}

	// parse values from string to float
	values = utils.ParseMACDValues(macdResponse.Values)

	return

}

type RSIResponse struct {
	Values []utils.RSIValueResponse
}

func FetchRSI(outputSize, timePeriod int, symbol, interval, tz, startDate, endDate string) (values []utils.RSIValueParsed, err error) {
	var key = config.EnvVars.TWELVE_DATA_KEY

	url := fmt.Sprintf(
		"%srsi?symbol=%s&interval=%s&outputsize=%d&timezone=%s&start_date=%s&end_date=%s&time_period=%d&apikey=%s",
		twelveDataUrl,
		utils.Symbols[symbol],
		utils.Intervals[interval],
		outputSize,
		tz,
		startDate,
		endDate,
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

	var macdResponse = RSIResponse{}
	err = json.Unmarshal(body, &macdResponse)

	if err != nil {
		utils.Log.ErrorLog(err, pack)
		err = fmt.Errorf("error fetching financial data, please try it later")
		return
	}

	// parse values from string to float
	values = utils.ParseRSIValues(macdResponse.Values)

	return

}
