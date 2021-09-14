package main

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"time"
)

const (
	BitLenTime      = 39
	BitLenSequence  = 8
	BitLenMachineID = 63 - BitLenTime - BitLenSequence
)

type Snowflake struct {
	mutex *sync.Mutex

	startTime   int64
	elapsedTime int64
	sequence    uint16
	machineID   uint16
}

// Settings configures Snowflake:
// StartTime is the time since which the Snowflake time is defined as the elapsed time.
//
// MachineID returns the unique ID of the Snowflake instance.
//
// Default MachineID returns the lower 16 bits of the private IP address.
// CheckMachineID validates the uniqueness of the machine ID.
type Options struct {
	StartTime      time.Time
	MachineID      func() (uint16, error)
	CheckMachineID func(uint16) bool
}

const flakeTimeUnit = 1e7

// NewSnowflake returns a new Snowflake configured with the given Option Settings
func NewSnowflake(opt Options) *Snowflake {
	sf := new(Snowflake)
	sf.sequence = uint16(1<<BitLenSequence - 1)
	sf.mutex = new(sync.Mutex)

	if opt.StartTime.After(time.Now()) {
		return nil
	}

	if opt.StartTime.IsZero() {
		sf.startTime = toFlakeTime(time.Date(2014, 9, 1, 0, 0, 0, 0, time.UTC))
	} else {
		sf.startTime = toFlakeTime(opt.StartTime)
	}

	var err error
	if opt.MachineID == nil {
		sf.machineID, err = lower16BitPrivateIP()
	} else {
		sf.machineID, err = opt.MachineID()
	}

	if err != nil || (opt.CheckMachineID != nil && !opt.CheckMachineID(sf.machineID)) {
		return nil
	}
	return sf
}

func (sf *Snowflake) NextID() (uint64, error) {
	const maskSeq = uint16(1<<BitLenSequence - 1)
	sf.mutex.Lock()

	defer sf.mutex.Unlock()

	current := elapsedTime(sf.startTime)
	if sf.elapsedTime < current {
		sf.elapsedTime = current
		sf.sequence = 0
	} else {
		sf.sequence = (sf.sequence + 1) & maskSeq
		if sf.sequence == 0 {
			sf.elapsedTime++
			ot := sf.elapsedTime - current
			time.Sleep(sleepTime(ot))
		}
	}

	return sf.toID() // id of this algo... time | seq | machine
}

func DeCoupled(id uint64) map[string]uint64 {
	const (
		maskSeq       = uint64((1<<BitLenSequence - 1) << BitLenMachineID)
		maskMachineID = uint64(1<<BitLenMachineID - 1)
	)

	seq := id & maskSeq >> BitLenMachineID
	msb := id >> 63
	machineID := id & maskMachineID
	time := id >> (BitLenSequence + BitLenMachineID)

	return map[string]uint64{
		"id":         id,
		"sequence":   seq,
		"msb":        msb,
		"machine-id": machineID,
		"time":       time,
	}
}

func (sf *Snowflake) toID() (uint64, error) {
	if sf.elapsedTime >= 1<<BitLenTime {
		return 0, errors.New("over time limit")
	}

	return uint64(sf.elapsedTime)<<(BitLenSequence+BitLenMachineID) | uint64(sf.sequence)<<BitLenMachineID | uint64(sf.machineID), nil
}

func lower16BitPrivateIP() (uint16, error) {
	ip, err := privateIPv4()
	if err != nil {
		return 0, err
	}
	return uint16(ip[2])<<8 + uint16(ip[3]), nil
}

func privateIPv4() (net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, a := range addrs {
		ipnet, ok := a.(*net.IPNet)
		if !ok {
			continue
		}

		ip := ipnet.IP
		if ip.IsLoopback() {
			continue
		}

		ipv4 := ip.To4()
		if isPrivateIPv4(ipv4) {
			return ipv4, nil
		}
	}
	return nil, errors.New("no private ip address")
}

func isPrivateIPv4(ip net.IP) bool {
	return ip != nil &&
		(ip[0] == 10 || ip[0] == 172 && (ip[1] >= 16 && ip[1] < 32) || ip[0] == 192 && ip[1] == 168)
}

func toFlakeTime(t time.Time) int64 {
	return t.UTC().UnixNano() / flakeTimeUnit
}

func elapsedTime(start int64) int64 {
	return toFlakeTime(time.Now()) - start
}

func sleepTime(overtime int64) time.Duration {
	return time.Duration(overtime)*10*time.Microsecond - time.Duration(time.Now().UTC().UnixNano()%flakeTimeUnit)*time.Nanosecond
}

func main() {
	sf := NewSnowflake(Options{})

	id, err := sf.NextID()
	if err != nil {
		panic(err)
	}

	fmt.Println("ID:->", id)
	fmt.Println("DeCoupled: ->", DeCoupled(id))
}
