package service

import (
	"go-testing/doubles/internal/city"
	"go-testing/doubles/internal/city/repository"
	"go-testing/doubles/internal/logger"
	"go-testing/doubles/internal/weather"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDefault_AddCity(t *testing.T) {
	t.Run("success to add a city", func(t *testing.T) {
		// arrange
		// - logger
		lgDummy := logger.NewDummy()
		// - weather
		weatherStub := weather.NewStub()
		weatherStub.FuncGetTemperature = func(city string) (float64, error) {
			return 20, nil
		}
		// - repository
		cityRpWriterMock := repository.NewMock()
		cityRpWriterMock.FuncSaveCity = func(c *city.City) error {
			(*c).ID = 1
			return nil
		}
		sv := NewDefault(cityRpWriterMock, weatherStub, lgDummy)

		// act
		c, err := sv.AddCity("city", "country", 1000, time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC))

		// assert
		require.NoError(t, err)
		require.Equal(t, city.City{
			ID: 	   1,
			Name:      "city",
			Country:   "country",
			Population: 1000,
			Temperature: 20,
			Date:      time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		}, c)
		require.Equal(t, city.City{
			ID: 	   0,
			Name:      "city",
			Country:   "country",
			Population: 1000,
			Temperature: 20,
			Date:      time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		}, cityRpWriterMock.Calls.SaveCityArgs)
		require.Equal(t, 1, cityRpWriterMock.Calls.SaveCityCount)
	})

	t.Run("fail to add a city - invalid params", func(t *testing.T) {
		// arrange

		// act

		// assert
	})

	t.Run("fail to add a city - temperature error", func(t *testing.T) {
		// arrange

		// act

		// assert
	})

	t.Run("fail to add a city - save city error", func(t *testing.T) {
		// arrange

		// act

		// assert
	})
}