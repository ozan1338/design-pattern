package main

import (
	"fmt"
	"strings"
)

type email struct {
	from, to, subject, body string
	err                     []error
}

type EmailBuilder struct {
	email email
}

func (this *EmailBuilder) ValidateEmail(email string) bool {
	if !strings.Contains(email, "@") {
		return false
	}

	return true
}

func (this *EmailBuilder) From(from string) *EmailBuilder {
	if !this.ValidateEmail(from) {
		this.email.err = append(this.email.err, fmt.Errorf("from email should be contain @"))
		return this
	}

	this.email.from = from
	return this
}

func (this *EmailBuilder) To(to string) *EmailBuilder {
	if !this.ValidateEmail(to) {
		this.email.err = append(this.email.err, fmt.Errorf("to email should be contain @"))
		return this
	}

	this.email.to = to
	return this
}

func (this *EmailBuilder) Subject(subject string) *EmailBuilder {
	this.email.subject = subject
	return this
}

func (this *EmailBuilder) Body(body string) *EmailBuilder {
	this.email.body = body
	return this
}

func sendMailImpl(email *email) {
	if len(email.err) > 0 {
		for _, val := range email.err {
			fmt.Println(val.Error())
		}
	} else {
		fmt.Printf("email send to %s from %s\n", email.to, email.from)
	}
}

type build func(*EmailBuilder)

func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder)
	sendMailImpl(&builder.email)
}

func main() {
	SendEmail(func(b *EmailBuilder) {
		b.From("foobar.com").To("barfoo.com").Subject("Test").Body("Hello")
	})

	SendEmail(func(b *EmailBuilder) {
		b.From("foo@bar.com").To("bar@foo.com").Subject("Test").Body("Hello")
	})
}
