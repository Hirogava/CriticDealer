package routresponse

type RouteResponse struct {
	Status  string      `json:"status" binding:"required"`
	Query   Query        `json:"query" binding:"required"`
	Type    string        `json:"type" binding:"required"`
	Message *string        `json:"message" binding:"omitempty"`
	Weather []string       `json:"weather" binding:"omitempty"`
	Result  []ResultItem `json:"result"`
}

type ResultItem struct {
	ID                      string           `json:"id"`
	Type                    string           `json:"type"`
	Algorithm               string           `json:"algorithm"`
	TotalDistance           int              `json:"total_distance"`
	TotalDuration           int              `json:"total_duration"`
	UITotalDistance         UIDistance       `json:"ui_total_distance"`
	UITotalDuration         string           `json:"ui_total_duration"`
	BeginPedestrianPath     GeometryWrapper  `json:"begin_pedestrian_path"`
	EndPedestrianPath       GeometryWrapper  `json:"end_pedestrian_path"`
	Maneuvers               []Maneuver       `json:"maneuvers"`
	Waypoints               []Waypoint       `json:"waypoints"`
	RouteID                 *string          `json:"route_id" binding:"omitempty"`
	Reliability             *float64         `json:"reliability" binding:"omitempty"`
	RequestedFilters        []string         `json:"requested_filters" binding:"omitempty"`
	ResultFilters           []string         `json:"result_filters" binding:"omitempty"`
	FilterRoadTypes         []string         `json:"filter_road_types" binding:"omitempty"`
	AltitudesInfo           *AltitudesInfo   `json:"altitudes_info" binding:"omitempty"`
	RoutePoints             []RoutePoint     `json:"route_points" binding:"omitempty"`
	AreLockedRoadsIgnored   *bool            `json:"are_locked_roads_ignored" binding:"omitempty"`
	RequestedHardFilters    []string         `json:"requested_hard_filters" binding:"omitempty"`
	AreTruckPassZonesIgnored *bool           `json:"are_truck_pass_zones_ignored" binding:"omitempty"`
	VisitedPassZoneIDs      []int            `json:"visited_pass_zone_ids" binding:"omitempty"`
	Features                *Features        `json:"features" binding:"omitempty"`
}

type Query struct {
	Exclude       []Exclude `json:"exclude"`
	NeedAltitudes bool      `json:"need_altitudes"`
	Points        []QueryPoint `json:"points" binding:"required"`
	RouteMode     string    `json:"route_mode"`
	TrafficMode   string    `json:"traffic_mode"`
	Transport     string    `json:"transport"`
	Output        *string    `json:"output" binding:"omitempty"`
	Utc           *int       `json:"utc" binding:"omitempty"`
	Filters  	 struct {
		Filter *string `json:"filter" binding:"omitempty"`
		Filters []string `json:"filters" binding:"omitempty"`
	} `json:"filters" binding:"omitempty"`
	AllowLockedRoads *bool `json:"allow_locked_roads" binding:"omitempty"`
	Locale *string `json:"locale" binding:"omitempty"`
	Params *Params `json:"params" binding:"omitempty"`
	Alternative *int `json:"alternative" binding:"omitempty"`
}

type Params struct {
	Pedestrian *Pedestrian `json:"pedestrian" binding:"omitempty"`
	Truck *Truck `json:"truck" binding:"omitempty"`
}

type Pedestrian struct {
	UseIndoor bool `json:"use_indoor"`
	UseInstructions bool `json:"use_instructions"`
}

type Truck struct {
	MaxPermMass *int `json:"max_perm_mass" binding:"omitempty"`
	Mass *int `json:"mass" binding:"omitempty"`
	AxleLoad *int `json:"axle_load" binding:"omitempty"`
	Height *int `json:"height" binding:"omitempty"`
	Width *int `json:"weight" binding:"omitempty"`
	Length *int `json:"length" binding:"omitempty"`
	DangerousCargo bool `json:"dangerous_cargo"`
	ExplosiveCargo bool `json:"explosive_cargo"`
	PassZonePassIds []int `json:"pass_zone_pass_ids" binding:"omitempty"`
}

type Exclude struct {
	Extent   *int        `json:"extent" binding:"omitempty"`
	Points   []LatLon   `json:"points" binding:"required"`
	Severity string     `json:"severity" binding:"required"`
	Type     string     `json:"type" binding:"required"`
}

type QueryPoint struct {
	Lat      float64  `json:"lat"`
	Lon      float64  `json:"lon"`
	Start    bool     `json:"start"`
	ObjectID *string  `json:"object_id" binding:"omitempty"`
	Type     string   `json:"type"`
}

type LatLon struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type UIDistance struct {
	Unit  string  `json:"unit"`
	Value string `json:"value"`
}

type Waypoint struct {
	OriginalPoint  LatLon `json:"original_point"`
	ProjectedPoint LatLon `json:"projected_point"`
	Transit        bool   `json:"transit"`
}

type RoutePoint struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type Features struct {
	Truck string `json:"truck"`
}


type AltitudesInfo struct {
	ElevationGain  int `json:"elevation_gain"`
	ElevationLoss  int `json:"elevation_loss"`
	MaxAltitude    int `json:"max_altitude"`
	MaxRoadAngle   int `json:"max_road_angle"`
	MinAltitude    int `json:"min_altitude"`
}

type GeometryWrapper struct {
	Geometry Geometry `json:"geometry"`
}

type Geometry struct {
	Selection string `json:"selection"`
}

type Maneuver struct {
	Comment               string        `json:"comment"`
	Icon                  string        `json:"icon"`
	ID                    string        `json:"id"`
	OutcomingPath         OutcomingPath `json:"outcoming_path"`
	OutcomingPathComment  string        `json:"outcoming_path_comment"`
	Type                  string        `json:"type"`
	TurnAngle             *int          `json:"turn_angle" binding:"omitempty"`
	TurnDirection         *string       `json:"turn_direction" binding:"omitempty"`
	RingroadExitNumber    *int          `json:"ringroad_exit_number" binding:"omitempty"`
	Critical              *bool         `json:"critical" binding:"omitempty"`
	CriticalProbability   *float32      `json:"critical_probability" binding:"omitempty"`
}

type OutcomingPath struct {
	Distance int                `json:"distance"`
	Duration int                `json:"duration"`
	Geometry []GeometryElement  `json:"geometry"`
	Names    []string           `json:"names"`
}

type GeometryElement struct {
	Angles    string `json:"angles"`
	Color     string `json:"color"`
	Length    int    `json:"length"`
	Selection string `json:"selection"`
	Style     string `json:"style"`
}
