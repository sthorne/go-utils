// Copyright 2014 Sean Thorne.  All rights reserved.
// Use of this code is governed by the license found in LICENSE
package hashed

import (

    "os"
    "testing"
    "io/ioutil"

)

func TestMd5String(t *testing.T) {
    s, e := md5("hello world")

    if e != nil {
        t.Error(e.Error())
    }

    t.Log(s)

    if s != "5eb63bbbe01eeed093cb22bb8f5acdc3" {
        t.Error("md5 string did not match expected")
    }
}

func TestMd5ByteSlice(t *testing.T) {
    s, e := md5([]byte("hello world"))

    if e != nil {
        t.Error(e.Error())
    }

    t.Log(s)

    if s != "5eb63bbbe01eeed093cb22bb8f5acdc3" {
        t.Error("md5 string did not match expected")
    }
}

func TestMd5File(t *testing.T) {

    ioutil.WriteFile("/tmp/md5File.txt", []byte(`hello world`), 0666)

    s, e := md5File("/tmp/md5File.txt")

    if e != nil {
        t.Error(e.Error())
    }

    t.Log(s)

    os.Remove("/tmp/md5File.txt")

    if s != "5eb63bbbe01eeed093cb22bb8f5acdc3" {
        t.Error("md5 string did not match expected")
    }
}

func TestSha256String(t *testing.T) {
    s, e := sha256("hello world")

    if e != nil {
        t.Error(e.Error())
    }

    t.Log(s)

    if s != "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9" {
        t.Error("sha256 string did not match expected")
    }
}

func TestSha256ByteSlice(t *testing.T) {
    s, e := sha256([]byte("hello world"))

    if e != nil {
        t.Error(e.Error())
    }

    t.Log(s)

    if s != "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9" {
        t.Error("sha256 string did not match expected")
    }
}

func TestSha256File(t *testing.T) {

    ioutil.WriteFile("/tmp/sha256File.txt", []byte(`hello world`), 0666)

    s, e := sha256File("/tmp/sha256File.txt")

    if e != nil {
        t.Error(e.Error())
    }

    t.Log(s)

    os.Remove("/tmp/sha256File.txt")

    if s != "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9" {
        t.Error("sha256 string did not match expected")
    }
}

func TestSha512String(t *testing.T) {
    s, e := sha512("hello world")

    if e != nil {
        t.Error(e.Error())
    }

    t.Log(s)

    if s != "309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f" {
        t.Error("sha512 string did not match expected")
    }
}

func TestSha512ByteSlice(t *testing.T) {
    s, e := sha512([]byte("hello world"))

    if e != nil {
        t.Error(e.Error())
    }

    t.Log(s)

    if s != "309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f" {
        t.Error("sha512 string did not match expected")
    }
}

func TestSha512File(t *testing.T) {

    ioutil.WriteFile("/tmp/sha512File.txt", []byte(`hello world`), 0666)

    s, e := sha512File("/tmp/sha512File.txt")

    if e != nil {
        t.Error(e.Error())
    }

    t.Log(s)

    os.Remove("/tmp/sha512File.txt")

    if s != "309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f" {
        t.Error("sha512 string did not match expected")
    }
}