package server

import (
	"encoding/json"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
	"time"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

type calendarRecord struct {
	UserId string
	date   time.Time
}

func getRecordsFromRows(rows pgx.Rows) []calendarRecord {
	calendarRecords := make([]calendarRecord, 0)
	for rows.Next() {
		record := calendarRecord{}
		err := rows.Scan(&record)
		if err != nil {
			log.Fatal(err)
		}
		calendarRecords = append(calendarRecords, record)
	}
	return calendarRecords
}

func (s *Server) createEvent(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		res.Header().Set("Content-Type", "application/json")

		userId := req.FormValue("user_id")
		date := req.FormValue("date")
		if userId == "" || date == "" {
			res.WriteHeader(400)
			data := map[string]string{"error": "Incorrect input"}
			err := json.NewEncoder(res).Encode(data)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		query := "INSERT INTO calendar(user_id, date) VALUES ('$1', '$2');"
		_, err := s.db.Query(s.ctx, query, userId, date)
		if err != nil {
			log.Fatal(err)
		}
		data := map[string]string{"result": "Success"}
		err = json.NewEncoder(res).Encode(data)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	res.WriteHeader(503)
}

func (s *Server) eventsForWeek(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		res.Header().Set("Content-Type", "application/json")

		params := req.URL.Query()
		dateStr := params.Get("date")
		if dateStr == "" {
			res.WriteHeader(400)
			data := map[string]string{"error": "Incorrect input"}
			err := json.NewEncoder(res).Encode(data)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		date, err := time.Parse("2019-09-09", dateStr)
		if err != nil {
			data := map[string]string{"error": "Incorrect date"}
			err := json.NewEncoder(res).Encode(data)
			if err != nil {
				log.Fatal(err)
			}
		}
		query := "SELECT user_id, date FROM calendar WHERE date >= '$1' AND date <= '$2'"
		rows, err := s.db.Query(s.ctx, query,
			date, date.AddDate(0, 0, 7))
		if err != nil {
			log.Fatal(err)
		}
		calendarRecords := getRecordsFromRows(rows)
		data := map[string][]calendarRecord{"result": calendarRecords}
		err = json.NewEncoder(res).Encode(data)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	res.WriteHeader(503)
}

func (s *Server) eventsForMonth(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		res.Header().Set("Content-Type", "application/json")
		params := req.URL.Query()
		dateStr := params.Get("date")
		if dateStr == "" {
			res.WriteHeader(400)
			data := map[string]string{"error": "Incorrect input"}
			err := json.NewEncoder(res).Encode(data)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		date, err := time.Parse("2019-09-09", dateStr)
		if err != nil {
			data := map[string]string{"error": "Incorrect date"}
			err := json.NewEncoder(res).Encode(data)
			if err != nil {
				log.Fatal(err)
			}
		}
		query := "SELECT user_id, date FROM calendar WHERE date >= '$1' AND date <= '$2'"
		rows, err := s.db.Query(s.ctx, query,
			date, date.AddDate(0, 0, 31))
		if err != nil {
			log.Fatal(err)
		}
		calendarRecords := getRecordsFromRows(rows)
		data := map[string][]calendarRecord{"result": calendarRecords}
		err = json.NewEncoder(res).Encode(data)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	res.WriteHeader(503)
}

func (s *Server) eventsForDay(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		res.Header().Set("Content-Type", "application/json")
		params := req.URL.Query()
		dateStr := params.Get("date")
		if dateStr == "" {
			res.WriteHeader(400)
			data := map[string]string{"error": "Incorrect input"}
			err := json.NewEncoder(res).Encode(data)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		date, err := time.Parse("2019-09-09", dateStr)
		if err != nil {
			data := map[string]string{"error": "Incorrect date"}
			err := json.NewEncoder(res).Encode(data)
			if err != nil {
				log.Fatal(err)
			}
		}
		query := "SELECT user_id, date FROM calendar WHERE date >= '$1' AND date <= '$2'"
		rows, err := s.db.Query(s.ctx, query,
			date, date.AddDate(0, 0, 1))
		if err != nil {
			log.Fatal(err)
		}
		calendarRecords := getRecordsFromRows(rows)
		data := map[string][]calendarRecord{"result": calendarRecords}
		err = json.NewEncoder(res).Encode(data)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	res.WriteHeader(503)
}

func (s *Server) updateEvent(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		res.Header().Set("Content-Type", "application/json")

		userId := req.FormValue("user_id")
		dateStr := req.FormValue("date")
		newDateStr := req.FormValue("new_date")
		if userId == "" || dateStr == "" || newDateStr == "" {
			res.WriteHeader(400)
			data := map[string]string{"error": "Incorrect input"}
			err := json.NewEncoder(res).Encode(data)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		date, err := time.Parse("2019-09-09", dateStr)
		newDate, err := time.Parse("2019-09-09", newDateStr)
		if err != nil {
			data := map[string]string{"error": "Incorrect date"}
			err := json.NewEncoder(res).Encode(data)
			if err != nil {
				log.Fatal(err)
			}
		}
		query := "UPDATE calendar SET date = '$2' WHERE user_id = '$1' AND date = '$3' "
		_, err = s.db.Query(s.ctx, query, userId, date, newDate)
		if err != nil {
			log.Fatal(err)
		}
		data := map[string]string{"result": "Success"}
		err = json.NewEncoder(res).Encode(data)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	res.WriteHeader(503)
}

func (s *Server) deleteEvent(res http.ResponseWriter, req *http.Request) {
	if req.Method == "DELETE" {
		params := req.URL.Query()
		userId := params.Get("uer_id")
		dateStr := params.Get("date")
		if dateStr == "" {
			res.WriteHeader(400)
			data := map[string]string{"error": "Incorrect input"}
			err := json.NewEncoder(res).Encode(data)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		date, err := time.Parse("2019-09-09", dateStr)
		if err != nil {
			data := map[string]string{"error": "Incorrect date"}
			err := json.NewEncoder(res).Encode(data)
			if err != nil {
				log.Fatal(err)
			}
		}
		query := "DELETE FROM calendar WHERE user_id = '$1' AND date = '$2' "
		_, err = s.db.Query(s.ctx, query, userId, date)
		if err != nil {
			log.Fatal(err)
		}
		res.WriteHeader(201)
		data := map[string]string{"result": "Success"}
		err = json.NewEncoder(res).Encode(data)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	res.WriteHeader(503)
}
