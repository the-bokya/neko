# Introduction

Frappe Neko (stylised `neko`) is an agent that can be put on any machine as a Go binary and be used to manage VMs.

# Prerequisites

- A proper [installation and configuration of libvirt and KVM/QEMU](https://wiki.archlinux.org/title/Libvirt)
- Root access (all commands should be run in root). This is mostly temporary.

# Installation

```bash
git clone github.com/the-bokya/neko.git
cd neko
go install
```

# Usage

1. Generate basic config:
   
   ```bash
   neko setup config
   ```
   
   This will create a config at `/etc/neko`

2. Download images:
   
   ```bash
   neko setup images
   ```
   
   Images specified in `/etc/neko/config.json` in the field `vm_images` will be downloaded and their sha256sum will be verified.

# Endpoints

1. Define new VMs

```bash
curl localhost:8000/new/vm --json '{"name": "new-vm", "vcpus": 1, "memory": 512, "image": "Ubuntu 24.04", "disk_size": 5}'
```

Output:

```bash
{"status":"Ok","message":"VM created successfully","data":{"name":"new-vm","uuid":"8d715798-663f-4cca-b6b3-0278ee772f1e"}}
```

# Note

This is currently a work in progress.
