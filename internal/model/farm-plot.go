package models

type FarmPlot struct {
	ID          int64
	Tag         string
	Coordinates string
	Area        float64
	WorkspaceId int64
}
