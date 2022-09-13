package main

type Project struct {
	Name           string `json:"name"`
	PathName       string `json:"pathName"`
	IsOpenSource   bool   `json:"isOpenSource"`
	RepositoryLink string `json:"repositoryLink"`
	ImageLocation  string `json:"imageLocation"`
}

