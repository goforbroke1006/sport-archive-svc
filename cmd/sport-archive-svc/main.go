package main

import (
	"bytes"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/google/logger"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
	"github.com/streadway/amqp"

	"github.com/goforbroke1006/sport-archive-svc/pkg/domain"
	"github.com/goforbroke1006/sport-archive-svc/pkg/endpoint"
	"github.com/goforbroke1006/sport-archive-svc/pkg/service"
	"github.com/goforbroke1006/sport-archive-svc/pkg/trace"
)

var (
	dbConnStr   = flag.String("db-conn", "./sport-archive.db", "")
	amqpConnStr = flag.String("amqp-conn", "amqp://guest:guest@localhost:5672/", "Define AMQP server address and credentials")
	inQueueName = flag.String("in-queue", "sport-archive-request", "Define AMQP server address and credentials")
	//outQueueName = flag.String("out-queue", "sport-archive-response", "Define AMQP server address and credentials")
	allowSave = flag.Bool("allow-save", true, "")
	logPath   = flag.String("log-path", "./access.log", "")
	verbose   = flag.Bool("verbose", true, "Print info level logs to stdout")

	zipkinAddr = flag.String("zipkin-addr", "http://127.0.0.1:9411", "Select zipkin host")
)

const serviceName = "sport-archive-svc"

func init() {
	flag.Parse()
}

func main() {
	logFile, err := os.OpenFile(*logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}
	defer finalizeCloser(logFile)

	defer logger.Init(serviceName, *verbose, false, logFile).Close()

	tracer, err := trace.NewTracer(*zipkinAddr, serviceName, 0)
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open("sqlite3", *dbConnStr)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}
	defer finalizeCloser(db)
	db.AutoMigrate(&domain.Sport{})
	db.AutoMigrate(&domain.Participant{})

	svc := service.NewSportArchiveService(db, *allowSave)
	eps := endpoint.NewSportArchiveService(svc)

	server := rpc.NewServer()
	server.RegisterCodec(json.NewCodec(), "application/json")
	server.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	err = server.RegisterService(eps, "SportArchiveService")
	if err != nil {
		log.Fatalf("Format of service SportArchiveServiceEndpoint isn't correct. %s", err)
	}

	conn, err := amqp.Dial(*amqpConnStr)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		*inQueueName, // name
		false, false, false, false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(1, 0, false)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(q.Name, "",
		false, false, false, false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			span := tracer.StartSpan("request")

			httpConn := &HttpConn{}

			logger.Infof("request: %s", string(d.Body))

			_, err := httpConn.Write(d.Body)
			failOnError(err, "can't create request")

			headers := http.Header{}
			headers.Set("Content-Type", d.ContentType)

			httpRequest := &http.Request{
				Method: "POST",
				Body:   httpConn,
				Header: headers,
			}

			httpResponseWtr := &MyResponseWriter{
				HeaderField: http.Header{},
				Body:        bytes.NewBufferString(""),
			}
			server.ServeHTTP(httpResponseWtr, httpRequest)

			respBytes, _ := ioutil.ReadAll(httpResponseWtr)

			logger.Infof("response: %s", string(respBytes))
			err = ch.Publish(
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          respBytes,
				})
			failOnError(err, "Failed to publish a message")

			span.Finish()

			d.Ack(false)
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever
}

func finalizeCloser(c io.Closer) {
	err := c.Close()
	if nil != err {
		logger.Fatalf("Failed to close descriptor: %v", err)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		logger.Fatalf("%s: %s", msg, err)
	}
}

type MyResponseWriter struct {
	HeaderField http.Header
	//Body        []byte
	Body *bytes.Buffer
}

func (rw MyResponseWriter) Header() http.Header {
	//panic("implement me")
	return rw.HeaderField
}

func (rw MyResponseWriter) Write(data []byte) (int, error) {
	//panic("implement me")
	//rw.BOdy = bytes.NewBuffer(data)
	return rw.Body.Write(data)
	//return len(data), nil
}

func (rw MyResponseWriter) WriteHeader(statusCode int) {
	//panic("implement me")
}

func (rw *MyResponseWriter) Read(p []byte) (n int, err error) {
	//return c.in.Read(p)
	return rw.Body.Read(p)
}

type HttpConn struct {
	//in  io.Reader
	//out io.Writer

	buf bytes.Buffer
}

func (c *HttpConn) Read(p []byte) (n int, err error) {
	//return c.in.Read(p)
	return c.buf.Read(p)
}

func (c *HttpConn) Write(d []byte) (n int, err error) {
	//return c.out.Write(d)
	return c.buf.Write(d)
}

func (c *HttpConn) Close() error {
	return nil
}
