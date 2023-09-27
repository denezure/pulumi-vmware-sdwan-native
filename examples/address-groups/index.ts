import * as pulumi from "@pulumi/pulumi";
import * as vsdw from "@pulumi/xyz";

let ag = new vsdw.resources.AddressGroup("pulumi-address-group", {
    addresses: ["192.168.10.1/32", "fd::/64"],
    domains: ["yahoo.com", "msn.com", "reddit.com", "google.com"],
}, {
    deleteBeforeReplace: true,
});

export const ag_id = ag.objectId;