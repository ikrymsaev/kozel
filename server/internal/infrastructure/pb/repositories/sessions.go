package pb_repositories

import (
	"fmt"
	"log"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
)

type SessionsRepository struct {
	pb *pocketbase.PocketBase
}

func NewSessionsRepository(pb *pocketbase.PocketBase) *SessionsRepository {
	return &SessionsRepository{
		pb: pb,
	}
}

// AddSessionLog - adds new auth session log.
func (m *SessionsRepository) AddSessionLog(
	user *models.Record,
	IpAddr string,
) error {
	collection, err := m.pb.Dao().FindCollectionByNameOrId("users_session")
	fmt.Printf("collection: %v\n", collection)
	if err != nil {
		fmt.Printf("collection error: %v\n", err)
		return err
	}

	params := dbx.Params{"user": user.GetId(), "IP": IpAddr}
	query := dbx.NewExp("user = {:user} AND IP = {:IP}", params)
	fmt.Printf("query: %v\n", query)
	records, err := m.pb.Dao().FindRecordsByExpr(collection.GetId(), query)
	if err != nil {
		log.Println("Failed to get user!:", err)
		return nil
	}

	// If allready exists
	if len(records) > 0 {
		log.Println("Session already exists!")
		records[0].Set("entry_count", records[0].GetInt("entry_count")+1)
		return m.pb.Dao().SaveRecord(records[0])
	}
	record := models.NewRecord(collection)
	record.Set("IP", IpAddr)
	record.Set("user", user.GetId())
	record.Set("entry_count", 1)
	return m.pb.Dao().SaveRecord(record)
}
