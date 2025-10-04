package routresponse

type WeatherResponse struct {
	Lat            float64    `json:"lat" binding:"required"`
	Lon            float64    `json:"lon" binding:"required"`
	Timezone       string     `json:"timezone" binding:"required"`
	TimezoneOffset int        `json:"timezone_offset" binding:"required"`
	Current        Current    `json:"current" binding:"required"`
	Minutely       []Minutely `json:"minutely" binding:"omitempty,dive"`
	Hourly         []Hourly   `json:"hourly" binding:"omitempty,dive"`
	Daily          []Daily    `json:"daily" binding:"omitempty,dive"`
	Alerts         []Alert    `json:"alerts" binding:"omitempty,dive"`
}

type Current struct {
	Dt         int64       `json:"dt" binding:"required"`
	Sunrise    *int64      `json:"sunrise,omitempty" binding:"omitempty"`
	Sunset     *int64      `json:"sunset,omitempty" binding:"omitempty"`
	Temp       float64     `json:"temp" binding:"required"`
	FeelsLike  float64     `json:"feels_like" binding:"required"`
	Pressure   int         `json:"pressure" binding:"required"`
	Humidity   int         `json:"humidity" binding:"required"`
	DewPoint   float64     `json:"dew_point" binding:"required"`
	UVI        float64     `json:"uvi" binding:"required"`
	Clouds     int         `json:"clouds" binding:"required"`
	Visibility int         `json:"visibility" binding:"required"`
	WindSpeed  float64     `json:"wind_speed" binding:"required"`
	WindDeg    int         `json:"wind_deg" binding:"required"`
	WindGust   *float64    `json:"wind_gust,omitempty" binding:"omitempty"`
	Rain       *Precip1h   `json:"rain,omitempty" binding:"omitempty"`
	Snow       *Precip1h   `json:"snow,omitempty" binding:"omitempty"`
	Weather    []Weather   `json:"weather" binding:"required,dive"`
}

type Precip1h struct {
	OneH float64 `json:"1h" binding:"required"`
}

type Minutely struct {
	Dt            int64   `json:"dt" binding:"required"`
	Precipitation float64 `json:"precipitation" binding:"required"`
}

type Hourly struct {
	Dt         int64       `json:"dt" binding:"required"`
	Temp       float64     `json:"temp" binding:"required"`
	FeelsLike  float64     `json:"feels_like" binding:"required"`
	Pressure   int         `json:"pressure" binding:"required"`
	Humidity   int         `json:"humidity" binding:"required"`
	DewPoint   float64     `json:"dew_point" binding:"required"`
	UVI        float64     `json:"uvi" binding:"required"`
	Clouds     int         `json:"clouds" binding:"required"`
	Visibility int         `json:"visibility" binding:"required"`
	WindSpeed  float64     `json:"wind_speed" binding:"required"`
	WindDeg    int         `json:"wind_deg" binding:"required"`
	WindGust   *float64    `json:"wind_gust,omitempty" binding:"omitempty"`
	Pop        float64     `json:"pop" binding:"required"`
	Rain       *Precip1h   `json:"rain,omitempty" binding:"omitempty"`
	Snow       *Precip1h   `json:"snow,omitempty" binding:"omitempty"`
	Weather    []Weather   `json:"weather" binding:"required,dive"`
}

type Daily struct {
	Dt        int64      `json:"dt" binding:"required"`
	Sunrise   *int64     `json:"sunrise,omitempty" binding:"omitempty"`
	Sunset    *int64     `json:"sunset,omitempty" binding:"omitempty"`
	Moonrise  *int64     `json:"moonrise,omitempty" binding:"omitempty"`
	Moonset   *int64     `json:"moonset,omitempty" binding:"omitempty"`
	MoonPhase *float64   `json:"moon_phase,omitempty" binding:"omitempty"`
	Summary   *string    `json:"summary,omitempty" binding:"omitempty"`
	Temp      TempRange  `json:"temp" binding:"required"`
	FeelsLike FeelsRange `json:"feels_like" binding:"required"`
	Pressure  int        `json:"pressure" binding:"required"`
	Humidity  int        `json:"humidity" binding:"required"`
	DewPoint  float64    `json:"dew_point" binding:"required"`
	WindSpeed float64    `json:"wind_speed" binding:"required"`
	WindDeg   int        `json:"wind_deg" binding:"required"`
	WindGust  *float64   `json:"wind_gust,omitempty" binding:"omitempty"`
	Clouds    int        `json:"clouds" binding:"required"`
	UVI       float64    `json:"uvi" binding:"required"`
	Pop       float64    `json:"pop" binding:"required"`
	Rain      *float64   `json:"rain,omitempty" binding:"omitempty"`
	Snow      *float64   `json:"snow,omitempty" binding:"omitempty"`
	Weather   []Weather  `json:"weather" binding:"required,dive"`
}

type TempRange struct {
	Day   float64 `json:"day" binding:"required"`
	Min   float64 `json:"min" binding:"required"`
	Max   float64 `json:"max" binding:"required"`
	Night float64 `json:"night" binding:"required"`
	Eve   float64 `json:"eve" binding:"required"`
	Morn  float64 `json:"morn" binding:"required"`
}

type FeelsRange struct {
	Day   float64 `json:"day" binding:"required"`
	Night float64 `json:"night" binding:"required"`
	Eve   float64 `json:"eve" binding:"required"`
	Morn  float64 `json:"morn" binding:"required"`
}

type Alert struct {
	SenderName  string   `json:"sender_name" binding:"required"`
	Event       string   `json:"event" binding:"required"`
	Start       int64    `json:"start" binding:"required"`
	End         int64    `json:"end" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Tags        []string `json:"tags" binding:"omitempty"`
}

type Weather struct {
	ID          int    `json:"id" binding:"required"`
	Main        string `json:"main" binding:"required"`
	Description string `json:"description" binding:"required"`
	Icon        string `json:"icon" binding:"required"`
}
