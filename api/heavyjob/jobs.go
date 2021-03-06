package heavyjob

import "github.com/bk7987/timecards/jobs"

// Job represents a single project.
type Job struct {
	ID             string `json:"id"`
	Code           string `json:"code"`
	Description    string `json:"description"`
	BusinessUnitID string `json:"businessUnitId"`
	Status         string `json:"status"`
}

// GetJobs returns all jobs owned by the company.
func (c *Client) GetJobs() ([]Job, error) {
	jobs := []Job{}
	_, err := c.get("/jobs", &jobs)
	return jobs, err
}

// transformJobs returns a slice of transformed Job models.
func transformJobs(hjJobs []Job) []jobs.Job {
	transformed := []jobs.Job{}
	for _, job := range hjJobs {
		transformed = append(transformed, jobs.Job{
			ID:          job.ID,
			Description: job.Description,
			JobNumber:   job.Code,
		})
	}
	return transformed
}
