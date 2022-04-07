# maven

## deploy

### 1. settings.xml
```
<servers>
    <server>
        <id>one</id>
        <username>admin</username>
        <password>123456</password>
    </server>

</servers>
```

### 2. pom.xml
```
<distributionManagement>
    <repository>
        <id>one</id>
        <name>Repository one</name>
        <url>http://one.com/repository/maven-public</url>
        <layout>default</layout>
        <uniqueVersion>false</uniqueVersion>
    </repository>

</distributionManagement>
```