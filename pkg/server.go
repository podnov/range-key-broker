package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type server struct {
	b broker
}

func createRangeValues() (result []string, err error) {
	rawRangeValue := viper.GetString(ConfigKeyRangeValueRange)

	err = json.Unmarshal([]byte(rawRangeValue), &result)
	if err != nil {
		message := fmt.Sprintf("Could not unmarshal range json [%s]", rawRangeValue)
		err = errors.Wrap(err, message)
	}

	return result, err
}

func (s *server) handleCheckin(c *gin.Context) {
	rangeValue := c.Param("rangeValue")
	err := s.b.CheckIn(rangeValue)
	if err != nil {
		glog.Warningf("Could not check in value [%s]", rangeValue)
	}
}

func (s *server) handleCheckout(c *gin.Context) {
	rangeValue, err := s.b.CheckOut()
	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, rangeValue)
}

func (s *server) handleHealthz(c *gin.Context) {
	c.Status(200)
}

func NewServer() (server, error) {
	rangeValues, err := createRangeValues()
	if err != nil {
		return server{}, err
	}

	broker := NewBroker(rangeValues)

	result := server{
		b: broker,
	}

	return result, nil
}

func (s *server) Start() {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/checkout", s.handleCheckout)
	router.DELETE("/checkout/:rangeValue", s.handleCheckin)
	router.GET("/healthz", s.handleHealthz)

	router.Run(":8080")
}
