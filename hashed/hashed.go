// Copyright 2014 Sean Thorne.  All rights reserved.
// Use of this code is governed by the license found in LICENSE
//
// some utility functions for dealing with hash-ing
package hashed

import (

    "os"
    "fmt"
    "hash"
    "bufio"
    "reflect"
    "encoding/hex"
    "encoding/binary"
    m "crypto/md5"
    sh "crypto/sha256"
    sh5 "crypto/sha512"

)

// convert the bytes or string
func readBytes(i interface{}) ([]byte, error) {

    var s []byte

    switch i.(type) {
    case []byte:
        s = reflect.ValueOf(i).Bytes()

    case string:
        s = []byte(reflect.ValueOf(i).String())

    case int:
        binary.PutVarint(s, reflect.ValueOf(i).Int())

    case uint:
        binary.PutUvarint(s, reflect.ValueOf(i).Uint())

    default:
        return s, fmt.Errorf("%s", `Cannot use interface type "` + reflect.TypeOf(i).Name() + `" given in hash`)
    }

    return s, nil
}

func readFile(f string, h hash.Hash) error {

    _, e := os.Stat(f)

    if os.IsNotExist(e) {
        return e
    }

    p, e := os.Open(f)

    if e != nil {
        return e
    }

    s := bufio.NewScanner(p)

    for s.Scan() {
        h.Write(s.Bytes())
    }

    if e := s.Err(); e != nil {
        return e
    }

    return nil
}

func work(h hash.Hash) (string) {

    b := h.Sum(nil)

    return hex.EncodeToString(b)
}

// create an md5 string
// will accept either a byte slice or string
func md5(i interface{}) (string, error) {

    s, e := readBytes(i)

    if e != nil {
        return "", e
    }

    h := m.New()

    h.Write(s)

    return work(h), nil
}

// create an md5 string from a file
// will accept either a byte slice or string
func md5File(f string) (string, error) {

    h := m.New()

    readFile(f, h)

    return work(h), nil
}

// create an sha256 string
// will accept either a byte slice or string
func sha256(i interface{}) (string, error) {

    s, e := readBytes(i)

    if e != nil {
        return "", e
    }

    h := sh.New()

    h.Write(s)

    return work(h), nil
}

// create an sha256 string from a file
// will accept either a byte slice or string
func sha256File(f string) (string, error) {

    h := sh.New()

    readFile(f, h)

    return work(h), nil
}

// create an sha256 string
// will accept either a byte slice or string
func sha512(i interface{}) (string, error) {

    s, e := readBytes(i)

    if e != nil {
        return "", e
    }

    h := sh5.New()

    h.Write(s)

    return work(h), nil
}

// create an sha512 string from a file
// will accept either a byte slice or string
func sha512File(f string) (string, error) {

    h := sh5.New()

    readFile(f, h)

    return work(h), nil
}