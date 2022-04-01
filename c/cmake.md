# cmake

## notes
```
# notes
```

## command
```
command(arg1 arg2 ...)
```

## var
```
set(FOO abc)
command(${FOO}) //command(abc)
command("${FOO}") //command("abc")
```

## flow control
```
if()...else()/elseif()...endif()
while()...endwhile()
foreach()...endforeach()
```

## command list
```
include_directories("dir1" "dir2" ...) //-Idir1 -Idir2
link_directories("dir1" "dir2")
aux_source_directory("source-dir" variable)
add_executable()
add_library()
add_custom_target()
add_dependencies(target1 t2 t3)
add_definitions("-Wall -ansi")
target_link_libraries(target-name lib1 lib2 ...)
link_libraries(lib1 lib2 ...)
set_target_properties(...)
message(...)
install(files "f1" "f2" destination)
set(variable value)
list(append|insert|length|get|remove_item|remove_at|sort ...)
string(toupper|tolower|length|substring|replace|regex ...)
separate_arguments(variable)
file(write|read|append|glob|glob_recurse|remove|make_directory ...)
find_file()
find_path()
find_library()
find_program()
find_package()
exec_program(bin [work_dir] args <..> [output_variable var] [return_value var])
option(option_var "description" [initial value])
```

## variable

### project path
```
CMAKE_SOURCE_DIR
PROJECT_SOURCE_DIR
<projectname>_SOURCE_DIR
```

### root path
```
CMAKE_BINARY_DIR
PROJECT_BINARY_DIR
<projectname>_BINARY_DIR
```

### current path
```
CMAKE_CURRENT_SOURCE_DIR
CMAKE_CURRENT_BINARY_DIR
CMAKE_CURRENT_LIST_FILE
```

### CMakeLists.txt path
```
CMAKE_BUILD_TYPE
```

## debug or release
```
set(CMAKE_BUILD_TYPE Debug)
cmake DCMAKE_BUILD_TYPE=Release
```

## build args
```
CMAKE_C_FLAGS
CMAKE_CXX_FLAGS
CMAKE_INCLUDE_PATH
CMAKE_LIBRARY_PATH
CMAKE_MODULE_PATH
CMAKE_INSTALL_PREFIX
BUILD_SHARED_LIBS
```

## example 1
```
cmake_minimum_required(VERSION 3.15)
project(c_demo)

include_directories(header)
set(SRC src/main.cpp)
aux_source_directory(src SRC)
add_executable(c_demo ${SRC})
```

# cmake
## INCLUDE_DIRECTORIES（添加头文件目录）
它相当于g++选项中的-I参数的作用，也相当于环境变量中增加路径到CPLUS_INCLUDE_PATH变量的作用。
```
include_directories("/opt/MATLAB/R2012a/extern/include")
export CPLUS_INCLUDE_PATH=CPLUS_INCLUDE_PATH:$MATLAB/extern/include
```

## LINK_DIRECTORIES（添加需要链接的库文件目录）
```
link_directories(directory1 directory2 ...)
```

它相当于g++命令的-L选项的作用，也相当于环境变量中增加LD_LIBRARY_PATH的路径的作用。

```
LINK_DIRECTORIES("/opt/MATLAB/R2012a/bin/glnxa64")
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:$MATLAB/bin/glnxa64
```

## LINK_LIBRARIES（添加需要链接的库文件路径，注意这里是全路径）
```
LINK_LIBRARIES("/opt/MATLAB/R2012a/bin/glnxa64/libeng.so")
LINK_LIBRARIES("/opt/MATLAB/R2012a/bin/glnxa64/libmx.so")
```
or
```
LINK_LIBRARIES("/opt/MATLAB/R2012a/bin/glnxa64/libeng.so"　"/opt/MATLAB/R2012a/bin/glnxa64/libmx.so")
```


## TARGET_LINK_LIBRARIES（设置要链接的库文件的名称）
```
TARGET_LINK_LIBRARIES(targetlibrary1 <debug | optimized> library2 ..)
```

```
TARGET_LINK_LIBRARIES(myProject hello)，连接libhello.so库
TARGET_LINK_LIBRARIES(myProject libhello.a)
TARGET_LINK_LIBRARIES(myProject libhello.so)
```
or
```
TARGET_LINK_LIBRARIES(myProject libeng.so)
TARGET_LINK_LIBRARIES(myProject eng)
TARGET_LINK_LIBRARIES(myProject -leng)
```

## example
```
cmake_minimum_required(VERSION 3.0 FATAL_ERROR)
 
include_directories("/opt/MATLAB/R2012a/extern/include")
 
#directly link to the libraries.
link_libraries("/opt/MATLAB/R2012a/bin/glnxa64/libeng.so")
link_libraries("/opt/MATLAB/R2012a/bin/glnxa64/libmx.so")
 
#equals to below
#link_libraries("/opt/MATLAB/R2012a/bin/glnxa64/libeng.so" "/opt/MATLAB/R2012a/bin/glnxa64/libmx.so")
 
add_executable(myProject main.cpp) 
```

```
cmake_minimum_required(VERSION 3.0 FATAL_ERROR)
 
include_directories("/opt/MATLAB/R2012a/extern/include")

link_directories("/opt/MATLAB/R2012a/bin/glnxa64")
 
add_executable(myProject main.cpp)
 
target_link_libraries(myProject eng mx)
 
#equals to below
#target_link_libraries(myProject -leng -lmx)
#target_link_libraries(myProject libeng.so libmx.so)
```
