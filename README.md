# Introduction

Frappe Neko (stylised `neko`) is an agent that can be put on any machine as a Go binary and be used to manage VMs.

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

# Note

This is currently a work in progress.
