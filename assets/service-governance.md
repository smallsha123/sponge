## 5 Service governance

### 5.1 Trace

#### 5.1.1 Starting jaeger and elasticsearch services

Use jaeger for tracing and elasticsearch for storage, and start both services locally using [docker-compose](https://github.com/docker/compose/releases).

**(1) elasticsearch service**

This is the [startup script for the elasticsearch service](https://github.com/zhufuyi/sponge/tree/main/test/server/elasticsearch), and the **.env** file is the startup configuration for elasticsearch, starting the elasticsearch service.

> docker-compose up -d

<br>

**(2) jaeger services**

This is the [jaeger service startup script](https://github.com/zhufuyi/sponge/tree/main/test/server/jaeger), the **.env** file is to configure the jaeger information and start the jaeger service at.

> docker-compose up -d

Visit the jaeger query home page in your browser [http://localhost:16686](http://localhost:16686) .

<br>

#### 5.1.2 Single Service Trace Example

Using the http service code created in **Section 3.1.2** as an example, modify the configuration file `configs/edusys.yml` to enable trace (field enableTrace) and fill in the jaeger configuration information.

If you want to track redis, enable redis caching, change the cache type field **cacheType** value to redis and configure the redis configuration, and start the redis service locally using docker, which is [redis service startup script](https://github.com/zhufuyi/ sponge/tree/main/test/server/redis).

Start the http service.

```bash
# Compile and run services
make run
```

Copy [http://localhost:8080/swagger/index.html](http://localhost:8080/apis/swagger/index.html) to the browser to access the swagger home page to request the get query as an example, requesting the same id twice in a row, with the trace shown in Figure 5-1.

![one-server-trace](https://raw.githubusercontent.com/zhufuyi/sponge/main/assets/one-server-trace.jpg)
*Figure 5-1 single service trace diagram*

<br>

From the figure you can see that the first request has 4 spans, which are.

- Request api `/api/v1/teacher/1`
- Querying redis
- Query mysql
- Setting up the redis cache

It means that the first request looks up from redis, does not hit the cache, then reads the data from mysql and finally places the cache.

The second request has only 2 spans, which are.

- Request api `/api/v1/teacher/1`
- Querying redis

It means that the second request hits the cache directly, with less querying mysql and setting up the cache process than the first.

These spans are automatically generated, many times you need to manually add custom spans, add span example.

```go
import "github.com/zhufuyi/sponge/pkg/tracer"

tags := map[string]interface{}{"foo": "bar"}
_, span := tracer.NewSpan(ctx, "spanName", tags)  
defer span.End()
```

<br>

#### 5.1.3 Multi-Service Trace Example

Take the rpc gateway service code generated by **section 4.3** as an example, a total of four services **shopgw**, **product**, **inventory**, **comment**, modify the configuration of each of the four services (in the configs directory), turn on tracing, and fill in the jaeger configuration information.

Find the template file in the **internal/service** directory of the **product**, **inventory**, and **comment** services, populate the code in place of `panic("improve me")` to make the code execute properly, and manually add a **span** that Add a random delay.

Start the four services **shopgw**, **product**, **inventory**, and **comment**, and visit [http://localhost:8080/apis/swagger/index.html](http://localhost:8080/apis/swagger/index.html) in the browser to execute a get request, and the trace page is shown in Figure 5-2.

![multi-servers-trace](https://raw.githubusercontent.com/zhufuyi/sponge/main/assets/multi-servers-trace.jpg)
*Figure 5-2 multi-service trace diagram*

As you can see from the figure there are 10 spans with the main links.

- Request api `/api/v1/detail`
- The shopgw service calls the product client
- product's rpc server
- mockDAO added manually in the product service
- shopgw service calls the inventory client
- inventory's rpc server
- manually added mockDAO in the inventory service
- The shopgw service calls the comment client
- comment's rpc server
- mockDAO added manually in the comment service

The shopgw service calls **product**, **inventory**, and **comment** services serially to get data, but in practice it would be more time efficient to call them in parallel, but be careful to control the number of concurrent processes.

<br>

### 5.2 Monitoring

#### 5.2.1 Starting Prometheus and Grafana Services

Use [Prometheus](https://prometheus.io/docs/introduction/overview) for collecting metrics and [Grafana](https://grafana.com/docs/) for display, and start both services locally using docker.

**(1) prometheus services**

This is the [prometheus service startup script](https://github.com/zhufuyi/sponge/tree/main/test/server/monitor/prometheus) that starts the prometheus service.

> docker-compose up -d

Visit the prometheus homepage in your browser [http://localhost:9090](http://localhost:9090/) .

<br>

**(2) grafana services**

This is the [grafana service startup script](https://github.com/zhufuyi/sponge/tree/main/test/server/monitor/grafana) that starts the grafana service.

> docker-compose up -d

Visit the grafana main page [http://localhost:33000](http://localhost:33000) in your browser, set the datasource for prometheus `http://192.168.3.37:9090`, remember the datasource name for prometheus (here it is **Prometheus**), and the **datasource** value for the json imported into the monitoring panel later should be the same.

<br>

#### 5.2.2 http service monitoring

As an example, the http service code generated by **Section 3.1.2** provides the indicator api [http://localhost:8080/metrics](http://localhost:8080/metrics) by default.

**(1) Add monitoring targets to prometheus**

Open the prometheus configuration file prometheus.yml and add the acquisition target.

```bash
  - job_name: 'http-edusys'
    scrape_interval: 10s
    static_configs:
      - targets: ['localhost:8080']
```

Note: If you use vim to modify the prometheus.yml file, you must change the file prometheus.yml permissions to `0777` before modifying it, otherwise the modification configuration file cannot be synchronized to the container.

Execute the request to make the prometheus configuration take effect `curl -X POST http://localhost:9090/-/reload`, wait a moment, and then visit [http://localhost:9090/targets](http://localhost:9090/targets) in your browser to check if the newly added capture target takes effect.

<br>

**(2) Adding a monitoring panel to grafana**

Import the [http monitoring panel](https://github.com/zhufuyi/sponge/blob/main/pkg/gin/middleware/metrics/gin_grafana.json) into grafana, if no data is displayed in the monitoring interface, check that the data source in the json name is the same as the grafana configuration prometheus data source name.

<br>

**(3) Compression test api, observation of monitoring data**

Using the [wrk](https://github.com/wg/wrk) tool to pressure test the api

```bash
# Interface 1
wrk -t2 -c10 -d10s http://192.168.3.27:8080/api/v1/teacher/1

# Interface 2
wrk -t2 -c10 -d10s http://192.168.3.27:8080/api/v1/course/1
```

The monitoring interface is shown in Figure 5-3.

![http-grafana](https://raw.githubusercontent.com/zhufuyi/sponge/main/assets/http-grafana.jpg)
*Figure 5-3 http service monitoring diagram*

<br>

#### 5.2.3 rpc service monitoring

As an example, the rpc service code generated by **Section 4.1.1** provides the indicator api [http://localhost:8283/metrics](http://localhost:8283/metrics) by default.

**(1) Add monitoring targets to prometheus**

Open the prometheus configuration file prometheus.yml and add the acquisition target.

```bash
  - job_name: 'rpc-server-edusys'
    scrape_interval: 10s
    static_configs:
      - targets: ['localhost:8283']
```

Execute the request to make the prometheus configuration take effect `curl -X POST http://localhost:9090/-/reload`, wait a moment, and then visit [http://localhost:9090/targets](http://localhost:9090/targets) in your browser to check if the newly added capture target is in effect.

<br>

**(2) Adding a monitoring panel to grafana**

Import the [rpc server monitoring panel](https://github.com/zhufuyi/sponge/blob/main/pkg/grpc/metrics/server_grafana.json) into grafana, and if no data is displayed in the monitoring interface, check that the data source name in the json is the same as the grafana configuration prometheus data source name is the same.

<br>

**(3) Compression test rpc methods, observation of monitoring data**

Open the `internal/service/teacher_client_test.go` file using **Goland** or **VS Code** and test each method under **Test_teacherService_methods** or **Test_teacherService_benchmark**.

The monitoring interface is shown in Figure 5-4.
![rpc-grafana](https://raw.githubusercontent.com/zhufuyi/sponge/main/assets/rpc-grafana.jpg)
*Figure 5-4 rpc server monitoring diagram*

<br>

The above is the monitoring of the rpc server, and the monitoring of the rpc client is similar, [rpc client monitoring panel](https://github.com/zhufuyi/sponge/blob/main/pkg/grpc/metrics/client_grafana.json) .

<br>

#### 5.2.4 Automatic addition and removal of monitoring targets in prometheus

Prometheus supports dynamic configuration using consul's service registration and discovery to automatically add and remove monitoring targets.

Start the consul service locally, this is the [consul service startup script](https://github.com/zhufuyi/sponge/tree/main/test/server/consul)

Open prometheus config prometheus.yml and add the consul configuration.

```yaml
  - job_name: 'consul-micro-exporter'
    consul_sd_configs:
      - server: 'localhost:8500'
        services: []  
    relabel_configs:
      - source_labels: [__meta_consul_tags]
        regex: . *edusys.*
        action: keep
      - regex: __meta_consul_service_metadata_(. +)
        action: labelmap
```

Execute the request to make the prometheus configuration take effect `curl -X POST http://localhost:9090/-/reload`.

After the consul service discovery is configured in prometheus, the address information of the service is then pushed to consul, and the push information edusys_exporter.json file reads as follows.

```json
{
  "ID": "edusys-exporter",
  "Name": "edusys",
  "Tags": [
    "edusys-exporter"
  ],
  "Address": "localhost",
  "Port": 8283,
  "Meta": {
    "env": "dev",
    "project": "edusys"
  },
  "EnableTagOverride": false,
  "Check": {
    "HTTP": "http://localhost:8283/metrics",
    "Interval": "10s"
  },
  "Weights": {
    "Passing": 10,
    "Warning": 1
  }
}
```

> curl -XPUT --data @edusys_exporter.json http://localhost:8500/v1/agent/service/register

Wait a moment, then open [http://localhost:9090/targets](http://localhost:9090/targets) in your browser to check if the newly added acquisition target is in effect. Then close the service, wait a while, and check if the acquisition target is automatically removed.

<br>

For your own services, you usually submit information to consul at the same time you start the service, convert edusys_exporter.json to a go struct, and call the http client inside the program to submit it to consul.

<br>

### 5.3 Collecting go program profiles

Usually use the pprof tool to find and locate program problems, especially online go program problems can automatically save the program run site (profile), and then use the tool pprof analysis to locate the problem.

The sponge generated service supports **http api** and **system signal notification** to collect profiles, the system signal notification method is enabled by default, just use one.

<br>

#### 5.3.1 Collecting profiles via http

Through the http api way to collect profile is closed by default, if you need to open, modify the configuration of the field `enableHTTPProfile` to true, usually used in development or testing, if the line open will have a little performance loss, according to the actual situation whether to open the use.

The default route `/debug/pprof`, combined with the **go tool pprof** tool, allows you to analyze the current running status of your program at any moment.

<br>

#### 5.3.2 Notification of acquisition profiles via system signals

Using the http api, the program background has been regularly recording profile-related information, etc., the vast majority of the time will not read these profiles, can be improved, only when needed and then start collecting profiles, automatically shut down after collection, sponge generated services support listening to the system signal to start and stop collecting profiles, the default uses **SIGTRAP**(5) system signals (suggested to be changed to SIGUSR1, not supported in windows environment), sends signals to the service.

```bash
# View service pid by name (second column)
ps aux | grep service name

# Send a signal to the service
kill -trap pid value

# kill -usr1 pid value
```

After the service receives the system signal notification, it starts to collect the profile and save it to the `/tmp/service_name_profile` directory, the default collection length is 60 seconds, after 60 seconds it automatically stops collecting the profile, if you only want to collect 30 seconds, send the first signal to start collecting, about 30 seconds later send the second signal to indicate that it stops collecting the profile, similar to switch. Default acquisition **cpu**, **memory**, **goroutine**, **block**, **mutex**, **threadcreate** six types of profiles, file format `date_time_pid_service_name_profile_type.out`, example.

```
xxx221809_58546_edusys_cpu.out
xxx221809_58546_edusys_mem.out
xxx221809_58546_edusys_goroutine.out
xxx221809_58546_edusys_block.out
xxx221809_58546_edusys_mutex.out
xxx221809_58546_edusys_threadcreate.out
```

Because the profile file of trace is relatively large, it is not captured by default and can be turned on to capture trace according to actual needs (call prof.EnableTrace() when the service starts).

Once the offline files are obtained, they are analysed using the pprof tool using an interactive or interface approach.

```bash
# Interactive
go tool pprof [options] source

# Interface
go tool pprof -http=[host]:[port] [options] source
```

<br>

#### 5.3.3 Automatic profile capture

All of the above are manual profile collection, and it is usually desirable to automatically collect profiles when problems occur. sponge-generated services support automatic profile collection by default, and are implemented in conjunction with the alerting feature of resource statistics. alerting conditions.

- Record the program's cpu usage 3 times in a row (once per minute by default) and trigger an alarm when the average usage exceeds 80% for 3 times.
- Record the program's use of physical memory for 3 consecutive times (once per minute by default), and trigger an alarm when the average system memory usage exceeds 80% for 3 times.
- If the alarm threshold is continuously exceeded, the default interval is 15 minutes between alarms.

When the alarm is triggered, the program internally calls the kill function to send x system signal to notify the acquisition profile, and the acquired profile file is saved to the `/tmp/service name_profile` directory, which is actually the basis of **notifying the acquisition profile by system signal** to change the manual trigger to automatic trigger, even in the middle of the night the program's cpu or memory is too high, the next day you can also analyze the profile to find out where the program is causing the cpu or memory to be too high.

Note: Automatic profile collection is not suitable for windows environment.

<br>

### 5.4 Registration Centre

The sponge-generated services support the [Nacos](https://nacos.io/zh-cn/docs/v2/what-is-nacos.html) configuration center by default. The role of the configuration center is to unify the configuration management of different environments and services, effectively solving the shortcomings of the ground static configuration.

Start the nacos service locally, this is the [nacos service startup configuration](https://github.com/zhufuyi/sponge/tree/main/test/server/nacos), after starting the nacos service, open the administration interface in your browser http://localhost:8848/nacos/index.html , login to the account password to enter the main interface.

Using the http service code generated by **Section 3.1.2** as an example using Configuration Center nacos, create a namespace `edusys` in the nacos interface, then create a new configuration with a Data ID value of `edusys.yml`, a Group value of `dev`, and a configuration content value of the `configs/edusys.yml` file contents, as shown in Figure 5-3.

![nacos-config](https://raw.githubusercontent.com/zhufuyi/sponge/main/assets/nacos-config.jpg)
*Figure 5-3 nacos add service configuration diagram*

Open the configuration center file `configs/edusys_cc.yml` in the edusys directory and fill in the nacos configuration information.

```yaml
# Generate the go struct command: sponge config --server-dir=./serverDir

# nacos settings
nacos:
  ipAddr: "192.168.3.37" # server address
  port: 8848 # listening port
  scheme: "http" # http or https
  contextPath: "/nacos" # path
  namespaceID: "ecfe0595-cae3-43a2- 9e47-216dc92207f9" # namespace id
  group: "dev" # group name: dev, prod, test
  dataID: "edusys.yml" # config file id
  format: "yaml" # configuration file type: json,yaml,toml
```

Compiling and starting edusys services.

```bash
# Switch to the main.go location
cdd cmd/edusys

# Compilation
go build

# Run
./edusys -enable-cc -c=../../configs/edusys_cc.yml
```

The start service parameter `-c` indicates that a configuration file is specified, and the parameter `-enable-cc` indicates that the configuration is obtained from the configuration center.

<br>

### 5.5 Rate limiter and circuit breaker

The service created by sponge supports rate limiter and circuit breaker, which is off by default. Open the service configuration file and modify the field **enableLimit** to a value of `true` to enable rate limiter, and modify the field **enableCircuitBreaker** to `true` to enable fusing.

Rate limiter and circuit breaker use a third-party library [aegis](https://github.com/go-kratos/aegis), which adapts according to system resources and error rates. Since different servers have different processing capabilities and parameters are not well set, using adaptive parameters avoids the trouble of manually setting parameters for each service.

<br>