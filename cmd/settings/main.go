package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/n7down/kuiper/internal/settings/listeners"
	"github.com/n7down/kuiper/internal/settings/persistence/mysql"
	"google.golang.org/grpc"

	commonServers "github.com/n7down/kuiper/internal/common/servers"
	settings_pb "github.com/n7down/kuiper/internal/pb/settings"
	settings "github.com/n7down/kuiper/internal/settings/servers"
	log "github.com/sirupsen/logrus"
)

const (
	ONE_MINUTE = 1 * time.Minute
)

var (
	Version         string
	Build           string
	showVersion     *bool
	listenersServer *commonServers.ListenersServer
	port            string
	server          *settings.SettingServer
)

func init() {
	showVersion = flag.Bool("v", false, "show version and build")
	flag.Parse()
	if !*showVersion {
		port = os.Getenv("PORT")
		dbConn := os.Getenv("DB_CONN")
		batCaveMQTTURL := os.Getenv("BAT_CAVE_MQTT_URL")

		settingsDB, err := mysql.NewSettingsMySqlDB(dbConn)
		if err != nil {
			log.Fatal(err)
		}

		server = settings.NewSettingServer(settingsDB)
		listenersServer = commonServers.NewListenersServer()
		settingsListenersEnv := listeners.NewSettingsListenersEnv(settingsDB)

		batCaveListener, err := settingsListenersEnv.NewBatCaveSettingsListener("bat_cave_listener", batCaveMQTTURL)
		if err != nil {
			log.Fatal(err)
		}
		listenersServer.AddListener(batCaveListener)
	}
}

func main() {
	if *showVersion {
		fmt.Printf("settings server: version %s build %s", Version, Build)
	} else {
		log.SetReportCaller(true)

		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
		if err != nil {
			log.Fatal(err)
		}

		listenersServer.Connect()

		log.Infof("Listening on port: %s\n", port)
		grpcServer := grpc.NewServer()
		settings_pb.RegisterSettingsServiceServer(grpcServer, server)
		grpcServer.Serve(lis)
	}
}
