package supbox

import (
	"fmt"
)

func GetDatabaseDSN(filePath string, encryptionKey string) string {
	dsn := fmt.Sprintf("file:"+filePath+"?_key=%s", encryptionKey)
	return dsn
}
