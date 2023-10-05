package api

import (
	"database/sql"
	"fmt"
	uuid "github.com/google/uuid"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"time"
	config "vid/config"
	src "vid/src"
	"vid/utils"
	ffmpeg "vid/utils/ffmpeg"
	jobqueue "vid/utils/jobqueue"
)

func UploadRoutes(r *mux.Router, db *sql.DB, pool *jobqueue.Pool) {
	r.HandleFunc("/api/upload", config.WithLogMiddleware(func(w http.ResponseWriter, r *http.Request) {
		valid := utils.NewRequestValidator(w, r)
		isValid := valid.ValidateRequest()
		if !isValid {
			fmt.Fprintf(w, "Invalid Request")
			return
		}
		handleUpload(w, r, pool)
	})).Methods("POST")

	r.HandleFunc("/api/upload", config.WithLogMiddleware(func(w http.ResponseWriter, r *http.Request) {
		valid := utils.NewRequestValidator(w, r)
		isValid := valid.ValidateRequest()
		if !isValid {
			fmt.Fprintf(w, "Invalid Request")
			return
		}
		handleGet(w, r, pool)
	})).Methods("GET")
}

func alphabets(name string) error {
	fmt.Println("<-------------- GENERATING UNIQUE ID ------------------->")
	for _, char := range name {
		fmt.Printf("%c", char)
		// print1 secs
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("\n <-------------- UNIQUE ID ENDS ------------------->")
	return nil
}

func handleGet(w http.ResponseWriter, r *http.Request, p *jobqueue.Pool) {
	queueName := r.URL.Hostname()
	queue := jobqueue.NewQueue(queueName, p.MAX_SIZE)
	jobs := []jobqueue.Job{}
	for i := 0; i < 3; i++ {
		job := jobqueue.Job{
			Name: uuid.New().String(),
			Action: func() error {
				var uuid string = uuid.New().String()
				//wait
				return alphabets(uuid)
			},
		}
		jobs = append(jobs, job)
	}
	queue.AddJobs(jobs)
	defaultWorker := jobqueue.NewWorker(queue)
	p.AddWorker(defaultWorker)
	p.Start()
	fmt.Fprintf(w, "Job added to queue")
}

func handleUpload(w http.ResponseWriter, r *http.Request, q *jobqueue.Pool) {
	r.ParseMultipartForm(10 << 20)
	multipart_file, file_headers, err := r.FormFile("file")
	if err != nil {
		slog.Error("Error Retrieving the File")
		return
	}

	file := src.File{
		OriginalName: file_headers.Filename,
		Size:         file_headers.Size,
		Header:       file_headers.Header.Get("Content-Type"),
	}
	src.ValidateDir("uploads")
	file.Upload(multipart_file)
	jobs := []jobqueue.Job{}

	for _, resolution := range []ffmpeg.VideoResolution{ffmpeg.LOW, ffmpeg.MID, ffmpeg.HIGH} {
		video := ffmpeg.NewVideo(&file, resolution)
		newAction := func() error {
			fmt.Println("Video", video)
			video.CreateHLS()
			fmt.Println("Video", video.GetResolution())
			return nil
		}
		job := jobqueue.Job{
			Name:      file.Filename,
			Completed: false,
			Action: func() error {
				newAction()
				return nil
			},
		}
		jobs = append(jobs, job)
	}

	queueName := r.URL.Hostname()
	queue := jobqueue.NewQueue(queueName, q.MAX_SIZE)
	queue.AddJobs(jobs)
	defaultWorker := jobqueue.NewWorker(queue)
	q.AddWorker(defaultWorker)
	q.Start()
	fmt.Fprintf(w, "Successfully Uploaded File at %s\n", file.Path)

}
