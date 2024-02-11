package entities

type Version struct {
	CommitId       string `json:"commit_id"`
	BuildTimestamp string `json:"build_timestamp"`
	BranchName     string `json:"branch_ref"`
	BranchNameRef  string `json:"branch"`
	BuildTag       string `json:"build_tag"`
}

var VERSION Version
