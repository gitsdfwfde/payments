/*
 * Copyright (C) 2020 The "MysteriumNetwork/payments" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package client

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReconnectableEthClientCreatesNewClient(t *testing.T) {
	client, err := NewReconnectableEthClient("http://127.0.0.1:1234", time.Second)
	assert.Nil(t, err)

	c1 := client.Client()
	c2 := client.Client()
	assert.Equal(t, c1, c2)

	a1 := client.Address()
	a2 := client.Address()
	assert.Equal(t, a1, a2)

	err = client.Reconnect(time.Second)
	assert.Nil(t, err)

	c3 := client.Client()
	assert.Equal(t, c1, c3)
	assert.Equal(t, c2, c3)

	a3 := client.Address()
	assert.Equal(t, a1, a3)
	assert.Equal(t, a2, a3)
}
