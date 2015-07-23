/* Copyright (C) 2015 by Alexandru Cojocaru */

/* This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>. */

package main // import "xojoc.pw/sss"

import (
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	p := ":8080"
	if len(os.Args) > 1 {
		p = os.Args[1]
	}
	log.SetPrefix(path.Base(os.Args[0]) + ": ")
	log.SetFlags(log.Lshortfile)
	log.Fatal(http.ListenAndServe(p, http.FileServer(http.Dir("."))))
}
