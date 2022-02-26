package shared

import "github.com/google/uuid"

func ConvertUUIDToString(bytes uuid.UUID) string {
	strs := []byte{}
	for _, b := range bytes {
		strs = append(strs, b)
	}

	uuid, _ := uuid.FromBytes(strs)

	return uuid.String()
}
