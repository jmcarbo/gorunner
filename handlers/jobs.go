package handlers

import (
	"html/template"
	"net/http"
	"github.com/jakecoffman/gorunner/db"
	"github.com/jakecoffman/gorunner/models"
)

const jobsFile = "jobs.json"


var jobs = template.Must(template.ParseFiles(
	"web/templates/_base.html",
	"web/templates/jobs.html",
))

func Jobs(w http.ResponseWriter, r *http.Request) {
	var jobList models.JobList
	db.Load(&jobList, jobsFile)

	if r.Method == "GET"{
		if err := jobs.Execute(w, jobList); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if r.Method == "POST" {
		name := r.FormValue("name")
		jobList.Append(models.Job{name, []models.Task{}})
		db.Save(&jobList, jobsFile)
		if err := jobs.Execute(w, jobList); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}