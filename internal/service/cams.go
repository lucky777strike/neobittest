package service

import (
	"camparser"
	"fmt"
	"log"
	"strconv"
)

func (s *Service) Parse(q string) error {
	go func() {
		cams := make([]camparser.Cam, 0, 10)
		taskid, err := s.CreateTask(camparser.Task{Query: q,
			Status: "in process",
			Count:  0})
		if err != nil {
			return
		}
		r, err := s.shodan.QuerySearch(q)
		if err != nil {
			err = s.UpdateTask(taskid, 0, 0, false)
			if err != nil {
				return
			}
			log.Println(err)
		}
		for _, v := range r.Matches {
			addr := fmt.Sprintf("http://%s:%d", v.Http.Host, v.Port)
			cam := camparser.Cam{Name: "webcam", Lon: v.Location.Lon, Lan: v.Location.Lan, Addr: addr}
			cams = append(cams, cam)
		}
		counter, err_counter := s.repo.InsertCams(cams)
		err = s.UpdateTask(taskid, counter, err_counter, false)
		if err != nil {
			return
		}

		fmt.Println(len(r.Matches))
	}()
	return nil
}

func (s *Service) GetAllCams() ([]camparser.Cam, error) {
	return s.repo.GetAllCams()
}

func (s *Service) RemoveCam(id string) error {
	idi, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return s.repo.RemoveCam(idi)
}

func (s *Service) RemoveAllCam() error {
	return s.repo.RemoveAllCam()
}
