package repository

import (
	"camparser"
	"log"
)

func (r *Repository) InsertCams(cams []camparser.Cam) (int, int) {
	counter := 0
	err_counter := 0
	for _, cam := range cams {
		var id int
		query := "INSERT INTO cams (camname, lon,lat,addr) values ($1, $2, $3,$4) RETURNING id"
		row := r.db.QueryRow(query, cam.Name, cam.Lon, cam.Lan, cam.Addr)
		if err := row.Scan(&id); err != nil {
			log.Println(err)
			err_counter++
			continue
		}
		counter++
		//TODO error handling
	}
	return counter, err_counter

}

func (r *Repository) GetAllCams() ([]camparser.Cam, error) {
	var cams []camparser.Cam
	query := "SELECT id,camname,lon,lat,addr FROM cams"
	err := r.db.Select(&cams, query)
	return cams, err
}

func (r *Repository) RemoveCam(id int) error {
	query := "DELETE FROM cams WHERE id = $1;"
	_, err := r.db.Exec(query, id)
	return err
}

func (r *Repository) RemoveAllCam() error {
	query := "TRUNCATE TABLE cams;"
	_, err := r.db.Exec(query)
	return err
}
