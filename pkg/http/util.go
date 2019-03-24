package http

import (
	"strconv"
	"strings"
)

func ParsePortFromAddr(addr string) uint16 {
	parts := strings.Split(addr, ":")
	if len(parts) == 1 {
		return 80
	}
	portNum, _ := strconv.Atoi(parts[1])
	return uint16(portNum)
}
