package fronius

import (
	u "ha-fronius-bm/src/utils"
	"time"

	"github.com/simonvetter/modbus"
)

var modbusClient *modbus.ModbusClient
var err error

func OpenModbusClient(url string) error {
	url = "tcp://" + url + ":" + "502"
	modbusClient, err = modbus.NewClient(&modbus.ClientConfiguration{
		URL:     url,
		Timeout: 1 * time.Second,
	})
	if err != nil {
		u.Log.Error("Someting goes wrong configuring Modbus Client")
		panic(err)
	}
	err = modbusClient.Open()
	if err != nil {
		u.Log.Error("Someting goes wrong opening Modbus Client")
		panic(err)
	}
	err = modbusClient.SetUnitId(1)
	if err != nil {
		u.Log.Error("Someting goes wrong setting Modbus Client SlaveID")
		panic(err)
	}

	return nil

}

func ClosemodbusClient() error {
	modbusClient.Close()
	if err != nil {
		u.Log.Error("Someting goes wrong closing Modbus Client")
		panic(err)
	}

	return nil
}
