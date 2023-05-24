package helper

import (
	"strings"

	"github.com/alemelomeza/cautious-journey/internal/domain"
)

func ParentsToIDs(parents ...domain.Parent) string {
	var ids []string
	for _, p := range parents {
		ids = append(ids, p.Name)
	}
	return strings.Join(ids, ",")
}