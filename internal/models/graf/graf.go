package models

import ()

// Узел графа
type Node struct {
 ID        string
 Shirota  float64
 Dolgota float64
 Danger  bool     // флаг, что этот узел отмечен как аварийный
}

// Ребро графа
type Rebro struct {
 From     string
 To       string
 Danger bool      // флаг, что этот участок дороги опасный
}

// Сам граф
type Graph struct {
 ID    string
 Nodes map[string]*Node   // узлы, по ID
 Rebra []Rebro            // список связей
}

// Аварийная точка
type CrashPoint struct {
 NodeID           string  // идентификатор узла
 LocalCrashIndex  float64 // локальный индекс аварийности
 
}