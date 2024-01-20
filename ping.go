package controller

import (
	"fmt"
	"net"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ping(ip net.IP, port string, timeout time.Duration) error {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", ip.String(), port), timeout)
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}
func Ping(c *fiber.Ctx) error {
	h := c.Params("host")
	host := net.ParseIP(h)
	port := c.Params("port")
	timeout := time.Duration(1 * time.Second)
	err := ping(host, port, timeout)
	if err != nil {
		fmt.Println("Ping failed:", err)
		return c.SendString("offline")
	}
	return c.SendString("online")
}
func Health(c *fiber.Ctx) error {
	return c.SendString(string("online"))
}
