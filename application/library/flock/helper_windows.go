// +build windows

/*
   Nging is a toolbox for webmasters
   Copyright (C) 2021-present Wenhui Shen <swh@admpub.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package flock

func LockEx(f *os.File) error {
	return nil
}

func LockExBlock(f *os.File) error {
	return nil
}

func Lock(f *os.File) error {
	return nil
}

func LockBlock(f *os.File) error {
	return nil
}

func Unlock(f *os.File) error {
	return nil
}

func UnlockAndClose(f *os.File) error {
	f.Close()
	return nil
}