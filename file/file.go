// Copyright 2014 Sean Thorne.  All rights reserved.
// Use of this code is governed by the license found in LICENSE
//
// some utility functions for dealing with files
package file

import (

    "os"

)

// utility function for seeing if a file exists
func FileExists(f string) bool {
    _, e := os.Stat(f)

    if os.IsNotExist(e) {
        return false
    }

    return true
}

// create an empty file
func Touch(f string) error {
    _, e := os.Create(f)

    if e != nil {
        return e
    }

    return nil
}