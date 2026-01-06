#Get contexts 

```
kubectl config get-contexts
```

```
kubectl config view --raw
```

#To run falco use 

```
falco -U | grep "Rule description"
```

Configuration is generally stored in /etc/falco/falco.yaml file
and you will see the rules files and folders in this file

You can add/edit rules present in the file
Usually it is the custom file which is responsible for adding these

#Supported fields for outputs and conditions
https://falco.org/docs/reference/rules/supported-fields/

#APIServer
The definition for this file is in 
/etc/kubernetes/manifests

The container can take various arguments and that can be used to configure the 
API server. As an example, --authorization-mode=RBAC
or --enable-admission-plugins=NodeRestriction

uipath 
red cat 
amd
amazon 
serve robotics
ouster 
tesla
uber 
============

kube-bench 

Run as binary 

kube-bench run --targets=master 
kube-bench run --targets=node

Use grep -A -B and also use 
kube-bench run --check=1.3.2

=========
Updating kubelet Configuration, say we are asked to update containerLogMaxFiles to 5 

There are two steps involved here
a) Update the kubelet configuration configmap
b) Then upgrade nodes(control plane and worker node ) kubelet config
kubeadm upgrade node phase kubelet-config

Similarly for cluster config updates or kube proxy, core dns update 
there are two steps involved 

==============
AppArmor / SELinux / Seccomp 
Container sandbox runtime security 


AppArmor is a mandatory access control
Implemented as a linux kernel module 

The profile needs to be loaded on the node
use sudo apparmor_parser -q <path-to-profile>

apparmor_status => Shows all loaded profiles 

Once the apparmor profile is loaded on the node
label the node, make sure that the pod lands on this node
One can use nodeSelector easily for this 

Now we also apparmor profile in the security context of either the pod or the container
We will most likely be using the localhost profile 
use the name of localhost profile 

Also unconfined and runtime are two other options 

======
vim ~/.vimrc
Now enter (in insert-mode activated with i) the following lines:

set expandtab
set tabstop=2
set shiftwidth=2