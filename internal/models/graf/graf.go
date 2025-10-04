package graf

// Узел графа
type Node struct {
	ID      string  `json:"id"`
	Shirota float64 `json:"shirota"`
	Dolgota float64 `json:"dolgota"`
	Danger  bool    `json:"danger"` // флаг, что этот узел отмечен как аварийный
}

// Ребро графа
type Rebro struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Danger bool   `json:"danger"` // флаг, что этот участок дороги опасный
}

// Сам граф
type Graph struct {
	ID    string           `json:"id"`
	Nodes map[string]*Node `json:"nodes"` // узлы, по ID
	Rebra []Rebro          `json:"rebra"` // список связей
    RiskScore float64          `json:"risk_score"`
}

// Аварийная точка
type CrashPoint struct {
	NodeID          string  `json:"node_id"`           // идентификатор узла
	LocalCrashIndex float64 `json:"local_crash_index"` // локальный индекс аварийности
}

