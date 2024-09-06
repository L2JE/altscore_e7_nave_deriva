package common

import "golang.org/x/exp/rand"

func PickRandomKey(m map[string]string) string {
	mapSize := len(m)
	keyIxToPick := rand.Intn(mapSize)
	keyToPick := ""

	ix := 0
	for key := range m {
		if keyIxToPick == ix {
			keyToPick = key
			break
		}
		ix++
	}

	return keyToPick
}
