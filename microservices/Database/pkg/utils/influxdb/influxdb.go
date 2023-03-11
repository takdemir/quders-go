package influxdb

import (
	"context"
	"database/pkg/utils/common"
	"errors"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"os"
	"strings"
	"time"
)

var client influxdb2.Client
var writeAPI api.WriteAPIBlocking
var queryAPI api.QueryAPI

type Point struct {
	Measurement string
	Tags        map[string]string
	Fields      map[string]interface{}
	Precision   string
	Time        time.Time
}

func Connect(influxDbEnvName string, batchSize uint) {
	if influxDbEnvName == "" {
		influxDbEnvName = "INFLUXDB_1"
	}
	if batchSize <= 0 {
		batchSize = 1
	}
	connectionStrings := ConnectionStringParser(os.Getenv(influxDbEnvName))
	if len(connectionStrings) < 4 {
		errorMessage := "influxDB connection string error"
		common.FailOnError(errors.New(errorMessage), errorMessage)
		panic(errorMessage)
	}
	_, isHostExist := connectionStrings["host"]
	if !isHostExist {
		errorMessage := "influxDB connection string host error"
		common.FailOnError(errors.New(errorMessage), errorMessage)
		panic(errorMessage)
	}

	_, isTokenExist := connectionStrings["token"]
	if !isTokenExist {
		errorMessage := "influxDB connection string token error"
		common.FailOnError(errors.New(errorMessage), errorMessage)
		panic(errorMessage)
	}

	_, isOrgExist := connectionStrings["org"]
	if !isOrgExist {
		errorMessage := "influxDB connection string organization error"
		common.FailOnError(errors.New(errorMessage), errorMessage)
		panic(errorMessage)
	}

	_, isBucketExist := connectionStrings["bucket"]
	if !isBucketExist {
		errorMessage := "influxDB connection string bucket error"
		common.FailOnError(errors.New(errorMessage), errorMessage)
		panic(errorMessage)
	}

	client := influxdb2.NewClientWithOptions(connectionStrings["host"], connectionStrings["token"], influxdb2.DefaultOptions().SetBatchSize(batchSize))
	writeAPI = client.WriteAPIBlocking(connectionStrings["org"], connectionStrings["bucket"])
	queryAPI = client.QueryAPI(connectionStrings["org"])
}

func Write(point Point) {
	if point.Precision == "" {
		point.Precision = "ns"
	}
	if point.Time.IsZero() {
		point.Time = time.Now()
	}
	newPoint := influxdb2.NewPointWithMeasurement(point.Measurement).SetTime(point.Time)
	for k, v := range point.Tags {
		newPoint.AddTag(k, v)
	}
	for k, v := range point.Fields {
		newPoint.AddField(k, v)
	}

	err := writeAPI.WritePoint(context.Background(), newPoint)
	if err != nil {
		common.FailOnError(err, "influxdb write point error")
		panic(err)
	}
	err = writeAPI.Flush(context.Background())
	if err != nil {
		return
	}
	client.Close()
}

func WriteBulk(points []Point) {
	for _, point := range points {
		if point.Precision == "" {
			point.Precision = "ns"
		}
		if point.Time.IsZero() {
			point.Time = time.Now()
		}
		newPoint := influxdb2.NewPointWithMeasurement(point.Measurement).SetTime(point.Time)
		for k, v := range point.Tags {
			newPoint.AddTag(k, v)
		}
		for k, v := range point.Fields {
			newPoint.AddField(k, v)
		}

		err := writeAPI.WritePoint(context.Background(), newPoint)
		if err != nil {
			common.FailOnError(err, "influxdb write point error")
			panic(err)
		}
	}
	err := writeAPI.Flush(context.Background())
	if err != nil {
		return
	}

	client.Close()
}

func Query(query string) {
	result, err := queryAPI.Query(context.Background(), query)
	if err == nil {
		// Use Next() to iterate over query result lines
		for result.Next() {
			// Observe when there is new grouping key producing new table
			if result.TableChanged() {
				fmt.Printf("table: %s\n", result.TableMetadata().String())
			}
			// read result
			fmt.Printf("row: %s\n", result.Record().String())
		}
		if result.Err() != nil {
			common.FailOnError(errors.New(""), fmt.Sprintf("Query error: %s\n", result.Err().Error()))
		}
	} else {
		panic(err)
	}
	client.Close()
}

func ConnectionStringParser(connectionString string) map[string]string {
	explodeFromAt := strings.Split(connectionString, "@")
	if len(explodeFromAt) < 2 {
		errorMessage := "no @ in connection string"
		common.FailOnError(errors.New(errorMessage), errorMessage)
		panic(errorMessage)
	}
	explodeFromComma := strings.Split(explodeFromAt[0], ":")
	if len(explodeFromAt) < 2 {
		errorMessage := "no : in connection string"
		common.FailOnError(errors.New(errorMessage), errorMessage)
		panic(errorMessage)
	}
	explodeFromQuestionMark := strings.Split(explodeFromAt[1], "?")
	if len(explodeFromAt) < 2 {
		errorMessage := "no ? in connection string"
		common.FailOnError(errors.New(errorMessage), errorMessage)
		panic(errorMessage)
	}

	return map[string]string{
		"org":    explodeFromComma[0],
		"bucket": explodeFromComma[1],
		"host":   explodeFromQuestionMark[0],
		"token":  explodeFromQuestionMark[1],
	}

}
