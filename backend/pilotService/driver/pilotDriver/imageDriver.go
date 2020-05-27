package pilotDriver

import (
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func WriteImage(oid int, data string) bool {

	db := openConnection();
	if db == nil {
		log.Error("Operator check failed")
		return false;
	}
	defer db.Close()

	cmd := fmt.Sprintf("BEGIN;SELECT pg_catalog.lo_open(%d, 393216);SELECT pg_catalog.lowrite(0,'\\x%s');SELECT pg_catalog.lo_close(0);COMMIT;", oid, data)

	db.Exec(cmd)
	db.Close()
	return true
}

func ReadImage(oid int, size int) []byte {

	var buf []string
	db := openConnection();
	if db == nil {
		log.Error("Operator check failed")
		return []byte(buf[0])
	}
	defer db.Close()

	tx := db.Begin()
	tx.Debug().Exec("SELECT pg_catalog.lo_open(?, 393216);", oid)
	tx.Debug().Raw("SELECT pg_catalog.loread(0, ?);", size).Pluck("loread", &buf)
	tx.Debug().Raw("SELECT pg_catalog.lo_close(0);")
	tx.Commit()

	db.Close()
	return []byte(buf[0])
}

func CreateOid() int {
	db := openConnection();
	if db == nil {
		log.Error("Operator check failed")
		return -1;
	}
	var lo [] int

	//24578
	db.Debug().Raw("SELECT pg_catalog.lo_creat(393216);").Pluck("lo_creat", &lo)

	log.Info(lo)
	db.Close()
	return lo[0];
}
