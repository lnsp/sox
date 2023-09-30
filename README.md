<p align="center">
<img src="docs/logo.png" height="100" alt="VirtM"><br>
Experimental virtual machine management system.
</p>

## Roadmap

- [x] Virtual machine creation and deletion
- [x] Aware of SSH keys, VM images and networks
- [x] Automatic image customization
- [x] IPv4 address assignment, internet access as well as routable from the host machine
- [x] Private IPv4 networks
- [ ] IPv6 address assignment
- [ ] Cross-node networking via VXLAN and dynamic peer resolution via l2miss, l3miss
- [ ] Aggregation mode with sox running on each node in a cluster, including draining and rebalancing nodes
- [ ] Advanced management of SSH keys and networks
- [ ] Advanced management of VM images, including creation of new images from running VMs
- [ ] Support for different users, projects and ACLs

## Design

The core idea is that sox exposes the same gRPC API, no matter if it's running on a single-node or in aggregation mode.

sox knows five core primitives: Machines, images, disks, networks and SSH keys.
Machines are made out of their image, the attached networks and disks and configured SSH keys.

There is a global IP space every machine gets a single IPv4/IPv6 from.