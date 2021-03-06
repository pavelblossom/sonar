package modules

import (
	"crypto/tls"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/bi-zone/sonar/internal/database"
	"github.com/bi-zone/sonar/internal/modules/api"
	"github.com/bi-zone/sonar/internal/modules/telegram"
)

type Controller interface {
	Start() error
}

type Notifier interface {
	Notify(*database.Event, *database.User, *database.Payload) error
}

func Init(cfg *Config, db *database.DB, log *logrus.Logger, tls *tls.Config, domain string) ([]Controller, []Notifier, error) {

	controllers := make([]Controller, 0)
	notifiers := make([]Notifier, 0)

	var (
		m   interface{}
		err error
	)

	for _, name := range cfg.Enabled {
		switch name {

		case "telegram":
			m, err = telegram.New(&cfg.Telegram, db, domain)

		case "api":
			m, err = api.New(&cfg.API, db, log, tls)

		}

		if err != nil {
			return nil, nil, fmt.Errorf("fail to create module %q: %w", name, err)
		}

		if c, ok := m.(Controller); ok {
			controllers = append(controllers, c)
		}

		if n, ok := m.(Notifier); ok {
			notifiers = append(notifiers, n)
		}

	}

	return controllers, notifiers, nil
}
