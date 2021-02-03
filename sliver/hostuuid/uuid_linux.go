//+ build linux

package hostuuid

/*
  Sliver Implant Framework
  Copyright (C) 2019  Bishop Fox

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"fmt"
	"io/ioutil"
)

func gethostuuid() string {
	uuid, err := ioutil.ReadFile("/var/lib/dbus/machine-id")
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s-%s-%s-%s-%s",
		uuid[0:8],
		uuid[8:12], uuid[12:16], uuid[16:20],
		uuid[20:32])
}
