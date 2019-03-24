package handlers

import (
	"strconv"
)

func parseSlugOrId(slugOrId string) (slug *string, id *uint64) {
	_id, err := strconv.ParseUint(slugOrId, 10, 64)
	if err == nil {
		return nil, &_id
	} else {
		return &slugOrId, nil
	}
}
