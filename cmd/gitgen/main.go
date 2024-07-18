package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/seymahandekli/git-gen/pkg/gitgen"
	"github.com/urfave/cli/v3"
)

var (
	PromptMap = map[string]gitgen.PromptType{
		"commit": gitgen.PromptCommitMessage,
		"review": gitgen.PromptCodeReview,
	}
)

func main() {
	var promptTypeStr string

	var maxTokens int64

	cmd := &cli.Command{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "prompt",
				Value:       "commit",
				Destination: &promptTypeStr,
			},
			&cli.IntFlag{
				Name:        "maxtokens",
				Value:       3500,
				Destination: &maxTokens,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			value, ok := PromptMap[promptTypeStr]
			if !ok {
				return fmt.Errorf("invalid prompt type - %s", promptTypeStr)
			}
			result, err := gitgen.Do(value, maxTokens)

			if err != nil {
				return err
			}

			log.Println(result)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
		return
	}
}
