# DaggerEnv

Load dotenv var/secrets within a container. 
Uses dagger Glob to find all `*.env`. 

If fil is named as `*.secret.env`, vars will be loaded as secrets 


*Usage*

```
c := dag.Container().From("alpine:latest")
c = dag.DaggerEnv().Load(c, src)
```

*Test* 

```
dagger call test --test-data=.dagger/test_envs
```