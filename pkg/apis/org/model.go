package org

type Org struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Orgs []Org
