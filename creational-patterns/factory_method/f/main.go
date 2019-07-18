package main

import "fmt"

type Interviewer interface {
	askQuestions()
}

type Developer struct{}

func (d *Developer) askQuestions() {
	fmt.Println("asking about design pattern")
}

type CommunityExecutive struct{}

func (c *CommunityExecutive) askQuestions() {
	fmt.Println("asking about community building")
}

type HiringManager struct{}

type Hiring interface {
	makeInterviewer() Interviewer
}

func (hm *HiringManager) takeInterview(h Hiring) {
	interviewer := h.makeInterviewer()
	interviewer.askQuestions()
}

type DevelopmentManager struct{}

func (d *DevelopmentManager) makeInterviewer() Interviewer {
	return &Developer{}
}

type MarketingManager struct{}

func (d *MarketingManager) makeInterviewer() Interviewer {
	return &CommunityExecutive{}
}

func main() {
	hiringManager := &HiringManager{}

	devManager := &DevelopmentManager{}
	marketingManager := &MarketingManager{}

	hiringManager.takeInterview(devManager)
	hiringManager.takeInterview(marketingManager)

}
