import * as pulumi from "@pulumi/pulumi";
import * as vsdw from "@pulumi/xyz";

const portDef = (start: number, stop?: number) => {
    return {portStart: start, portEnd: stop ?? start} as vsdw.types.input.resources.ServiceGroupTcpArgs
}

let testAg = new vsdw.resources.AddressGroup("pulumi-addresses-group", {
    prefixes: ["192.168.10.1/32", "fd::/64"],
    domains: ["yahoo.com", "msn.com", "reddit.com", "google.com"],
}, {
    deleteBeforeReplace: true,
});

let webSg = new vsdw.resources.ServiceGroup("pulumi-web-service-group", {
    tcp: [
        portDef(80), portDef(443)
    ]
})

export const testAgId = testAg.addressGroupId;
export const webSgId = webSg.serviceGroupId;