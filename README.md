# go-box

- [go-box](#go-box)
  - [Install](#install)
  - [Usage](#usage)
  - [Commands](#commands)
    - [**hash**](#hash)
    - [**mkdir**](#mkdir)
    - [**mv**](#mv)
    - [**cp**](#cp)

Provides several Unix like utilities in a single executable file

## Install

```terminal
go install github.com/NexizOne/go-box@latest
```

## Usage

```terminal
go-box mkdir path/to/directory
```

## Commands

<!-- [BusyBox](https://busybox.net/downloads/BusyBox.html) -->

### **hash**

hash `OPTION`

Returns hash of FILE or STRING

Options:

        -a      hashing algorithm (sha256, sha1, md5) (default: "sha256")
        -f      file to hash
        -s      string to hash

### **mkdir**

mkdir `OPTIONS` DIRECTORY...

Create DIRECTORY

Options:

        -m      mode

### **mv**

### **cp**
