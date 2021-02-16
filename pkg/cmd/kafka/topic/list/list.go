package list

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/flag"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/kas"
	strimziadminclient "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/strimzi-admin/client"

	"gopkg.in/yaml.v2"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/dump"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/iostreams"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"
	"github.com/spf13/cobra"
)

type Options struct {
	Config     config.IConfig
	IO         *iostreams.IOStreams
	Connection func() (connection.Connection, error)
	Logger     func() (logging.Logger, error)

	kafkaID string
	output  string
}

type topicRow struct {
	Name            string `json:"name,omitempty" header:"Name"`
	PartitionsCount int    `json:"partitions_count,omitempty" header:"Partitions"`
}

// NewListTopicCommand gets a new command for getting kafkas.
func NewListTopicCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
	}

	localizer.LoadMessageFiles("cmd/kafka/topic/common", "cmd/kafka/topic/list", "cmd/kafka/common")

	cmd := &cobra.Command{
		Use:     localizer.MustLocalizeFromID("kafka.topic.list.cmd.use"),
		Short:   localizer.MustLocalizeFromID("kafka.topic.list.cmd.shortDescription"),
		Long:    localizer.MustLocalizeFromID("kafka.topic.list.cmd.longDescription"),
		Example: localizer.MustLocalizeFromID("kafka.topic.list.cmd.example"),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.output != "" {
				if err := flag.ValidateOutput(opts.output); err != nil {
					return err
				}
			}

			return runCmd(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.output, "output", "o", "", localizer.MustLocalize(&localizer.Config{
		MessageID:   "kafka.topic.common.flag.output.description",
		PluralCount: 2,
	}))

	return cmd
}

func runCmd(opts *Options) error {
	conn, err := opts.Connection()
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	api := conn.API()

	ctx := context.Background()
	kafkaInstance, _, apiErr := api.Kafka().GetKafkaById(ctx, opts.kafkaID).Execute()

	if kas.IsErr(apiErr, kas.ErrorNotFound) {
		return errors.New(localizer.MustLocalize(&localizer.Config{
			MessageID: "kafka.common.error.notFoundByIdError",
			TemplateData: map[string]interface{}{
				"ID": opts.kafkaID,
			},
		}))
	} else if apiErr.Error() != "" {
		return apiErr
	}

	a := api.TopicAdmin(opts.kafkaID).GetTopicsList(context.Background())
	topicData, _, topicErr := a.Execute()

	if topicErr.Error() != "" {
		return topicErr
	}

	if topicData.GetCount() == 0 {
		logger.Info(localizer.MustLocalize(&localizer.Config{
			MessageID: "kafka.topic.list.log.info.noTopics",
			TemplateData: map[string]interface{}{
				"InstanceName": kafkaInstance.GetName(),
			},
		}))

		return nil
	}

	stdout := opts.IO.Out
	switch opts.output {
	case "json":
		data, _ := json.Marshal(topicData)
		_ = dump.JSON(stdout, data)
	case "yaml", "yml":
		data, _ := yaml.Marshal(topicData)
		_ = dump.YAML(stdout, data)
	default:
		topics := topicData.GetTopics()
		rows := mapTopicResultsToTableFormat(topics)
		dump.Table(stdout, rows)
	}

	return err
}

func mapTopicResultsToTableFormat(topics []strimziadminclient.Topic) []topicRow {
	var rows []topicRow = []topicRow{}

	for _, t := range topics {
		row := topicRow{
			Name:            t.GetName(),
			PartitionsCount: len(t.GetPartitions()),
		}
		rows = append(rows, row)
	}

	return rows
}