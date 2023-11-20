package internal

type struct PlanificadorViajes {
	camioneros		[]Camionero
	viajes			[]Viaje
	asignaciones 	map[camioneros]viajes
}
