# full-auto-gke

Google Kubernetes Engine cluster running in both node auto-provisioning
autoscaling mode and horizontal pod autoscaling.

As a result, the scaling should scale **extremely** well.

Pods are being created/destroyed based on CPU usage (could also be GPU, since
it is a Standard, non-Autopilot cluster) and nodes are scaled automatically
based on the usage.

The cluster could just as well configured to use the Autopilot, yet I am a fan
of specifying the node pools on one's own.

## Disclaimer

Same as [here](https://github.com/piotrostr/gke-gpu-tf/).

## Installation

Provisioning of the resources happens using Terraform

```sh
terraform apply
```

in the `full-auto-gke/terraform/` directory, followed by

```sh
kubectl apply -f manifest.yaml
```

to create the resources in the GKE cluster.

## License

MIT
