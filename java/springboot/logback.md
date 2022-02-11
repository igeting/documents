# springboot logback

springboot默认的日志文件是不会自动按天分割的，所以生产环境的日志文件越来越大，很不利于排错。查了很多资料最终配置如下，可以完美按天按错误等级分割日志文件，配置如下。

由于springboot采用约定优先于配置的方式，日志文件也是，SpringBoot项目中在官方文档中https://docs.spring.io/spring-boot/docs/current/reference/html/boot-features-logging.html 说明，默认已经依赖了一些日志框架。而其中推荐使用的就是Logback，说明一下，SpringBoot已经依赖了Logback所以不需要手动添加依赖。

首先不同环境下的logback配置肯定是不一样的，所以我的解决办法是：

项目中的application.properties已经通过spring.profiles.active来分割成不同环境下使用不同的properties配置，比如application-dev.properties(开发环境)，application-test.properties(测试环境)，application-prod.properties（生产环境）,再加上application.properties（这个文件可能只包含spring.profiles.active就够了，真正的配置可能都在带-的文件里，因为springboot默认会加载它，然后通过它来指定使用哪个文件） 共有4个properties 文件，然后在application.properties 通过spring.profiles.active指定要使用的真正的properties

ok 上面是简单的聊了一下application.properties配置的问题，正戏来了

接下来就要创建出来 logback-spring-dev.xml，logback-spring-test.xml，logback-spring-prod.xml三个文件，然后在每个对应环境的properties通过 logging.config 来指定logback 的xml文件

logging.config=classpath:logback-spring-dev.xml

至于logback的xml文件内容如下：

logback-spring-dev.xml 开发环境下，不需要输出到文件，只需要打印在控制台就行了。
```
<?xml version="1.0" encoding="UTF-8"?>

<!-- ALL < TRACE < DEBUG < INFO < WARN < ERROR < FATAL < OFF -->
<configuration scan="true" scanPeriod="60 seconds" debug="false">
    <contextName>d1money-web-ys-ems</contextName>

    <appender name="STDOUT" class="ch.qos.logback.core.ConsoleAppender">
        <encoder>
            <pattern>
                %d{yyyy-MM-dd HH:mm:ss} [%thread] %-5level %logger -%msg%n
            </pattern>
        </encoder>
    </appender>

    <logger name="java.sql.PreparedStatement" value="DEBUG"/>
    <logger name="java.sql.Connection" value="DEBUG"/>
    <logger name="java.sql.Statement" value="DEBUG"/>
    <logger name="com.ibatis" value="DEBUG"/>
    <logger name="com.ibatis.common.jdbc.SimpleDataSource" value="DEBUG"/>
    <logger name="com.ibatis.common.jdbc.ScriptRunner" level="DEBUG"/>
    <logger name="com.ibatis.sqlmap.engine.impl.SqlMapClientDelegate" value="DEBUG"/>
    <logger name="com.apache.ibatis" level="TRACE"/>

    <root level="debug">
        <appender-ref ref="STDOUT"/>
    </root>
</configuration>
```

然后是 logback-spring-test.xml 和logback-spring-prod.xml，test环境与prod的只是文件位置不同所以只贴一份了，改下路径就行了，并且测试和生产是不需要控制台输出的，不然catalina.out文件就要爆炸了！
```
<?xml version="1.0" encoding="UTF-8"?>

<!-- 从高到地低 OFF 、 FATAL 、 ERROR 、 WARN 、 INFO 、 DEBUG 、 TRACE 、 ALL -->
<!-- 日志输出规则 根据当前ROOT 级别，日志输出时，级别高于root默认的级别时 会输出 -->
<!-- 以下 每个配置的 filter 是过滤掉输出文件里面，会出现高级别文件，依然出现低级别的日志信息，通过filter 过滤只记录本级别的日志 -->
<!-- 属性描述 scan：性设置为true时，配置文件如果发生改变，将会被重新加载，默认值为true scanPeriod:设置监测配置文件是否有修改的时间间隔，如果没有给出时间单位，默认单位是毫秒。当scan为true时，此属性生效。默认的时间间隔为1分钟。 
	debug:当此属性设置为true时，将打印出logback内部日志信息，实时查看logback运行状态。默认值为false。 -->
<configuration scan="true" scanPeriod="60 seconds" debug="false">
    <contextName>d1money-web-ys-ems</contextName>
    <!-- 定义日志文件 输入位置 -->
    <property name="log_dir" value="/soft/apache-tomcat-8.5.30-ems/logs"/>
    <!-- 日志最大的历史 30天 -->
    <property name="maxHistory" value="30"/>
    <property name="maxFileSize" value="10MB"/>
    <!-- ERROR级别日志 -->
    <!-- 滚动记录文件，先将日志记录到指定文件，当符合某个条件时，将日志记录到其他文件 RollingFileAppender -->
    <appender name="ERROR" class="ch.qos.logback.core.rolling.RollingFileAppender">
        <!-- 过滤器，只记录WARN级别的日志 -->
        <filter class="ch.qos.logback.classic.filter.LevelFilter">
            <level>ERROR</level>
            <onMatch>ACCEPT</onMatch>
            <onMismatch>DENY</onMismatch>
        </filter>
        <!-- 最常用的滚动策略，它根据时间来制定滚动策略.既负责滚动也负责出发滚动 -->
        <rollingPolicy class="ch.qos.logback.core.rolling.TimeBasedRollingPolicy">

            <!--日志输出位置 可相对、和绝对路径 -->
            <fileNamePattern>
                ${log_dir}/app_error.%d{yyyy-MM-dd}.%i.log
            </fileNamePattern>
            <!-- 可选节点，控制保留的归档文件的最大数量，超出数量就删除旧文件假设设置每个月滚动，且<maxHistory>是6， 则只保存最近6个月的文件，删除之前的旧文件。注意，删除旧文件是，那些为了归档而创建的目录也会被删除 -->
            <maxHistory>${maxHistory}</maxHistory>
            <timeBasedFileNamingAndTriggeringPolicy class="ch.qos.logback.core.rolling.SizeAndTimeBasedFNATP">
                <maxFileSize>${maxFileSize}</maxFileSize>
            </timeBasedFileNamingAndTriggeringPolicy>

        </rollingPolicy>
        <!-- 按照固定窗口模式生成日志文件，当文件大于20MB时，生成新的日志文件。窗口大小是1到3，当保存了3个归档文件后，将覆盖最早的日志。-->
        <!--<rollingPolicy class="ch.qos.logback.core.rolling.FixedWindowRollingPolicy">
            <fileNamePattern>${log_dir}/%d{yyyy-MM-dd}/.log.zip</fileNamePattern>
            <minIndex>1</minIndex>
            <maxIndex>3</maxIndex>
        </rollingPolicy>-->
        <!-- 查看当前活动文件的大小，如果超过指定大小会告知RollingFileAppender 触发当前活动文件滚动-->
        <!--<triggeringPolicy
                class="ch.qos.logback.core.rolling.SizeBasedTriggeringPolicy">
            <maxFileSize>5MB</maxFileSize>
        </triggeringPolicy>-->
        <encoder>
            <pattern>%d{yyyy-MM-dd HH:mm:ss.SSS} [%thread] %-5level %logger - %msg%n</pattern>
        </encoder>
    </appender>

    <!-- WARN级别日志 appender -->
    <appender name="WARN" class="ch.qos.logback.core.rolling.RollingFileAppender">
        <!-- 过滤器，只记录WARN级别的日志 -->
        <filter class="ch.qos.logback.classic.filter.LevelFilter">
            <level>WARN</level>
            <onMatch>ACCEPT</onMatch>
            <onMismatch>DENY</onMismatch>
        </filter>
        <rollingPolicy class="ch.qos.logback.core.rolling.TimeBasedRollingPolicy">
            <!-- 按天回滚 daily -->
            <fileNamePattern>
                ${log_dir}/app_warn.%d{yyyy-MM-dd}.%i.log
            </fileNamePattern>
            <!-- 日志最大的历史 30天 -->
            <maxHistory>${maxHistory}</maxHistory>
            <timeBasedFileNamingAndTriggeringPolicy class="ch.qos.logback.core.rolling.SizeAndTimeBasedFNATP">
                <maxFileSize>${maxFileSize}</maxFileSize>
            </timeBasedFileNamingAndTriggeringPolicy>
        </rollingPolicy>
        <encoder>
            <pattern>%d{yyyy-MM-dd HH:mm:ss.SSS} [%thread] %-5level %logger - %msg%n</pattern>
        </encoder>
    </appender>

    <!-- INFO级别日志 appender -->
    <appender name="INFO" class="ch.qos.logback.core.rolling.RollingFileAppender">
        <!-- 过滤器，只记录INFO级别的日志 -->
        <filter class="ch.qos.logback.classic.filter.LevelFilter">
            <level>INFO</level>
            <onMatch>ACCEPT</onMatch>
            <onMismatch>DENY</onMismatch>
        </filter>
        <rollingPolicy class="ch.qos.logback.core.rolling.TimeBasedRollingPolicy">
            <!-- 按天回滚 daily -->
            <fileNamePattern>
                ${log_dir}/app_info.%d{yyyy-MM-dd}.%i.log
            </fileNamePattern>
            <!-- 日志最大的历史 30天 -->
            <maxHistory>${maxHistory}</maxHistory>
            <timeBasedFileNamingAndTriggeringPolicy class="ch.qos.logback.core.rolling.SizeAndTimeBasedFNATP">
                <maxFileSize>${maxFileSize}</maxFileSize>
            </timeBasedFileNamingAndTriggeringPolicy>
        </rollingPolicy>
        <encoder>
            <pattern>%d{yyyy-MM-dd HH:mm:ss.SSS} [%thread] %-5level %logger - %msg%n</pattern>
        </encoder>
    </appender>

    <!-- DEBUG级别日志 appender -->
    <appender name="DEBUG" class="ch.qos.logback.core.rolling.RollingFileAppender">
        <!-- 过滤器，只记录DEBUG级别的日志 -->
        <filter class="ch.qos.logback.classic.filter.LevelFilter">
            <level>DEBUG</level>
            <onMatch>ACCEPT</onMatch>
            <onMismatch>DENY</onMismatch>
        </filter>
        <rollingPolicy class="ch.qos.logback.core.rolling.TimeBasedRollingPolicy">
            <!-- 按天回滚 daily -->
            <fileNamePattern>
                ${log_dir}/app_debug.%d{yyyy-MM-dd}.%i.log
            </fileNamePattern>
            <!-- 日志最大的历史 30天 -->
            <maxHistory>${maxHistory}</maxHistory>
            <timeBasedFileNamingAndTriggeringPolicy class="ch.qos.logback.core.rolling.SizeAndTimeBasedFNATP">
                <maxFileSize>${maxFileSize}</maxFileSize>
            </timeBasedFileNamingAndTriggeringPolicy>
        </rollingPolicy>
        <encoder>
            <pattern>%d{yyyy-MM-dd HH:mm:ss.SSS} [%thread] %-5level %logger - %msg%n</pattern>
        </encoder>
    </appender>

    <!-- TRACE级别日志 appender -->
    <appender name="TRACE" class="ch.qos.logback.core.rolling.RollingFileAppender">
        <!-- 过滤器，只记录ERROR级别的日志 -->
        <filter class="ch.qos.logback.classic.filter.LevelFilter">
            <level>TRACE</level>
            <onMatch>ACCEPT</onMatch>
            <onMismatch>DENY</onMismatch>
        </filter>
        <rollingPolicy class="ch.qos.logback.core.rolling.TimeBasedRollingPolicy">
            <!-- 按天回滚 daily -->
            <fileNamePattern>
                ${log_dir}/app_trace.%d{yyyy-MM-dd}.%i.log
            </fileNamePattern>
            <!-- 日志最大的历史 30天 -->
            <maxHistory>${maxHistory}</maxHistory>

            <timeBasedFileNamingAndTriggeringPolicy class="ch.qos.logback.core.rolling.SizeAndTimeBasedFNATP">
                <maxFileSize>${maxFileSize}</maxFileSize>
            </timeBasedFileNamingAndTriggeringPolicy>
        </rollingPolicy>
        <encoder>
            <pattern>%d{yyyy-MM-dd HH:mm:ss.SSS} [%thread] %-5level %logger - %msg%n</pattern>
        </encoder>
    </appender>

    <logger name="java.sql.PreparedStatement" value="DEBUG"/>
    <logger name="java.sql.Connection" value="DEBUG"/>
    <logger name="java.sql.Statement" value="DEBUG"/>
    <logger name="com.ibatis" value="DEBUG"/>
    <logger name="com.ibatis.common.jdbc.SimpleDataSource" value="DEBUG"/>
    <logger name="com.ibatis.common.jdbc.ScriptRunner" level="DEBUG"/>
    <logger name="com.ibatis.sqlmap.engine.impl.SqlMapClientDelegate" value="DEBUG"/>
    <logger name="com.apache.ibatis" level="TRACE"/>

    <!-- root级别 DEBUG -->
    <root level="debug">
        <!-- 文件输出 -->
        <appender-ref ref="ERROR"/>
        <appender-ref ref="INFO"/>
        <appender-ref ref="WARN"/>
        <appender-ref ref="DEBUG"/>
        <appender-ref ref="TRACE"/>
    </root>
</configuration>
```

最后达到的效果是按大小日期生成日志文件，超过设定大小就会重新生成新文件，由0开始累积。