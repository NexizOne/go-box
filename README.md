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

```
NAME:
   go-box hash - Hash string or file

USAGE:
   go-box hash [options]

OPTIONS:
   --algorithm,  -a  hashing algorithm (sha256, sha1, md5) (default: "sha256")
   --file,       -f  file to hash
   --string,     -s  string to hash
   --help,       -h  show help
```

### **mkdir**

```
NAME:
   go-box mkdir - Creates directory

USAGE:
   go-box mkdir [directory] [options]

OPTIONS:
   --mode,  -m  mode (default: 777)
   --help,  -h  show help
```

### **mv**

```
NAME:
   go-box mv - Mode file or directory

USAGE:
   go-box mv [from] [to] [options]

DESCRIPTION:
   from    source (masks supported, example: *.txt)
   to      destination

OPTIONS:
   --force,  -f  force (default: false)
   --help,   -h  show help
```

### **cp**

```
NAME:
   go-box cp - Copy file or directory

USAGE:
   go-box cp [from] [to] [options]

DESCRIPTION:
   from    source (masks supported, example: *.txt)
   to      destination

OPTIONS:
   --force,  -f  force (default: false)
   --help,   -h  show help
```
