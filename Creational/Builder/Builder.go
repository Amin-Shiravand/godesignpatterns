package main

import "fmt"

type EmailBuilder struct {
	title       string
	address     string
	description string
	attachments []string
}

func newEmailBuilder() *EmailBuilder {
	return &EmailBuilder{}
}

func (e *EmailBuilder) SetTitle(title string) {
	e.title = title
}

func (e *EmailBuilder) SetAddress(address string) {
	e.address = address
}

func (e *EmailBuilder) SetDescription(description string) {
	e.description = description
}

func (e *EmailBuilder) SetAttachments(attachments []string) {
	e.attachments = attachments
}

func (e *EmailBuilder) Build() (*EmailBuilder, error) {

	if e.title == "" {
		return nil, fmt.Errorf("email title is empty")
	}

	if e.address == "" {
		return nil, fmt.Errorf("title is empty")
	}
	return &EmailBuilder{
		title:       e.title,
		address:     e.address,
		description: e.description,
		attachments: e.attachments,
	}, nil
}
