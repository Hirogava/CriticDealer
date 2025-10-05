package routresponse

type WeatherResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []WeatherElement `json:"weather"`
	Base     string           `json:"base"`
	Main     MainInfo         `json:"main"`
	Visibility int            `json:"visibility"`
	Wind     WindInfo         `json:"wind"`
	Rain     *RainInfo        `json:"rain,omitempty"`
	Clouds   CloudInfo        `json:"clouds"`
	Dt       int64            `json:"dt"`
	Sys      SysInfo          `json:"sys"`
	Timezone int              `json:"timezone"`
	ID       int              `json:"id"`
	Name     string           `json:"name"`
	Cod      int              `json:"cod"`
}

type WeatherElement struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type MainInfo struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
}

type WindInfo struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}

type RainInfo struct {
	OneH float64 `json:"1h"`
}

type CloudInfo struct {
	All int `json:"all"`
}

type SysInfo struct {
	Type    int    `json:"type"`
	ID      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}
