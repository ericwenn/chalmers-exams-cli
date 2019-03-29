package main

import (
	"fmt"
	"github.com/ericwenn/chalmers-exams-cli/internal/bowald"
	"github.com/fatih/color"
	"github.com/urfave/cli"
	"log"
	"os"
	"strings"
)

func main() {
	bold := color.New(color.Bold).SprintFunc()
	notPassed := color.New(color.FgRed).SprintFunc()
	three := color.New(color.FgMagenta).SprintFunc()
	four := color.New(color.FgYellow).SprintFunc()
	five := color.New(color.FgGreen).SprintFunc()
	block := "▇"
	barWidth := 50

	app := cli.NewApp()
	app.Name = "exams"
	app.Usage = "Check exam results at Chalmers University of Technology"
	app.UsageText = "exams [options] search-term"
	app.Action = func(c *cli.Context) error {

		courses, err := bowald.Search(c.Args().Get(0))
		if err != nil {
			return err
		}
		for _, course := range courses {
			fmt.Printf("%s: %s\n", bold(course.Code), course.Name)
			for i, exam := range course.Exams {
				if i > 5 {
					break
				}

				participants := exam.NotPassed + exam.Three + exam.Four + exam.Five
				if participants == 0 {
					continue
				}
				mult := float32(barWidth) / float32(participants)
				fmt.Printf("%s: %s%s%s%s (%s/%s/%s/%s)\n",
					exam.Date.Format("2006-01-02"),
					notPassed(strings.Repeat(block, int(float32(exam.NotPassed)*mult))),
					three(strings.Repeat(block, int(float32(exam.Three)*mult))),
					four(strings.Repeat(block, int(float32(exam.Four)*mult))),
					five(strings.Repeat(block, int(float32(exam.Five)*mult))),
					notPassed(exam.NotPassed),
					three(exam.Three),
					four(exam.Four),
					five(exam.Five),
				)
			}
			fmt.Println()
		}
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}