# kversion
Small experiment with kubectl versions binaries

# Usage
Create a file in your $HOME folder named
> .kversion.yaml

with the following format:
```yaml
---
binaries:
 - /usr/local/bin/kubectl-1.14.8
 - /usr/local/bin/kubectl-1.14.6
 - /usr/local/bin/kubectl-1.14.5
 - /usr/local/bin/kubectl-1.13.7
```

