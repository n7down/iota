package influxpersistence

import (
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	sensors "github.com/n7down/kuiper/internal/sensors/persistence/devicesensors"
)

func (i InfluxPersistence) CreateBMP280(sensor *sensors.BMP280Sensor) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  i.Database,
		Precision: "s",
	})
	if err != nil {
		return err
	}

	pressureFloat, err := sensor.GetPressureFloat()
	if err != nil {
		return err
	}

	temperatureFloat, err := sensor.GetTemperatureFloat()
	if err != nil {
		return err
	}

	// indexed
	tags := map[string]string{
		"mac": sensor.Mac,
	}

	// not indexed
	fields := map[string]interface{}{
		"pressure": pressureFloat,
		"temp":     temperatureFloat,
	}

	point, err := client.NewPoint(
		"bmp280_listene,",
		tags,
		fields,
		time.Now().UTC(),
	)

	bp.AddPoint(point)

	err = i.Client.Write(bp)
	if err != nil {
		return err
	}

	return nil
}
