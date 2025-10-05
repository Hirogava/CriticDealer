package routresponse

// Query представляет основной запрос для построения маршрута
type QueryResponse struct {
	Points          []QueryPointResponse   `json:"points" binding:"required,min=2"`
	Transport       string         `json:"transport,omitempty" binding:"omitempty,oneof=driving walking taxi bicycle scooter emergency truck motorcycle"`
	Output          string         `json:"output,omitempty" binding:"omitempty,oneof=summary detailed"`
	Locale          string         `json:"locale,omitempty" binding:"omitempty,oneof=en ru uk es it cs ar az uz kk hy ka tg ky"`
	RouteMode       string         `json:"route_mode,omitempty" binding:"omitempty,oneof=fastest shortest"`
	TrafficMode     string         `json:"traffic_mode,omitempty" binding:"omitempty,oneof=jam statistics"`
	Alternative     *int           `json:"alternative,omitempty" binding:"omitempty,min=0,max=5"`
	NeedAltitudes   *bool          `json:"need_altitudes,omitempty"`
	AllowLockedRoads *bool         `json:"allow_locked_roads,omitempty"`
	UTCOffset       *int64         `json:"utc,omitempty" binding:"omitempty"`
	Filters  	 struct {
		Filter *string `json:"filter" binding:"omitempty"`
		Filters []string `json:"filters" binding:"omitempty"`
	} `json:"filters" binding:"omitempty"`
	Exclude         []ExcludeObject `json:"exclude,omitempty" binding:"omitempty,max=25"`
	Params          *ParamsResponse        `json:"params,omitempty"`
}

// QueryPoint представляет точку маршрута
type QueryPointResponse struct {
	Type     string   `json:"type" binding:"required,oneof=stop walking pref"`
	Lon      float64  `json:"lon" binding:"required,min=-180,max=180"`
	Lat      float64  `json:"lat" binding:"required,min=-90,max=90"`
	ObjectID *string  `json:"object_id,omitempty"`
	Start    *bool    `json:"start,omitempty"`
	Azimuth  *int     `json:"azimuth,omitempty" binding:"omitempty,min=0,max=360"`
	ZLevel   *int     `json:"zlevel,omitempty"`
	FloorID  *int     `json:"floor_id,omitempty"`
	EDirection *int   `json:"e_direction,omitempty"`
}

// ExcludeObject представляет исключаемую область
type ExcludeObject struct {
	Type     string    `json:"type" binding:"required,oneof=point polyLine polygon"`
	Points   []WGS84Point `json:"points" binding:"required,min=1"`
	Severity string    `json:"severity" binding:"required,oneof=soft hard"`
	Extent   *int      `json:"extent,omitempty" binding:"omitempty,min=1"`
}

// WGS84Point представляет точку в формате WGS84
type WGS84Point struct {
	Lon float64 `json:"lon" binding:"required,min=-180,max=180"`
	Lat float64 `json:"lat" binding:"required,min=-90,max=90"`
}

// Params представляет дополнительные параметры
type ParamsResponse struct {
	Pedestrian *PedestrianParams `json:"pedestrian,omitempty"`
	Truck      *TruckParams      `json:"truck,omitempty"`
}

// PedestrianParams представляет параметры для пешеходного маршрута
type PedestrianParams struct {
	UseIndoor      *bool `json:"use_indoor,omitempty"`
	UseInstructions *bool `json:"use_instructions,omitempty"`
}

// TruckParams представляет параметры для грузового транспорта
type TruckParams struct {
	MaxPermMass    *float64 `json:"max_perm_mass,omitempty" binding:"omitempty,min=0"`
	Mass           *float64 `json:"mass,omitempty" binding:"omitempty,min=0"`
	AxleLoad       *float64 `json:"axle_load,omitempty" binding:"omitempty,min=0"`
	Height         *float64 `json:"height,omitempty" binding:"omitempty,min=0"`
	Width          *float64 `json:"width,omitempty" binding:"omitempty,min=0"`
	Length         *float64 `json:"length,omitempty" binding:"omitempty,min=0"`
	DangerousCargo *bool    `json:"dangerous_cargo,omitempty"`
	ExplosiveCargo *bool    `json:"explosive_cargo,omitempty"`
	PassZoneIds    []int    `json:"pass_zone_pass_ids,omitempty"`
}