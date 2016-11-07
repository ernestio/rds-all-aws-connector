/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	ecc "github.com/ernestio/ernest-config-client"
	"github.com/nats-io/nats"
)

var nc *nats.Conn
var natsErr error

func eventHandler(m *nats.Msg) {
	var e Event

	err := e.Process(m.Subject, m.Data)
	if err != nil {
		println(err.Error())
		return
	}

	if err = e.Validate(); err != nil {
		e.Error(err)
		return
	}

	parts := strings.Split(m.Subject, ".")
	switch parts[1] {
	case "create":
		err = createRds(&e)
	case "update":
		err = updateRds(&e)
	case "delete":
		err = deleteRds(&e)
	}

	if err != nil {
		e.Error(err)
		return
	}

	e.Complete()
}

func createRds(ev *Event) (err error) {
	return err
}

func updateRds(ev *Event) (err error) {
	return err
}

func deleteRds(ev *Event) (err error) {
	return err
}

func main() {
	nc = ecc.NewConfig(os.Getenv("NATS_URI")).Nats()

	fmt.Println("listening for rds.create.aws")
	nc.Subscribe("rds.create.aws", eventHandler)

	fmt.Println("listening for rds.update.aws")
	nc.Subscribe("rds.update.aws", eventHandler)

	fmt.Println("listening for rds.delete.aws")
	nc.Subscribe("rds.delete.aws", eventHandler)

	runtime.Goexit()
}
