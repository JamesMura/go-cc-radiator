## GO CC Radiator

Build Status Radiator compatible with cctray xml

### Goals 
* It should be easily configurable
* It should be easy to setup
* It should work well in resource constrained environments




[![wercker status](https://app.wercker.com/status/99137d90dfa982512fb3f103407f760e/s "wercker status")](https://app.wercker.com/project/bykey/99137d90dfa982512fb3f103407f760e)



## RUN USING DOCKER

docker pull jamesmura/cctray-radiator

mkdir /var/cctray

* create a config file /var/cctray/app.ini to match [this example configuration](conf/app.ini.sample)

docker run --name=cctray-radiator -d -p 8080:8080 -v /var/cctray:/go/src/app/conf jamesmura/cctray-radiator