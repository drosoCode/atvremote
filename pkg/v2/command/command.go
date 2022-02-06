package command

import (
	"crypto/tls"
	"errors"
	"strconv"
	"time"

	"github.com/drosocode/atvremote/pkg/common"
	pb "github.com/drosocode/atvremote/pkg/v2/proto"
	"google.golang.org/protobuf/proto"
)

type Command struct {
	Buffer       []byte
	BufferSize   int
	Connection   *tls.Conn
	Certificates *tls.Certificate
	Address      string
	Port         int
	Read         bool
	RemoteData   RemoteData
	Error        error
}

// struct of available infos for the android tv device
type RemoteData struct {
	Powered struct {
		Powered    bool      `json:"powered"`
		UpdateDate time.Time `json:"updateDate"`
	} `json:"powered"`
	App struct {
		CurrentApp string    `json:"currentApp"`
		UpdateDate time.Time `json:"updateDate"`
	} `json:"app"`
	Volume struct {
		Level       uint32    `json:"level"`
		Maximum     uint32    `json:"maximum"`
		Muted       bool      `json:"muted"`
		PlayerModel string    `json:"playerModel"`
		UpdateDate  time.Time `json:"updateDate"`
	} `json:"volume"`
	Device struct {
		Model      string    `json:"model"`
		Vendor     string    `json:"vendor"`
		Version    string    `json:"version"`
		UpdateDate time.Time `json:"updateDate"`
	} `json:"device"`
}

// recreate a complete protobuf message from multiple chunks
// each new message is sent to handleData
func (c *Command) processData(read_size int, size_to_read int, buffer *[]byte, total_data *[]byte) int {

	if read_size < size_to_read {
		// not enough data
		size_to_read -= read_size
		data := make([]byte, read_size)
		copy(data, (*buffer)[0:read_size])
		(*total_data) = append((*total_data), data...)

	} else {

		data := make([]byte, size_to_read)
		copy(data, (*buffer)[0:size_to_read])

		c.handleData(append((*total_data), data...))

		if read_size > size_to_read {
			// too much data
			remaining := read_size - size_to_read - 1
			size_to_read = int(c.Buffer[size_to_read]) - remaining

			(*total_data) = make([]byte, remaining)
			copy((*total_data), (*buffer)[size_to_read+1:read_size])

		} else {
			// exact amount of data
			(*total_data) = make([]byte, 0)
			size_to_read = 0
		}
	}

	return size_to_read
}

// reading loop to read data from the android tv device
func (c *Command) readLoop() {
	buffer := make([]byte, 512)
	var total_data []byte
	size_to_read := 0

	for c.Read {

		read_size, err := c.Connection.Read(buffer)
		if err != nil {
			c.Error = err
			break
		}

		if size_to_read == 0 {
			size_to_read = int(buffer[0])
			if read_size > 1 {
				// if there is more than 1 byte, the size is the first byte of the data
				buffer = buffer[1:]
				size_to_read = c.processData(read_size-1, size_to_read, &buffer, &total_data)
			}
		} else {
			size_to_read = c.processData(read_size, size_to_read, &buffer, &total_data)
		}

	}
}

// send a remote message to the android tv device
func (c *Command) write(data pb.RemoteMessage) error {
	raw, err := proto.Marshal(&data)
	if err != nil {
		return errors.New("v2 - write message - marshal:" + err.Error())
	}

	_, err = c.Connection.Write([]byte{byte(len(raw))})
	if err != nil {
		return errors.New("v2 - write message - send 1:" + err.Error())
	}
	_, err = c.Connection.Write(raw)
	if err != nil {
		return errors.New("v2 - write message - send 2:" + err.Error())
	}

	return nil
}

// Create a new command object with the ip and port of the android tv device and a certificate
func New(addr string, port int, certs *tls.Certificate) Command {
	return Command{Buffer: make([]byte, 512), Address: addr, Port: port, Certificates: certs, Read: true, RemoteData: RemoteData{}, Error: nil}
}

// Connect to the android tv device
func (c *Command) Connect() error {
	config := &tls.Config{Certificates: []tls.Certificate{*c.Certificates}, InsecureSkipVerify: true}

	conn, err := tls.Dial("tcp", c.Address+":"+strconv.Itoa(c.Port), config)
	if err != nil {
		return errors.New("command - connexion: " + err.Error())
	}
	c.Connection = conn

	state := conn.ConnectionState()
	for !state.HandshakeComplete {
	}

	go c.readLoop()

	c.sendConfiguration()

	return nil
}

// handle message from the android tv device
func (c *Command) handleData(raw []byte) {
	data := pb.RemoteMessage{}
	err := proto.Unmarshal(raw, &data)
	if err != nil {
		// ignore invalid protobuf messages
		return
	}

	switch {
	case data.RemoteSetActive != nil:
		err = c.write(pb.RemoteMessage{
			RemoteSetActive: &pb.RemoteSetActive{
				Active: 622,
			},
		})
	case data.RemotePingRequest != nil:
		err = c.write(pb.RemoteMessage{
			RemotePingResponse: &pb.RemotePingResponse{
				Val1: data.RemotePingRequest.Val1,
			},
		})
	case data.RemoteImeKeyInject != nil:
		c.RemoteData.App.CurrentApp = data.RemoteImeKeyInject.AppInfo.AppPackage
		c.RemoteData.App.UpdateDate = time.Now()
	case data.RemoteStart != nil:
		c.RemoteData.Powered.Powered = data.RemoteStart.Started
		c.RemoteData.Powered.UpdateDate = time.Now()
	case data.RemoteSetVolumeLevel != nil:
		c.RemoteData.Volume.Level = data.RemoteSetVolumeLevel.VolumeLevel
		c.RemoteData.Volume.Maximum = data.RemoteSetVolumeLevel.VolumeMax
		c.RemoteData.Volume.Muted = data.RemoteSetVolumeLevel.VolumeMuted
		c.RemoteData.Volume.PlayerModel = data.RemoteSetVolumeLevel.PlayerModel
		c.RemoteData.Volume.UpdateDate = time.Now()
	case data.RemoteConfigure != nil:
		c.RemoteData.Device.Model = data.RemoteConfigure.DeviceInfo.Model
		c.RemoteData.Device.Vendor = data.RemoteConfigure.DeviceInfo.Vendor
		c.RemoteData.Device.Version = data.RemoteConfigure.DeviceInfo.AppVersion
		c.RemoteData.Device.UpdateDate = time.Now()
	}

	if err != nil {
		c.Error = err
	}
}

// Return informations on the current android tv device
func (c *Command) GetData() (RemoteData, error) {
	start := time.Now()
	c.sendConfiguration()
	// wait until data is available
	for start.After(c.RemoteData.Powered.UpdateDate) || start.After(c.RemoteData.Device.UpdateDate) || start.After(c.RemoteData.App.UpdateDate) {
		// timeout after 3 seconds
		if time.Since(start) > time.Second*3 {
			return c.RemoteData, errors.New("unable to fully retreive data")
		}
	}
	return c.RemoteData, nil
}

// send a configuration message to the android tv device
func (c *Command) sendConfiguration() error {
	err := c.write(pb.RemoteMessage{
		RemoteConfigure: &pb.RemoteConfigure{
			Code1: 622,
			DeviceInfo: &pb.RemoteDeviceInfo{
				Model:       "androidtvremote",
				Vendor:      "drosocode",
				Unknown1:    1,
				Unknown2:    "1",
				PackageName: "androitv-remote",
				AppVersion:  "1.0.0",
			},
		},
	})
	if err != nil {
		return err
	}
	err = c.write(pb.RemoteMessage{
		RemoteSetActive: &pb.RemoteSetActive{
			Active: 622,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

// emulate a key press on the android tv device
func (c *Command) SendKey(keycode common.RemoteKeyCode) error {
	return c.write(pb.RemoteMessage{
		RemoteKeyInject: &pb.RemoteKeyInject{
			KeyCode:   pb.RemoteKeyCode(keycode),
			Direction: pb.RemoteDirection_SHORT,
		},
	})
}

// open a link on the android tv device
// if the links matches the one defined in an app's manifest, the app will be launched
// else, the link will open in a browser
func (c *Command) OpenLink(link string) error {
	return c.write(pb.RemoteMessage{
		RemoteAppLinkLaunchRequest: &pb.RemoteAppLinkLaunchRequest{
			AppLink: link,
		},
	})
}
