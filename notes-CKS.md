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

