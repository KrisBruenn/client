package main
/*
A driver that runs the client postgresql wrapper.
    Copyright (C) 2023    Kris Bruenn     

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

Contact: klbruenn@gmail.com, or PO Box 2357, Santa Clara, CA, 95055.
*/

import "fmt"
import "klbrun.com/client"

func Disclaimer() {
    fmt.Println("****************************************************************************")
    fmt.Println("               Client  Copyright (C) 2023  Kris Bruenn")
    fmt.Println("****************************************************************************")
    fmt.Println("This program comes with ABSOLUTELY NO WARRANTY; for details see the GNU GPL.") 
    fmt.Println("This is free software, and you are welcome to redistribute it under")
    fmt.Println("certain conditions; see GNU General Public License for details.")
    fmt.Println("****************************************************************************")
}

func main() {
    Disclaimer()
    client.Client()
}
