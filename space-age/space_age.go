package space

/*
*
Mercury  汞	0.2408467
Venus  金星	0.61519726
Earth  地球	1.0
Mars  火星	1.8808158
Jupiter  木星	11.862615
Saturn  土星	29.447498
Uranus  天王星	84.016846
Neptune  海王星	164.79132
*/
var mp = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

type Planet string

func Age(seconds float64, planet Planet) float64 {
	if _, ok := mp[planet]; !ok {
		return -1
	}
	return seconds / mp[planet] / 31557600
}
