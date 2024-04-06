package timezone

import "os"

func SetTimeZoneAsiaTokyo() {
	os.Setenv("TZ", "Asia/Tokyo")
}
