package github

import (
	"sync"

	"app/service/github/payload"
	"app/service/github/request"
	"app/service/github/response"
)

// GetUserInfo - Get from github information about user, user's repos and organizations.
func GetUserInfo(userName string) response.Default {
	n := 3
	jobs := make(chan request.Job, n)
	results := make(chan request.Job, n)
	wg := sync.WaitGroup{}
	wg.Add(n)

	for i := 0; i < n; i++ {
		go fulfillJob(jobs, results, &wg)
	}
	createJobs(jobs, userName)
	wg.Wait()
	close(results)

	return prepareResponse(results)
}

func createJobs(jobs chan<- request.Job, userName string) {
	jobs <- request.GetProfileJob(userName)
	jobs <- request.GetOrgsJob(userName)
	jobs <- request.GetReposJob(userName)
	close(jobs)
}

func fulfillJob(jobs <-chan request.Job, results chan<- request.Job, wg *sync.WaitGroup) {
	for job := range jobs {
		job.Error = UnmarshalData(GetURI(job.UserName, job.Name), &job.Data)
		results <- job
		wg.Done()
	}
}

func prepareResponse(results <-chan request.Job) response.Default {
	res := response.Default{}
	for result := range results {
		err := ""
		if result.Error != nil {
			err = result.Error.Error()
		}
		res[result.Name] = payload.Default{Data: result.Data, Error: err}
	}

	return res
}
